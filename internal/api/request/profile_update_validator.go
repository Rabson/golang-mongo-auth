package request

import (
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// type HttpFile struct {
// 	Profile   string `json:"profile"`
// 	Filename  string `json:"filename"`
// 	Header    string `json:"header"`
// 	Size      int64  `json:"size"`
// 	Content   string `json:"content"`
// 	Tmpfile   string `json:"tmpfile"`
// 	TmpOff    int64  `json:"tmpoffset"`
// 	TmpName   string `json:"tmpname"`
// 	TmpShared bool   `json:"tmpshared"`
// }

type UpdateProfileValidator struct {
	Name    string               `json:"name"`
	Profile multipart.FileHeader `json:"profile"`
}

func (v UpdateProfileValidator) Validate(data map[string]interface{}) error {
	if name, ok := data["name"]; ok {
		if nameStr, ok := name.(string); ok {
			v.Name = nameStr
		}
	}

	// if profile, ok := data["profile"]; ok {
	// 	if profileStr, ok := profile.(string); ok {
	// 		v.Profile = profileStr
	// 	}
	// }

	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Required.Error("Name is required")),
		// validation.Field(&v.Profile, validation.Required.Error("Profile is required")),
		// validation.Field(&v.Profile, validation.Length(0, 255).Error("Profile URL is too long")),
		// validation.Field(&v.Profile, validation.Match(validation.NewStringPattern("^(http|https)://.*")).Error("Profile URL must be a valid URL")),
		// validation.Field(&v.Profile, validation.Match(validation.NewStringPattern(".*\\.(jpg|jpeg|png|gif)$")).Error("Profile URL must be an image (jpg, jpeg, png, gif)")),
	)
}

func (v UpdateProfileValidator) GetKeys() []string {
	return []string{"name", "profile"}
}
