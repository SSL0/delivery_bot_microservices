package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Version               float32 `mapstructure:"version"`
	DBUrl                 string  `mapstructure:"db_url"`
	MigrationsPath        string  `mapstructure:"migrations_path"`
	ListeningAddress      string  `mapstructure:"listening_address"`
	CatalogServiceAddress string  `mapstructure:"catalog_service_address"`
}

func LoadConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "$", ""))

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		if strings.Contains(value, "$") {
			expanded := os.ExpandEnv(value)
			viper.Set(key, expanded)
		}
	}

	var config Config
	err = viper.Unmarshal(&config)
	return config, err
}
