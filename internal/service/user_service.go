package service

import (
	"errors"
	"go-workflow-rnd/internal/models"
	"go-workflow-rnd/internal/repository"

	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	return s.userRepo.Create(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) UpdateUser(user *models.User) error {
	if user.ID == 0 {
		return errors.New("user ID is required")
	}
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("user ID is required")
	}
	return s.userRepo.Delete(id)
}

func NewUserServiceDI(injector do.Injector) (UserService, error) {
	userRepo := do.MustInvoke[repository.UserRepository](injector)
	return NewUserService(userRepo), nil
}