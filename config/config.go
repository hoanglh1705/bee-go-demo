package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppName     string `mapstructure:"APP"`
	Stage       string `mapstructure:"UP_STAGE"`
	ConfigStage string `mapstructure:"CONFIG_STAGE"`
	Host        string `mapstructure:"HOST"`
	Port        int    `mapstructure:"PORT"`
	DbDsn       string `mapstructure:"DB_DSN"`
}

var defaultConfig = []byte(`
APP: bee-go-demo
UP_STAGE: dev
HOST: localhost
PORT: 8188
DB_DSN: postgres://user:user123@localhost:5443/myuser?sslmode=disable&connect_timeout=5
`)

func Load() (*Configuration, error) {
	var cfg = &Configuration{}
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
