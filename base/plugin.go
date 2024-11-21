package base

import (
	"fmt"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/proto"
	"github.com/wbsr9876/streamdocksdk/session"
)

type PluginInf interface {
	AgentInf
	SetConnection(conn *session.ConnectionManager)
	Init()
}

type Plugin struct {
	Agent
	Actions map[string]ActionInf
	Info    *proto.Info
}

func (p *Plugin) Init(plugin PluginInf) {
	p.Agent.Init(plugin)
}

func (p *Plugin) RegisterAction(name string, action ActionInf) {
	if p.Actions == nil {
		p.Actions = make(map[string]ActionInf)
	}
	action.Init()
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

func (p *Plugin) OnMessage(message *session.Message) {
	if message.Header != nil {
		if message.Header.Action != "" {
			act, ok := p.Actions[message.Header.Action]
			if !ok {
				log.Message(fmt.Sprintf("unknown action %v", message.Header.Action))
				return
			}
			act.OnMessage(message)
		} else {
			p.Agent.OnMessage(message)
		}
	}
}

// TxBegin TODO implement
func (p *Plugin) TxBegin(message *session.Message) {
	p.Agent.TxBegin(message)
}

// Tick TODO implement
func (p *Plugin) Tick(tick int) {
	//Do nothing
}
