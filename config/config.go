package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
)

var Config *YamlConfig
var Version string
var BestdoriAPIUrl *url.URL

type YamlConfig struct {
	RunAddr string        `yaml:"run-addr"`
	Debug   bool          `yaml:"debug"`
	API     YamlConfigAPI `yaml:"api"`
}

type YamlConfigAPI struct {
	BestdoriProxy string `yaml:"bestdori-proxy"`
}

func NewYamlConfig() *YamlConfig {
	return &YamlConfig{
		RunAddr: "0.0.0.0:8080",
		Debug:   true,
		API: YamlConfigAPI{
			BestdoriProxy: "https://proxy.bestdori.com",
		},
	}
}

func init() {
	Version = "2.2.0"
	Config = NewYamlConfig()
	var err error
	if envEnabled() {
		initEnv()
	} else {
		yamlFile, err := os.ReadFile("conf.yaml")
		if err != nil {
			yamlConfig, _ := yaml.Marshal(Config)
			err = os.WriteFile("conf.yaml", yamlConfig, 0666)
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Printf("conf.yaml not found, one is generate!")
			os.Exit(0)
		}
		err = yaml.Unmarshal(yamlFile, Config)
		if err != nil {
			log.Fatal("Check the conf.yaml, Cannot Read!")
		}
	}
	BestdoriAPIUrl, err = url.Parse(Config.API.BestdoriProxy)
	if err != nil {
		log.Fatal("Cannot parse BestdoriAPI!")
	}
}
