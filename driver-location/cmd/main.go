package main

import (
	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/cmd/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	
	app.Run()
}