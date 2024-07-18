package user

import "fmt"

// Define the error messages for required parameters.
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateContentRequest represents the request for creating content.
type CreateUserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`

}

// Validate checks the request fields and returns an error if any required field is missing.
func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}
