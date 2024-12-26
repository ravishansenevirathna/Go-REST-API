package service

import (
	"rest-api/dto"
	"rest-api/repository"
	"rest-api/utils"
)

// UserService handles business logic for users.
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(userDto *dto.User) (*dto.User, error) {
	user, err := s.Repo.Save(userDto)
	if err != nil {
		return nil, err // Return error if save fails
	}

	// Map the saved user model back to a DTO
	createdUserDto := &dto.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	// Return the created user DTO
	return createdUserDto, nil
}

func (s *UserService) LoginUser(userDto *dto.User) (*dto.User, string, error) {
	// Authenticate the user
	user, err := s.Repo.Login(userDto)
	if err != nil {
		return nil, "", err
	}

	createdUserDto := &dto.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, "", err
	}

	return createdUserDto, token, nil
}

// GetAllUsers gets all users and returns them as DTOs
func (s *UserService) GetAllUsers() ([]dto.User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userDTOs []dto.User
	for _, user := range users {
		userDTO := dto.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs, nil
}
