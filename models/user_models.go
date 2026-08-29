package models

import (
	"fmt"
	"gin-and-tonic/config"
	"gin-and-tonic/database"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRole represents custom roles for users
type UserRole string

const (
	Admin  UserRole = "admin"
	Staff  UserRole = "staff"
	Normal UserRole = "normal"
)

// User defines database schema with constraints
type User struct {
	gorm.Model
	Username string   `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Email    string   `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string   `gorm:"type:varchar(255);not null" json:"-"`
	Role     UserRole `gorm:"type:varchar(30); default:'normal'; not null" json:"role"`
}

// Implementing BeforeSave (from GORM).
// It automatically runs before inserting or updating a record.
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	// Validate the user role
	if u.Role != Admin && u.Role != Staff && u.Role != Normal {
		u.Role = Normal // Set the role to default
	}

	// Hash the password with bcrypt if it was changed or is new
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), config.BcryptCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword checks if the provided password matches the hashed database password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// createUser create a User.
// args: name, email, password and role as arguments
// returns: bool & error
func createUser(username string, email string, password string, role UserRole) (bool, error) {
	database.ConnectDatabase()

	newUser := User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	fmt.Println("\n Saving user to database...")
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		fmt.Printf("Database Error: %v\n", result.Error)
		return false, result.Error
	}
	return true, nil
}

// CreateAdminUser create a Admin User.
// args: name, email and password as arguments
// returns: bool & error
func CreateAdminUser(username string, email string, password string) (bool, error) {
	created, err := createUser(username, email, password, Admin)
	if err != nil || !created {
		log.Printf("Error creating Admin User.")
		return false, err
	}
	return true, nil
}

// CreateStaffUser create a Admin User.
// args: name, email and password as arguments
// returns: bool & error
func CreateStaffUser(username string, email string, password string) (bool, error) {
	created, err := createUser(username, email, password, Staff)
	if err != nil || !created {
		log.Printf("Error creating Staff User.")
		return false, err
	}
	return true, nil
}

// CreateAdminUser create a Admin User.
// args: name, email and password as arguments
// returns: bool & error
func CreateNormalUser(username string, email string, password string) (bool, error) {
	created, err := createUser(username, email, password, Normal)
	if err != nil || !created {
		log.Printf("Error creating User.")
		return false, err
	}
	return true, nil
}
