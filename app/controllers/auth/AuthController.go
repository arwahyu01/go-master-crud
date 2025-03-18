package auth

import (
	"encoding/json"
	"github.com/arwahyu01/go-jwt/database"
	"github.com/arwahyu01/go-jwt/helpers"
	"net/http"
	"os"
	"time"

	"github.com/arwahyu01/go-jwt/app/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		return
	}

	var user models.Users
	if err := database.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		helpers.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"message": "Username or password is incorrect"})
		return
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		helpers.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"message": "Username or password is incorrect"})
		return
	}

	// Buat token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":       user.ID.String(),
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Error generating token"})
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Login successful",
		"user": map[string]string{
			"id":       user.ID.String(),
			"username": user.Username,
		},
		"meta": map[string]interface{}{
			"token":   tokenString,
			"expired": expirationTime.Format(time.RFC3339),
		},
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Nama     string `json:"nama" validate:"required"`
		Username string `json:"username" validate:"required,min=3,max=20"`
		Password string `json:"password" validate:"required,min=6"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		return
	}

	errors := helpers.ValidateStruct(request)
	if len(errors) > 0 {
		helpers.ResponseJSON(w, http.StatusBadRequest, errors)
		return
	}

	// Cek username
	var existingUser models.Users
	if err := database.DB.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Username already taken"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
		return
	}

	// Simpan user
	user := models.Users{
		Nama:     request.Nama,
		Username: request.Username,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to register user"})
		return
	}

	helpers.ResponseJSON(w, http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"status":  "success",
		"message": "Register successful",
		"user": map[string]interface{}{
			"id":       user.ID.String(),
			"nama":     user.Nama,
			"username": user.Username,
		},
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Tidak ada cara untuk "mematikan" token JWT di backend secara langsung
	// Biasanya token dihapus dari client-side atau menggunakan blocklist di server (opsional)

	helpers.ResponseJSON(w, http.StatusOK, map[string]string{"message": "Logout successful. Please delete your token from client."})
}
