package base

import (
	"github.com/wbsr9876/streamdocksdk/proto"
	"github.com/wbsr9876/streamdocksdk/session"
	"strings"
	"time"
)

type ActionInf interface {
	AgentInf
	SetConnection(conn *session.ConnectionManager)
	OnSettingsChanged()
	Init()
}

type Action[T any] struct {
	Agent
	Conn              *session.ConnectionManager
	Context           string
	Name              string
	State             int
	OnSettingsChanged func()

	settings       *T
	backupSettings []byte
	ticker         *time.Ticker
	stop           chan struct{}
	tick           int
	messageChan    chan *session.Message
	dirtyTick      int
}

func (p *Action[T]) Init(action ActionInf) {
	p.Agent.Init(action)
	p.OnSettingsChanged = action.OnSettingsChanged
}

func (p *Action[T]) GetSettings() *T {
	return p.settings
}

func (p *Action[T]) Tick(tick int) {
	if p.dirtyTick > 0 {
		p.dirtyTick--
		if p.dirtyTick == 0 {
			if p.OnSettingsChanged != nil {
				p.OnSettingsChanged()
			}
		}
	}
}

func (p *Action[T]) TxBegin(message *session.Message) {
	p.Context = message.Header.Context
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
			p.State = msg.Payload.State
			p.dirtyTick = 0
			if p.OnSettingsChanged != nil {
				p.OnSettingsChanged()
			}
		})
	}
	p.Agent.TxBegin(message)
}

func (p *Action[T]) SetConnection(conn *session.ConnectionManager) {
	p.Conn = conn
}

func (p *Action[T]) SetState(state int) error {
	msg := proto.NewSetState()
	msg.Payload.State = state
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action[T]) SetTitle(title string, target proto.ESDSDKTarget, state int) error {
	msg := proto.NewSetTitle()
	msg.Payload.Title = title
	msg.Payload.Target = target
	msg.Payload.State = state
	msg.Context = p.Context
	return p.Conn.Send(msg)
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
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action[T]) SetSettings(settings interface{}) error {
	msg := proto.NewSetSettings()
	msg.Payload = settings
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

//func (p *Action[T]) SetGlobalSettings(settings interface{}) error {
//	msg := proto.NewSetGlobalSettings()
//	msg.Payload = settings
//	msg.Context = p.Context
//	return p.Conn.Send(msg)
//}

func (p *Action[T]) ShowAlert() error {
	msg := proto.NewShowAlert()
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action[T]) ShowOk() error {
	msg := proto.NewShowOk()
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action[T]) SendToPropertyInspector(payload interface{}) error {
	msg := proto.NewSendToPropertyInspector()
	msg.Context = p.Context
	msg.Action = p.Name
	msg.Payload = payload
	return p.Conn.Send(msg)
}

func (p *Action[T]) SendToPlugin(payload *T) error {
	msg := proto.NewSendToPlugin[T]()
	msg.Context = p.Context
	msg.Action = p.Name
	msg.Payload = payload
	return p.Conn.Send(msg)
}
