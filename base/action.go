package base

import (
	"github.com/wbsr9876/streamdocksdk/proto"
	"github.com/wbsr9876/streamdocksdk/session"
	"strings"
)

type ActionInterface interface {
	SetContext(string)
	SetConnection(*session.ConnectionManager)
	SetState(int) error
	SetTitle(string, proto.ESDSDKTarget, int) error
	SetImage(string, proto.ESDSDKTarget, int) error
	SetSettings(interface{}) error
	SetGlobalSettings(interface{}) error
	ShowAlert() error
	ShowOk() error
	SendToPropertyInspector(interface{}) error
	SendToPlugin(interface{}) error
	HandleAction(header *proto.MessageHeader, body []byte) error
}

type Action struct {
	Conn    *session.ConnectionManager
	Context string
	Name    string
}

func (p *Action) SetContext(context string) {
	p.Context = context
}

func (p *Action) SetConnection(conn *session.ConnectionManager) {
	p.Conn = conn
}

func (p *Action) SetState(state int) error {
	msg := proto.NewSetState()
	msg.Payload.State = state
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action) SetTitle(title string, target proto.ESDSDKTarget, state int) error {
	msg := proto.NewSetTitle()
	msg.Payload.Title = title
	msg.Payload.Target = target
	msg.Payload.State = state
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action) SetImage(image string, target proto.ESDSDKTarget, state int) error {
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

func (p *Action) SetSettings(settings interface{}) error {
	msg := proto.NewSetSettings()
	msg.Payload = settings
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action) SetGlobalSettings(settings interface{}) error {
	msg := proto.NewSetGlobalSettings()
	msg.Payload = settings
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action) ShowAlert() error {
	msg := proto.NewShowAlert()
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action) ShowOk() error {
	msg := proto.NewShowOk()
	msg.Context = p.Context
	return p.Conn.Send(msg)
}

func (p *Action) SendToPropertyInspector(payload interface{}) error {
	msg := proto.NewSendToPropertyInspector()
	msg.Context = p.Context
	msg.Action = p.Name
	msg.Payload = payload
	return p.Conn.Send(msg)
}

func (p *Action) SendToPlugin(payload interface{}) error {
	msg := proto.NewSendToPlugin()
	msg.Context = p.Context
	msg.Action = p.Name
	msg.Payload = payload
	return p.Conn.Send(msg)
}

func (p *Action) HandleAction(header *proto.MessageHeader, body []byte) error {
	return nil
}
