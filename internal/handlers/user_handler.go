package handlers

import (
	"net/http"
	"strconv"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// UserHandler gerencia as requisições HTTP relacionadas a usuários
type UserHandler struct {
	service *services.UserService
}

// NewUserHandler cria uma nova instância do handler de usuários
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll retorna todos os usuários
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAll()
	render.JSON(w, r, models.Response{
		Success: true,
		Data:    users,
	})
}

// GetByID retorna um usuário pelo ID
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.service.GetByID(id)
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
		Data:    user,
	})
}

// Create cria um novo usuário
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "Dados inválidos",
		})
		return
	}

	user, err := h.service.Create(req)
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
		Message: "Usuário criado com sucesso",
		Data:    user,
	})
}

// Update atualiza um usuário existente
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var req models.UserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, models.Response{
			Success: false,
			Error:   "Dados inválidos",
		})
		return
	}

	user, err := h.service.Update(id, req)
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
		Message: "Usuário atualizado com sucesso",
		Data:    user,
	})
}

// Delete remove um usuário
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
		Message: "Usuário removido com sucesso",
	})
}
