package session

import "github.com/wbsr9876/streamdocksdk/proto"

type Plugin interface {
	SetInfo(info *proto.Info)
	SetConnection(conn *ConnectionManager)
	HandleEvent(event Event) error
	HandleAction(action Action) error
}
