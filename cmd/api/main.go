package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/handlers"
	customMiddleware "github.com/CristianSsousa/go-api-actions-ci-cd/internal/middleware"
	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/repositories"
	"github.com/CristianSsousa/go-api-actions-ci-cd/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	chiCors "github.com/go-chi/cors"
	chiRender "github.com/go-chi/render"
)

func main() {
	// Inicializa repositórios
	userRepo := repositories.NewUserRepository()
	productRepo := repositories.NewProductRepository()

	// Inicializa serviços
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	// Inicializa handlers
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	healthHandler := handlers.NewHealthHandler()

	// Configura router
	r := chi.NewRouter()

	// Middlewares globais
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.Logger)
	r.Use(chiRender.SetContentType(chiRender.ContentTypeJSON))

	// CORS
	r.Use(chiCors.Handler(chiCors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Rotas de health check
	r.Get("/health", healthHandler.Check)
	r.Get("/", healthHandler.Check)

	// Rotas de usuários
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)
		r.Get("/{id}", userHandler.GetByID)
		r.Post("/", userHandler.Create)
		r.Put("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)
	})

	// Rotas de produtos
	r.Route("/api/products", func(r chi.Router) {
		r.Get("/", productHandler.GetAll)
		r.Get("/{id}", productHandler.GetByID)
		r.Get("/category/{category}", productHandler.GetByCategory)
		r.Post("/", productHandler.Create)
		r.Put("/{id}", productHandler.Update)
		r.Delete("/{id}", productHandler.Delete)
	})

	// Obtém a porta do ambiente ou usa 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s", port)
	log.Printf("Health check: http://localhost:%s/health", port)
	log.Printf("API de usuários: http://localhost:%s/api/users", port)
	log.Printf("API de produtos: http://localhost:%s/api/products", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

