package auth

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
)

func (s *serv) Create(ctx context.Context, info *model.AuthInfo) (string, error) {
	var login string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		login, errTx = s.authRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}
		curr, err := s.authRepository.Get(ctx, info.Login)
		if err != nil {
			return errTx
		}
		fmt.Println(curr.Info.Login)
		return nil
	})

	if err != nil {
		return "nothing", err
	}

	return login, nil
}
