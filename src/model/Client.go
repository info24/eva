package model

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	id        string
	token     string
	ice       map[string]any
	conn      *websocket.Conn
	lastHeart int64
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{id: id, conn: conn, lastHeart: time.Now().UnixMilli()}
}

func (client *Client) Send(msg any) {
	client.conn.WriteJSON(msg)
}

func (client *Client) SetIce(msg map[string]any) {
	client.ice = msg
}
