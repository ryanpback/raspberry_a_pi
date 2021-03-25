package helpers

import (
	bs "github.com/ryanpback/raspberry_a_pi/bootstrap"
	"github.com/sirupsen/logrus"
)

var isTesting bool

// AppConfig holds all application configuration
var AppConfig bs.Config

// Log this is here so we can share the same logger with the main package
var Log *logrus.Logger
