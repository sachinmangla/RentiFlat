package rentiflat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sachinmangla/rentiflat/database"
	"golang.org/x/crypto/bcrypt"
)

type response struct {
	Message string `json:"message"`
}

func HashPassword(password string) (string, error) {
	// Generate a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func checkUserAlreadyExist(email string) (bool, error) {
	var owner database.OwnerDetails

	// Perform the database query
	result := database.GetDb().Where("email = ?", email).First(&owner)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // User not found
		}
		return false, result.Error // Some other error occurred
	}

	return true, nil // User found
}

// @Summary Create a new owner
// @Description Register a new owner in the system
// @Tags owners
// @Accept json
// @Produce json
// @Param owner body database.OwnerDetails true "Owner details"
// @Success 201 {object} response "User created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 409 {string} string "Conflict"
// @Failure 500 {string} string "Internal server error"
// @Router /owners [post]
func OwnerDetailCreatePost(w http.ResponseWriter, r *http.Request) {
	var owner database.OwnerDetails

	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if hashedPass, err := HashPassword(owner.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		owner.Password = hashedPass
	}

	userExist, err := checkUserAlreadyExist(owner.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userExist {
		http.Error(w, "User already exist", http.StatusInternalServerError)
		return
	}

	result := database.GetDb().Create(&owner).Commit()
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response{Message: "User created successfully"})
}

// @Summary Login
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginDetail body database.LoginDetail true "Login credentials"
// @Success 202 {object} map[string]string "JWT token"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginDetail database.LoginDetail
	var owner database.OwnerDetails

	err := json.NewDecoder(r.Body).Decode(&loginDetail)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.GetDb()
	result := db.Where("EMAIL = ?", loginDetail.Email).First(&owner)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusBadRequest)
		}
		return
	}

	if VerifyPassword(owner.Password, loginDetail.Password) != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	if token, err := CreateJwtToken(int(owner.ID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintf(w, `{"token":"%s"}`, token)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
