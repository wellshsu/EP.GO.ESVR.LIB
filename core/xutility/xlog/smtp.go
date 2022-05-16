//go:binary-only-package
// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package xlog

import (
	_ "crypto/tls"
	_ "encoding/json"
	_ "fmt"
	_ "net"
	"net/smtp"
	_ "strings"
	"time"
)

// SMTPWriter implements LoggerInterface and is used to send emails via given SMTP-server.
type SMTPWriter struct {
	Username           string   `json:"username"`
	Password           string   `json:"password"`
	Host               string   `json:"host"`
	Subject            string   `json:"subject"`
	FromAddress        string   `json:"fromAddress"`
	RecipientAddresses []string `json:"sendTos"`
	Level              int      `json:"level"`
}

// NewSMTPWriter create smtp writer.
func newSMTPWriter() Logger {
	return nil
}

// Init smtp writer with json config.
// config like:
//	{
//		"username":"example@gmail.com",
//		"password:"password",
//		"host":"smtp.gmail.com:465",
//		"subject":"email title",
//		"fromAddress":"from@example.com",
//		"sendTos":["email1","email2"],
//		"level":LevelError
//	}
func (s *SMTPWriter) Init(jsonconfig string) error {
	return nil
}
func (s *SMTPWriter) getSMTPAuth(host string) smtp.Auth {
	return nil
}
func (s *SMTPWriter) sendMail(hostAddressWithPort string, auth smtp.Auth, fromAddress string, recipients []string, msgContent []byte) error {
	return nil
}

// WriteMsg write message in smtp writer.
// it will send an email with subject and only this message.
func (s *SMTPWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}

// Flush implementing method. empty.
func (s *SMTPWriter) Flush() {
	return
}

// Destroy implementing method. empty.
func (s *SMTPWriter) Destroy() {
	return
}
func init() {
	return
}
