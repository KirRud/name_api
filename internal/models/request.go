package models

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	operatorList = [6]string{"eq", "ne", "gt", "ge", "lt", "le"}
)

type Operator struct {
	code string
}

func (o Operator) GetSymbol() string {
	var symbol string
	switch o.code {
	case "eq":
		symbol = "="
	case "ne":
		symbol = "<>"
	case "gt":
		symbol = ">"
	case "ge":
		symbol = ">="
	case "lt":
		symbol = "<"
	case "le":
		symbol = "<="
	}
	return symbol
}

func (o Operator) IsOperator() bool {
	for _, operator := range operatorList {
		if o.code == operator {
			return true
		}
	}
	return false
}

type AgeParam struct {
	age int
	Operator
}

type GenderParam struct {
	gender string
}

func (gp GenderParam) IsGender() bool {
	if gp.gender == "male" || gp.gender == "female" {
		return true
	}
	return false
}

type RequestParam struct {
	AgeParam
	GenderParam
	page int
}

func NewRequestParam(ageStr, genderStr, pageStr string) (RequestParam, error) {
	ageSplitStr := strings.Split(ageStr, ":")
	if len(ageSplitStr) != 2 {
		return RequestParam{}, fmt.Errorf("wrong age params")
	}
	operator := ageSplitStr[0]
	age, err := strconv.Atoi(ageSplitStr[1])
	if err != nil {
		return RequestParam{}, fmt.Errorf("wrong age params")
	}
	ageParam := AgeParam{age, Operator{code: operator}}
	if !ageParam.IsOperator() {
		return RequestParam{}, fmt.Errorf("wrong age operator")
	}

	genderParam := GenderParam{genderStr}
	if !genderParam.IsGender() {
		return RequestParam{}, fmt.Errorf("wrong gender")
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return RequestParam{}, fmt.Errorf("wrong page")
	}

	return RequestParam{ageParam, genderParam, page}, nil
}

func (rp RequestParam) GetSqlWhere() string {
	return "age " + rp.GetSymbol() + " ? AND gender = ?"
}

func (rp RequestParam) GetParams() (int, string) {
	return rp.age, rp.gender
}

func (rp RequestParam) GetPage() int {
	return rp.page
}
