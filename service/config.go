package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Debug     bool     `json:"debug"`
	Appkey    string   `json:"appkey"`
	Secret    string   `json:"secret"`
	AllowHost []string `json:"allow_host"`
}

func ReadConfigFromFile(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
