package response

import (
	"encoding/json"
	"net/http"
)

// Json mengirimkan response JSON ke client
func Json(writer http.ResponseWriter, status int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return
	}
}
