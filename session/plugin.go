package session

import "github.com/wbsr9876/streamdocksdk/proto"

type Plugin interface {
	SetInfo(info *proto.Info)
	SetConnection(conn *ConnectionManager)
	HandleEvent(message *proto.MessageHeader, content []byte) error
	HandleAction(message *proto.MessageHeader, content []byte) error
}
