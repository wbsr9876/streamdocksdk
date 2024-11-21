package session

import "github.com/wbsr9876/streamdocksdk/proto"

type Plugin interface {
	SetInfo(info *proto.Info)
	SetConnection(conn *ConnectionManager)
	OnMessage(message *Message)
}
