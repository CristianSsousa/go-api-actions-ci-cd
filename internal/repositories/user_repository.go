package repositories

import (
	"errors"
	"sync"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
)

var (
	ErrUserNotFound = errors.New("usuário não encontrado")
)

// UserRepository gerencia os dados de usuários em memória
type UserRepository struct {
	mu     sync.RWMutex
	users  []models.User
	nextID int
}

// NewUserRepository cria uma nova instância do repositório com dados pré-prontos
func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		users: []models.User{
			{
				ID:       1,
				Name:     "João Silva",
				Email:    "joao.silva@example.com",
				Role:     "admin",
				Active:   true,
				CreateAt: "2024-01-15T10:00:00Z",
			},
			{
				ID:       2,
				Name:     "Maria Santos",
				Email:    "maria.santos@example.com",
				Role:     "user",
				Active:   true,
				CreateAt: "2024-01-16T11:30:00Z",
			},
			{
				ID:       3,
				Name:     "Pedro Oliveira",
				Email:    "pedro.oliveira@example.com",
				Role:     "user",
				Active:   false,
				CreateAt: "2024-01-17T14:20:00Z",
			},
			{
				ID:       4,
				Name:     "Ana Costa",
				Email:    "ana.costa@example.com",
				Role:     "manager",
				Active:   true,
				CreateAt: "2024-01-18T09:15:00Z",
			},
			{
				ID:       5,
				Name:     "Carlos Ferreira",
				Email:    "carlos.ferreira@example.com",
				Role:     "user",
				Active:   true,
				CreateAt: "2024-01-19T16:45:00Z",
			},
		},
		nextID: 6,
	}
	return repo
}

// GetAll retorna todos os usuários
func (r *UserRepository) GetAll() []models.User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.users
}

// GetByID retorna um usuário pelo ID
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range r.users {
		if r.users[i].ID == id {
			return &r.users[i], nil
		}
	}
	return nil, ErrUserNotFound
}

// Create cria um novo usuário
func (r *UserRepository) Create(user models.User) models.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user
}

// Update atualiza um usuário existente
func (r *UserRepository) Update(id int, user models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.users {
		if r.users[i].ID == id {
			user.ID = id
			r.users[i] = user
			return &r.users[i], nil
		}
	}
	return nil, ErrUserNotFound
}

// Delete remove um usuário
func (r *UserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.users {
		if r.users[i].ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return ErrUserNotFound
}
