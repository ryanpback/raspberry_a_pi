package main

import (
	bs "github.com/ryanpback/raspberry_a_pi/bootstrap"
	"github.com/ryanpback/raspberry_a_pi/helpers"
	"github.com/ryanpback/raspberry_a_pi/models"
)

var appConfig bs.Config

func bootstrap() {
	c, err := bs.InitConfig()
	appConfig = c

	// let's now hydrate a few things
	models.DBConn = appConfig.DBConn
	helpers.Logger = appConfig.Logger

	if err != nil {
		helpers.LogError(err)
		helpers.LogInfo(appConfig)
		panic("Something went wrong, check your environment variables")
	}

	helpers.LogInfo("Application bootstrapped with the following settings:")
	helpers.LogInfo("Port: " + appConfig.AppPort)
	helpers.LogInfo("Database: " + appConfig.DBDatabase)
	helpers.LogInfo("DB Username: " + appConfig.DBUsername)
}
