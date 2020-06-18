package main

import (
	"fmt"
	"owm-telegram/owm"
	"owm-telegram/telegram"
)

var coord owm.Coordinates
var reqLoc string = "Share location? I actually need it"

func main() {

	fmt.Println("Telegram bot started")
	owm.WeatherIcons, _ = owm.InitMapWeather()

	telegram.RunBot()

}
