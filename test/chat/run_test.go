//go:build integration && fixtures

package chat

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jmoiron/sqlx"
	"github.com/justtrackio/gosoline/pkg/apiserver"
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"
	"github.com/justtrackio/gosoline/pkg/fixtures"
	"github.com/justtrackio/gosoline/pkg/mdl"
	"github.com/justtrackio/gosoline/pkg/test/suite"
	"github.com/nikplx/task-matched/cmd/chat/api"
	"github.com/nikplx/task-matched/pkg/chat"
	"github.com/nikplx/task-matched/pkg/message"
	"github.com/nikplx/task-matched/pkg/user"
	"net/http"
	"testing"
)

type ApiServerTestSuite struct {
	suite.Suite

	client *resty.Client
}

func TestApiServerTestSuite(t *testing.T) {
	suite.Run(t, new(ApiServerTestSuite))
}

func (a *ApiServerTestSuite) SetupApiDefinitions() apiserver.Definer {
	return api.DefineRouter
}

func (a *ApiServerTestSuite) SetupSuite() []suite.Option {
	return []suite.Option{
		suite.WithLogLevel("info"),
		suite.WithConfigFile("../../cmd/chat/config.dist.yml"),
		// suite.WithConfigFile("./config.test.yml"),
		suite.WithFixtures(fixtureSets),
		suite.WithSharedEnvironment(),
	}
}

func (a *ApiServerTestSuite) CreateChat(init, sec uint) uint {
	inp := chat.CreateInput{
		Initiator: init,
		Second:    sec,
	}

	var out chat.Output
	resp, err := a.Call(http.MethodPost, "/v0/chat", inp, nil, &out)
	a.NoError(err)
	a.Equal(http.StatusOK, resp.StatusCode())

	a.Equal(out.Initiator, init)
	a.Equal(out.Second, sec)

	return out.Id
}

func (a *ApiServerTestSuite) TestCreateChat(_ suite.AppUnderTest, client *resty.Client) error {
	a.client = client
	a.CreateChat(1, 2)

	return nil
}

func (a *ApiServerTestSuite) TestCreateMessage(_ suite.AppUnderTest, client *resty.Client) error {
	a.client = client
	chatId := a.CreateChat(1, 2)

	inp := message.CreateInput{
		Sender:  1,
		Chat:    chatId,
		Message: "hello world",
	}

	var out message.Output
	resp, err := a.Call(http.MethodPost, "/v0/message", inp, nil, &out)
	a.NoError(err)
	a.Equal(http.StatusOK, resp.StatusCode())

	a.Equal(out.FromName, "kevin")
	a.Equal(out.Message, "hello world")

	return nil
}

func (a *ApiServerTestSuite) TestListMessages(_ suite.AppUnderTest, client *resty.Client) error {
	a.client = client
	chatId := a.CreateChat(1, 2)

	// send a new message
	message1 := message.CreateInput{
		Sender:  1,
		Chat:    chatId,
		Message: "hello world",
	}

	var messageOut1 message.Output
	resp, err := a.Call(http.MethodPost, "/v0/message", message1, nil, &messageOut1)
	a.NoError(err)
	a.Equal(http.StatusOK, resp.StatusCode())

	// answer the message
	message2 := message.CreateInput{
		Sender:  2,
		Chat:    chatId,
		Message: "hello world 2",
	}

	var messageOut2 message.Output
	resp, err = a.Call(http.MethodPost, "/v0/message", message2, nil, &messageOut2)
	a.NoError(err)
	a.Equal(http.StatusOK, resp.StatusCode())

	var messages messagesListOut
	filter := fmt.Sprintf("{\"filter\":{\"matches\":[{\"dimension\":\"chat.id\",\"operator\":\"=\",\"values\":[%d]}]}}", chatId)
	resp, err = a.Call(http.MethodPost, "/v0/messages", filter, nil, &messages)
	a.NoError(err)
	a.Equal(http.StatusOK, resp.StatusCode())
	a.Equal(messages.Total, 2)

	var expected = []message.Output{
		{
			From:     1,
			FromName: "kevin",
			Chat:     chatId,
			Message:  "hello world",
		},
		{
			From:     2,
			FromName: "lucy",
			Chat:     chatId,
			Message:  "hello world 2",
		},
	}
	a.ElementsMatch(messages.Results, expected)

	return nil
}

type messagesListOut struct {
	Total   int              `json:"total"`
	Results []message.Output `json:"results"`
}

func (a *ApiServerTestSuite) Call(method string, urlPath string, body interface{}, additionalHeaders map[string]string, result interface{}) (*resty.Response, error) {
	req := a.client.R().
		SetHeaders(additionalHeaders).
		SetBody(body)

	if result != nil {
		req = req.SetResult(result)
	}

	return req.Execute(method, urlPath)
}

func (a *ApiServerTestSuite) sql() *sqlx.DB {
	return a.Env().MySql("default").Client()
}

var fixtureSets = []*fixtures.FixtureSet{
	{
		Enabled: true,
		Purge:   false,
		Writer:  fixtures.MysqlOrmFixtureWriterFactory(&user.TableMetadata),
		Fixtures: []interface{}{
			&user.User{
				Model: db_repo.Model{Id: mdl.Uint(1)},
				Name:  "kevin",
			},
			&user.User{
				Model: db_repo.Model{Id: mdl.Uint(2)},
				Name:  "lucy",
			},
		},
	},
}
