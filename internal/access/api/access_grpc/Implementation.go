package access_grpc

import (
	"github.com/vitaliysev/mts_go_project/internal/access/service"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/access_v1"
)

const (
	authPrefix = "Bearer "

	accessTokenSecretKey = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="
)

type Implementation struct {
	descAuth.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
