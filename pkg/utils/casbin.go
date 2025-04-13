package utils

import (
	"golang-mongo-auth/pkg/common/types"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
)

var e *casbin.Enforcer

func CasbinLoad(uri string, dbName string) {
	pwd, _ := os.Getwd()
	modalPath := pwd + "/casbin/rbac_model.conf"
	policyPath := pwd + "/casbin/policy.csv"

	// clientOptions := options.Client().ApplyURI(uri)
	// adapter, err := mongodbAdapter.NewAdapterWithCollectionName(clientOptions, dbName, "casbin_rule")
	// if err != nil {
	// 	log.Fatalf("CasbinLoad: Failed to load to casbin mongodb: %v", err.Error())
	// }

	// e, err = casbin.NewEnforcer(path, adapter)
	// if err != nil {
	// 	log.Fatalf("CasbinLoad: Failed to load to casbin model: %v", err.Error())
	// }

	e, err := casbin.NewEnforcer(modalPath, policyPath)
	if err != nil {
		log.Fatalf("CasbinLoad: Failed to create enforcer: %v", err.Error())
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatalf("CasbinLoad: Failed to load policy: %v", err.Error())
	}
}

func CasbinValidateRole(role types.Role, module types.Module, action types.Action) bool {
	allowed, err := e.Enforce(role, module, action)
	if err != nil {
		log.Fatalf("CasbinValidateRole: Failed to enforce policy: %v", err.Error())
		return false
	}
	return allowed
}
