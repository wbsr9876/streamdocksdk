package session

type Action interface {
	GetEvent() string
	GetAction() string
	GetContext() string
}

type Event interface {
	GetEvent() string
}
