package dawn

import (
	"bytes"
	"fmt"
)

type UnmarshalFunc func(b []byte) (bs []byte, err error)

// handlerUnmarshaler is a combination of unmarshal and handle functions for message.
type handlerUnmarshaler struct {
	//handler     HandlerFunc
	unmarshaler UnmarshalFunc
}

var (
	buf *bytes.Buffer
	// messageRegistry is the registry of all
	// message-related unmarshal and handle functions.
	messageRegistry map[int32]handlerUnmarshaler
)

func init() {
	buf = new(bytes.Buffer)
	messageRegistry = map[int32]handlerUnmarshaler{}
}

func Register(fid int32, fn func(b []byte) (bs []byte, err error)) {
	if _, ok := messageRegistry[fid]; ok {
		panic(fmt.Sprintf("trying to register message %d twice", fid))
	}

	messageRegistry[fid] = handlerUnmarshaler{
		unmarshaler: fn,
	}
}

func getUnMarshalFunc(msgType int32) UnmarshalFunc {
	entry, ok := messageRegistry[msgType]
	if !ok {
		return nil
	}
	return entry.unmarshaler
}
