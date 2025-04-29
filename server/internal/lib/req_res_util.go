package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type HandlerResponse struct {
	Message string
}

// WriteJSONResponse writes a response to the client
func WriteJSONResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if v != nil {
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Failed to encode JSON response: %v", err)
			errResp := []byte(`{"message": "Failed to Encode when preparing Response Object"}`)
			w.Write(errResp)
		}
	}

}

// DecodeAndValidateRequest takes in a request object (r) and a struct (v)
// and decodes the request body into the struct interface and then validates
// the fields in the struct
func DecodeAndValidateRequest(r *http.Request, v interface{}) error {
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		errMsg := fmt.Errorf("%s: %w", ERR_DECODE_REQ_FAILURE, err)
		log.Println(errMsg.Error())
		return fmt.Errorf("%s: %w", ERR_DECODE_REQ_FAILURE, err)
	}

	if err := validate.Struct(v); err != nil {
		errMsg := fmt.Errorf("%s: %w", ERR_VALIDATE_REQ_FAILURE, err)
		log.Println(errMsg.Error())
		return fmt.Errorf("%s: %w", ERR_VALIDATE_REQ_FAILURE, err)
	}

	return nil
}

func SanitizeInput(data interface{}) {
	v := reflect.ValueOf(data)

	// Ensure the input is a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return // Do nothing if it's not a pointer to a struct
	}

	// Get the underlying struct
	v = v.Elem()

	// Iterate through the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Only process fields that are settable (exported fields)
		if !field.CanSet() {
			continue
		}

		// Handle string fields
		if field.Kind() == reflect.String {
			trimmed := strings.TrimSpace(field.String())
			trimmed = strings.ReplaceAll(trimmed, " ", "")
			field.SetString(trimmed)
		}
	}
}
