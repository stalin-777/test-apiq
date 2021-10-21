package config

import (
	"runtime"
	"time"

	"github.com/juju/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stalin-777/test-apiq/internal/logger"
)

type Config struct {
	Web struct {
		Host string `validate:"required"`
		Port int    `validate:"required"`
	}
	Logger     logger.Config `validate:"required"`
	WorkersNum int
	TTL        time.Duration `validate:"required"`
}

func New() (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	cfg := &Config{
		Logger: logger.Config{},
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Annotate(err, "Failed to load configuration file")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Annotate(err, "Failed to unmarshal configuration file")
	}

	pflag.IntVarP(&cfg.WorkersNum, "workersNum", "w", runtime.GOMAXPROCS(0), "Number of workers. Default number of processors")
	pflag.StringVarP(&cfg.Web.Host, "host", "h", cfg.Web.Host, "Hostname. Default value from config file")
	pflag.IntVarP(&cfg.Web.Port, "port", "p", cfg.Web.Port, "Port. Default value from config file")
	pflag.Parse()

	return cfg, nil
}
