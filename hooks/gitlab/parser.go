package gitlab

import (
	"github.com/dtkkki/cement/hooks"
	"github.com/tidwall/gjson"
)

var DefaultEventParser = new(EventParser)

// ParsePushEvent event from a gitlab push event payload.
func ParsePushEvent(payload []byte) (hooks.PushEvent, error) {
	return DefaultEventParser.ParsePushEvent(payload)
}

// EventParser is a payload parser for gitlab.
type EventParser struct {
}

// ParsePushEvent parse event from a gitlab push event payload.
func (parser *EventParser) ParsePushEvent(payload []byte) (hooks.PushEvent, error) {
	s := string(payload)
	ref := gjson.Get(s, "ref").String()
	repo := hooks.Repo{
		Name:        gjson.Get(s, "repository.name").String(),
		HomePage:    gjson.Get(s, "repository.homepage").String(),
		Description: gjson.Get(s, "repository.description").String(),
	}
	var commits = make([]hooks.Commit, 0)
	gjson.Get(s, "commits").ForEach(func(_, value gjson.Result) bool {
		author := hooks.Author{
			Name:  value.Get("author").Get("name").String(),
			Email: value.Get("author").Get("email").String(),
		}
		commit := hooks.Commit{
			ID:        value.Get("id").String(),
			Message:   value.Get("message").String(),
			Timestamp: value.Get("timestamp").String(),
			URL:       value.Get("url").String(),
			Author:    &author,
		}
		commits = append(commits, commit)
		return true
	})

	return PushEvent{
		ref:     ref,
		repo:    repo,
		commits: commits,
	}, nil

}
