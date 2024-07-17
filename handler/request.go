package handler

import (
	"fmt"
	"time"
)

// Define the error messages for required parameters.
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateContentRequest represents the request for creating content.
type CreateContentRequest struct {
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	ContentType string    `json:"contentType"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
}

// Validate checks the request fields and returns an error if any required field is missing.
func (r *CreateContentRequest) Validate() error {
	if r.Title == "" && r.ContentType == "" && r.Desc == "" && r.Genre == "" && r.ReleaseDate.IsZero() {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Desc == "" {
		return errParamIsRequired("desc", "string")
	}
	if r.ContentType == "" {
		return errParamIsRequired("contentType", "string")	
	}
	if r.Genre == "" {
		return errParamIsRequired("genre", "string")
	}
	if r.ReleaseDate.IsZero() {
		return errParamIsRequired("releaseDate", "Time")
	}

	return nil
}
