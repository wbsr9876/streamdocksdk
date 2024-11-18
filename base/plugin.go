package base

import (
	"fmt"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/session"
)

type Plugin struct {
	Actions map[string]ActionInterface
}

func (p *Plugin) RegisterAction(name string, action ActionInterface) {
	if p.Actions == nil {
		p.Actions = make(map[string]ActionInterface)
	}
	p.Actions[name] = action
}

func (p *Plugin) SetConnection(conn *session.ConnectionManager) {
	for _, action := range p.Actions {
		action.SetConnection(conn)
	}
}

func (p *Plugin) HandleAction(a session.Action) error {
	act, ok := p.Actions[a.GetAction()]
	if !ok {
		log.Message(fmt.Sprintf("unknown action %v", a))
		return nil
	}
	act.SetContext(a.GetContext())
	return act.HandleAction(a)
}
