package service

import (
	"github.com/gorilla/websocket"
	"github.com/info24/eva/model"
)

var conn map[string]*model.Client

func init() {
	conn = make(map[string]*model.Client)
}

func GetClient(id string) *model.Client {
	return conn[id]
}

var joinMsg string = " join the room."

func SetClient(id string, connection *websocket.Conn) {
	println("add id:　" + id)
	// set client uuid.
	connection.WriteJSON(&model.Body{Type: model.ResponseUuid, Id: id, Msg: id})
	// welcome new friend.
	for s := range conn {
		conn[s].Send(&model.Body{Type: model.ResponseConnection, Id: s, Msg: id + joinMsg})
	}
	client := model.NewClient(id, connection)
	conn[id] = client
}

func RemoveClient(id string) {
	println("remove id: " + id)
	delete(conn, id)
}

func RequestAllClient(id string, connection *websocket.Conn) {
	var ids []string
	for key := range conn {
		if key != id {
			ids = append(ids, key)
		}
	}
	conn[id].Send(&model.Body{Type: model.ResponseAllInstance, Id: id, Msg: ids})
}

func SendMsg(id string, msg any) {
	client := GetClient(id)
	client.Send(msg)
}

func ClientHopeConnection(body *model.Body) {
	_from := body.Id
	_to := body.To
	SendMsg(_to, model.Body{Type: model.RequestHopeConnection, Id: _to, To: _from})
}

func ClientIce(body *model.Body) {
	_from := body.Id
	_to := body.To
	GetClient(_from).SetIce(body.Msg.(map[string]any))

	SendMsg(_to, model.Body{Type: model.RequestHopeIce, Id: _to, To: _from, Msg: body.Msg})
}

func RequestIceAnswer(body *model.Body) {
	SendMsg(
		body.To,
		model.Body{Type: model.RequestIceAnswer, Id: body.To, To: body.Id, Msg: body.Msg},
	)
}

func RequestIceCandidateLocal(body *model.Body) {
	SendMsg(
		body.To,
		model.Body{Type: model.RequestIceCandidateLocal, Id: body.To, To: body.Id, Msg: body.Msg},
	)
}

func RequestIceCandidateRemote(body *model.Body) {
	SendMsg(
		body.To,
		model.Body{Type: model.RequestIceCandidateRemote, Id: body.To, To: body.Id, Msg: body.Msg},
	)
}

func RequestShareFile(body *model.Body) {
	SendMsg(
		body.To,
		model.Body{Type: model.RequestShareFile, Id: body.To, To: body.Id, Msg: body.Msg},
	)
}

func RequestRelay(body *model.Body) {
	SendMsg(
		body.To,
		model.Body{Type: body.Type, Id: body.To, To: body.Id, Msg: body.Msg},
	)
}
