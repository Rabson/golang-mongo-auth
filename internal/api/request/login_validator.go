package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type LoginValidator struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (v LoginValidator) Validate(data map[string]interface{}) error {
	if email, ok := data["email"]; ok {
		if emailStr, ok := email.(string); ok {
			v.Email = emailStr
		}
	}
	if password, ok := data["password"]; ok {
		if passwordStr, ok := password.(string); ok {
			v.Password = passwordStr
		}
	}
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required.Error("Email is required")),
		validation.Field(&v.Password, validation.Required.Error("Password is required")),
	)
}
