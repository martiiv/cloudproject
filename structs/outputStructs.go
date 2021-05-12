package structs

import "time"

type OutputCharge struct {
	Charger    string
	Address    string
	Phone      string
	Connectors []Connectors
}

type Connectors struct {
	ConnectorType string  `json:"connectorType"`
	RatedPowerKW  float64 `json:"ratedPowerKW"`
}

type OutputPetrol struct {
	StationName  string
	StationBrand string
	Address      string
}

type OutIncident struct {
	Start time.Time
	End   time.Time
	From  string
	To    string
	Event string
}

type OutputWeather struct {
	Main       MainStruct
	Rain1h     float64
	Snow1h     float64
	Temp       TempStruct
	FeelsLike  FeelsLikeStruct
	TempMin    TempMinStruct
	TempMax    TempMaxStruct
	Humidity   HumidityStruct
	Visibility VisibilityStruct
	WindSpeed  WindSpeedStruct
	WindDeg    WindDegStruct
	Sunrise    SunriseStruct
	Sunset     SunsetStruct
}

type MainStruct struct {
	Main    string
	Message string
}

type TempStruct struct {
	Temp    float64
	Message string
}

type FeelsLikeStruct struct {
	FeelsLike float64
	Message   string
}

type TempMinStruct struct {
	TempMin float64
	Message string
}

type TempMaxStruct struct {
	TempMax float64
	Message string
}

type HumidityStruct struct {
	Humidity int
	Message  string
}

type VisibilityStruct struct {
	Visibility int
	Message    string
}

type WindSpeedStruct struct {
	WindSpeed float64
	Message   string
}

type WindDegStruct struct {
	WindDeg int
	Message string
}

type SunriseStruct struct {
	Sunrise int
	Message string
}

type SunsetStruct struct {
	Sunset  int
	Message string
}

type Route struct {
	Street       string
	Maneuver     string
	RoadNumber   string
	JunctionType string
}

type RoadInformation struct {
	EstimatedArrival string
	LengthKM         int
	Route            []Route
}

type Webhook struct {
	Url                 string
	DepartureLocation   string
	ArrivalDestination  string
	Weather             string
	ArrivalTime         string
	Repeat              string
	EstimatedTravelTime int
	Id                  string
}

type NotificationInput struct {
	URL string
}

type NotificationResponse struct {
	Text string
}

type OutputPoi struct {
	Name        string
	PhoneNumber string
	Address     string
}

type LocationLonLat struct {
	Longitude float64
	Latitude  float64
}
