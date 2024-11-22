package base

import (
	"github.com/wbsr9876/streamdocksdk/proto"
	"github.com/wbsr9876/streamdocksdk/session"
	"strings"
)

type ActionConstraint[T any] interface {
	*T
	ActionInf
}

func NewAction[T any, P ActionConstraint[T]](name string, context string, conn *session.ConnectionManager) ActionInf {
	p := P(new(T))
	p.Init(name, context, conn, p)
	return p
}

type ActionCreator func(string, string, *session.ConnectionManager) ActionInf

type ActionInf interface {
	AgentInf
	OnSettingsChanged()
	Init(name, context string, conn *session.ConnectionManager, impl ActionInf)
}

type Action[T any] struct {
	Agent
	conn              *session.ConnectionManager
	context           string
	name              string
	settings          *T
	state             int
	messageChan       chan *session.Message
	dirtyTick         int
	onSettingsChanged func()
}

func (p *Action[T]) Init(name, context string, conn *session.ConnectionManager, action ActionInf) {
	p.name = name
	p.context = context
	p.conn = conn
	p.Agent.Init(action)
	p.onSettingsChanged = action.OnSettingsChanged
}

func (p *Action[T]) GetSettings() *T {
	return p.settings
}

func (p *Action[T]) GetState() int {
	return p.state
}

func (p *Action[T]) Tick() {
	if p.dirtyTick > 0 {
		p.dirtyTick--
		if p.dirtyTick == 0 {
			if p.onSettingsChanged != nil {
				p.onSettingsChanged()
			}
		}
	}
}

func (p *Action[T]) TxBegin(message *session.Message) {
	p.context = message.Header.Context
	switch message.Header.Event {
	case "sendToPlugin":
		p.settings = new(T)
		msg := &proto.SendToPlugin[T]{}
		msg.Payload = p.settings
		p.SetTx(msg, func() {
			_ = p.SetSettings(p.settings)
			p.dirtyTick = 5
		})
	case "willAppear":
		msg := &proto.WillAppear{}
		p.settings = new(T)
		msg.Payload.Settings = p.settings
		p.SetTx(msg, func() {
			p.state = msg.Payload.State
			p.dirtyTick = 0
			if p.onSettingsChanged != nil {
				p.onSettingsChanged()
			}
		})
	}
	p.Agent.TxBegin(message)
}

func (p *Action[T]) SetState(state int) error {
	p.state = state
	msg := proto.NewSetState()
	msg.Payload.State = state
	msg.Context = p.context
	return p.conn.Send(msg)
}

func (p *Action[T]) SetTitle(title string, target proto.ESDSDKTarget) error {
	msg := proto.NewSetTitle()
	msg.Payload.Title = title
	msg.Payload.Target = target
	msg.Payload.State = p.state
	msg.Context = p.context
	return p.conn.Send(msg)
}

func (p *Action[T]) SetImage(image string, target proto.ESDSDKTarget, state int) error {
	const prefix string = "data:image/png;base64,"
	msg := proto.NewSetImage()
	if image == "" || strings.HasPrefix(image, prefix) {
		msg.Payload.Image = image
	} else {
		msg.Payload.Image = prefix + image
	}
	msg.Payload.Target = target
	msg.Payload.State = state
	msg.Context = p.context
	return p.conn.Send(msg)
}

func (p *Action[T]) SetSettings(settings interface{}) error {
	msg := proto.NewSetSettings()
	msg.Payload = settings
	msg.Context = p.context
	return p.conn.Send(msg)
}

//func (p *Action[T]) SetGlobalSettings(settings interface{}) error {
//	msg := proto.NewSetGlobalSettings()
//	msg.Payload = settings
//	msg.context = p.context
//	return p.conn.Send(msg)
//}

func (p *Action[T]) ShowAlert() error {
	msg := proto.NewShowAlert()
	msg.Context = p.context
	return p.conn.Send(msg)
}

func (p *Action[T]) ShowOk() error {
	msg := proto.NewShowOk()
	msg.Context = p.context
	return p.conn.Send(msg)
}

func (p *Action[T]) SendToPropertyInspector(payload interface{}) error {
	msg := proto.NewSendToPropertyInspector()
	msg.Context = p.context
	msg.Action = p.name
	msg.Payload = payload
	return p.conn.Send(msg)
}

func (p *Action[T]) SendToPlugin(payload *T) error {
	msg := proto.NewSendToPlugin[T]()
	msg.Context = p.context
	msg.Action = p.name
	msg.Payload = payload
	return p.conn.Send(msg)
}
