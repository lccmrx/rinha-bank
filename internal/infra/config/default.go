package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("app.name", "rinha-bank")
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.environment", "local")

	viper.SetDefault("server.port", 5000)

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.database", "rinha-bank")
}
