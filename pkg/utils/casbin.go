package utils

import (
	"golang-mongo-auth/pkg/common/types"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
)

var enforcer *casbin.Enforcer

func LoadCasbinModel(uri string) {

	var err error

	adapter, err := mongodbadapter.NewAdapter(uri)
	if err != nil {
		panic(err)
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/casbin/rbac_model.conf"
	println(path)

	enforcer, err = casbin.NewEnforcer(path, adapter)
	if err != nil {
		panic(err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		if err.Error() == "slice bounds out of range" {
			println("No policies found in the database. Ensure policies are added.")
		} else {
			panic(err)
		}
	}
}

func ValidateCasbin(sub string, obj types.Module, act types.Action) bool {
	if enforcer == nil {
		panic("Casbin enforcer is not initialized")
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		log.Println("ValidateCasbin: Error enforcing policy:", err)
		return false
	}

	return ok
}
