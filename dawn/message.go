package dawn

import(
	"net"
	"io"
	"bytes"
	"errors"
	"encoding/binary"
	"fmt"
)

const (
	MSG_TYPE_SIZE = 4
	MSG_BODY_SIZE = 4
	MSG_BODY_MAX = 1<<23	//8M
)

type Header struct {
	mt int32	//消息类型
}

type Context struct {
	conn net.Conn
	Header Header
	Body []byte
}

type Codec interface {
	Decode(net.Conn) (Context, error)
	Encode(Context) ([]byte, error)
}

type CodecHeader struct {

}

func Decode(conn net.Conn) (*Context, error){
	msgType, err := ReadMsgLen(conn, MSG_TYPE_SIZE)
	if err != nil {
		return nil, errors.New("")
	}
	fmt.Println("======>>.3001:\t", msgType)
	msgLen, err := ReadMsgLen(conn, MSG_BODY_SIZE)
	if err != nil {
		return nil, errors.New("")
	}
	fmt.Println("======>>.3002:\t", msgLen)
	msgBody := make([]byte, msgLen)
	if _, err = io.ReadFull(conn, msgBody); err != nil {
		return nil, errors.New("")
	}

	return &Context{
		conn: conn,
		Header: Header{
			mt: msgType,
		},
		Body: msgBody,

	}, nil
}

func ReadMsgLen(conn net.Conn, size int) (len int32, err error) {
	bs := make([]byte, size)
	if _, err := io.ReadFull(conn, bs); err != nil {
		return 0, err
	}

	if binary.Read(bytes.NewReader(bs), binary.LittleEndian, &len); err != nil{
		return 0, err
	}

	return len, nil
}



