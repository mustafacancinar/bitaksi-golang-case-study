package main

import (
	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/cmd/app"
)

// @title Driver Location API
// @version 1.0
// @description This is a sample API.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}