package config

import (
	"github.com/spf13/viper"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     uint16
	User     string
	Dbname   string
	Password string
	Sslmode  string
}

type Config struct {
	Mode string

	Logger struct {
		Level string
	}

	Http struct {
		Addr string
		Port int
	}

	DB struct {
		Private    DatabaseConfig
		UserAccess DatabaseConfig
	}
}

func (c *Config) IsDevMode() bool {
	return c.Mode == "dev"
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
