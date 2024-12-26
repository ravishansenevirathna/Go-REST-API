package repository

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"rest-api/dto"
	"rest-api/models"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Save persists a user in the database.
func (r *UserRepository) Save(userDto *dto.User) (*models.User, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	userDto.Password = string(hashedPassword)

	user := models.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err // Return the error if save fails
	}

	return &user, nil
}

func (r *UserRepository) Login(userDto *dto.User) (*models.User, error) {
	var user models.User

	// Fetch the user from the database by email
	if err := r.DB.Where("email = ?", userDto.Email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Return the user details if the credentials are valid
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User                         // Declare a slice to hold all users
	if err := r.DB.Find(&users).Error; err != nil { // Use Find to fetch all users
		return nil, err
	}

	return users, nil
}
