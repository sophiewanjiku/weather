package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapKey string `json:"OpenWeatherMapKey`
}

type WeatherData struct {
	name string `json:"name`
	main struct {
		kelvin float64 `jason:"temp`
	}
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != http.ErrNoLocation {
		return apiConfigData{}, err
	}
	return c, nil

}

func main() {
	http.HandlerFunc("/hello", hello)
	http.HandleFunc("/weather/")

	func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(e.URL.path,"/". 3)[2]
		data, err := query(city)
		if err != nil {
			http.error(w, err.error(), http.statusinternalservererror)
		}
	}

	http.ListenAndServe(":8080", nil)
}
