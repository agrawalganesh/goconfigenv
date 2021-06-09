package main

import (
	"fmt"

	"github.com/iamganeshagrawal/goconfigenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port   int    `configenv:"PORT"`
	Secret string `configenv:"JWT_SECRET"`
	HTTPS  bool   `configenv:"HTTPS"`
}

func main() {
	var config Config
	goconfigenv.Load(&config)
	fmt.Println(config)
}
