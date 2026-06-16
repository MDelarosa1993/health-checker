package config

import (
	"flag"
	"os"
)

type Config struct {
	URLToCheck string
}

func Load() Config {
	defaultURL := os.Getenv("URL_TO_CHECK")

	if defaultURL == "" {
		defaultURL = "https://example.com"
	}

	url := flag.String("url", defaultURL, "URL to check")
	flag.Parse()

	return Config{
		URLToCheck: *url,
	}
}
