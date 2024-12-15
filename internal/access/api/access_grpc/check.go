package access_grpc

import (
	"context"
	"errors"
	"github.com/vitaliysev/mts_go_project/internal/auth/utils"
	descAccess "github.com/vitaliysev/mts_go_project/pkg/access_v1"
	"google.golang.org/grpc/metadata"
	"strings"
)

func (i *Implementation) Check(ctx context.Context, req *descAccess.CheckRequest) (*descAccess.CheckResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return nil, errors.New("access token is invalid")
	}

	accessibleMap, err := i.accessService.AccessibleRoles(ctx)
	if err != nil {
		return nil, errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[req.GetEndpointAddress()]
	if !ok {
		return nil, nil
	}

	if role == claims.Role {
		return &descAccess.CheckResponse{Username: claims.Username}, nil
	}

	return nil, errors.New("access denied")
}
