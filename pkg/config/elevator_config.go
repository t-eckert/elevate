package config

import (
	"fmt"

	"github.com/google/uuid"
)

type Config struct {
	Id                string
	TLS               bool
	Host              string
	Port              int
	Floors            int
	MaxSpeed          float64
	ControllerAddress string
}

// NewConfig creates a new instance of elevator config with default settings.
func NewConfig() *Config {
	return &Config{
		Id:                uuid.NewString(),
		TLS:               false,
		Host:              "127.0.0.1",
		Port:              4000,
		Floors:            30,
		MaxSpeed:          1,
		ControllerAddress: "http://127.0.0.1:3000",
	}
}

func (c *Config) WithId(id string) *Config {
	c.Id = id
	return c
}

func (c *Config) WithTLS(tls bool) *Config {
	c.TLS = tls
	return c
}

func (c *Config) WithHost(host string) *Config {
	c.Host = host
	return c
}

func (c *Config) WithPort(port int) *Config {
	c.Port = port
	return c
}

func (c *Config) WithFloors(floors int) *Config {
	c.Floors = floors
	return c
}

func (c *Config) WithMaxSpeed(maxSpeed float64) *Config {
	c.MaxSpeed = maxSpeed
	return c
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s://%s:%d", cond(c.TLS, "https", "http"), c.Host, c.Port)
}

func cond[T string | int](test bool, a, b T) T {
	if test {
		return a
	}
	return b
}
