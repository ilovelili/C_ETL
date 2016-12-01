package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Credential user / password / database info
type Credential struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// DatabaseConfig client / credential info
type DatabaseConfig struct {
	Client     string `json:"client"`
	Server     string `json:"server"`
	Port       string `json:"port"`
	Database   string `json:"database"`
	Credential `json:"credential"`
}

// OnPremisesConfig On-Premises side config info
type OnPremisesConfig struct {
	DatabaseConfig `json:"db"`
}

// Config config info
type Config struct {
	OnPremisesConfig `json:"onpremises"`
}

// GetConfig parse config info from config.json
func GetConfig() (config *Config) {
	path, _ := filepath.Abs("../config.json")
	configFile, err := os.Open(path)
	defer configFile.Close()

	if err != nil {
		panic("opening config file: " + err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic("parsing config file " + err.Error())
	}

	return
}

// GetConnectionString Get db connection string by config
func (config *DatabaseConfig) GetConnectionString() string {
	client := config.Client

	// user:password@tcp(127.0.0.1:3306)/hello
	if client == "mysql" {
		return config.User + ":" +
			config.Password + "@tcp(" +
			config.Server + ":" +
			config.Port + ")/" +
			config.Database
	}

	// server=localhost;port=1433;user id=sa;password=123
	if client == "mssql" {
		return "server=" + config.Server + ";" +
			"port=" + config.Port + ";" +
			"user id=" + config.User + ";" +
			"password=" + config.Password
	}

	// postgres://pqgotest:password@localhost/pqgotest
	if client == "postgres" {
		return "postgres://" + config.User + ":" +
			config.Password + "@" +
			config.Server + ":" +
			config.Port + "/" +
			config.Database
	}

	return ""
}
