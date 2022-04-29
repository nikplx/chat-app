package chat

import (
	"context"
	"github.com/justtrackio/gosoline/pkg/apiserver/crud"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/justtrackio/gosoline/pkg/log"
)

type CreateInput struct {
	Initiator uint `json:"initiator"`
	Second    uint `json:"second"`
}

type Output struct {
	Id        uint `json:"id"`
	Initiator uint `json:"initiator"`
	Second    uint `json:"second"`
}

type handler struct {
	repo db_repo.Repository
}

func (h *handler) GetModel() db_repo.ModelBased {
	return &Chat{}
}

func (h *handler) GetRepository() crud.Repository {
	return h.repo
}

func (h *handler) GetCreateInput() interface{} {
	return &CreateInput{}
}

func (h *handler) TransformCreate(ctx context.Context, inp interface{}, model db_repo.ModelBased) error {
	c := inp.(*CreateInput)

	m := model.(*Chat)
	m.InitiatorId = &c.Initiator
	m.SecondId = &c.Second

	return nil
}

func (h *handler) TransformOutput(_ context.Context, m db_repo.ModelBased, apiView string) (interface{}, error) {
	model := m.(*Chat)
	return &Output{
		Id:        *model.Id,
		Initiator: *model.InitiatorId,
		Second:    *model.SecondId,
	}, nil
}

func NewCrudHandler(ctx context.Context, config cfg.Config, logger log.Logger) (*handler, error) {
	repo, err := NewRepo(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	return &handler{repo: repo}, nil
}
