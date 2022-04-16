package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config config
type Config struct {
	DB struct {
		Host     string
		Port     uint16
		User     string
		Password string
	}
	LDAP struct {
		Enable      bool
		Placeholder struct {
			Username string
			Password string
		}
		Addr       string
		BaseDN     string `yaml:"baseDN"`
		Level      uint32
		OrgID      uint32 `yaml:"orgID"`
		AttrMapKey struct {
			Name  string
			Email string
		} `yaml:"attrMapKey"`
	}
	Teams struct {
		Enable     bool
		Entrypoint string
	}
}

// ParseConfig parses a file to config
func ParseConfig(configFile string) (*Config, error) {
	var (
		data []byte
		err  error
	)
	c := new(Config)
	if data, err = ioutil.ReadFile(configFile); err == nil {
		if err = yaml.Unmarshal(data, c); err == nil {
			return c, nil
		}
	}
	return nil, err
}
