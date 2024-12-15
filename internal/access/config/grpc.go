package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	grpcAccessHostEnvName = "GRPC_ACCESS_HOST"
	grpcAccessPortEnvName = "GRPC_ACCESS_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcAccessConfig struct {
	host string
	port string
}

func NewGRPCAccessConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcAccessHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcAccessPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcAccessConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcAccessConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
