package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Logger struct {
		Level string
	}
	Http struct {
		Addr string
		Port int
	}
}

type Reader struct {
	v   *viper.Viper
	cfg *Config
}

func NewReader() *Reader {
	v := viper.New()
	v.AutomaticEnv()

	configFile := os.Getenv("CONFIG_FILE")
	if len(configFile) == 0 {
		configFile = "config.yaml"
	}

	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	return &Reader{
		v: v,
	}
}

func (c *Reader) Read() error {
	err := c.v.ReadInConfig()
	if err != nil {
		return err
	}

	c.cfg = new(Config)
	err = c.v.Unmarshal(c.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Reader) Get() *Config {
	return c.cfg
}
