/**
 * 监听信号量/其他信号，实现优雅的停止，避免消息丢失
 */
package dawn

import (
	"net"
	"flag"
	"zhao.com/dawn/util"
	"fmt"
)

func init() {
	flag.Parse()
	netIdentifier = util.NewAtomicInt64(0)
}

var (
	netIdentifier *util.AtomicInt64
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
		connID := netIdentifier.Increment()
		fmt.Println("====>>.1001:\t", connID)
		sc := NewServerConn(connID, this, conn)
		this.conns.Put(connID, sc)

		go func() {
			//开始监听消息
			sc.process()
		}()
	}

}