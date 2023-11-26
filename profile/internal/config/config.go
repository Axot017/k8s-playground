package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port  int
	Debug bool
}

func NewConfig() (*Config, error) {
	port, err := getPort()
	if err != nil {
		return nil, err
	}
	return &Config{
		Port:  port,
		Debug: os.Getenv("DEBUG") == "true",
	}, nil
}

func getPort() (int, error) {
	portStr := os.Getenv("PORT")
	var port int
	if portStr == "" {
		port = 8001
	} else {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			return 0, err
		}
		port = p
	}

	return port, nil
}
