package main

import (
	"github.com/justtrackio/gosoline/pkg/application"
	"github.com/nikplx/task-matched/cmd/chat/api"
)

func main() {
	application.RunApiServer(api.DefineRouter)
}
