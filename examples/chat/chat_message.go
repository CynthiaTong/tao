package chat

import (
	"context"

	"github.com/cynthiatong/tao"
	"github.com/leesper/holmes"
)

const (
	// ChatMessageNumber is the message number of chat message.
	ChatMessageNumber int32 = 1
)

// Message defines the chat message.
type Message struct {
	Content string
}

// MessageNumber returns the Message number.
func (cm Message) MessageNumber() int32 {
	return ChatMessageNumber
}

// MessageName returns the Message name.
func (cm Message) MessageName() string {
	return ""
}

// Serialize Serializes Message into bytes.
func (cm Message) Serialize() ([]byte, error) {
	return []byte(cm.Content), nil
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	content := string(data)
	msg := Message{
		Content: content,
	}
	return msg, nil
}

// ProcessMessage handles the Message logic.
func ProcessMessage(ctx context.Context, conn tao.WriteCloser) {
	holmes.Infof("ProcessMessage")
	s, ok := tao.ServerFromContext(ctx)
	if ok {
		msg := tao.MessageFromContext(ctx)
		s.Broadcast(msg)
	} else {
		holmes.Errorf("ServerFromContext error")
	}
}
