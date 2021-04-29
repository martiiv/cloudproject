package extra

import "time"

type geoLocation struct {
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

type OutputCharge struct {
	Charger    string
	Address    string
	Phone      string
	Connectors string
	PowerKW    float64
}

type OutputPetrol struct {
	StationName  string
	StationBrand string
	Address      string
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
