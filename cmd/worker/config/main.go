package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type WorkerConfig struct {
	MasterHost string `yaml:"masterHost"`
}

var AppConfig WorkerConfig

func LoadCfg() error {
	var (
		cfgFilePath string
	)
	cfgFilePath = os.Getenv("WORKER_CFG")
	if cfgFilePath == "" {
		cfgFilePath = "config.yml"
	}
	bytes, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return err
	}
	errUnmarshal := yaml.Unmarshal(bytes, &AppConfig)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return nil
}
