package config

import (
	"encoding/json"
	"flag"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

// Represents database server and credentials
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat              string
	PostgresqlConnectionString string
}

// TODO application.Config vs config.Get()
var config Config

// read and parse the configuration file
func (c *Config) read() {
	var configPath string
	// TODO override with environment variables
	flag.StringVar(&configPath, "config-path", "./config.json", "Config file")
	flag.Parse()
	// TODO relevant path to the runner (app)
	file, e := ioutil.ReadFile(configPath)
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
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config = Config{}
	config.read()

	return Get()
}

func Get() *Config {
	return &config
}
