package config

import (
	"fmt"
	"net/url"
)

// MySQLConfig mysql configuration
type MySQLConfig struct {
	Username string `json:"username" mapstructure:"username" yaml:"username"`
	Password string `json:"password" mapstructure:"password" yaml:"password"`
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Port     int    `json:"port" mapstructure:"port" yaml:"port"`
	Database string `json:"database" mapstructure:"database" yaml:"database"`
	Options  string `json:"options" mapstructure:"options" yaml:"options"`
}

// MySQLDefaultConfig creates mysql default config
func MySQLDefaultConfig() MySQLConfig {
	return MySQLConfig{
		Username: "root",
		Password: "nopassword",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "noname",
		Options:  "?charset=utf8mb4&parseTime=true",
	}
}

// DSN returns mysql data source name
func (c MySQLConfig) DSN() string {
	options := c.Options
	if options != "" {
		if options[0] != '?' {
			options = "?" + options
		}
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.Username,
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Database,
		options,
	)
}

func (c MySQLConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.DSN())
}
