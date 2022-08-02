package weatherstack

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Location struct {
	Name string `json:"name"`
}

type Weather struct {
	Temperature int `json:"temperature"`
}

type Response struct {
	Location Location `json:"location"`
	Current  Weather  `json:"current"`
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

func processingResponse(body []byte, cytiName string) (string, error) {
	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil || apiResponse.Location.Name != cytiName {

		var apiResponseErr ResponseError
		if err = json.Unmarshal(body, &apiResponseErr); err != nil || apiResponseErr.Error.Type == "" {
			return "", errors.New("api decode error")
		}

		return "", errors.New(apiResponseErr.Error.Info)
	}
	return fmt.Sprintf("Current temperature in %s is %d℃", apiResponse.Location.Name, apiResponse.Current.Temperature), nil

}
