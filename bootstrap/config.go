package bootstrap

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config holds the structure of the config
type Config struct {
	AppPort string
	Logger  *logrus.Logger
}

// InitConfig sets the app config
func InitConfig(envFile string) (Config, error) {
	godotenv.Load("../" + envFile)

	c := Config{
		AppPort: os.Getenv("APP_PORT"),
	}

	// initialize logging to stdout
	c.Logger = logrus.New()
	c.Logger.Out = os.Stdout

	if c.AppPort == "" {
		return c, fmt.Errorf("APP_PORT is required")
	}

	return c, nil
}
