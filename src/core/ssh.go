package core

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"time"
)

const PongTimeOut = 60

type SshClient struct {
	User    string
	Passwd  string
	Timeout int
}

type SshSession struct {
	reader  io.Reader
	writer  io.WriteCloser
	stderr  io.Reader
	session *ssh.Session
	pong    time.Time
}

func CreateSshConfig(config SshClient) *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User:            config.User,
		Auth:            []ssh.AuthMethod{ssh.Password(config.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(config.Timeout) * time.Second,
	}

	sshConfig.SetDefaults()
	return sshConfig
}

func CreateSshTcp(address string, config *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial("tcp", address, config)
}

func create(ip, username, password, pty string, row, col int) (*SshSession, error) {
	config := SshClient{
		User:    username,
		Passwd:  password,
		Timeout: 100,
	}
	sshConfig := CreateSshConfig(config)

	tcp, err := CreateSshTcp(ip, sshConfig)
	if err != nil {
		log.Printf("open tcp error: %s", err)
		return nil, err
	}

	session, err := NewSshSession(tcp, pty, true, row, col)

	return session, err
}

func NewSshSession(client *ssh.Client, pty string, echo bool, row, col int) (*SshSession, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	if echo {
		modes := ssh.TerminalModes{
			ssh.ECHO: 1,
		}
		//err := session.RequestPty("vt100", 0, 2000, modes)
		err := session.RequestPty(pty, row, col, modes)
		if err != nil {
			return nil, err
		}
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = session.Shell()
	if err != nil {
		return nil, err
	}

	return &SshSession{
		reader:  stdout,
		writer:  stdin,
		stderr:  stderr,
		session: session,
		pong:    time.Now(),
	}, nil
}

func (session *SshSession) TimeOut() bool {
	sub := time.Now().Sub(session.pong)
	if sub.Seconds() > PongTimeOut {
		return true
	}
	return false
}

func (session *SshSession) UpdatePong() {
	session.pong = time.Now()
}
