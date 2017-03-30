package hook

import (
	"io/ioutil"

	"github.com/dtkkki/cement/hooks"
	"github.com/dtkkki/cement/hooks/github"
	"github.com/dtkkki/cement/hooks/gitlab"
	"github.com/dtkkki/cement/toolkits/log"
	"github.com/labstack/echo"
)

func newHookServer(parser hooks.EventParser) echo.HandlerFunc {
	return func(c echo.Context) error {
		bodyReader := c.Request().Body
		defer bodyReader.Close()
		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			return err
		}

		event, _ := parser.ParsePushEvent(body)
		log.Info(event.Ref())
		log.Info(event.Repo())
		log.Info(event.Commits())
		return nil
	}
}

var (
	GithubHookServer = newHookServer(github.DefaultEventParser)
	GitlabHookServer = newHookServer(gitlab.DefaultEventParser)
)
