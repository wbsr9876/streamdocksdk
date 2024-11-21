package base

import (
	"encoding/json"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/session"
	"time"
)

type AgentInf interface {
	Tick(tick int)
	TxBegin(message *session.Message)
	OnMessage(message *session.Message)
}
type Agent struct {
	ticker      *time.Ticker
	agent       AgentInf
	tick        int
	messageChan chan *session.Message
	Tx          interface{}
	TxEnd       func()
}

func (p *Agent) SetTx(tx interface{}, end func()) {
	p.Tx = tx
	p.TxEnd = end
}

func (p *Agent) TxBegin(message *session.Message) {
	_ = json.Unmarshal(message.Body, p.Tx)
	if p.TxEnd != nil {
		p.TxEnd()
	}
	p.SetTx(nil, nil)
}

func (p *Agent) Init(agent AgentInf) {
	p.agent = agent
	p.messageChan = make(chan *session.Message, 128)
	p.ticker = time.NewTicker(time.Second)
	go p.loop()
}

func (p *Agent) Destroy() {
	p.OnMessage(&session.Message{})
}

func (p *Agent) OnMessage(message *session.Message) {
	p.messageChan <- message
}

func (p *Agent) loop() {
	for {
		select {
		case <-p.ticker.C:
			p.tick++
			p.agent.Tick(p.tick)
		case m := <-p.messageChan:
			if m.Header == nil {
				return
			}
			log.Message("message:%+v", m.Header)
			p.agent.TxBegin(m)
		}
	}
}
