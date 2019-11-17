package toggler

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	ProjectDir string `env:"PROJECT_DIR,required"`
}

var Cfg Config

func Load() Config {
	cfg := Config{}
	err := env.Parse(&cfg)

	if err == nil {
		log.Printf("config/load : Environment variables loaded from shell")
		Cfg = cfg
		return Cfg
	}

	godotenv.Load()

	err = env.Parse(&cfg)
	if err == nil {
		log.Printf("config/load : Environment variables loaded from .env file")
		Cfg = cfg
		return Cfg
	}

	if err != nil {
		log.Printf("config/load : Failed to load environment variables")
		panic(fmt.Sprintf("%+v\n", err))
	}

	return Config{}
}
