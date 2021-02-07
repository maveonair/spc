package config

import (
	"log"
	"os"
	"path"
	"time"

	"git.sr.ht/~maveonair/spc/internal/release"
	"github.com/BurntSushi/toml"
)

const (
	defaultInterval = 24 * time.Hour
)

type Config struct {
	Interval time.Duration              `toml:"interval"`
	Releases map[string]release.Release `toml:"releases"`
}

func NewConfig() (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configFile(), &config); err != nil {
		return nil, err
	}

	config.Interval = defaultInterval

	return &config, nil
}

func configFile() string {
	if envConfigFile := os.Getenv("SPC_CONFIG_FILE"); envConfigFile != "" {
		return envConfigFile
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get user home directory")
	}

	return path.Join(homeDir, ".spc", "config.toml")
}
