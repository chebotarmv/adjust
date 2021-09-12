package helpers

type Event string

const (
	PushEvent        Event = "PushEvent"
	WatchEvent       Event = "WatchEvent"
	PullRequestEvent Event = "PullRequestEvent"
	CreateEvent      Event = "CreateEvent"
	IssuesEvent      Event = "IssuesEvent"
	DeleteEvent      Event = "DeleteEvent"
)
