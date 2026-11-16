package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

var configuration *Config

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

	dbHost := os.Getenv("HOST")

	if dbHost == "" {
		fmt.Println("Database Host is required")
		os.Exit(1)
	}

	dbPrt := os.Getenv("PORT")

	if dbPrt == "" {
		fmt.Println("Database Port is required")
		os.Exit(1)
	}

	dbPort, err := strconv.ParseInt(dbPrt, 10, 64)

	if err != nil {
		fmt.Println("Database Port must be number")
		os.Exit(1)
	}

	dbName := os.Getenv("NAME")

	if dbName == "" {
		fmt.Println("Database name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("USER")

	if dbUser == "" {
		fmt.Println("Database user is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("PASSWORD")

	if dbPassword == "" {
		fmt.Println("Database password is required")
		os.Exit(1)
	}

	enableSslMode := os.Getenv("ENABLE_SSL_MODE")
	enbSslMode, err := strconv.ParseBool(enableSslMode)

	if err != nil {
		fmt.Println("Invalid enable ssl mode value.", err)
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbPort),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		EnableSSLMode: enbSslMode,
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port), // type casting
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
