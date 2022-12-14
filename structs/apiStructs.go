package structs

import (
	"time"
)

type GeoLocation struct {
	Results []struct {
		Locations []struct {
			LatLng struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"latLng"`
		} `json:"locations"`
	} `json:"results"`
}

type Charger struct {
	Summary struct {
		QueryType    string `json:"queryType"`
		QueryTime    int    `json:"queryTime"`
		NumResults   int    `json:"numResults"`
		Offset       int    `json:"offset"`
		TotalResults int    `json:"totalResults"`
		FuzzyLevel   int    `json:"fuzzyLevel"`
		GeoBias      struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"geoBias"`
	} `json:"summary"`
	Results []struct {
		Type  string  `json:"type"`
		ID    string  `json:"id"`
		Score float64 `json:"score"`
		Dist  float64 `json:"dist"`
		Info  string  `json:"info"`
		Poi   struct {
			Name        string `json:"name"`
			Phone       string `json:"phone"`
			CategorySet []struct {
				ID int `json:"id"`
			} `json:"categorySet"`
			Categories      []string `json:"categories"`
			Classifications []struct {
				Code  string `json:"code"`
				Names []struct {
					NameLocale string `json:"nameLocale"`
					Name       string `json:"name"`
				} `json:"names"`
			} `json:"classifications"`
		} `json:"poi,omitempty"`
		Address struct {
			StreetNumber       string `json:"streetNumber"`
			StreetName         string `json:"streetName"`
			Municipality       string `json:"municipality"`
			CountrySubdivision string `json:"countrySubdivision"`
			PostalCode         string `json:"postalCode"`
			CountryCode        string `json:"countryCode"`
			Country            string `json:"country"`
			CountryCodeISO3    string `json:"countryCodeISO3"`
			FreeformAddress    string `json:"freeformAddress"`
			LocalName          string `json:"localName"`
		} `json:"address,omitempty"`
		Position struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"position"`
		Viewport struct {
			TopLeftPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"topLeftPoint"`
			BtmRightPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"btmRightPoint"`
		} `json:"viewport"`
		EntryPoints []struct {
			Type     string `json:"type"`
			Position struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"position"`
		} `json:"entryPoints"`
		ChargingPark struct {
			Connectors []struct {
				ConnectorType string  `json:"connectorType"`
				RatedPowerKW  float64 `json:"ratedPowerKW"`
				VoltageV      int     `json:"voltageV"`
				CurrentA      int     `json:"currentA"`
				CurrentType   string  `json:"currentType"`
			} `json:"connectors"`
		} `json:"chargingPark,omitempty"`
		DataSources struct {
			ChargingAvailability struct {
				ID string `json:"id"`
			} `json:"chargingAvailability"`
		} `json:"dataSources,omitempty"`
	} `json:"results"`
}

type Petrol struct {
	Summary struct {
		Query        string `json:"query"`
		QueryType    string `json:"queryType"`
		QueryTime    int    `json:"queryTime"`
		NumResults   int    `json:"numResults"`
		Offset       int    `json:"offset"`
		TotalResults int    `json:"totalResults"`
		FuzzyLevel   int    `json:"fuzzyLevel"`
		GeoBias      struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"geoBias"`
	} `json:"summary"`
	Results []struct {
		Type  string  `json:"type"`
		ID    string  `json:"id"`
		Score float64 `json:"score"`
		Dist  float64 `json:"dist"`
		Info  string  `json:"info"`
		Poi   struct {
			Name   string `json:"name"`
			Brands []struct {
				Name string `json:"name"`
			} `json:"brands"`
			CategorySet []struct {
				ID int `json:"id"`
			} `json:"categorySet"`
			URL             string   `json:"url"`
			Categories      []string `json:"categories"`
			Classifications []struct {
				Code  string `json:"code"`
				Names []struct {
					NameLocale string `json:"nameLocale"`
					Name       string `json:"name"`
				} `json:"names"`
			} `json:"classifications"`
		} `json:"poi,omitempty"`
		Address struct {
			StreetNumber       string `json:"streetNumber"`
			StreetName         string `json:"streetName"`
			Municipality       string `json:"municipality"`
			CountrySubdivision string `json:"countrySubdivision"`
			PostalCode         string `json:"postalCode"`
			CountryCode        string `json:"countryCode"`
			Country            string `json:"country"`
			CountryCodeISO3    string `json:"countryCodeISO3"`
			FreeformAddress    string `json:"freeformAddress"`
			LocalName          string `json:"localName"`
		} `json:"address,omitempty"`
		Position struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"position"`
		Viewport struct {
			TopLeftPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"topLeftPoint"`
			BtmRightPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"btmRightPoint"`
		} `json:"viewport"`
		EntryPoints []struct {
			Type     string `json:"type"`
			Position struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"position"`
		} `json:"entryPoints"`
	} `json:"results"`
}

type AutoGenerated struct {
	Type     string `json:"type"`
	Features []struct {
		Bbox       []float64 `json:"bbox"`
		Type       string    `json:"type"`
		Properties struct {
			Segments []struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
				Steps    []struct {
					Distance    float64 `json:"distance"`
					Duration    float64 `json:"duration"`
					Type        int     `json:"type"`
					Instruction string  `json:"instruction"`
					Name        string  `json:"name"`
					WayPoints   []int   `json:"way_points"`
					ExitNumber  int     `json:"exit_number,omitempty"`
				} `json:"steps"`
			} `json:"segments"`
			Summary struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
			} `json:"summary"`
			WayPoints []int `json:"way_points"`
		} `json:"properties"`
		Geometry struct {
			Coordinates [][]float64 `json:"coordinates"`
			Type        string      `json:"type"`
		} `json:"geometry"`
	} `json:"features"`
	Bbox     []float64 `json:"bbox"`
	Metadata struct {
		Attribution string `json:"attribution"`
		Service     string `json:"service"`
		Timestamp   int64  `json:"timestamp"`
		Query       struct {
			Coordinates [][]float64 `json:"coordinates"`
			Profile     string      `json:"profile"`
			Format      string      `json:"format"`
		} `json:"query"`
		Engine struct {
			Version   string    `json:"version"`
			BuildDate time.Time `json:"build_date"`
			GraphDate time.Time `json:"graph_date"`
		} `json:"engine"`
	} `json:"metadata"`
}

type Incidents struct {
	Incidents []struct {
		Type       string `json:"type"`
		Properties struct {
			ID               string        `json:"id"`
			IconCategory     int           `json:"iconCategory"`
			MagnitudeOfDelay int           `json:"magnitudeOfDelay"`
			StartTime        time.Time     `json:"startTime"`
			EndTime          time.Time     `json:"endTime"`
			From             string        `json:"from"`
			To               string        `json:"to"`
			Length           float64       `json:"length"`
			Delay            int           `json:"delay"`
			RoadNumbers      []interface{} `json:"roadNumbers"`
			Events           []struct {
				Code        int    `json:"code"`
				Description string `json:"description"`
			} `json:"events"`
			Aci interface{} `json:"aci"`
		} `json:"properties"`
		Geometry struct {
			Type        string      `json:"type"`
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"incidents"`
}

// WeatherData Used to store weather related data from the API, in the format the API returns the data
type WeatherData struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Snow struct {
		OneH   float64 `json:"1h"`
		ThreeH float64 `json:"3h"`
	} `json:"snow"`
	Rain struct {
		OneH   float64 `json:"1h"`
		ThreeH float64 `json:"3h"`
	} `json:"rain"`
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`
}

type AutoGenerated3 struct {
	FormatVersion string `json:"formatVersion"`
	Routes        []struct {
		Summary struct {
			LengthInMeters        int       `json:"lengthInMeters"`
			TravelTimeInSeconds   int       `json:"travelTimeInSeconds"`
			TrafficDelayInSeconds int       `json:"trafficDelayInSeconds"`
			TrafficLengthInMeters int       `json:"trafficLengthInMeters"`
			DepartureTime         time.Time `json:"departureTime"`
			ArrivalTime           time.Time `json:"arrivalTime"`
		} `json:"summary"`
		Legs []struct {
			Summary struct {
				LengthInMeters        int       `json:"lengthInMeters"`
				TravelTimeInSeconds   int       `json:"travelTimeInSeconds"`
				TrafficDelayInSeconds int       `json:"trafficDelayInSeconds"`
				TrafficLengthInMeters int       `json:"trafficLengthInMeters"`
				DepartureTime         time.Time `json:"departureTime"`
				ArrivalTime           time.Time `json:"arrivalTime"`
			} `json:"summary"`
			Points []struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"points"`
		} `json:"legs"`
		Sections []struct {
			StartPointIndex int    `json:"startPointIndex"`
			EndPointIndex   int    `json:"endPointIndex"`
			SectionType     string `json:"sectionType"`
			TravelMode      string `json:"travelMode"`
		} `json:"sections"`
		Guidance struct {
			Instructions []struct {
				RouteOffsetInMeters int `json:"routeOffsetInMeters"`
				TravelTimeInSeconds int `json:"travelTimeInSeconds"`
				Point               struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"point"`
				PointIndex                int      `json:"pointIndex"`
				InstructionType           string   `json:"instructionType"`
				Street                    string   `json:"street,omitempty"`
				CountryCode               string   `json:"countryCode"`
				PossibleCombineWithNext   bool     `json:"possibleCombineWithNext"`
				DrivingSide               string   `json:"drivingSide"`
				Maneuver                  string   `json:"maneuver"`
				JunctionType              string   `json:"junctionType,omitempty"`
				TurnAngleInDecimalDegrees int      `json:"turnAngleInDecimalDegrees,omitempty"`
				RoadNumbers               []string `json:"roadNumbers,omitempty"`
				SignpostText              string   `json:"signpostText,omitempty"`
				RoundaboutExitNumber      int      `json:"roundaboutExitNumber,omitempty"`
			} `json:"instructions"`
			InstructionGroups []struct {
				FirstInstructionIndex int `json:"firstInstructionIndex"`
				LastInstructionIndex  int `json:"lastInstructionIndex"`
				GroupLengthInMeters   int `json:"groupLengthInMeters"`
			} `json:"instructionGroups"`
		} `json:"guidance"`
	} `json:"routes"`
}

type PointsOfInterest struct {
	Summary struct {
		Query        string `json:"query"`
		Querytype    string `json:"queryType"`
		Querytime    int    `json:"queryTime"`
		Numresults   int    `json:"numResults"`
		Offset       int    `json:"offset"`
		Totalresults int    `json:"totalResults"`
		Fuzzylevel   int    `json:"fuzzyLevel"`
		Geobias      struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"geoBias"`
	} `json:"summary"`
	Results []struct {
		Type  string  `json:"type"`
		ID    string  `json:"id"`
		Score float64 `json:"score"`
		Dist  float64 `json:"dist"`
		Info  string  `json:"info"`
		Poi   struct {
			Name        string `json:"name"`
			Phone       string `json:"phone"`
			Categoryset []struct {
				ID int `json:"id"`
			} `json:"categorySet"`
			Categories      []string `json:"categories"`
			Classifications []struct {
				Code  string `json:"code"`
				Names []struct {
					Namelocale string `json:"nameLocale"`
					Name       string `json:"name"`
				} `json:"names"`
			} `json:"classifications"`
		} `json:"poi,omitempty"`
		Address struct {
			Streetnumber                string `json:"streetNumber"`
			Streetname                  string `json:"streetName"`
			Municipalitysubdivision     string `json:"municipalitySubdivision"`
			Municipality                string `json:"municipality"`
			Countrysecondarysubdivision string `json:"countrySecondarySubdivision"`
			Countrysubdivision          string `json:"countrySubdivision"`
			Countrysubdivisionname      string `json:"countrySubdivisionName"`
			Postalcode                  string `json:"postalCode"`
			Extendedpostalcode          string `json:"extendedPostalCode"`
			Countrycode                 string `json:"countryCode"`
			Country                     string `json:"country"`
			Countrycodeiso3             string `json:"countryCodeISO3"`
			Freeformaddress             string `json:"freeformAddress"`
			Localname                   string `json:"localName"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"position"`
		Viewport struct {
			Topleftpoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"topLeftPoint"`
			Btmrightpoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"btmRightPoint"`
		} `json:"viewport"`
		Entrypoints []struct {
			Type     string `json:"type"`
			Position struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"position"`
		} `json:"entryPoints"`
	} `json:"results"`
}
