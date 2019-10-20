package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// todo required camelcase, do not now yet why
type PG struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

// Represents database server and credentials
type Config struct {
	Port string
	Pg   PG
}

// TODO application.Config vs config.Get()
var config Config

// read and parse the configuration file
func (c *Config) read() {
	// TODO relevant path to the runner (app)
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}

	var err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}
}

func Load() *Config {
	config = Config{}
	config.read()

	return Get()
}

func Get() *Config {
	return &config
}
