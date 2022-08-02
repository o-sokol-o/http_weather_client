package main

import (
	"fmt"
	"log"
	"time"
	"weatherstack/weatherstack"
)

// После регистрации получен ключ доступа к АПИ e0aeb5680a2db97d039d4370f8979728

func main() {
	weatherClient, err := weatherstack.NewClient("e0aeb5680a2db97d039d4370f8979728", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}

	if str, err := weatherClient.GetWeather("Lon_don"); err != nil {
		log.Println(err)
	} else {
		fmt.Println(str)
	}

	if str, err := weatherClient.GetWeather("London"); err != nil {
		log.Println(err)
	} else {
		fmt.Println(str)
	}

	if str, err := weatherClient.GetWeather("Kiev"); err != nil {
		log.Println(err)
	} else {
		fmt.Println(str)
	}

	if str, err := weatherClient.GetWeatherForecast("Kiev", 3); err != nil {
		log.Println(err)
	} else {
		fmt.Println(str)
	}
}
