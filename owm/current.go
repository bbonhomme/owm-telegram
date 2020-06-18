package owm

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	bottl "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	// KelZ 0°C in Kelvin to convert in a more human friendly format
	KelZ float64 = 273.15
)

// WeatherIcons maps the icon with the weather description
var WeatherIcons map[string]string

// CurrentWeatherInfo global abstraction spi response for weather at t time in a certain place
type CurrentWeatherInfo struct {
	GeoPos   Coordinates `json:"coord"`
	Sys      Sys         `json:"sys"`
	Base     string      `json:"base"`
	Weather  []Weather   `json:"weather"`
	Main     Main        `json:"main"`
	Wind     Wind        `json:"wind"`
	Clouds   Clouds      `json:"clouds"`
	Rain     Rain        `json:"rain"`
	Snow     Snow        `json:"snow"`
	Dt       int         `json:"dt"`
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Cod      int         `json:"cod"`
	Timezone int         `json:"timezone"`
	Unit     string
	Lang     string
	Key      string
	*Settings
}

// Coordinates Struct for Geo coordinates of the resquest
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Sys Struct for Geo sys informations
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// Wind struct containing speed and angle
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Weather description with imgs and comment
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main struct containing the main informations
type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

// Clouds struct
type Clouds struct {
	All int `json:"all"`
}

// Rain struct contains 3 hour data diff
type Rain struct {
	OneH   float64 `json:"1h,omitempty"`
	ThreeH float64 `json:"3h"`
}

// Snow struct contains 3 hour data diff
type Snow struct {
	OneH   float64 `json:"1h,omitempty"`
	ThreeH float64 `json:"3h"`
}

// CurrentByCoordinates get the weather by geo loc
func (w *CurrentWeatherInfo) CurrentByCoordinates(location *bottl.Location) error {

	urlRequest := fmt.Sprintf(fmt.Sprintf(baseURL, "appid=%s&lat=%f&lon=%f&units=%s&lang=%s"), w.Key, location.Latitude, location.Longitude, defUnit, defLang)
	response, err := w.client.Get(urlRequest)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&w); err != nil {
		return err
	}
	w.TempCel()

	return nil
}

// NewCurrent returns a new CurrentWeatherInfo pointer
func NewCurrent(key string) (*CurrentWeatherInfo, error) {

	curr := CurrentWeatherInfo{
		Settings: NewSettings(),
	}
	var err error
	curr.Key, err = setKey(key)
	if err != nil {
		return nil, err
	}
	return &curr, nil
}

// TempCel convert the received units in Kelvin into Celcius
func (w *CurrentWeatherInfo) TempCel() error {
	w.Main.Temp -= KelZ
	w.Main.TempMax -= KelZ
	w.Main.TempMin -= KelZ
	w.Main.FeelsLike -= KelZ
	return nil
}

// GetCurrent gets the current weather for the provided localisation
func GetCurrent(location *bottl.Location) (*CurrentWeatherInfo, error) {
	w, err := NewCurrent(os.Getenv("OWM_API_KEY"))
	if err != nil {
		return nil, err
	}

	w.CurrentByCoordinates(location)
	return w, nil
}

// BuildAnswer is response for user's template
func (w *CurrentWeatherInfo) BuildAnswer() (weatherMessage string) {
	weatherMessage = `Location: ` + w.Name + `
🕘 Local Time: ` + strconv.Itoa(w.Dt) + `
🌡 Temperature: ` + strconv.FormatFloat(w.Main.Temp, 'f', 2, 32) + `°C
💧 Humidity: ` + strconv.Itoa(w.Main.Humidity) + `%
🌀 Sky: ` + WeatherIcons[w.Weather[0].Icon] + `
💨 Wind Speed: ` + strconv.FormatFloat(w.Wind.Speed, 'f', 2, 32) + ` km/h
🔃 Wind direction: ` + strconv.FormatFloat(w.Wind.Deg, 'f', 2, 32) + `°
🌅 Sunrise: ` + strconv.Itoa(w.Sys.Sunrise) + `
🌄 Sunset: ` + strconv.Itoa(w.Sys.Sunset) + `
`
	return weatherMessage
}

// InitMapWeather Will populate the map containing the icons
func InitMapWeather() (m map[string]string, err error) {

	mapW := make(map[string]string)

	mapW["01d"] = "☀️"
	mapW["01n"] = "☀️"
	mapW["02d"] = "🌤"
	mapW["02n"] = "🌤"
	mapW["03d"] = "🌥"
	mapW["03n"] = "🌥"
	mapW["04d"] = "🌥"
	mapW["04d"] = "🌥"
	mapW["09d"] = "⛈"
	mapW["09d"] = "⛈"
	mapW["10d"] = "🌦"
	mapW["10d"] = "🌦"
	mapW["11d"] = "🌩"
	mapW["11d"] = "🌩"
	mapW["13d"] = "❄️"
	mapW["13d"] = "❄️"
	mapW["50d"] = "🌫"
	mapW["50d"] = "🌫"

	m = mapW
	return m, nil
}