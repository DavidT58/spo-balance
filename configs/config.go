package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type PoolConfig struct {
	Name   string `json:"name" yaml:"name"`
	PoolID string `json:"poolID" yaml:"poolID"`
}

type Config struct {
	Pools []PoolConfig `json:"pools"`
}

var storedConfig Config

func LoadConfigFromYAML(filePath string) (Config, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}

func (c Config) getPools() []map[string]string {
	var result []map[string]string
	for _, pool := range c.Pools {
		result = append(result, map[string]string{
			"name":   pool.Name,
			"poolID": pool.PoolID,
		})
	}
	return result
}
