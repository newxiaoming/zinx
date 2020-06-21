package ziface

import "net"

//IConnection 定义连接接口
type IConnection interface {
	//启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接状态M
	Stop()
	//从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获取当前连接ID
	GetConnID() uint32
	//获取远程客户端地址信息
	RemoteAddr() net.Addr
}

//HandleFunc 定义一个统一处理链接业务的接口,第一参数是 socket 原生链接，第二个参数是客户端请求的数据，第三个参数是客户端请求的数据长度
type HandleFunc func(*net.TCPConn, []byte, int) error
