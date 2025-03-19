package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Get(r *http.Request, key string, defaultValue int) int {
	valueStr := r.URL.Query().Get(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// ParseJSONToMap mengubah request JSON menjadi map[string]interface{}
func ParseJSONToMap(request *http.Request) (map[string]interface{}, error) {
	// Simpan isi body sebelum parsing
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	// Reset request.Body agar bisa dibaca ulang jika perlu
	request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Decode JSON ke map
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
