package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/info24/eva/model"
	"github.com/info24/eva/service"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}} // use default option

func RegisterDevice(c *gin.Context) {
	//c.JSON(200, gin.H{})
	defer func() {
		println("done")
		err := recover()
		if err != nil {
			println(err)
		}

	}()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		println(err.Error())
		c.JSON(500, response(500, err.Error()))
		return
	}
	defer conn.Close()
	//id := c.Params.ByName("id")
	id := createUUid().String()
	service.SetClient(id, conn)
	for {
		var body model.Body
		err := conn.ReadJSON(&body)
		if err != nil {
			println("read json error: " + err.Error())
			closeError, exist := err.(*websocket.CloseError)
			if exist {
				println("conn close: " + id + ", error: " + closeError.Error())
				service.RemoveClient(id)
				return
			}
			err := conn.WriteJSON(&model.Body{Type: model.ResponseJsonErr, Msg: err.Error()})
			if err != nil {
				println("conn close. " + err.Error())
				service.RemoveClient(id)
				return
			}
			continue
		}
		parse(&body, conn)

	}
	println("all done.")
}

func createUUid() uuid.UUID {
	return uuid.New()
}

func parse(body *model.Body, conn *websocket.Conn) {
	switch body.Type {
	case model.RequestJoin:
		println("connect")
		service.SetClient(body.Id, conn)
		break
	case model.RequestConnection:
		// deal with client connect other client
		// to request other client init
		service.ClientHopeConnection(body)
		conn.WriteJSON(&model.Body{
			Type: model.ResponseConnection,
		})
		break
	case model.RequestIce:
		println("ice")
		service.ClientIce(body)
		err := conn.WriteJSON(&model.Body{Type: model.ResponseIce})
		if err != nil {
			return
		}
		break
	case model.RequestIceAnswer:
		conn.WriteJSON(&model.Body{
			Type: model.ResponseIceAnswer,
			Msg:  body.Msg,
		})
		service.RequestIceAnswer(body)
		break
	case model.RequestAllInstance:
		// to get all alive instance list
		service.RequestAllClient(body.Id, conn)
		break
	case model.RequestIceCandidateRemote:
		service.RequestIceCandidateRemote(body)
		break
	case model.RequestIceCandidateLocal:
		service.RequestIceCandidateLocal(body)
		break
	case model.RequestShareFile,
		model.RequestShareFileCopy:
		service.RequestRelay(body)
		break
	}

}
