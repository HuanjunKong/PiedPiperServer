package conf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Passwd string `yaml:"passwd"`
}

var ServiceConfig struct {
	HttpPort string      `yaml:"http_port"`
	Redis    RedisConfig `yaml:"redis"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	dir, _ := os.Getwd()
	configFile := filepath.Join(dir, "conf/config.yaml")

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
