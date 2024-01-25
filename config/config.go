package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database       DatabaseConfig
		Server         ServerConfig
		Authentication AuthenticationConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
		TimeZone string
	}

	ServerConfig struct {
		Name         string
		Version      string
		RPCPort      string
		RESTPort     string
		Debug        bool
		Environment  string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	AuthenticationConfig struct {
		Key       string
		SecretKey string
		SaltKey   string
	}
)

func getProjectPath() (string, error) {
	// Get the file of the calling function (main in this case).
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("unable to get caller information")
	}

	// Traverse up to find the project root.
	projectPath := filepath.Dir(filename)
	for {
		if _, err := os.Stat(filepath.Join(projectPath, "go.mod")); err == nil {
			return projectPath, nil
		}

		// Move up one directory level.
		parentDir := filepath.Dir(projectPath)
		if parentDir == projectPath {
			return "", fmt.Errorf("unable to find project root")
		}
		projectPath = parentDir
	}
}

func LoadConfig(app string) (Config, error) {
	var folder string
	env := os.Getenv("APPLICATION_ENV")

	switch env {
	case "master", "dev", "uat", "sandbox", "localhost":
		folder = env
	default:
		folder = "dev"
	}

	projectPath, err := getProjectPath()
	if err != nil {
		fmt.Println("Error:", err)
		return Config{}, err
	}

	configFile := ""
	if(app == "auth") {
		configFile = "config_auth.yaml"
	}
	if(app == "web") {
		configFile = "config_web.yaml"
	}

	path := filepath.Join(projectPath, "config", folder, configFile)
	fmt.Println("config path dir: ", path)

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return Config{}, err
	}

	return config, nil
}
