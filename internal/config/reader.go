package config

import (
	"log"

	dotenv "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envLocal = ".env.local"

type Config struct {
	Postgres Postgres
	Service  Service
	Site     Site
	//Auth     Auth
}

type Postgres struct {
	User    string `envconfig:"POSTGRES_USER"`
	Pass    string `envconfig:"POSTGRES_PASS"`
	Address string `envconfig:"POSTGRES_ADDR"`
	Port    string `envconfig:"POSTGRES_PORT"`
	Db      string `envconfig:"POSTGRES_DB"`
}

type Service struct {
	Port string `envconfig:"SERVICE_PORT"`
}

type Site struct {
	Site []string `envconfig:"SITES"`
}

//type Auth struct {
//	Auth string `envconfig:"AUTH"`
//}

func Read() *Config {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	_ = dotenv.Overload(envLocal)

	var cfg Config

	envconfig.MustProcess("", &cfg)

	return &cfg
}
