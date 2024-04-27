package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
}

type Weather struct {
	Temperature int `json:"temperature"`
}

type Response struct {
	Location Location `json:"location"`
	Current  Weather  `json:"current"`
}

func GetTemperature(city string) string {
	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("access_key", "API_KEY")
	q.Add("query", city)
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var apiResponse Response
	json.NewDecoder(res.Body).Decode(&apiResponse)

	return fmt.Sprintf("Current temperature in %s is %d℃", apiResponse.Location.Name, apiResponse.Current.Temperature)
}
