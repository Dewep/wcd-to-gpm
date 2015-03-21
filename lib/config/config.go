package config

import "code.google.com/p/gcfg"

type ConfigUser struct {
	Username string
	Password string
}

type Config struct {
	Whatcd ConfigUser
	Transmission ConfigUser
	Googleplaymusic ConfigUser
}

func Get(file string) Config {
	var cfg Config

	err := gcfg.ReadFileInto(&cfg, file)

	if err != nil {
		panic(err)
	}

	return cfg
}
