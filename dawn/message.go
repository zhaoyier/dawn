package dawn

import(
	"net"
	"io"
	"bytes"
	"errors"
	"encoding/binary"
	//"fmt"
	"zhao.com/dawn/util"
	"zhao.com/examples/proto3"
	"github.com/golang/protobuf/proto"

)

const (
	MSG_TYPE_SIZE = 4
	MSG_BODY_SIZE = 4
	MSG_BODY_MAX = 1<<23	//8M
)




type Header struct {
	rid int32	//请求类型
}

type Context struct {
	conn net.Conn
	Header Header
	Body []byte
}

//type Codec interface {
//	Decode(net.Conn) (Context, error)
//	Encode(Context) ([]byte, error)
//}
//
//type CodecHeader struct {
//
//}

func Decode(conn net.Conn) (*Context, error){
	msgType, err := ReadMsgLen(conn, MSG_TYPE_SIZE)
	if err != nil {
		return nil, errors.New("")
	}
	msgLen, err := ReadMsgLen(conn, MSG_BODY_SIZE)
	if err != nil {
		return nil, errors.New("")
	}
	if msgLen > MSG_BODY_MAX {
		return nil, errors.New("")
	}
	msgBody := make([]byte, msgLen)
	if _, err = io.ReadFull(conn, msgBody); err != nil {
		return nil, errors.New("")
	}

	temp := &proto3.Page{}
	_ = proto.Unmarshal(msgBody, temp)

	return &Context{
		conn: conn,
		Header: Header{
			rid: msgType,
		},
		Body: msgBody,

	}, nil
}

func Encode(c *Context) []byte {
	//执行接口函数
	fn := getUnMarshalFunc(c.Header.rid)
	if fn == nil {

	}
	data, err := fn(c.Body)
	if err != nil {

	}
	//打包消息头
	var buf bytes.Buffer
	buf.Write(util.Int32ToBytes(c.Header.rid))
	buf.Write(util.Int32ToBytes(int32(len(data))))
	//打包消息体
	buf.Write(data)
	return buf.Bytes()
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



