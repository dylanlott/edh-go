package config

type Config struct {
	Port        int
	Environment string
}

func NewConfig() Config {
	return Config{}
}
