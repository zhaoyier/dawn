package dawn

import (
	"net"
	"flag"
	"fmt"
)

func init() {
	flag.Parse()
	netIdentifier = NewAtomicInt64(0)

}

var (
	netIdentifier *AtomicInt64
	tlsWrapper    func(net.Conn) net.Conn
)

type Server struct {
	conns  *ConnMap
}

//初始化服务信息
func NewServer() *Server {
	return &Server{
		conns: NewConnMap(),
	}
}

//开始接受连接请求
func (this *Server) Start(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {

	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {	//连接失败,等待重新连接

		}
		if this.conns.Size() > 100 {
			conn.Close()
			continue
		}
		netid := netIdentifier.GetAndIncrement()
		sc := NewServerConn(netid, this, conn)
		this.conns.Put(netid, sc)

		//开始监听消息
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("=====>>001:\t", reqLen, string(buf))
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}