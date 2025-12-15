package session

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/wbsr9876/streamdocksdk/log"
	"github.com/wbsr9876/streamdocksdk/proto"
)

type ConnectionManager struct {
	port          int
	pluginUUID    string
	registerEvent string
	plugin        Plugin
	conn          *websocket.Conn
	closed        bool
	sendBuff      chan []byte
}

func NewConnectionManager(port int, pluginUUID string, registerEvent string, plugin Plugin) *ConnectionManager {
	c := &ConnectionManager{
		port:          port,
		pluginUUID:    pluginUUID,
		registerEvent: registerEvent,
		plugin:        plugin,
		sendBuff:      make(chan []byte, 512),
	}
	//log.SetConnection(c)
	plugin.SetConnection(c)
	return c
}

func (c *ConnectionManager) OnOpen() error {
	msg := &proto.Register{UUID: c.pluginUUID, Event: c.registerEvent}
	return c.Send(msg)
}

func (c *ConnectionManager) LogMessage(message string) {
	if message != "" {
		msg := proto.NewLogMessage()
		msg.Payload.Message = message
		_ = c.Send(msg)
	}
}

func (c *ConnectionManager) OnClose(code int, text string) error {
	c.closed = true
	log.Message("Close with reason: " + text)
	return nil
}

func (c *ConnectionManager) receiveLoop() {
	for {
		if c.closed {
			break
		}
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			c.closed = true
			c.LogMessage(fmt.Sprintf("Read error:%s", err.Error()))
			break
		}
		if messageType != websocket.TextMessage {
			continue
		}
		log.Message(string(message))
		m, err := NewMessage(message)
		if err != nil {
			continue
		}
		c.plugin.OnMessage(m)
	}
}

func (c *ConnectionManager) sendLoop() {
	for {
		if c.closed {
			break
		}
		select {
		case msg := <-c.sendBuff:
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				//log.Message("Write error:" + err.Error())
				c.closed = true
			}
		default:
		}
	}
}

func (c *ConnectionManager) Send(v interface{}) error {
	if c.closed || c.conn == nil {
		return errors.New("no connection")
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.sendBuff <- bytes
	//log.Message("Send:" + string(bytes))
	//err = c.conn.WriteMessage(websocket.TextMessage, bytes)
	//if err != nil {
	//	log.Message("Write error:" + err.Error())
	//	c.closed = true
	//}
	return err
}

func (c *ConnectionManager) Run() {
	url := fmt.Sprintf("ws://127.0.0.1:%d", c.port)
	var err error
	//var res *http.Response
	c.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Message("Dial error:" + err.Error())
		return
	}
	err = c.OnOpen()
	if err != nil {
		log.Message("Open failed.")
		return
	}
	//cache, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Message("res.Body:" + err.Error())
	//	return
	//}
	//log2.Println(string(cache))

	c.conn.SetCloseHandler(c.OnClose)
	go c.sendLoop()
	c.receiveLoop()
	//message := []byte("Hello, WebSocket!")
	//if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
	//	log.Message("Write error:" + err.Error())
	//}
}
