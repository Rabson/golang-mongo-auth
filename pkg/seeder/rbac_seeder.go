package seeder

import (
	"context"
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/user/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func SeedRbac(db *mongo.Database, isDryRun bool, shouldReset bool) {

	repository.SetRepositories(db)

	repository.RbacRepo.Collection.DeleteMany(context.TODO(), nil)

	for r, modules := range constants.RoleModuleActions {
		for m, actions := range modules {
			for _, a := range actions {
				rbac := &models.Rbac{
					PType: "p",
					V0:    r,
					V1:    m,
					V2:    a,
				}
				if isDryRun {
					log.Printf("[rbac] Dry-run mode: would insert %v\n", rbac)
					continue
				}

				if _, err := repository.RbacRepo.InsertOne(context.TODO(), rbac); err != nil {
					log.Fatal("Seeding RBAC failed:", err)
				}
			}
		}
	}

}
