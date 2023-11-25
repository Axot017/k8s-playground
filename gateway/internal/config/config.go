package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Port              int
	Debug             bool
	ProfileServiceUrl url.URL
}

func NewConfig() (*Config, error) {
	port, err := getPort()
	if err != nil {
		return nil, err
	}
	profileServiceUrl, err := getProfileServiceUrl()
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:              port,
		ProfileServiceUrl: profileServiceUrl,
		Debug:             os.Getenv("DEBUG") == "true",
	}, nil
}

func getProfileServiceUrl() (url.URL, error) {
	profileServiceUrl := os.Getenv("PROFILE_SERVICE_DOMAIN_NAME")
	if profileServiceUrl == "" {
		return url.URL{}, errors.New("PROFILE_SERVICE_DOMAIN_NAME is not set")
	}
	u, err := url.Parse(fmt.Sprintf("http://%s", profileServiceUrl))
	if err != nil {
		return url.URL{}, fmt.Errorf("failed to parse PROFILE_SERVICE_DOMAIN_NAME: %w", err)
	}
	return *u, nil
}

func getPort() (int, error) {
	portStr := os.Getenv("PORT")
	var port int
	if portStr == "" {
		port = 8000
	} else {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			return 0, err
		}
		port = p
	}

	return port, nil
}
