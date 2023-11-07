package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Stage       string `mapstructure:"UP_STAGE"`
	ConfigStage string `mapstructure:"CONFIG_STAGE"`
	Region      string `mapstructure:"REGION"`
	Host        string `mapstructure:"HOST"`
	Port        int    `mapstructure:"PORT"`
	DbDsn       string `mapstructure:"DB_DSN"`
}

var defaultConfig = []byte(`
app: pm-back-office
env: dev
PORT: 8088
DB_DSN: postgres://postgres:postgres123@authdb.cozhhzyndzwt.ap-southeast-1.rds.amazonaws.com:5432/authdb?sslmode=disable&connect_timeout=5
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
