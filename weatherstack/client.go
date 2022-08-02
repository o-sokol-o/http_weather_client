package weatherstack

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// User: toxin-air
// Password: rOTlZqgDseCk

// The domains for your account are:
// Weather API: https://pfa.foreca.com
// Weather Map API: https://map-eu.foreca.com

// Example queries:
// curl -d '{"user": "toxin-air", "password": "rOTlZqgDseCk"}' 'https://pfa.foreca.com/authorize/token?expire_hours=2'
// curl -H 'Authorization: Bearer <token>' 'https://pfa.foreca.com/api/v1/location/search/Barcelona?lang=es'
// curl -H 'Authorization: Bearer <token>' 'https://pfa.foreca.com/api/v1/forecast/daily/103128760'
// curl -H 'Authorization: Bearer <token>' 'https://map-eu.foreca.com/api/v1/capabilities'

type Client struct {
	token  string
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) Login(user, password string) error {

	req, err := http.NewRequest("POST", "https://pfa.foreca.com/authorize/token?expire_hours=2", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("user", user)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.token, err = respGetToken(body)
	return err
}

func (c *Client) Logout() error {

	url := fmt.Sprintf("https://pfa.foreca.com/authorize/key/%s", c.token)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.token = ""
	return nil
}

func (c *Client) GetLocations(cityName string) (*ResponseLocations, error) {

	url := fmt.Sprintf("https://pfa.foreca.com/api/v1/location/search/%s", cityName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lang", "en")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respGetLocations(body)
}

func (c *Client) GetWeather(cityID int) (*ResponseForecastDaily, error) {

	url := fmt.Sprintf("https://pfa.foreca.com/api/v1/forecast/daily/%d", cityID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lang", "en")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respGetForecastDaily(body)
}
