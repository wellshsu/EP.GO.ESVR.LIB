//go:binary-only-package
package xlog

import (
	_ "encoding/json"
	_ "fmt"
	_ "net/http"
	_ "net/url"
	"time"
)

// JLWriter implements beego LoggerInterface and is used to send jiaoliao webhook
type JLWriter struct {
	AuthorName  string `json:"authorname"`
	Title       string `json:"title"`
	WebhookURL  string `json:"webhookurl"`
	RedirectURL string `json:"redirecturl,omitempty"`
	ImageURL    string `json:"imageurl,omitempty"`
	Level       int    `json:"level"`
}

// newJLWriter create jiaoliao writer.
func newJLWriter() Logger {
	return nil
}

// Init JLWriter with json config string
func (s *JLWriter) Init(jsonconfig string) error {
	return nil
}

// WriteMsg write message in smtp writer.
// it will send an email with subject and only this message.
func (s *JLWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}

// Flush implementing method. empty.
func (s *JLWriter) Flush() {
	return
}

// Destroy implementing method. empty.
func (s *JLWriter) Destroy() {
	return
}
func init() {
	return
}
