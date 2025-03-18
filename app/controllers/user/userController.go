package user

import (
	"github.com/arwahyu01/go-jwt/app/models"
	"github.com/arwahyu01/go-jwt/database"
	"github.com/arwahyu01/go-jwt/helpers"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	page := helpers.GetParam(r, "page", 1)
	pageSize := helpers.GetParam(r, "page_size", 10)

	pagination, err := models.GetAllUser(database.DB, page, pageSize)
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Error fetching users"})
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, pagination)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var user models.Users
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "User profile",
		"user": map[string]interface{}{
			"id":       user.ID.String(),
			"username": user.Username,
			"nama":     user.Nama,
		},
	})
}
