package constants

import "golang-mongo-auth/pkg/common/types"

const (
	MODULE_USER types.Module = "USER"
)

const (
	ACTION_CREATE types.Action = "CREATE"
	ACTION_UPDATE types.Action = "UPDATE"
	ACTION_READ   types.Action = "READ"
	ACTION_DELETE types.Action = "DELETE"
)

const (
	ROLE_USER  types.Role = "USER"
	ROLE_ADMIN types.Role = "ADMIN"
)

var RoleModuleActions = types.RoleModuleAction{
	ROLE_ADMIN: {
		MODULE_USER: {ACTION_READ, ACTION_UPDATE},
	},
	ROLE_USER: {
		MODULE_USER: {ACTION_READ, ACTION_UPDATE},
	},
}
