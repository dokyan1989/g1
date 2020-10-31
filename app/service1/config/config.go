package config

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/dokyan1989/g1/app/service1/config/db"
	"github.com/dokyan1989/g1/app/service1/config/server"
	"github.com/spf13/viper"
)

// Config application
type Config struct {
	MySQL  db.MySQLConfig      `json:"mysql" mapstructure:"mysql"`
	Server server.ServerConfig `json:"server" mapstructure:"server"`
}

// Load loads system env config
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./app/service1")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	c := loadDefaultConfig()
	if configBuffer, err := json.Marshal(c); err != nil {
		log.Println("Oops! Marshal config is failed. ", err)
		return nil, err
	} else if err := viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
		log.Println("Oops! Read default config is failed. ", err)
		return nil, err
	}

	if err := viper.MergeInConfig(); err != nil {
		log.Println("Read config file failed.", err)
	}

	err := viper.Unmarshal(c)

	log.Printf("%+v\n", c)

	return c, err
}

func loadDefaultConfig() *Config {
	return &Config{
		MySQL:  db.MySQLDefaultConfig(),
		Server: server.DefaultServerConfig(),
	}
}
