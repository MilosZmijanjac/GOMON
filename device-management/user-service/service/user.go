package service

import (
	"context"
	"log"
	"os"
	"time"
	"user-service/models"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go-micro.dev/v5"
	"go-micro.dev/v5/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{ db *gorm.DB }

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint8  `json:"role"`
	IsActive bool   `json:"isActive"`
}

type RegisterResponse struct{}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type ListRequest struct{}

type ListResponse struct {
	Users []models.User `json:"users"`
}
type UpdateRequest struct {
	Username string `json:"username"`
	Role     uint8  `json:"role"`
	IsActive bool   `json:"isActive"`
}

type Updateesponse struct{}

func (s *UserService) Register(ctx context.Context, req *RegisterRequest, rsp *RegisterResponse) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.InternalServerError("UserService.Register", "Error hashing password: %v", err)
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
		IsActive: req.IsActive,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return errors.InternalServerError("UserService.Register", "Error saving user: %v", err)
	}

	return nil
}

func (s *UserService) List(ctx context.Context, req *ListRequest, rsp *ListResponse) error {
	s.db.Find(&rsp.Users)
	return nil
}
func (s *UserService) Update(ctx context.Context, req *UpdateRequest, rsp *UpdateRequest) error {

	s.db.Model(&models.User{Username: req.Username}).Select("role", "is_active").Updates(models.User{IsActive: req.IsActive, Role: req.Role})
	return nil
}
func (s *UserService) Login(ctx context.Context, req *LoginRequest, rsp *LoginResponse) error {
	var user models.User

	if err := s.db.Where("is_active=true and username = ?", req.Username).First(&user).Error; err != nil {
		return errors.NotFound("UserService.Login", "User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.Unauthorized("UserService.Login", "Invalid password")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return errors.InternalServerError("UserService.Login", "Error generating token: %v", err)
	}

	rsp.Token = tokenString
	return nil
}

func RegisterUserServiceHandler(service *micro.Service, db *gorm.DB) {
	micro.RegisterHandler((*service).Server(), &UserService{db: db})
}
