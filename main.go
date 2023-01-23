package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	API_KEY           = os.Getenv("OPENWEATHER_API")
	ENDPOINT          = os.Getenv("MINIO_ENDPOINT")
	ACCESS_KEY_ID     = os.Getenv("MINIO_ACCESS_KEY_ID")
	SECRET_ACCESS_KEY = os.Getenv("MINIO_SECRET_ACCESS_KEY")
)

func getWeather(url string) (wi *WeahterInfo, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	wi = &WeahterInfo{}
	if err := json.Unmarshal(byteArray, wi); err != nil {
		return nil, err
	}
	return wi, nil
}

func postWeather(wi *WeahterInfo) (byteArray []byte, err error) {
	byteArray, _ = json.Marshal(wi)
	postResp, err := http.Post("http://localhost:9880/weather.hiroshima", "application/json", bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	defer postResp.Body.Close()
	return byteArray, nil
}

func main() {
	if API_KEY == "" {
		log.Fatal("API_KEY is empty")
	}

	// open weather APIのデータの取得
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=34.3999649&lon=132.7135802&units=metric&APPID=%s", API_KEY)
	wi, err := getWeather(url)
	if err != nil {
		log.Fatalf("Get Request Error\n%v", err)
	}

	// データの検証を行う
	v := validator.New()
	if err := v.Struct(wi); err != nil {
		log.Fatalf("Invalid Struct\n%v\n", err)
	}

	// fluentdにPOSTする
	byteArray, err := postWeather(wi)
	if err != nil {
		log.Fatalf("Post Request Error %v", err)
	}
	fmt.Println(string(byteArray))
}
