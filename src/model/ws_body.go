package model

import "github.com/info24/eva/common"

type WsBody struct {
	Code common.WsType `json:"code"`
	Msg  string        `json:"msg"`
}
