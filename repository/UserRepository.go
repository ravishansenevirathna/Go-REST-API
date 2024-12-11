package repository

import (
	"gorm.io/gorm"
	"rest-api/dto"
	"rest-api/models"
)

// UserRepository handles database operations for users.
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Save persists a user in the database.
func (r *UserRepository) Save(userDto *dto.User) (*models.User, error) {
	// Convert userDto to models.User
	user := models.User{
		Name:  userDto.Name,
		Email: userDto.Email,
	}

	// Save the user in the database
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err // Return the error if save fails
	}

	// Return the created user object
	return &user, nil
}

func (r *UserRepository) GetAllUsers() (*models.User, error) {
	// Convert userDto to models.User
	user := models.User{
		Name:  userDto.Name,
		Email: userDto.Email,
	}

	if err := r.DB.FindAllUsers(&user).Error; err != nil {
		return nil, err // Return the error if save fails
	}

	// Return the created user object
	return &user, nil
}
