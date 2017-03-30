package gitlab

import (
	"github.com/dtkkki/cement/hooks"
)

// PushEvent is contains the info from a github push event payload.
type PushEvent struct {
	ref     string
	repo    hooks.Repo
	commits []hooks.Commit
}

// Ref returns the ref info.
func (event PushEvent) Ref() string {
	return event.ref
}

// Repo returns the repo info.
func (event PushEvent) Repo() hooks.Repo {
	return event.repo
}

// Commits returns the commits info.
func (event PushEvent) Commits() []hooks.Commit {
	return event.commits
}
