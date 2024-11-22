package base

import (
	"encoding/json"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/session"
	"time"
)

type AgentInf interface {
	Tick()
	TxBegin(message *session.Message)
	OnMessage(message *session.Message)
}
type Agent struct {
	ticker      *time.Ticker
	agent       AgentInf
	messageChan chan *session.Message
	tx          interface{}
	txEnd       func()
}

func (p *Agent) SetTx(tx interface{}, end func()) {
	p.tx = tx
	p.txEnd = end
}

func (p *Agent) TxBegin(message *session.Message) {
	_ = json.Unmarshal(message.Body, p.tx)
	if p.txEnd != nil {
		p.txEnd()
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
			p.agent.Tick()
		case m := <-p.messageChan:
			if m.Header == nil {
				p.ticker.Stop()
				return
			}
			log.Message("message:%+v", m.Header)
			p.agent.TxBegin(m)
		}
	}
}
