package chat

import (
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/nikplx/task-matched/pkg/user"
)

type Chat struct {
	db_repo.Model
	InitiatorId *uint
	Initiator   *user.User
	SecondId    *uint
	Second      *user.User
}
