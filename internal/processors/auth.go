package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/repository"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

const (
	salt       = "afdafadfadfadf"
	signingKey = "qrkjk#4#35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	User_UUID uuid.UUID `json:"user_UUID"`
}

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetToken(username, password string) (string, error)
}

type AuthService struct {
	repo  repository.Authorization
	cache repository.Cache
}

// Функция NewAuthService является конструктором структуры AuthService. Принимает на вход переменную типа repository.Authorization и возвращает AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo,
		cache: *repository.NewCache()}
}

/*
	Функция CreateUser проверяет на корректность полученные от пользователя данные и вызывает функцию repo.CreateUser для создания пользователя. Принимает на вход структуру User,

возвращает id типа int и ошибку типа error
*/
func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.User_UUID = uuid.NewV4()
	if len(user.Password) == 0 || len(user.Username) == 0 {
		return 0, fmt.Errorf("имя пользователя или пароль не могут быть пустыми")
	}
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("GenerateToken: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.User_UUID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetToken(username, password string) (string, error) {
	token, err := s.GenerateToken(username, password)
	if err != nil {
		return "", fmt.Errorf("GetToken 1: %v", err)
	}

	password = generatePasswordHash(password)

	value, err := json.Marshal(&entity.UserSessions{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return "", fmt.Errorf("GetToken 2: %v", err)
	}

	expiration := time.Hour * 12

	err = s.cache.SetCache(token, value, int32(expiration))
	if err != nil {
		return "", fmt.Errorf("GetToken 2: %v", err)
	}
	return token, nil
}

/*
Функция generatePasswordHash создаёт хеш пороля. Принимает на вход переменную password типа string возвращает хешированный пароль типа string
*/
func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
