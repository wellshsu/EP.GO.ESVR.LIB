//go:binary-only-package
package xconn

import (
	_ "esvr/core/utility/xlog"
	"esvr/core/utility/xobj"
	_ "esvr/core/utility/xrun"
	_ "esvr/core/utility/xtime"
	_ "fmt"
	"net"
	_ "time"
)

type Client struct {
	xobj.OBJECT
	Server   *Server
	Conn     net.Conn
	ID       int
	Expired  int
	Tick     int
	Load     int
	Sender   *Sender   // 消息发送对象
	Receiver *Receiver // 消息接收对象
	ChClose  chan bool
}

func NewClient(server *Server) *Client {
	return nil
}
func (this *Client) CTOR(CHILD interface{}) {
	return
}
func (this *Client) GetServer() *Server {
	return nil
}
func (this *Client) SetServer(server *Server) {
	return
}
func (this *Client) GetConn() net.Conn {
	return nil
}
func (this *Client) SetConn(conn net.Conn) {
	return
}
func (this *Client) GetID() int {
	return 0
}
func (this *Client) SetID(id int) {
	return
}
func (this *Client) GetExpired() int {
	return 0
}
func (this *Client) SetExpired(expired int) {
	return
}
func (this *Client) GetTick() int {
	return 0
}
func (this *Client) SetTick(tick int) {
	return
}
func (this *Client) GetLoad() int {
	return 0
}
func (this *Client) SetLoad(load int) {
	return
}
func (this *Client) GetChClose() chan bool {
	return nil
}
func (this *Client) SetChClose(ch chan bool) {
	return
}
func (this *Client) GetSender() *Sender {
	return nil
}
func (this *Client) SetSender(sender *Sender) {
	return
}
func (this *Client) GetReceiver() *Receiver {
	return nil
}
func (this *Client) SetReceiver(receiver *Receiver) {
	return
}
func (this *Client) Reset(id int, conn net.Conn) {
	return
}
func (this *Client) Clear() {
	return
}
func (this *Client) Write(bytes []byte) error {
	return nil
}
func (this *Client) Kick() {
	return
}
func (this *Client) Update() bool {
	return false
}
