package constants

import "golang-mongo-auth/pkg/common/types"

const (
	HTTP_READ   types.HttpMethod = "GET"
	HTTP_CREATE types.HttpMethod = "POST"
	HTTP_PUT    types.HttpMethod = "PUT"
	HTTP_PATCH  types.HttpMethod = "PATCH"
	HTTP_DELETE types.HttpMethod = "DELETE"
)

const (
	ACTION_READ   types.Action = "read"
	ACTION_CREATE types.Action = "create"
	ACTION_UPDATE types.Action = "update"
	ACTION_DELETE types.Action = "delete"
)

var HttpMethodToAction = map[types.HttpMethod]types.Action{
	HTTP_READ:   ACTION_READ,
	HTTP_CREATE: ACTION_CREATE,
	HTTP_PUT:    ACTION_UPDATE,
	HTTP_PATCH:  ACTION_UPDATE,
	HTTP_DELETE: ACTION_DELETE,
}
