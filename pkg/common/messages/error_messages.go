package messages

const (
	// failed
	ErrSomethingWentWrong = "Something went wrong" // 500
	ErrOperationFailed    = "Operation failed"     // 500
	// unauthorized
	ErrMissingToken       = "Missing token"       // 401
	ErrInvalidToken       = "Invalid token"       // 401
	ErrTokenExpired       = "Token expired"       // 401
	ErrInvalidCredentials = "Invalid credentials" // 401
	ErrInvalidRole        = "Invalid role"        // 401
	// Forbidden
	ErrAccessDenied     = "Access denied"     // 403
	ErrPermissionDenied = "Permission denied" // 403
	ErrForbidden        = "Forbidden"         // 403
	// Notfound
	ErrResourceNotFound                     = "Resource not found"                           // 404
	ErrUserNotFound                         = "User not found"                               // 404
	ErrEmailAlreadyExists                   = "Email already exists"                         // 409
	ErrUserNotAuthorized                    = "User not authorized"                          // 403
	ErrUserNotFoundById                     = "User not found by ID"                         // 404
	ErrUserNotFoundByEmail                  = "User not found by email"                      // 404
	ErrUserNotFoundByUsername               = "User not found by username"                   // 404
	ErrUserNotFoundByPhone                  = "User not found by phone"                      // 404
	ErrUserNotFoundByUsernameOrEmail        = "User not found by username or email"          // 404
	ErrUserNotFoundByUsernameOrPhone        = "User not found by username or phone"          // 404
	ErrUserNotFoundByEmailOrPhone           = "User not found by email or phone"             // 404
	ErrUserNotFoundByUsernameOrEmailOrPhone = "User not found by username or email or phone" // 404
	// Invalid
	ErrInvalidData              = "Invalid data"                 // 400
	ErrInvalidEmailFormat       = "Invalid email format"         // 400
	ErrInvalidPhoneFormat       = "Invalid phone format"         // 400
	ErrInvalidUsernameFormat    = "Invalid username format"      // 400
	ErrInvalidPasswordFormat    = "Invalid password format"      // 400
	ErrInvalidRoleFormat        = "Invalid role format"          // 400
	ErrInvalidStatusFormat      = "Invalid status format"        // 400
	ErrInvalidCreatedAtFormat   = "Invalid created at format"    // 400
	ErrInvalidUpdatedAtFormat   = "Invalid updated at format"    // 400
	ErrInvalidDeletedAtFormat   = "Invalid deleted at format"    // 400
	ErrInvalidIdFormat          = "Invalid id format"            // 400
	ErrInvalidObjectIdFormat    = "Invalid object id format"     // 400
	ErrInvalidObjectId          = "Invalid object id"            // 400
	ErrInvalidObjectIdHex       = "Invalid object id hex"        // 400
	ErrInvalidObjectIdHexFormat = "Invalid object id hex format" // 400
	//
)
