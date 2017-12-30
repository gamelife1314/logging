package logging

import (
	"io"
)

type MessageHandler interface {
	Write(message []byte)
}

// logging.StreamMessageHandler
type StreamMessageHandler struct {
	Level       MessageLevel
	Filter      MessageFilter
	Formatter   *MessageFormatter
	Destination io.Writer
}

func (handler StreamMessageHandler) Write(message []byte) {
	handler.Destination.Write(message)
}

// logging.FileMessageHandler
type FileMessageHandler struct {
	Level       MessageLevel
	Filter      MessageFilter
	Formatter   *MessageFormatter
	Destination io.Writer
}

func (handler FileMessageHandler) Write(message []byte) {
	handler.Destination.Write(message)
}
