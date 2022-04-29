package message

import (
	"github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/nikplx/task-matched/pkg/chat"
	"github.com/nikplx/task-matched/pkg/user"
)

type Message struct {
	db_repo.Model

	ChatId   *uint
	Chat     *chat.Chat
	SenderId *uint
	Sender   *user.User

	Message string
}
