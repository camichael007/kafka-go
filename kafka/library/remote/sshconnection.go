package remote

import (
	"time"
	"fmt"
	"golang.org/x/crypto/ssh"
)

/**
*
*author : daozai
*time   : 2017/9/18 21:27
*file   : sshconnection.go
*
*/

func SshConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err error
		addr string
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
