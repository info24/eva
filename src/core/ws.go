package core

import (
	"github.com/info24/eva/common"
	"github.com/olahol/melody"
	"net/http"
	"sync"
	"time"
)

var sshInstance map[*http.Request]*SshSession
var mux sync.Mutex
var instance *melody.Melody
var done = make(chan bool)

func init() {
	sshInstance = make(map[*http.Request]*SshSession, 8)
	instance = newWsInstance()
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				deleteList := make([]*http.Request, 0)
				for request, session := range sshInstance {
					if session.TimeOut() {
						deleteList = append(deleteList, request)
					}
				}
				if len(deleteList) > 0 {
					mux.Lock()
					for _, request := range deleteList {
						delete(sshInstance, request)
					}
					mux.Unlock()
				}
			}
		}

	}()
}

func newWsInstance() *melody.Melody {
	m := melody.New()
	m.HandleConnect(func(session *melody.Session) {
		instan, err := create(
			session.Keys[common.IP].(string),
			session.Keys[common.USERNAME].(string),
			session.Keys[common.PASSWORD].(string),
			session.Keys[common.PTY].(string),
			session.Keys[common.ROW].(int),
			session.Keys[common.COL].(int),
		)
		if err != nil {
			session.Write([]byte("open ssh error: " + err.Error()))
			session.Close()
			return
		}
		mux.Lock()
		sshInstance[session.Request] = instan
		mux.Unlock()
		go func() {
			buffer := make([]byte, 102400)
			for {
				n, err := instan.reader.Read(buffer)
				if err != nil {
					session.Write([]byte("\r\nread error: " + err.Error()))
					break
				}
				session.Write(buffer[:n])
			}
		}()
	})

	m.HandleDisconnect(func(session *melody.Session) {
		conn := sshInstance[session.Request]
		if conn != nil {
			delete(sshInstance, session.Request)
		}
	})
	m.HandleMessage(func(session *melody.Session, msg []byte) {
		conn := sshInstance[session.Request]
		if conn != nil {
			conn.writer.Write(msg)
		}
	})
	m.HandlePong(func(session *melody.Session) {
		conn := sshInstance[session.Request]
		if conn != nil {
			conn.UpdatePong()
		}
	})

	return m
}

func RegisterWsInstance(r *http.Request, w http.ResponseWriter, keys map[string]any) {
	instance.HandleRequestWithKeys(w, r, keys)
}
