package config

import (
	"errors"
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	Connection    string        `json:"connection" yaml:"connection"`
	LoggingConfig LoggingConfig `json:"logging" yaml:"logging"`
	Accounts      []struct {
		Name    string `json:"name" yaml:"name"`
		Key     string `json:"key" yaml:"key"`
		Address string `json:"address" yaml:"address"`
	} `json:"accounts" yaml:"accounts"`
	TwitterUserAddr string `json:"twitterUserAddr" yaml:"twitterUserAddr"`
	TwitterAddr     string `json:"twitterAddr" yaml:"twitterAddr"`
}

type LoggingConfig struct {
	Development bool `json:"development" yaml:"development"`
}

func NewConfig(configFile string) (*Config, error) {
	if configFile == "" {
		return nil, errors.New("you must provide configFile")
	}
	k := koanf.New(".")

	if err := k.Load(file.Provider(configFile), json.Parser()); err != nil {
		fmt.Println("Failed to load configration from json file ", configFile)
		return nil, err
	}
	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		fmt.Printf("failed to unmarshal with conf. err: %v", err)
	}

	return &config, nil
}
