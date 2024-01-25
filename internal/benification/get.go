package benification

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"name_api/internal/models"
	"net/http"
	"strings"
)

type ProviderInterface interface {
	GetInfo(name string) models.PersonInfo
	GetAgeByName(ch chan int, name string)
	GetGenderByName(ch chan string, name string)
	GetNationByName(ch chan string, name string)
}

type Provider struct {
	client *http.Client
}

func NewProvider(client *http.Client) ProviderInterface {
	return Provider{client: client}
}

func (p Provider) GetInfo(name string) models.PersonInfo {
	ageCh := make(chan int)
	genderCh := make(chan string)
	nationCh := make(chan string)

	go p.GetAgeByName(ageCh, name)
	go p.GetGenderByName(genderCh, name)
	go p.GetNationByName(nationCh, name)

	pInfo := models.PersonInfo{
		Name:   name,
		Age:    <-ageCh,
		Gender: <-genderCh,
		Nation: <-nationCh,
	}
	log.Debugf("benification.GetInfo; person info: %v", pInfo)

	return pInfo
}

func (p Provider) GetAgeByName(ageCh chan int, name string) {
	defer close(ageCh)
	url := buildAgeURL(name)
	resp, err := p.client.Get(url)
	if err != nil {
		log.Errorf("Error get request for age: %s\n", err)
		return
	}
	defer resp.Body.Close()

	var data models.AgeInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Errorf("Error decode json for age: %s\n", err)
		return
	}

	ageCh <- data.Age
}

func (p Provider) GetGenderByName(genderCh chan string, name string) {
	defer close(genderCh)
	url := buildGenderURL(name)
	resp, err := p.client.Get(url)
	if err != nil {
		log.Errorf("Error get request for gender: %s\n", err)
		return
	}
	defer resp.Body.Close()

	var data models.GenderInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Errorf("Error decode json for gender: %s\n", err)
		return
	}

	genderCh <- data.Gender
}

func (p Provider) GetNationByName(nationCh chan string, name string) {
	defer close(nationCh)
	url := buildNationURL(name)
	resp, err := p.client.Get(url)
	if err != nil {
		log.Errorf("Error get request for nation: %s\n", err)
		return
	}
	defer resp.Body.Close()

	var data models.NationInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Errorf("Error decode json for nation: %s\n", err)
		return
	}

	nationCh <- getCountryWithMaxProbability(data.Countries)
}

func getCountryWithMaxProbability(countries []models.CountryInfo) string {
	var maxProbability float64
	var countryCode string
	for _, countryInfo := range countries {
		if countryInfo.Probability > maxProbability {
			maxProbability = countryInfo.Probability
			countryCode = countryInfo.Country
		}
	}
	return countryCode
}

func buildAgeURL(name string) string {
	var url strings.Builder
	url.WriteString("https://api.agify.io/?name=")
	url.WriteString(name)
	return url.String()
}

func buildGenderURL(name string) string {
	var url strings.Builder
	url.WriteString("https://api.genderize.io/?name=")
	url.WriteString(name)
	return url.String()
}

func buildNationURL(name string) string {
	var url strings.Builder
	url.WriteString("https://api.nationalize.io/?name=")
	url.WriteString(name)
	return url.String()
}
