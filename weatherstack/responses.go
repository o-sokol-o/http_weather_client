package weatherstack

import (
	"encoding/json"
	"errors"
)

type ResponseToken struct {
	Token      string `json:"access_token" validate:"required"`
	Expires_in int    `json:"expires_in" validate:"required"`
	Token_type string `json:"token_type" validate:"required"`
}

type ResponseLocation struct {
	ID        int     `json:"id"`        // Identifier to be used with weather queries
	Name      string  `json:"name"`      // Location name eg. London
	Country   string  `json:"country"`   // Location country name eg. United Kingdom
	State     string  `json:"state"`     // U.S. state abbreviation (if applicable)
	AdminArea string  `json:"adminArea"` // Administrative region name (if applicable)
	Timezone  string  `json:"timezone"`  // Location timezone eg. Europe/London
	Lon       float32 `json:"lon"`       // Location longitude
	Lat       float32 `json:"lat"`       // Location latitude
}

type ResponseLocations struct {
	Locations []ResponseLocation `json:"locations"`
}
type ResponseForecast struct {
	Date             string  `json:"date"`             // ISO 8601 date
	Symbol           string  `json:"symbol"`           // Weather symbol code representing daytime conditions (7am to 7pm local time)
	SymbolPhrase     string  `json:"symbolPhrase"`     // Description of weather during the day (dataset=full)
	MaxTemp          int     `json:"maxTemp"`          // 24h maximum temperature (°C)
	MinTemp          int     `json:"minTemp"`          // 24h minimum temperature (°C)
	MaxFeelsLikeTemp float32 `json:"maxFeelsLikeTemp"` // 24h maximum feels-like temperature (°C) (dataset=full)
	MinFeelsLikeTemp float32 `json:"minFeelsLikeTemp"` // 24h minimum feels-like temperature (°C) (dataset=full)
	MaxRelHumidity   float32 `json:"maxRelHumidity"`   // 24h maximum relative humidity (°C) (dataset=full)
	MinRelHumidity   float32 `json:"minRelHumidity"`   // 24h minimum relative humidity (°C) (dataset=full)
	MaxDewPoint      float32 `json:"maxDewPoint"`      // maximum dew point (°C) (dataset=full)
	MinDewPoint      float32 `json:"minDewPoint"`      // minimum dew point (°C) (dataset=full)
	PrecipAccum      float32 `json:"precipAccum"`      // Accumulated precipitation (as liquid water, mm)
	MaxWindSpeed     float32 `json:"maxWindSpeed"`     // Maximum wind speed (m/s or requested unit)
	WindDir          float32 `json:"windDir"`          // Wind direction at maximum wind speed (degrees)
	MaxWindGust      float32 `json:"maxWindGust"`      // Maximum wind gust speed (m/s or requested unit) (dataset=full)
	PrecipProb       float32 `json:"precipProb"`       // Precipitation probability during the day (%) (dataset=full)
	Cloudiness       float32 `json:"cloudiness"`       // Daytime cloudiness average (%) (dataset=full)
	Sunrise          string  `json:"sunrise"`          // Sunrise time of day in local time (dataset=full)
	Sunset           string  `json:"sunset"`           // Sunset time of day in local time (dataset=full)
	SunriseEpoch     float32 `json:"sunriseEpoch"`     // Sunrise time in Unix time (dataset=full)
	SunsetEpoch      float32 `json:"sunsetEpoch"`      // Sunset time in Unix time (dataset=full)
	Moonrise         string  `json:"moonrise"`         // Moonrise time of day in local time (dataset=full)
	Moonset          string  `json:"moonset"`          // Moonset time of day in local time (dataset=full)
	MoonPhase        float32 `json:"moonPhase"`        // Phase of moon at midday (degrees) (dataset=full)
	UvIndex          float32 `json:"uvIndex"`          // Maximum ultraviolet index (dataset=full)
	MinVisibility    float32 `json:"minVisibility"`    // Minimum visibility (m) (dataset=full)
	Pressure         float32 `json:"pressure"`         // Daytime average sea level pressure (hPa) (dataset=full)
	SnowAccum        float32 `json:"snowAccum"`        // Accumulated snow (thickness of precipiated snow, cm)
	Confidence       string  `json:"confidence"`       // Confidence in forecast: g (green, good), y (yellow, normal), o (orange, low)
}

type ResponseForecastDaily struct {
	ForecastDaily []ResponseForecast `json:"forecast"`
}

type WetherError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}

type ResponseError struct {
	Success bool        `json:"success"`
	Error   WetherError `json:"error"`
}

func respGetToken(body []byte) (string, error) {
	var rt ResponseToken
	if err := json.Unmarshal(body, &rt); err != nil || rt.Token == "" || rt.Token_type != "bearer" {

		return "", errors.New("api decode error")
	}
	return rt.Token, nil
}

func respGetForecastDaily(body []byte) (*ResponseForecastDaily, error) {
	var rfd ResponseForecastDaily
	if err := json.Unmarshal(body, &rfd); err != nil {
		return nil, errors.New("forecast daily api decode error:  " + string(body))
	}
	return &rfd, nil
}
func respGetLocations(body []byte) (*ResponseLocations, error) {
	var rl ResponseLocations

	if err := json.Unmarshal(body, &rl); err != nil {
		return nil, errors.New("locations api decode error:  " + string(body))
	}
	return &rl, nil
}
