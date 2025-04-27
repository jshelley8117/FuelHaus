package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type HandlerResponse struct {
	Message string
}

// WriteJSONResponse writes a response to the client
func WriteJSONResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		errResp := []byte(`{"message": "Failed to Encode when preparing Response Object"}`)
		w.Write(errResp)
	}
}

// DecodeAndValidateRequest takes in a request object (r) and a struct (v)
// and decodes the request body into the struct interface and then validates
// the fields in the struct
func DecodeAndValidateRequest(r *http.Request, v interface{}) error {
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("%s: %w", ERR_DECODE_REQ_FAILURE, err)
	}

	if err := validate.Struct(v); err != nil {
		return fmt.Errorf("%s: %w", ERR_VALIDATE_REQ_FAILURE, err)
	}

	return nil
}
