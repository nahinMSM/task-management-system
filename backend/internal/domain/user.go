package domain

// User representa um usuário do sistema
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Não retornaremos a senha na API por questões de segurança
}

// UserRepository define a interface para operações de persistência relacionadas a usuários
type UserRepository interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id int64) error
}
