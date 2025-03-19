package auth

import (
	"encoding/json"
	"github.com/arwahyu01/go-jwt/database"
	"github.com/arwahyu01/go-jwt/helpers/response"
	"github.com/arwahyu01/go-jwt/helpers/validation"
	"net/http"
	"os"
	"time"

	"github.com/arwahyu01/go-jwt/app/models/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"message": "Invalid req"})
		return
	}

	var userData user.Users
	if err := database.DB.Where("email = ?", req.Email).First(&userData).Error; err != nil {
		response.Json(w, http.StatusUnauthorized, map[string]string{"message": "Username or password is incorrect"})
		return
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password)); err != nil {
		response.Json(w, http.StatusUnauthorized, map[string]string{"message": "password is incorrect"})
		return
	}

	// Buat token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":    userData.ID.String(),
		"email": userData.Email,
		"exp":   expirationTime.Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response.Json(w, http.StatusInternalServerError, map[string]string{"message": "Error generating token"})
		return
	}

	response.Json(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Login successful",
		"userData": map[string]string{
			"id":    userData.ID.String(),
			"email": userData.Email,
		},
		"meta": map[string]interface{}{
			"token":   tokenString,
			"expired": expirationTime.Format(time.RFC3339),
		},
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nama     string `json:"nama" validate:"required"`
		Email    string `json:"email" validate:"required,min=3,max=20"`
		Password string `json:"password" validate:"required,min=6"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		return
	}

	errors, valid := validation.Request(req, nil)
	if !valid {
		response.Json(w, http.StatusBadRequest, errors)
		return
	}

	var existingUser user.Users
	if err := database.DB.Where("username = ?", req.Email).First(&existingUser).Error; err == nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"message": "Username already taken"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Json(w, http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
		return
	}

	userData := user.Users{
		FirstName: req.Nama,
		Email:     req.Email,
		Password:  string(hashedPassword),
	}

	if err := database.DB.Create(&userData).Error; err != nil {
		response.Json(w, http.StatusInternalServerError, map[string]string{"message": "Failed to register user"})
		return
	}

	response.Json(w, http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"status":  "success",
		"message": "Register successful",
		"userData": map[string]interface{}{
			"id":       userData.ID.String(),
			"nama":     userData.FirstName,
			"username": userData.Email,
		},
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Tidak ada cara untuk "mematikan" token JWT di backend secara langsung
	// Biasanya token dihapus dari client-side atau menggunakan blocklist di server (opsional)

	response.Json(w, http.StatusOK, map[string]string{"message": "Logout successful. Please delete your token from client."})
}
