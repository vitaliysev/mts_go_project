package config

import (
	"errors"
	"net"
	"os"
)

const (
	restHostEnvName = "REST_HOST"
	restPortEnvName = "REST_PORT"
)

type RESTConfig interface {
	Address() string
}

type restConfig struct {
	host string
	port string
}

func NewRESTConfig() (RESTConfig, error) {
	host := os.Getenv(restHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("rest host not found")
	}

	port := os.Getenv(restPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("rest port not found")
	}

	return &restConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *restConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
