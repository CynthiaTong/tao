package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cynthiatong/tao"
	"github.com/cynthiatong/tao/examples/chat"
	"github.com/leesper/holmes"
)

var (
	banlist = map[string]bool{"127.0.0.1": true}
)

// ChatServer is the chatting server.
type ChatServer struct {
	*tao.Server
}

// NewChatServer returns a ChatServer.
func NewChatServer() *ChatServer {
	onConnectOption := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})
	onErrorOption := tao.OnErrorOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on error")
	})
	onCloseOption := tao.OnCloseOption(func(conn tao.WriteCloser) {
		holmes.Infoln("close chat client")
	})
	banlistOption := tao.BanlistOption(banlist)
	return &ChatServer{
		tao.NewServer(onConnectOption, onErrorOption, onCloseOption, banlistOption),
	}
}

func main() {
	defer holmes.Start().Stop()

	tao.Register(chat.ChatMessageNumber, chat.DeserializeMessage, chat.ProcessMessage)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 12345))
	if err != nil {
		holmes.Fatalln("listen error", err)
	}
	chatServer := NewChatServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		chatServer.Stop()
	}()

	holmes.Infoln(chatServer.Start(l))
}
