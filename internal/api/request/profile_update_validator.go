package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateProfileValidator struct {
	Name    string `json:"name"`
	Profile string `form:"profile" binding:"required"`
}

func (v UpdateProfileValidator) Validate(data map[string]interface{}) error {
	if name, ok := data["name"]; ok {
		if nameStr, ok := name.(string); ok {
			v.Name = nameStr
		}
	}

	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Required.Error("Name is required")),
	)
}
