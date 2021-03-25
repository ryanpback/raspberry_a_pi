package handlers

import (
	"database/sql"

	bs "github.com/ryanpback/raspberry_a_pi/bootstrap"
	"github.com/sirupsen/logrus"
)

// AppConfig holds all application configuration
var AppConfig bs.Config

// DBConn this is here so that we can hydrate all the web handlers with the same
// DB connection that we are using in the main package
var DBConn *sql.DB

// Log this is here so we can share the same logger with the main package
var Log *logrus.Logger

// Response is the response type to return from requests
type Response map[string]interface{}
