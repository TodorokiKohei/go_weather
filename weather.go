package main

type Coord struct {
	Lon float64 `json:"lon" validate:"required"`
	Lat float64 `json:"lat" validate:"required"`
}

type Weather struct {
	Id          int    `json:"id" validate:"required"`
	Main        string `json:"main" validate:"required"`
	Description string `json:"description" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type Main struct {
	Temp      float64 `json:"temp" validate:"required"`
	FeelsLike float64 `json:"feels_like" validate:"required"`
	TempMin   float64 `json:"temp_min" validate:"required"`
	TempMax   float64 `json:"temp_max" validate:"required"`
	Pressure  float64 `json:"pressure" validate:"required"`
	Humidity  float64 `json:"humidity" validate:"required"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed" validate:"required"`
	Deg   int     `json:"deg" validate:"required"`
	Gust  float64 `json:"gust" validate:"required"`
}

type Rain struct {
	H1 float64 `json:"1h"`
}

type Snow struct {
	H1 float64 `json:"1h"`
}

type Clouds struct {
	All int `json:"all" validate:"required"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Message string `json:"message"`
	Country string `json:"country" validate:"required"`
	Sunrise int    `json:"sunrise" validate:"required"`
	Sunset  int    `json:"sunset" validate:"required"`
}

type WeahterInfo struct {
	Coord      Coord     `json:"coord" validate:"required"`
	Weather    []Weather `json:"weather" validate:"required"`
	Base       string    `json:"base" validate:"required"`
	Main       Main      `json:"main" validate:"required"`
	Visibility int       `json:"visibility" validate:"required"`
	Rain       Rain      `json:"rain"`
	Snow       Snow      `json:"snow"`
	Clouds     Clouds    `json:"clouds" validate:"required"`
	Dt         int       `json:"dt" validate:"required"`
	Sys        Sys       `json:"sys" validate:"required"`
	Timezon    int       `json:"timezone" validate:"required"`
	Id         int       `json:"id" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Cod        int       `json:"cod" validate:"required"`
}
