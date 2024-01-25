package models

type AgeInfo struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type GenderInfo struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type NationInfo struct {
	Count     int           `json:"count"`
	Name      string        `json:"name"`
	Countries []CountryInfo `json:"country"`
}

type CountryInfo struct {
	Country     string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
