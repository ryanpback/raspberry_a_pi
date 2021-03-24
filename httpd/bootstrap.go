package main

import (
	bs "github.com/ryanpback/raspberry_a_pi/bootstrap"
	"github.com/ryanpback/raspberry_a_pi/helpers"
	"github.com/ryanpback/raspberry_a_pi/httpd/handlers"
	"github.com/ryanpback/raspberry_a_pi/models"
)

var appConfig bs.Config

func bootstrap() {
	c, err := bs.InitConfig(helpers.GetEnv(isTesting))
	if err != nil {
		panic(err)
	}

	appConfig = c
	helpers.AppConfig = c
	models.AppConfig = c
	handlers.AppConfig = c
}
