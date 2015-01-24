package main

import (
	"code.google.com/p/go.crypto/ssh"
	"io"
)


type Host struct {
	Name     string
	User     string
	Password string
	Port     string
}

type Script struct {
	Id     int
	Name   string
	Script string
}

func (self *Script) Execute(host Host, out io.Writer) error {
	cfg := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(host.Password),
		},
	}
	client, err := ssh.Dial("tcp", host.Name+":"+host.Port, cfg)
	if err != nil {
		return err
	}
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = out
	session.Stderr = out
	if err := session.Run(self.Script); err != nil {
		return err
	}
	return nil
}
