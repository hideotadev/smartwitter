package config

import (
	"os"

	"github.com/remeh/smartwitter/log"
)

type Config struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string

	Conn string
}

// EnvConfig returns the config set in
// environment variables.
// Could end the process if any of the env variable
// is unavailable.
func Env() Config {
	return Config{
		ConsumerKey:       readEnvVar("CONSUMER_KEY", true, ""),
		ConsumerSecret:    readEnvVar("CONSUMER_SECRET", true, ""),
		AccessToken:       readEnvVar("ACCESS_TOKEN", true, ""),
		AccessTokenSecret: readEnvVar("ACCESS_TOKEN_SECRET", true, ""),
		Conn:              readEnvVar("CONN", false, "host=/var/run/postgresql sslmode=disable user=smartwitter dbname=smartwitter password=smartwitter"),
	}
}

func readEnvVar(v string, mandatory bool, def string) string {
	var rv string
	if rv = os.Getenv(v); len(rv) == 0 {
		if mandatory {
			log.Error("Can't find the environment variable:", v)
			os.Exit(1)
		} else {
			log.Warning("Using default value for:", v)
			rv = def
		}
	}
	return rv
}