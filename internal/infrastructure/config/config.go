package config

import (
	"os"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     uint16
	User     string
	Dbname   string
	Password string
	Schema   string
	Sslmode  string
}

type Config struct {
	Mode string

	Logger struct {
		Level string
	}

	HTTP struct {
		Addr string
		Port int
	}

	DB struct {
		Private DatabaseConfig
		Public  DatabaseConfig
	}

	Jwt struct {
		Secret     string
		TTLSeconds int32
	}

	CORS struct {
		AllowedOrigins []string
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
		v:   v,
		cfg: new(Config),
	}
}

func (c *Reader) Read() error {
	err := c.v.ReadInConfig()
	if err != nil {
		return err
	}

	err = c.v.Unmarshal(c.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Reader) Get() *Config {
	return c.cfg
}
