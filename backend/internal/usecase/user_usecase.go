package usecase

import (
	"context"
	"errors"

	// "github.com/seu_projeto/internal/domain"
	"backend/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

// UserUseCase define as operações relacionadas a usuários
type UserUseCase struct {
	UserRepo domain.UserRepository
}

// NewUserUseCase cria uma nova instância de UserUseCase
func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo}
}

// Register registra um novo usuário
func (u *UserUseCase) Register(ctx context.Context, user *domain.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email e senha são obrigatórios")
	}

	// Hash da senha antes de salvar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.UserRepo.Create(user)
}

// Login autentica um usuário
func (u *UserUseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := u.UserRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	// Verifica a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	return user, nil
}

// GetByID retorna um usuário pelo ID
func (u *UserUseCase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.UserRepo.GetByID(id)
}
