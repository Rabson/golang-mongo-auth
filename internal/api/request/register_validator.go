package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type RegisterValidator struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	// Password string `json:"password" binding:"required,min=8,max=20,eqfield=ConfirmPassword"`
	// ConfirmPassword string `json:"confirm_password" binding:"required,min=8,max=20"`
	// ConfirmPassword string `json:"confirm_password" binding:"required,min=8,max=20,eqfield=Password"`
	Name string `json:"name" binding:"required,min=4,max=20"`
}

func (v RegisterValidator) Validate(data map[string]interface{}) error {
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

	if name, ok := data["name"]; ok {
		if nameStr, ok := name.(string); ok {
			v.Name = nameStr
		}
	}

	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required.Error("Email is required")),
		validation.Field(&v.Password, validation.Required.Error("Password is required")),
		validation.Field(&v.Name, validation.Required.Error("Name is required")),
	)
}

func (v RegisterValidator) GetKeys() []string {
	return []string{"email", "password", "name"}
}
