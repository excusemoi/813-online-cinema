package config

type Config struct {
	DbHost     string `mapstructure:"db_host"`
	DBPort     uint16 `mapstructure:"db_port"`
	DbUser     string `mapstructure:"db_user"`
	DbPassword string `mapstructure:"db_password"`
	DbName     string `mapstructure:"db_name"`
}
