package models

import bs "github.com/ryanpback/raspberry_a_pi/bootstrap"

// AppConfig holds all application configuration
var AppConfig bs.Config

// Response is how all requests and responses are wrapped
type Response map[string]interface{}
