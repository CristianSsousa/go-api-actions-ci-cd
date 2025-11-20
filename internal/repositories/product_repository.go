package repositories

import (
	"errors"
	"sync"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
)

var (
	ErrProductNotFound = errors.New("produto não encontrado")
)

// ProductRepository gerencia os dados de produtos em memória
type ProductRepository struct {
	mu       sync.RWMutex
	products []models.Product
	nextID   int
}

// NewProductRepository cria uma nova instância do repositório com dados pré-prontos
func NewProductRepository() *ProductRepository {
	repo := &ProductRepository{
		products: []models.Product{
			{
				ID:          1,
				Name:        "Notebook Dell XPS 15",
				Description: "Notebook de alta performance com processador Intel i7",
				Price:       8999.99,
				Stock:       15,
				Category:    "Eletrônicos",
				Active:      true,
			},
			{
				ID:          2,
				Name:        "Mouse Logitech MX Master 3",
				Description: "Mouse sem fio ergonômico para produtividade",
				Price:       599.90,
				Stock:       50,
				Category:    "Periféricos",
				Active:      true,
			},
			{
				ID:          3,
				Name:        "Teclado Mecânico Keychron K8",
				Description: "Teclado mecânico sem fio com switches Gateron",
				Price:       799.00,
				Stock:       30,
				Category:    "Periféricos",
				Active:      true,
			},
			{
				ID:          4,
				Name:        "Monitor LG UltraWide 34",
				Description: "Monitor ultrawide 34 polegadas 4K",
				Price:       3499.99,
				Stock:       8,
				Category:    "Monitores",
				Active:      true,
			},
			{
				ID:          5,
				Name:        "Webcam Logitech C920",
				Description: "Webcam Full HD para videoconferências",
				Price:       499.90,
				Stock:       0,
				Category:    "Periféricos",
				Active:      false,
			},
		},
		nextID: 6,
	}
	return repo
}

// GetAll retorna todos os produtos
func (r *ProductRepository) GetAll() []models.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.products
}

// GetByID retorna um produto pelo ID
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range r.products {
		if r.products[i].ID == id {
			return &r.products[i], nil
		}
	}
	return nil, ErrProductNotFound
}

// GetByCategory retorna produtos por categoria
func (r *ProductRepository) GetByCategory(category string) []models.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []models.Product
	for i := range r.products {
		if r.products[i].Category == category {
			filtered = append(filtered, r.products[i])
		}
	}
	return filtered
}

// Create cria um novo produto
func (r *ProductRepository) Create(product models.Product) models.Product {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	r.nextID++
	r.products = append(r.products, product)
	return product
}

// Update atualiza um produto existente
func (r *ProductRepository) Update(id int, product models.Product) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.products {
		if r.products[i].ID == id {
			product.ID = id
			r.products[i] = product
			return &r.products[i], nil
		}
	}
	return nil, ErrProductNotFound
}

// Delete remove um produto
func (r *ProductRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.products {
		if r.products[i].ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return ErrProductNotFound
}
