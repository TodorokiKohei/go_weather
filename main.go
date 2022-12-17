package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	API_KEY = os.Getenv("API_KEY")
)

func main() {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=34.3999649&lon=132.7135802&units=metric&APPID=%s", API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request Error, %v", err)
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
