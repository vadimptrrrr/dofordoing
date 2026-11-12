package config

type Config struct {
	BackendMode string
	DataPath    string
}

func Load() (Config, error)
