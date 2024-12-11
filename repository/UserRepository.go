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

// GetAllUsers fetches all users from the database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User                         // Declare a slice to hold all users
	if err := r.DB.Find(&users).Error; err != nil { // Use Find to fetch all users
		return nil, err // Return the error if any
	}

	return users, nil // Return the users list
}
