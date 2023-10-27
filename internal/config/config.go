package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	App struct {
		Port string `env:"PORT" env-default:"8000"`
	}
	Sqlite struct {
		StoragePath string `env:"STORAGE_PATH"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			panic(err)
		}
	})

	return instance
}
