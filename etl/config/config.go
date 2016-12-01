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
	Database string `json:"database"`
}

// DatabaseConfig client / credential info
type DatabaseConfig struct {
	Client     string `json:"client"`
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
