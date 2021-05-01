package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var apiKey = "92721f2c7ecab4f083189daef6b7f146"

func CurrentWeather(rw http.ResponseWriter, request *http.Request /*, latitude string, longitude string*/) {
	rw.Header().Set("Content-type", "application/json")

	address := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers

	latitude, longitude, err := GetLocation(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	url := ""

	if latitude != "" && longitude != "" {
		url = "https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + apiKey
	} else {
		fmt.Fprint(rw, "Check formatting of lat and lon")
	}
	currentWeatherHandler(rw, url)
}

func currentWeatherHandler(rw http.ResponseWriter, url string) {
	// Uses request URL
	resp, err := http.Get(url)
	if err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Reads the data from the resp.Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var weather weatherData

	// Unmarshalling the body into the weatherData struct/fields
	if err := json.Unmarshal(body, &weather); err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var data []outputWeather

	main := weather.Weather[0].Main
	rain1H := weather.Rain.OneH
	snow1H := weather.Snow.OneH
	tempActual := weather.Main.Temp
	tempFeelsLike := weather.Main.FeelsLike
	tempMin := weather.Main.TempMin
	tempMax := weather.Main.TempMax
	humidity := weather.Main.Humidity
	visibility := weather.Visibility
	windSpeed := weather.Wind.Speed
	windDeg := weather.Wind.Deg
	sunrise := weather.Sys.Sunrise
	sunset := weather.Sys.Sunset

	jsonStruct := outputWeather{Main: main, Rain1h: rain1H, Snow1h: snow1H, Temp: tempActual,
		FeelsLike: tempFeelsLike, TempMin: tempMin, TempMax: tempMax, Humidity: humidity, Visibility: visibility,
		WindSpeed: windSpeed, WindDeg: windDeg, Sunrise: sunrise, Sunset: sunset}

	data = append(data, jsonStruct)

	response(rw, data)
	/*// Displays the data to the user
	mes, err := fmt.Fprintf(rw, `{
			"The main weather category at the given location is": "%v",
			"The actual temperature is": "%v",
			"You can expect the temperature to feel like": "%v",
			"The minimum temperature to be expected is": "%v",
			"The maximum temperature to be expected is": "%v",
			"The humidity is": "%v"
		}`, jsonStruct.Main, jsonContinent, scope, jsonConfirmed, jsonRecovered, math.Round(jsonPopulationPercentage*100)/100)
	if err == nil {
		fmt.Print(mes)
	}*/

	output, err := json.Marshal(data) //Marshalling the array to JSON
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "%v", string(output)) //Outputs the weather
}

func response(rw http.ResponseWriter, data []outputWeather) {

	mainMessage := ""
	tempMessage := ""
	feelsLikeMessage := ""
	tempMinMessage := ""
	tempMaxMessage := ""
	humidityMessage := ""
	visibilityMessage := ""
	windSpeedMessage := ""
	windDegMessage := ""
	sunriseMessage := ""
	sunsetMessage := ""

	dataTemp := data[0]
	// Converts the temperatures from Kelvin to Celsius
	// Rounds it of to two decimals
	dataTemp.Temp = math.Round((dataTemp.Temp-273.15)*100) / 100
	dataTemp.FeelsLike = math.Round((dataTemp.FeelsLike-273.15)*100) / 100
	dataTemp.TempMin = math.Round((dataTemp.TempMin-273.15)*100) / 100
	dataTemp.TempMax = math.Round((dataTemp.TempMax-273.15)*100) / 100

	switch dataTemp.Main {
	case "Rain":
		switch true {
		case dataTemp.Rain1h <= 2:
			mainMessage = "It has been light rain the last hour, consider bringing a umbrella."
			break
		case dataTemp.Rain1h > 2 && dataTemp.Rain1h <= 7:
			mainMessage = "It has been moderate rain the last hour, consider bringing rainwear."
			break
		case dataTemp.Rain1h > 7 && dataTemp.Rain1h <= 50:
			mainMessage = "It has been heavy rain the last hour, you should bring rainwear."
			break
		case dataTemp.Rain1h > 50:
			mainMessage = "It has been violent rain the last hour, prepare your anus."
			break
		default:
			mainMessage = "It is raining, bring appropriate clothing."
			break
		}
	case "Snow":
		switch true {
		case dataTemp.Snow1h <= 1:
			mainMessage = "It has been snowing lightly the last hour, consider winter tires, " +
				"wear appropriate clothing and set of a couple of minutes to clear the snow of your car."
			break
		case dataTemp.Snow1h > 1 && dataTemp.Snow1h <= 3:
			mainMessage = "It has been snowing moderately the last hour, make sure to have winter tires, " +
				"wear appropriate clothing and expect at least 10 minutes to clear the snow of your car."
			break
		case dataTemp.Snow1h > 3:
			mainMessage = "It has been snowing heavily the last hour, you must have winter tires, " +
				"wear appropriate clothing and expect at least 20 minutes to clear the snow around- and of your car."
			break
		default:
			mainMessage = "It is snowing, drive carefully, bring appropriate clothing and turn on the heater."
			break
		}
	case "Clear":
		mainMessage = "The sky is clear, wear appropriate clothing with respect to terrain and temperature."
	default:
		mainMessage = "The weather of your destination is: " + dataTemp.Main
	}

	s := fmt.Sprintf("%.1f", dataTemp.Temp)
	switch true {
	case dataTemp.Temp <= 0:
		tempMessage = "The temperature is: " + s + " degrees Celsius. It is below the freezing point outside, drive carefully " +
			"and we recommend you to use winter tires. Bring warm clothes and something warm to drink."
		break
	case dataTemp.Temp > 0 && dataTemp.Temp <= 20:
		tempMessage = "The temperature is: " + s + " degrees Celsius. The temperature outside is moderate. Consider whether " +
			"winter tires is needed. Most likely summer tires are recommended. Wear a moderate amount of clothing."
		break
	case dataTemp.Temp > 20:
		tempMessage = "The temperature is: " + s + " degrees Celsius. The temperature outside is high, and summer tires are highly " +
			"recommended! Put on some sunscreen, some light clothing, blast your air condition at max and put a slush in the " +
			"cup holder. Enjoy the temperature!"
		break
	default:
		tempMessage = "The temperature is: " + s + " degrees Celsius. Consider whether winter tires is needed and wear appropriate clothing."
		break
	}
	t := fmt.Sprintf("%.1f", dataTemp.FeelsLike)
	switch true {
	case dataTemp.FeelsLike <= 10:
		feelsLikeMessage = "It feels cold outside, with a feels like temperature of: " + t + " degrees Celsius."
		break
	case dataTemp.FeelsLike > 10 && dataTemp.FeelsLike <= 20:
		feelsLikeMessage = "It is a great temperature outside for jeans and jumper, with a feels like temperature of: " + t + " degrees Celsius."
		break
	case dataTemp.FeelsLike > 20:
		feelsLikeMessage = "It is warm outside today, light clothes are strongly recommended. Take a trip to the beach " +
			"and enjoy the day. The temperature feels like: " + t + " degrees Celsius."
		break
	default:
		feelsLikeMessage = "The temperature outside feels like: " + t + " degrees Celsius. Dress accordingly."
		break
	}

	u := fmt.Sprintf("%.1f", dataTemp.TempMin)
	tempMinMessage = "The minimum temperature for the day is expected to be: " + u + " degrees Celsius. Consider this both with regards " +
		"to tires used and clothes you plan to wear."

	v := fmt.Sprintf("%.1f", dataTemp.TempMax)
	tempMaxMessage = "The maximum temperature for the day is expected to be: " + v + " degrees Celsius. Consider this both with regards " +
		"to tires used and clothes you plan to wear."

	switch true {
	case dataTemp.Humidity <= 30:
		humidityMessage = "The humidity value at the moment is:" + strconv.Itoa(dataTemp.Humidity) + " percent. " +
			"This is relatively low humidity which can result in health issues."
		break
	case dataTemp.Humidity > 30 && dataTemp.Humidity <= 75:
		humidityMessage = "The humidity values at the moment is: " + strconv.Itoa(dataTemp.Humidity) + " percent. " +
			"This is an ideal humidity value and you should be comfortable going outside."
		break
	case dataTemp.Humidity > 75:
		humidityMessage = "The humidity values at the moment is: " + strconv.Itoa(dataTemp.Humidity) + " percent. " +
			"If it ain't raining already, there is a high change it will start raining soon."
		break
	default:
		humidityMessage = "The humidity values at the moment is: " + strconv.Itoa(dataTemp.Humidity) + " percent."
		break
	}

	visibilityMessage = "The visibility outside is: " + strconv.Itoa(dataTemp.Visibility) + " meters."

	x := fmt.Sprintf("%.2f", dataTemp.WindSpeed)
	windSpeedMessage = "The wind speed today is " + x + " m/s."

	windDegMessage = "The wind degree direction is: " + strconv.Itoa(dataTemp.WindDeg)

	// Convert from epoch to human readable date
	// Inspired by:
	//	- https://play.golang.org/p/6h0A0WPxtq
	//	- https://www.epochconverter.com
	sunriseEpoch := dataTemp.Sunrise
	sunsetEpoch := dataTemp.Sunset

	sunriseMessage = "The time for sunrise is: " + epochToHumanReadable(int64(sunriseEpoch)).String()

	sunsetMessage = "The time for sunset is: " + epochToHumanReadable(int64(sunsetEpoch)).String()

	// Displays the data to the user
	mes, err := fmt.Fprintf(rw,
		mainMessage+"\n"+
			tempMessage+"\n"+
			feelsLikeMessage+"\n"+
			tempMinMessage+"\n"+
			tempMaxMessage+"\n"+
			humidityMessage+"\n"+
			visibilityMessage+"\n"+
			windSpeedMessage+"\n"+
			windDegMessage+"\n"+
			sunriseMessage+"\n"+
			sunsetMessage,
	)
	if err == nil {
		fmt.Print(mes)
	}
}

/** Inspired by and taken from:
 *	- https://play.golang.org/p/6h0A0WPxtq
 * 	- https://www.epochconverter.com
 *  We found this useful to use because we got the data from the API in Epoch and Unit format and this is not readable
 *  for the user, when the format is for instance: "1619838624". After conversion we get the following format:
 *  "2021-05-01 05: 10: 24 +0200 CEST"
 */
func epochToHumanReadable(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}
