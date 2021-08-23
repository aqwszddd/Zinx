package znet

import (
	"fmt"
	"net"

	"github.com/zinx/ziface"
)

type Server struct {
	// 服务区名称
	Name string
	// IP版本
	IpVersion string
	// IP地址
	Ip string
	// 端口
	Port int
}

// 初始化服务器
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp",
		Ip:        "0.0.0.0",
		Port:      8888,
	}

	return s
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP :%s, Port %d, is starting.\n", s.Ip, s.Port)

	go func() {
		// 1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("net.ResolveTcpAddr err", err)
			return
		}

		// 2.监听服务器的地址
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("net.ListenTCP err", err)
			return
		}

		fmt.Println("start Zinx server name, ", s.Name, " succ, Listenning...")

		// 3.堵塞的等待客户端链接，处理客户端链接业务（读写）
		for {
			// 如果有客户端链接过来，堵塞会返回
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listener.Accept err", err)
				return
			}

			// 已经与客户端建立业务，做一些业务,做一个最基础的最大512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read back buferr", err)
						continue
					}

					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}

	}()
}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或回收
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务器后的额外业务

	// 堵塞状态
	select {}
}
