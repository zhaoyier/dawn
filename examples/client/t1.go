package main
import (
	"net"
	"os"
	"bytes"
	"encoding/binary"
	"zhao.com/examples/proto3"
	"github.com/golang/protobuf/proto"
	//"fmt"
	//"fmt"
	"fmt"
)

func main() {
	strEcho := "Halo"
	servAddr := "localhost:11000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	//_, err = conn.Write([]byte(strEcho))
	_, err = conn.Write(msg())
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	//data := read(conn)
	temp := &proto3.Page{}
	proto.Unmarshal(reply, temp)

	fmt.Println("reply from server=", temp)
	//println("======>>001:\t", temp)

	conn.Close()
}

func msg() []byte {
	var buf bytes.Buffer
	buf.Write(Int32ToBytes(1000))
	data, _ := proto.Marshal(&proto3.Page{
		PageSize: 20,
		PageNumber: 2,
	})
	buf.Write(Int32ToBytes(int32(len(data))))
	buf.Write(data)
	return buf.Bytes()
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}