package main

import (
	"fmt"
	"log"
	"time"
	"weatherstack/weatherstack"

	"github.com/spf13/viper"
)

func main() {

	weatherClient, err := weatherstack.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация читалки конфига
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("configs/") // path to look for the config file in
	err = viper.ReadInConfig()      // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Fatalf("fatal error config file: %+v", err)
	}

	// Получаем JWT токен (если регистрация валидна)
	l := viper.GetString("weather.login")
	p := viper.GetString("weather.password")
	err = weatherClient.Login(l, p)
	if err != nil {
		log.Fatalf("Authorization error")
	}

	// Получаем списки локаций, и находим ID Киева
	fmt.Println()
	var ID int
	lct, err := weatherClient.GetLocations("Kiev")
	if err != nil {
		log.Fatal(err)
	} else {
		for _, lo := range lct.Locations {
			fmt.Printf("Name: %s (%s)\t\t ID = %d\n", lo.Name, lo.Country, lo.ID)
			if lo.Name == "Kyiv" && lo.Country == "Ukraine" {
				ID = lo.ID
			}
		}
	}
	if ID == 0 {
		log.Fatal("Kyiv Ukraine not found")
	}

	// Получаем погоду в Киеве
	fmt.Println()
	if fd, err := weatherClient.GetWeather(ID); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Forecast: Kyiv (Ukraine)")
		for _, forecast := range fd.ForecastDaily {
			fmt.Printf("%s:\t t.min = %d  t.max = %d\n", forecast.Date, forecast.MinTemp, forecast.MaxTemp)

		}
	}

	// Завершение (Удаление токена)
	fmt.Println()
	weatherClient.Logout()
}
