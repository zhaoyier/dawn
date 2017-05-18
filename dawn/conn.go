package dawn

import (
	"net"
	"time"
)

type ServerConn struct {
	cid    int64	 //连接编号
	belong *Server   //归属服务
	conn   net.Conn	 //socket句柄
	heart  int64	 //心跳
	status int32     //连接状态, 0: 正常, 1:超时
	latest int64 //最近一次发送消息的时间戳(非心跳)
	name    string //连接名称
}

func NewServerConn(id int64, s *Server, c net.Conn) *ServerConn {
	sc := &ServerConn{
		cid: id,
		belong: s,
		conn: c,
		heart:time.Now().UnixNano(),
		latest: 0,
		name: c.RemoteAddr().String(),
	}
	return sc
}

func Test() {
	time.Now()
}
