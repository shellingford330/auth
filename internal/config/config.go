package config

import (
	"github.com/kelseyhightower/envconfig"
)

var Config *Configuration

func init() {
	var err error
	Config, err = NewConfiguration()
	if err != nil {
		panic(err)
	}
}

type config struct {
	Env string `default:"development"`

	// HTTP Server
	HTTPServer struct {
		Host string
		Port string `default:"4545"`
	} `envconfig:"HTTP_SERVER"`

	GRPCServer struct {
		Host string `default:"0.0.0.0"`
		Port string `default:"50051"`
	} `envconfig:"GRPC_SERVER"`

	// DB
	DB struct {
		Name     string `default:"shellinford"`
		Host     string `default:"localhost"`
		Port     string `default:"3306"`
		User     string `default:"root"`
		Password string `default:"password"`
	}
}

type Configuration struct {
	config
}

func NewConfiguration() (*Configuration, error) {
	config := &Configuration{}
	err := envconfig.Process("", &config.config)
	return config, err
}

func (c *Configuration) Env() string {
	return c.config.Env
}

func (c *Configuration) HTTPServerHost() string {
	return c.config.HTTPServer.Host
}

func (c *Configuration) HTTPServerPort() string {
	return c.config.HTTPServer.Port
}

func (c *Configuration) GRPCServerHost() string {
	return c.config.GRPCServer.Host
}

func (c *Configuration) GRPCServerPort() string {
	return c.config.GRPCServer.Port
}

func (c *Configuration) DBName() string {
	return c.config.DB.Name
}

func (c *Configuration) DBHost() string {
	return c.config.DB.Host
}

func (c *Configuration) DBPort() string {
	return c.config.DB.Port
}

func (c *Configuration) DBUser() string {
	return c.config.DB.User
}

func (c *Configuration) DBPassword() string {
	return c.config.DB.Password
}
