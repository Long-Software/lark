package config

type Config struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	AppPort       int    `mapstructure:"APP_PORT"`
	ApiQuota      int    `mapstructure:"API_QUOTA"`
}

