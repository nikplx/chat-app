package message

import (
	"context"
	"github.com/justtrackio/gosoline/pkg/apiserver/crud"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/justtrackio/gosoline/pkg/log"
)

type CreateInput struct {
	Sender  uint   `json:"sender"`
	Chat    uint   `json:"chat"`
	Message string `json:"message"`
}

type Output struct {
	From     uint   `json:"sender"`
	FromName string `json:"fromName"`
	Chat     uint   `json:"chat"`
	Message  string `json:"message"`
}

type handler struct {
	repo db_repo.Repository
}

func (h *handler) GetModel() db_repo.ModelBased {
	return &Message{}
}

func (h *handler) GetRepository() crud.Repository {
	return h.repo
}

func (h *handler) GetCreateInput() interface{} {
	return &CreateInput{}
}

func (h *handler) TransformCreate(ctx context.Context, inp interface{}, model db_repo.ModelBased) error {
	c := inp.(*CreateInput)

	m := model.(*Message)
	m.ChatId = &c.Chat
	m.SenderId = &c.Sender
	m.Message = c.Message

	return nil
}

func (h *handler) TransformOutput(_ context.Context, m db_repo.ModelBased, apiView string) (interface{}, error) {
	model := m.(*Message)
	return &Output{
		From:     *model.Sender.Id,
		FromName: model.Sender.Name,
		Chat:     *model.ChatId,
		Message:  model.Message,
	}, nil
}

func NewCrudHandler(ctx context.Context, config cfg.Config, logger log.Logger) (*handler, error) {
	repo, err := NewRepo(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	return &handler{repo: repo}, nil
}

func (h *handler) List(ctx context.Context, qb *db_repo.QueryBuilder, apiView string) (interface{}, error) {
	result := make([]*Message, 0)
	err := h.repo.Query(ctx, qb, &result)
	if err != nil {
		return nil, err
	}

	out := make([]interface{}, 0)
	for _, res := range result {
		d, err := h.TransformOutput(ctx, res, apiView)
		if err != nil {
			return nil, err
		}

		out = append(out, d)
	}

	return out, nil
}
