package model

type Body struct {
	Type MsgType `json:"type"`
	Id   string  `json:"id"`
	To   string  `json:"to"`
	Msg  any     `json:"msg"`
}

type MsgType int

const RequestJoin MsgType = 100
const RequestConnection MsgType = 101
const RequestHopeConnection MsgType = 102
const RequestIce MsgType = 103
const RequestHopeIce MsgType = 104
const RequestIceAnswer MsgType = 105
const RequestAllInstance MsgType = 106
const RequestIceCandidateLocal MsgType = 107
const RequestIceCandidateRemote MsgType = 108
const RequestShareFile MsgType = 110
const RequestShareFileCopy MsgType = 111
const ResponseJoin MsgType = 200
const ResponseConnection MsgType = 201
const ResponseHopeConnection MsgType = 202
const ResponseIce MsgType = 203
const ResponseHopeIce MsgType = 204
const ResponseIceAnswer MsgType = 205
const ResponseAllInstance MsgType = 206
const ResponseIceCandidateRemote MsgType = 208
const ResponseJsonErr MsgType = -1
const ResponseUuid MsgType = 300
