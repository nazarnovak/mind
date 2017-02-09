package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"errors"
)

type Config struct {
	App `json:"app"`
	DB `json:"db"`
	Redis `json:"redis"`
}

type App struct {
	Greet string `json:"greet"`
	Port string
}

type DB struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Name string `json:"name"`
}

type Redis struct {
	URL string `json:"url"`
}

var conf Config

func LoadConfig() error {
	port := flag.String("port", "8080", "Server port")
	cFile := flag.String("conf", "config.json", "Config file")
	flag.Parse()

	raw, err := ioutil.ReadFile(*cFile)
	if err != nil {
		return err
	}

	json.Unmarshal(raw, &conf)

	if conf.App.Greet == "" || conf.DB.Host == "" || conf.Redis.URL == "" {
		return errors.New("Loading config failed")
	}

	conf.App.Port = *port

	return nil
}
