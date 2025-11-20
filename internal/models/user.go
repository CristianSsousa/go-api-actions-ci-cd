package models

// User representa um usuário no sistema
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Active   bool   `json:"active"`
	CreateAt string `json:"created_at"`
}

// UserRequest representa a requisição para criar/atualizar um usuário
type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

