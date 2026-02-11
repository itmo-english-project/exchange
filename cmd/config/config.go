package config

import (
	"bytes"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/itmo-english-project/common/pkg/storage/postgres"
)

type Config struct {
	General  GeneralConfig  `yaml:"general"`
	Database DatabaseConfig `yaml:"db"`
}

type GeneralConfig struct {
	Development bool `yaml:"development"`
	Port        int  `yaml:"port"`
}
type DatabaseConfig struct {
	Exchanges *postgres.Config `yaml:"exchanges"`
}

func InitConfig(path string) (Config, error) {
	var (
		cfgRaw []byte
		err    error
	)

	if fp := os.Getenv(path); fp != "" {
		if cfgRaw, err = os.ReadFile(filepath.Clean(fp)); err != nil {
			return Config{}, err
		}
	}

	var cfg Config
	env := os.ExpandEnv(string(cfgRaw))
	d := yaml.NewDecoder(bytes.NewReader([]byte(env)))
	d.KnownFields(true)
	if err = d.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
