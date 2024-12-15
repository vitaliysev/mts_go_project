package converter

import (
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/wrapperspb"

	desc "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
)

func ToCreateFromSignin(info *desc.SigninRequest) (auth *model.AuthInfo) {
	return &model.AuthInfo{
		Login:           info.GetUsername(),
		Hashed_password: HashPassword(info.GetPassword()),
		Role:            info.GetRole(),
	}
}

func unwrapStringValue(value *wrapperspb.StringValue) string {
	if value != nil {
		return value.Value
	}
	return ""
}

func HashPassword(password string) string {
	// Генерация хэша пароля с использованием bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
