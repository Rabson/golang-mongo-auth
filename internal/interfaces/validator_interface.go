package interfaces

type Validator interface {
	Validate(data map[string]interface{}) error

	GetKeys() []string
}
