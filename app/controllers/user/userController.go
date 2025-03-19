package user

import (
	"github.com/arwahyu01/go-jwt/app/models/user"
	"github.com/arwahyu01/go-jwt/database"
	"github.com/arwahyu01/go-jwt/helpers/auth"
	"github.com/arwahyu01/go-jwt/helpers/request"
	"github.com/arwahyu01/go-jwt/helpers/response"
	"github.com/arwahyu01/go-jwt/helpers/validation"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	page := request.Get(req, "page", 1)
	pageSize := request.Get(req, "page_size", 10)

	pagination, err := user.FetchAll(database.DB, page, pageSize)
	if err != nil {
		response.Json(w, http.StatusInternalServerError, map[string]string{"message": "Error fetching users"})
		return
	}

	response.Json(w, http.StatusOK, pagination)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		response.Json(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
		return
	}

	var userData user.Users
	if err := database.DB.First(&userData, "id = ?", userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	response.Json(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "User profile",
		"user":    userData,
	})
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	var requestData map[string]interface{}
	requestData, _ = request.ParseJSONToMap(req)

	rules := map[string]string{
		"first_name": "required",
		"last_name":  "required",
		"email":      "omitempty,email,min=3,max=20",
		"password":   "omitempty,min=6",
	}

	errors, valid := validation.Request(requestData, rules)
	if !valid {
		response.Json(w, http.StatusBadRequest, errors)
		return
	}

	userID := mux.Vars(req)["id"]
	var usr user.Users
	if err := database.DB.First(&usr, "id = ?", userID).Error; err != nil {
		response.Json(w, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	if password, ok := requestData["password"].(string); ok && password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		requestData["password"] = string(hashedPassword)
	}

	if err := database.DB.Model(&usr).Updates(requestData).Error; err != nil {
		response.Json(w, http.StatusInternalServerError, map[string]string{"message": "Error updating user"})
		return
	}

	response.Json(w, http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user": map[string]interface{}{
			"id":    usr.ID.String(),
			"nama":  usr.FirstName + " " + usr.LastName,
			"email": usr.Email,
		},
	})
}
