package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// "gopkg.in/yaml.v2"

type AppConfig struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Token    string   `yaml:"tokenF"`
}

type Server struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	ModeRun  string `yaml:"mode_run"`  // prod or debug
	LevelLog int    `yaml:"level_log"` // level log from 0 -> 6
}

type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type S3AWS struct {
	BucketName string `yaml:"bucket_name"`
	Region     string `yaml:"region"`
	ApiKey     string `yaml:"api_key"`
	Secret     string `yaml:"secret"`
	Domain     string `yaml:"domain"`
}

func NewAppConfig(configPath string) (*AppConfig, error) {
	var appCfg AppConfig
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}
	err = yaml.Unmarshal(file, &appCfg)

	return &appCfg, nil
}
