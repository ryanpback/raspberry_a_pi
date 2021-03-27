package helpers

import (
	"github.com/sirupsen/logrus"
)

var isTesting bool

// Log this is here so we can share the same logger with the main package
var Logger *logrus.Logger
