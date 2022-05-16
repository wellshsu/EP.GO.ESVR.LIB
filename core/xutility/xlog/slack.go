//go:binary-only-package
package xlog

import (
	_ "encoding/json"
	_ "fmt"
	_ "net/http"
	_ "net/url"
	"time"
)

// SLACKWriter implements beego LoggerInterface and is used to send jiaoliao webhook
type SLACKWriter struct {
	WebhookURL string `json:"webhookurl"`
	Level      int    `json:"level"`
}

// newSLACKWriter create jiaoliao writer.
func newSLACKWriter() Logger {
	return nil
}

// Init SLACKWriter with json config string
func (s *SLACKWriter) Init(jsonconfig string) error {
	return nil
}

// WriteMsg write message in smtp writer.
// it will send an email with subject and only this message.
func (s *SLACKWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}

// Flush implementing method. empty.
func (s *SLACKWriter) Flush() {
	return
}

// Destroy implementing method. empty.
func (s *SLACKWriter) Destroy() {
	return
}
func init() {
	return
}
