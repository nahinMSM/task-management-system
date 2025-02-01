package usecase

import (
	"backend/internal/domain"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo}
}

func (u *UserUseCase) Register(ctx context.Context, user *domain.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email e senha são obrigatórios")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.UserRepo.Create(user)
}

func (u *UserUseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := u.UserRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	return user, nil
}

func (u *UserUseCase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.UserRepo.GetByID(id)
}
