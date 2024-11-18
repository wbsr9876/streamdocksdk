package session

type Plugin interface {
	SetConnection(conn *ConnectionManager)
	HandleEvent(event Event) error
	HandleAction(action Action) error
}
