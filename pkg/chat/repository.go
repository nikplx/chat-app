package chat

import (
	"context"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/justtrackio/gosoline/pkg/mdl"
)

type repo struct {
	db_repo.Repository
}

func NewRepo(ctx context.Context, config cfg.Config, logger log.Logger) (db_repo.Repository, error) {
	tableMetadata := db_repo.Metadata{
		ModelId: mdl.ModelId{
			Application: "chat",
			Name:        "chat",
		},
		TableName:  "chats",
		PrimaryKey: "chat.id",
	}

	settings := db_repo.Settings{
		Metadata: tableMetadata,
	}

	return db_repo.New(config, logger, settings)
}
