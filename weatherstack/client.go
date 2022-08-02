package weatherstack

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	key    string
	client *http.Client
}

func NewClient(key string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		key: key,
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetWeather(cytiName string) (string, error) {

	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("access_key", c.key)
	q.Add("query", cytiName)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return processingResponse(body, cytiName)
}

// Weather Forecast API Endpoint
func (c Client) GetWeatherForecast(cytiName string, days int) (string, error) {

	if days > 14 {
		return "", errors.New("maximum 14 days")
	}

	if days <= 0 {
		days = 7
	}

	req, err := http.NewRequest("GET", "http://api.weatherstack.com/forecast", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("access_key", c.key)
	q.Add("query", cytiName)
	q.Add("forecast_days", strconv.Itoa(days))

	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))

	return processingResponse(body, cytiName)
}
