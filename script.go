package main

import (
	"code.google.com/p/go.crypto/ssh"
	"io"
	"strconv"
)

type Host struct {
	Name     string
	User     string
	Password string
	Port     int
}

type Script struct {
	Name    string
	Host    string
	Content string
}

func (self *Script) Execute(host *Host, out io.Writer) error {
	cfg := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(host.Password),
		},
	}
	client, err := ssh.Dial("tcp", host.Name+":"+strconv.Itoa(host.Port), cfg)
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
	if err := session.Run(self.Content); err != nil {
		return err
	}
	return nil
}
