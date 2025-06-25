package handler

import (
	"database/sql"
	"encoding/json"
	"go-auth-api/internal/user/dto"
	"go-auth-api/internal/user/repository"
	"go-auth-api/internal/user/service"
	"go-auth-api/pkg/config"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(repo)
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	err := h.authService.Signup(&req)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário criado com sucesso"))
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
	}

	token, err := h.authService.Login(&req, config.LoadConfig())

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // ⚠️ usar true em produção (https)
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400, // 1 dia
		Domain:   config.LoadConfig().CookieDomain,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Autenticado com sucesso"}`))
}
