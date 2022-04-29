package user

import (
	"github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/justtrackio/gosoline/pkg/mdl"
)

type User struct {
	db_repo.Model
	Name string
}

var TableMetadata = db_repo.Metadata{
	ModelId: mdl.ModelId{
		Application: "chat",
		Name:        "users",
	},
	TableName:  "users",
	PrimaryKey: "users.id",
	Mappings: db_repo.FieldMappings{
		"user.id": db_repo.NewFieldMapping("users.id"),
	},
}
