package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-auth-api/internal/user/dto"
	"go-auth-api/internal/user/model"
	"go-auth-api/internal/user/repository"
	"go-auth-api/pkg/config"
	"go-auth-api/pkg/utils"
	"time"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Signup(req *dto.SignupRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	return s.repo.Create(user)
}

func (s *AuthService) Login(req *dto.LoginRequest, cfg *config.Config) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return "", errors.New("senha incorreta")
	}

	// Criar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return tokenString, nil
}
