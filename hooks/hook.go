package hooks

// EventParser is the hook payload parser.
type EventParser interface {
	ParsePushEvent([]byte) (PushEvent, error)
}

// PushEvent is a interface to return push event payload.
type PushEvent interface {
	Ref() string
	Repo() Repo
	Commits() []Commit
}

// Repo is the base information of the repo.
type Repo struct {
	Name        string
	HomePage    string
	Description string
}

// Commit is a commit.
type Commit struct {
	ID        string
	Message   string
	Timestamp string
	URL       string
	Author    *Author
}

// Author is the commit author.
type Author struct {
	Name  string
	Email string
}
