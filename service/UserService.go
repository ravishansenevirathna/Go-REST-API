package service

import (
	"rest-api/dto"
	"rest-api/repository"
)

// UserService handles business logic for users.
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// CreateUser validates and saves a user, then returns the created user DTO.
func (s *UserService) CreateUser(userDto *dto.User) (*dto.User, error) {
	// Call the repository Save method to save the user in the database
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

// GetAllUsers gets all users and returns them as DTOs
func (s *UserService) GetAllUsers() ([]dto.User, error) {
	users, err := s.Repo.GetAllUsers() // Get all users from the repository
	if err != nil {
		return nil, err // Return error if something goes wrong
	}

	// Convert the models to DTOs
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
