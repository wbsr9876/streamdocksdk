package base

import (
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
	actionCreators map[string]ActionCreator
	actions        map[string]ActionInf
	info           *proto.Info
	conn           *session.ConnectionManager
}

func (p *Plugin) Init(plugin PluginInf) {
	p.actionCreators = make(map[string]ActionCreator)
	p.actions = make(map[string]ActionInf)
	p.Agent.Init(plugin)
}

func (p *Plugin) RegisterActionCreator(name string, f ActionCreator) {
	p.actionCreators[name] = f
}

func (p *Plugin) SetInfo(info *proto.Info) {
	p.info = info
}

func (p *Plugin) SetConnection(conn *session.ConnectionManager) {
	p.conn = conn
}

func (p *Plugin) OnMessage(message *session.Message) {
	if message.Header == nil {
		return
	}
	if message.Header.Action == "" {
		p.Agent.OnMessage(message)
		return
	}
	if message.Header.Context == "" {
		return
	}
	if message.Header.Event == "willDisappear" {
		if act, ok := p.actions[message.Header.Context]; ok {
			delete(p.actions, message.Header.Context)
			act.OnMessage(&session.Message{})
		}
		return
	}
	act, ok := p.actions[message.Header.Context]
	if !ok {
		if f, ok := p.actionCreators[message.Header.Action]; ok {
			if act = f(message.Header.Action, message.Header.Context, p.conn); act == nil {
				return
			}
			p.actions[message.Header.Context] = act
		}
	}
	act.OnMessage(message)
}

// TxBegin TODO implement
func (p *Plugin) TxBegin(message *session.Message) {
	p.Agent.TxBegin(message)
}

// Tick TODO implement
func (p *Plugin) Tick() {
	//Do nothing
}
