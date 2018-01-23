package testhelper

import (
	"net"
	"sync/atomic"
	"fmt"

	"gopcp.v2/helper/log"
)

// 日志记录器。
var logger = log.DLogger()

// ServerReq 表示服务器请求的结构。
type ServerReq struct {
	ID int64
	Operands []int
	Operator string
}

// ServerResp 表示服务器响应的结构。
type ServerResp struct {
	ID int64
	Formula string
	Result int
	Err error
}

// reqHandler 会把参数sresp代表的请求转换为数据并发送给连接。
func reqHandler(conn net.Conn) {
	// var errMsg string
	// var sresp ServerResp
}

// TCPServer tcp 协议的服务器
type TCPServer struct {
	listener net.Listener
	active	 uint32
}

// NewTCPServer 创建一个服务器实体
func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (server *TCPServer) init(addr string) error {
	if !atomic.CompareAndSwapUint32(&server.active, 0, 1) {
		return nil
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server.listener = ln
	return nil
}

// Listen 监听 server 服务	
func (server *TCPServer) Listen(addr string) error {
	err := server.init(addr)

	if err != nil {
		return err
	}

	go func() {
		for {
			// 原子级读取 tpc 状态
			if atomic.LoadUint32(&server.active) != 1 {
				fmt.Println("tcp server is'n active")
				break
			}

			conn, err := server.listener.Accept()

			if err != nil {
				if atomic.LoadUint32(&server.active) == 1 {
					logger.Errorf("Server: Request Acception Error: %s\n", err)
				} else {
					logger.Warnf("Server: Broken acception because of closed network connection.")
				}
				continue
			}
			go reqHandler(conn)
		}
	}()

	return nil
}