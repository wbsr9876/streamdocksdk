package base

import (
	"fmt"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/proto"
	"github.com/wbsr9876/streamdocksdk/session"
)

type Plugin struct {
	Actions map[string]ActionInterface
	Info    *proto.Info
}

func (p *Plugin) RegisterAction(name string, action ActionInterface) {
	if p.Actions == nil {
		p.Actions = make(map[string]ActionInterface)
	}
	p.Actions[name] = action
}

func (p *Plugin) SetInfo(info *proto.Info) {
	p.Info = info
}

func (p *Plugin) SetConnection(conn *session.ConnectionManager) {
	for _, action := range p.Actions {
		action.SetConnection(conn)
	}
}

func (p *Plugin) HandleAction(header *proto.MessageHeader, body []byte) error {
	if header == nil {
		return fmt.Errorf("header is nil")
	}
	if header.Action != "" {
		act, ok := p.Actions[header.Action]
		if !ok {
			log.Message(fmt.Sprintf("unknown action %v", header.Action))
			return nil
		}
		act.SetContext(header.Context)
		return act.HandleAction(header, body)
	}
	return nil
}

func (p *Plugin) HandleEvent(header *proto.MessageHeader, body []byte) error {
	return nil
}
