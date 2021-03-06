package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

// IServer接口实现，定义一个服务类
type server struct {

	// 服务器名称
	Name string
	// tcp4 or other
	IPVersion string
	// 服务绑定的IP地址
	IP string
	// 服务端口号
	Port int
}

//========================实现 ziface.IServer 全部方法===========================

// Start 开启网络服务
func (s *server) Start() {

	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	//开启一个goroutine去做服务端Linster业务
	go func() {
		//1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve to ipaddr,err:", err)
			return
		}

		//2.监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen:", s.IPVersion, "err:", err)
			return
		}

		// 已成功监听
		fmt.Println("start zinx server", s.Name, "succ, now listenning...")

		//3 启动server网络连接业务
		for {
			//3.1阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			//3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接

			//3.3 TODO Server.Start() 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的

			//3.4 编写一个最大512字节的回显服务
			go func() {
				//监听客户端数据
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read buf err:", err)
						continue
					}
					//回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write conn err:", err)
						continue
					}
				}
			}()
		}
	}()

}

// Stop 停止服务
func (s *server) Stop() {
	fmt.Println("stop zinx server , name:", s.Name)
	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

func (s *server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞，否则主goroutine退出，Linster也会退出
	for {
		time.Sleep(time.Second * 10)
	}
}

/**
创建一个服务器句柄
*/

func NewServer(name string) ziface.IServer {
	s := &server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
	return s
}
