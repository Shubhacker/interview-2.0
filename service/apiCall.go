package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName map[string]struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nativeName"`
	} `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Capital    []string `json:"capital"`
	Population int64    `json:"population"`
}

// Using fmt for logging errors, later can add logger
func ExternalAPI(name string) ([]Data, error) {
	url := "https://restcountries.com/v3.1/name/" + name + "?fields=name,capital,currencies,population"

	resp, err := http.Get(url)
	if err != nil {
		// error handling
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// error handling
		fmt.Println("Something went wrong...")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Error handling
		fmt.Println(err.Error())
		return nil, err
	}

	var c []Country
	err = json.Unmarshal(body, &c)
	if err != nil {
		// Error handling
		fmt.Println("unmarshaling error :>", err.Error())
		return nil, err
	}
	return formatData(c), nil
}

func formatData(c []Country) []Data {
	var response []Data
	for _, dt := range c {
		var singleCountry Data
		singleCountry.Name = dt.Name.Common
		singleCountry.Capital = dt.Capital[0]
		singleCountry.Population = int(dt.Population)

		for _, cur := range dt.Currencies {
			// for now will only get single currency, In future can add feature to return all currency
			singleCountry.Currency = cur.Name
		}
		response = append(response, singleCountry)
	}
	return response
}
