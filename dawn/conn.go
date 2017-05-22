package dawn

import (
	"net"
	"time"
	"fmt"
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

/**
 * 开始接收消息
 */
func (s *ServerConn) process() {
	//var codec Codec
	ctx, err := Decode(s.conn)
	if err != nil {
		return
	}

	fmt.Println("======>>>2001:\t", ctx.Header)
	fmt.Println("======>>>2002:\t", string(ctx.Body))


	//buf := make([]byte, 1024)
	//// Read the incoming connection into the buffer.
	//reqLen, err := s.conn.Read(buf)
	//if err != nil {
	//	fmt.Println("Error reading:", err.Error())
	//}
	//fmt.Println("=====>>001:\t", reqLen, string(buf))


	// Send a response back to person contacting us.
	s.conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	s.conn.Close()
}

/**
 * 更新连接信息
 */
func (s *ServerConn) update() {
	//更新心跳
	//更新消息接收时间
	//更新消息接收数量
}

/**
 * 处理请求消息
 * 处理心跳包/回调消息
 */
