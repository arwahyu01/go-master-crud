package helpers

import (
	"context"
	"errors"
	"net/http"
)

type ContextKey string

const UserContextKey ContextKey = "user"

// Simpan User ID ke context request
func SetUserToContext(r *http.Request, userID string) *http.Request {
	ctx := context.WithValue(r.Context(), UserContextKey, userID)
	return r.WithContext(ctx)
}

// Ambil User ID dari context
func GetUserID(r *http.Request) (string, error) {
	userID, ok := r.Context().Value(UserContextKey).(string)
	if !ok || userID == "" {
		return "", errors.New("user ID not found in context")
	}
	return userID, nil
}
