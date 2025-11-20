package services

import (
	"errors"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/repositories"
)

var (
	ErrInvalidProductData = errors.New("dados do produto inválidos")
	ErrInsufficientStock  = errors.New("estoque insuficiente")
)

// ProductService contém a lógica de negócio para produtos
type ProductService struct {
	repo *repositories.ProductRepository
}

// NewProductService cria uma nova instância do serviço de produtos
func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetAll retorna todos os produtos
func (s *ProductService) GetAll() []models.Product {
	return s.repo.GetAll()
}

// GetByID retorna um produto pelo ID
func (s *ProductService) GetByID(id int) (*models.Product, error) {
	if id <= 0 {
		return nil, ErrInvalidProductData
	}
	return s.repo.GetByID(id)
}

// GetByCategory retorna produtos por categoria
func (s *ProductService) GetByCategory(category string) []models.Product {
	return s.repo.GetByCategory(category)
}

// Create cria um novo produto
func (s *ProductService) Create(req models.ProductRequest) (*models.Product, error) {
	if req.Name == "" || req.Price <= 0 {
		return nil, ErrInvalidProductData
	}

	if req.Stock < 0 {
		return nil, ErrInvalidProductData
	}

	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Active:      true,
	}

	if product.Category == "" {
		product.Category = "Geral"
	}

	created := s.repo.Create(product)
	return &created, nil
}

// Update atualiza um produto existente
func (s *ProductService) Update(id int, req models.ProductRequest) (*models.Product, error) {
	if id <= 0 {
		return nil, ErrInvalidProductData
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	product := models.Product{
		ID:          existing.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Active:      existing.Active,
	}

	if product.Name == "" {
		product.Name = existing.Name
	}
	if product.Description == "" {
		product.Description = existing.Description
	}
	if product.Price <= 0 {
		product.Price = existing.Price
	}
	if product.Stock < 0 {
		product.Stock = existing.Stock
	}
	if product.Category == "" {
		product.Category = existing.Category
	}

	return s.repo.Update(id, product)
}

// Delete remove um produto
func (s *ProductService) Delete(id int) error {
	if id <= 0 {
		return ErrInvalidProductData
	}
	return s.repo.Delete(id)
}
