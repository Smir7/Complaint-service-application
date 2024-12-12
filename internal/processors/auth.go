package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/repository"
	"crypto/sha1"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const salt = "afdafadfadfadf"

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	user.User_UUID = uuid.NewV4()
	if len(user.Password) == 0 || len(user.Username) == 0 {
		return 0, fmt.Errorf("Username or password can't be empty")
	}
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
