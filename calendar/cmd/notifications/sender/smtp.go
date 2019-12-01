/*
 * HomeWork-14: RabbitMQ receiver
 * Created on 30.11.2019 23:30
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

// EmailTemplate for emails.
const EmailTemplate = `
From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

const (
	userName = "cap@kirk.com"
	password = "hrenvam"
	server   = "192.168.137.2"
	port     = 25
)

// EmailMessage base struct.
type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type loginAuth struct {
	username, password string
}

var tmplEmail *template.Template

func init() {
	tmplEmail = template.New("email")
	if _, err := tmplEmail.Parse(EmailTemplate); err != nil {
		log.Println(err)
	}
}

func sendEmail(from, subject, body string, to []string) error {
	message := &EmailMessage{
		From:    from,
		Subject: subject,
		Body:    body,
		To:      to,
	}

	var bodyB bytes.Buffer
	if err := tmplEmail.Execute(&bodyB, message); err != nil {
		return err
	}

	auth := authLogin(userName, password)

	sp := fmt.Sprintf("%s:%d", server, port)
	if err := smtp.SendMail(sp,
		auth,
		message.From,
		message.To,
		bodyB.Bytes(),
	); err != nil {
		return err
	}

	return nil
}

func authLogin(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}
