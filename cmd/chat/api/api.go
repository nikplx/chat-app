package api

import (
	"context"
	"fmt"
	"github.com/justtrackio/gosoline/pkg/apiserver"
	"github.com/justtrackio/gosoline/pkg/apiserver/crud"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/nikplx/task-matched/pkg/chat"
	"github.com/nikplx/task-matched/pkg/message"
)

func DefineRouter(ctx context.Context, config cfg.Config, logger log.Logger) (*apiserver.Definitions, error) {
	d := &apiserver.Definitions{}

	chatHandler, err := chat.NewCrudHandler(ctx, config, logger)
	if err != nil {
		return nil, fmt.Errorf("could not create chat handler: %w", err)
	}

	crud.AddCreateHandler(logger, d, 0, "chat", chatHandler)

	messageHandler, err := message.NewCrudHandler(ctx, config, logger)

	crud.AddCreateHandler(logger, d, 0, "message", messageHandler)
	crud.AddListHandler(logger, d, 0, "message", messageHandler)

	return d, nil
}
