package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"app_name"`
	Version string `mapstructure:"version"`
}

func WriteJsonConfig[T interface{}](config T, path string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	return err
}
func LoadConfig[T interface{}](config T, path string) error {
	// viper.AddConfigPath(path)
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	// read from environment variablef
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	return nil
}
func LoadJsonConfig[T interface{}](config T, path string) error {
	viper.SetConfigFile(path)

	// Read the config from the JSON file
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Allow system environment variables to override config values
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Unmarshal the final config into the struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	return nil
}
func LoadENVConfig[T interface{}](config T, path string) error {
	viper.AddConfigPath(path)
	// Load the .env file (environment variables config)
	viper.SetConfigName(".env")  // Look for an .env file
	viper.SetConfigType("env")   // ENV file format
	err := viper.MergeInConfig() // Merge the .env config with the already-loaded JSON config
	if err != nil {
		return err
	}

	// Allow system environment variables to override config values
	viper.AutomaticEnv()

	// Unmarshal the final config into the struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	return nil
}

func WriteJSON(v interface{}, w io.Writer) error {
	return nil;
}