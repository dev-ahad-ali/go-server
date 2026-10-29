package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

var configuration Config

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load the env variables", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	configuration = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port), // type casting
		JwtSecretKey: jwtSecretKey,
	}

}

func GetConfig() Config {
	loadConfig()
	return configuration
}
