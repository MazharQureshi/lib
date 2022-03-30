package config

import "fmt"

var Shark = "Testor"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "blog",
			Charset:  "utf8",
		},
	}
}

func Hello() {
	fmt.Println("Hello, World Testor22!")
}
