package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	grpcAuthHostEnvName = "GRPC_AUTH_HOST"
	grpcAuthPortEnvName = "GRPC_AUTH_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcAuthConfig struct {
	host string
	port string
}

func NewGRPCAuthConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcAuthHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcAuthPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}
	return &grpcAuthConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcAuthConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
