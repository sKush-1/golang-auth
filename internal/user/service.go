package user

import (
	"context"
	"fmt"
	"strings"
	"time"

	"golang-auth/internal/auth"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repo
	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.TrimSpace(input.Email)
	password := strings.TrimSpace(input.Password)

	if email == "" || password == "" {
		return AuthResult{}, fmt.Errorf("email and password are required")
	}

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, fmt.Errorf("user with email %s already exists", email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to hash password: %w", err)
	}

	u := User{
		Email:     email,
		Password:  string(hashedPassword),
		Role:      "user",
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	user, err := s.repo.Create(ctx, u)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := auth.CreateToken(s.jwtSecret, user.ID.Hex(), user.Role)
	if err != nil {
		return AuthResult{}, fmt.Errorf("failed to create token: %w", err)
	}

	return AuthResult{
		Token: token,
		User:  *ToPublicUser(&user),
	}, nil
}
