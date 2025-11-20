package services

import (
	"errors"
	"time"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/repositories"
)

var (
	ErrInvalidUserData = errors.New("dados do usuário inválidos")
	ErrEmailExists     = errors.New("email já cadastrado")
)

// UserService contém a lógica de negócio para usuários
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService cria uma nova instância do serviço de usuários
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAll retorna todos os usuários
func (s *UserService) GetAll() []models.User {
	return s.repo.GetAll()
}

// GetByID retorna um usuário pelo ID
func (s *UserService) GetByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, ErrInvalidUserData
	}
	return s.repo.GetByID(id)
}

// Create cria um novo usuário
func (s *UserService) Create(req models.UserRequest) (*models.User, error) {
	if req.Name == "" || req.Email == "" {
		return nil, ErrInvalidUserData
	}

	// Verifica se o email já existe
	users := s.repo.GetAll()
	for _, u := range users {
		if u.Email == req.Email {
			return nil, ErrEmailExists
		}
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Active:   true,
		CreateAt: time.Now().UTC().Format(time.RFC3339),
	}

	if user.Role == "" {
		user.Role = "user"
	}

	created := s.repo.Create(user)
	return &created, nil
}

// Update atualiza um usuário existente
func (s *UserService) Update(id int, req models.UserRequest) (*models.User, error) {
	if id <= 0 {
		return nil, ErrInvalidUserData
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:       existing.ID,
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Active:   existing.Active,
		CreateAt: existing.CreateAt,
	}

	if user.Name == "" {
		user.Name = existing.Name
	}
	if user.Email == "" {
		user.Email = existing.Email
	}
	if user.Role == "" {
		user.Role = existing.Role
	}

	return s.repo.Update(id, user)
}

// Delete remove um usuário
func (s *UserService) Delete(id int) error {
	if id <= 0 {
		return ErrInvalidUserData
	}
	return s.repo.Delete(id)
}

