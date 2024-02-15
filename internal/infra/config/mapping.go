package config

type Config struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Cache    Cache    `mapstructure:"cache"`
}

type App struct {
	Name        string `mapstructure:"name"`
	Debug       bool   `mapstructure:"debug"`
	Environment string `mapstructure:"environment"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Cache struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DB     int    `mapstructure:"db"`
	Pass   string `mapstructure:"pass"`
	Prefix string `mapstructure:"prefix"`
}
