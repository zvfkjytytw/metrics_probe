package metricsserver

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func GetConfig(configFile string) (config *AppConfig, err error) {
	filename, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err //nolint // wraped higher
	}

	yamlConfig, err := os.ReadFile(filename)
	if err != nil {
		return nil, err //nolint // wraped higher
	}

	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		return nil, err //nolint // wraped higher
	}

	return config, nil
}
