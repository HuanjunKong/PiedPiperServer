package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	host   string `yaml:"host"`
	port   int    `yaml:"port"`
	passwd string `yaml:"passwd"`
}

var ServiceConfig struct {
	httpPort int         `yaml:"http_port"`
	redis    RedisConfig `yaml:"redis"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	dir, _ := os.Getwd()
	configFile := filepath.Join(dir, "config.yaml")

	confFile, err := filepath.Abs(configFile)
	if err != nil {
		log.Printf("No correct config file: %s - %s", configFile, err.Error())
		os.Exit(1)
	}

	confBs, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Printf("Failed to read config fliel <%s> : %s", confFile, err.Error())
		os.Exit(1)
	}

	err = yaml.Unmarshal(confBs, &ServiceConfig)
	if err != nil {
		log.Printf("Failed to parse config fliel <%s> : %s", confFile, err.Error())
		os.Exit(1)
	}
}
