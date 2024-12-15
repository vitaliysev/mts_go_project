package converter

import (
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
)

func ToAuthFromRepo(auth *model.Auth) *model.Auth {
	return &model.Auth{
		Info:      ToAuthInfoFromRepo(auth.Info),
		CreatedAt: auth.CreatedAt,
	}
}

func ToAuthInfoFromRepo(info model.AuthInfo) model.AuthInfo {
	return model.AuthInfo{
		Login:           info.Login,
		Hashed_password: info.Hashed_password,
		Role:            info.Role,
	}
}
