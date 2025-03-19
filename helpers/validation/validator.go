package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Request(data interface{}, rules map[string]string) (map[string]string, bool) {
	errors := make(map[string]string)

	// 🔹 Jika data adalah struct, gunakan validator.v10
	if reflect.TypeOf(data).Kind() == reflect.Struct {
		err := validate.Struct(data)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				field := strings.ToLower(err.Field())
				errors[field] = fmt.Sprintf("Field %s %s", field, err.Tag())
			}
		}
	} else if reflect.TypeOf(data).Kind() == reflect.Map {
		// 🔹 Jika data adalah map, gunakan aturan manual
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			errors["error"] = "Invalid data format"
			return errors, false
		}

		for field, rule := range rules {
			value, exists := dataMap[field]

			// 🔹 Validasi "required"
			if strings.Contains(rule, "required") && (!exists || fmt.Sprintf("%v", value) == "") {
				errors[field] = fmt.Sprintf("Field %s is required", field)
				continue
			}

			// 🔹 Validasi panjang string (min & max)
			if exists {
				strValue := fmt.Sprintf("%v", value)

				if strings.Contains(rule, "min=") {
					minLength := extractRuleValue(rule, "min=")
					if len(strValue) < minLength {
						errors[field] = fmt.Sprintf("Field %s must be at least %d characters", field, minLength)
					}
				}

				if strings.Contains(rule, "max=") {
					maxLength := extractRuleValue(rule, "max=")
					if len(strValue) > maxLength {
						errors[field] = fmt.Sprintf("Field %s must not exceed %d characters", field, maxLength)
					}
				}
			}
		}
	}

	return errors, len(errors) == 0
}

func extractRuleValue(rule string, key string) int {
	rules := strings.Split(rule, ",")
	for _, r := range rules {
		if strings.HasPrefix(r, key) {
			var value int
			fmt.Sscanf(r, key+"%d", &value)
			return value
		}
	}
	return 0
}
