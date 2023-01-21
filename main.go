package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	API_KEY = os.Getenv("OPENWEATHER_API")
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          string
	Main        string
	Description string
	Icon        string
}

type WeahterInfo struct {
	Coord Coord `json:"coord"`
	Weather
}

func main() {
	if API_KEY == "" {
		log.Fatal("API_KEY is empty")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=34.3999649&lon=132.7135802&units=metric&APPID=%s", API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request Error, %v", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	postResp, err := http.Post("http://localhost:9880/weather.test", "application/json", bytes.NewBuffer(byteArray))
	if err != nil {
		log.Fatalf("Post Request Error %v", err)
	}
	defer postResp.Body.Close()
	byteArray, _ = ioutil.ReadAll(postResp.Body)
	fmt.Println(byteArray)

	// //
	// var weatherInfo WeahterInfo
	// if err := json.Unmarshal(byteArray, &weatherInfo); err != nil {
	// 	log.Fatal(err)
	// }

}
