package handlers

import (
	"net/http"
	"strconv"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// ProductHandler gerencia as requisições HTTP relacionadas a produtos
type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler cria uma nova instância do handler de produtos
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetAll retorna todos os produtos
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAll()
	render.JSON(w, r, models.Response{
		Success: true,
		Data:    products,
	})
}

// GetByID retorna um produto pelo ID
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "ID inválido",
		})
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	render.JSON(w, r, models.Response{
		Success: true,
		Data:    product,
	})
}

// GetByCategory retorna produtos por categoria
func (h *ProductHandler) GetByCategory(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	if category == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "Categoria inválida",
		})
		return
	}

	products := h.service.GetByCategory(category)
	render.JSON(w, r, models.Response{
		Success: true,
		Data:    products,
	})
}

// Create cria um novo produto
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.ProductRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "Dados inválidos",
		})
		return
	}

	product, err := h.service.Create(req)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, models.Response{
		Success: true,
		Message: "Produto criado com sucesso",
		Data:    product,
	})
}

// Update atualiza um produto existente
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "ID inválido",
		})
		return
	}

	var req models.ProductRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "Dados inválidos",
		})
		return
	}

	product, err := h.service.Update(id, req)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	render.JSON(w, r, models.Response{
		Success: true,
		Message: "Produto atualizado com sucesso",
		Data:    product,
	})
}

// Delete remove um produto
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "ID inválido",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	render.JSON(w, r, models.Response{
		Success: true,
		Message: "Produto removido com sucesso",
	})
}

