package session

import (
	"encoding/json"
	"github.com/wbsr9876/streamdocksdk/proto"
)

type Action interface {
	GetEvent() string
	GetAction() string
	GetContext() string
}

type Event interface {
	GetEvent() string
}

type Message struct {
	Header *proto.MessageHeader
	Body   []byte
}

func NewMessage(body []byte) (*Message, error) {
	m := &Message{
		Header: &proto.MessageHeader{},
		Body:   body,
	}
	return m, json.Unmarshal(m.Body, m.Header)
}
