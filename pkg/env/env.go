package env

import "github.com/spf13/viper"

func Load(config interface{}, path string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	err := viper.MergeInConfig()
	if err != nil {
		return err
	}

	// Allow system environment variables to override config values
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return err
}
