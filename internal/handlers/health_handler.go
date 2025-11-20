package handlers

import (
	"net/http"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/models"
	"github.com/go-chi/render"
)

// HealthHandler gerencia as requisições de health check
type HealthHandler struct{}

// NewHealthHandler cria uma nova instância do handler de health
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check retorna o status da API
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, models.Response{
		Success: true,
		Message: "API está funcionando corretamente",
		Data: map[string]string{
			"status":  "healthy",
			"service": "go-api-actions-ci-cd",
		},
	})
}
