package config

type Config struct {
	Port        int
	Environment string
}

func NewConfig() Config {
	return Config{
		Port:        9090,
		Environment: "development",
	}
}
