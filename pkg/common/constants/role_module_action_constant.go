package constants

import (
	"golang-mongo-auth/pkg/common/types"
)

const (
	MODULE_USER types.Module = "users"
)

const (
	ROLE_USER  types.Role = "user"
	ROLE_ADMIN types.Role = "admin"
)

var RoleModuleActions = types.RoleModuleAction{
	ROLE_ADMIN: {
		MODULE_USER: {
			ACTION_READ,
			ACTION_UPDATE,
		},
	},
	ROLE_USER: {
		MODULE_USER: {
			ACTION_READ,
			ACTION_UPDATE,
		},
	},
}
