package service

// modelに含めないビジネスロジックを記述
// 重複するメールアドレスの登録を拒否する

import (
	"fmt"
	// "server/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (us *UserService) Duplicated(email string) error {
	user, err := us.repo.GetByEmail(email)
	if user != nil {
		return fmt.Errorf("%s already exists", email)
	}
	if err != nil {
		return err
	}
	return nil
}
