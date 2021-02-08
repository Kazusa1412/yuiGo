package main

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Ip string
	Port int
}

// Server 构造器
func NewServer(ip string,port int) *Server {
	return &Server{
		ip,
		port,
	}
}

func (this *Server) Handler(connection net.Conn){
	fmt.Println("connect success")
}

// 启动服务
func (this *Server) Start() {
	// socket listen
	var listener,err = net.Listen("tcp",this.Ip + ":" + strconv.Itoa(this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ",err)
		return
	}

	// finally close listen socket
	defer listener.Close()

	// accept
	for {
		var connection,err = listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ",err)
			continue
		}
		// do handler
		go this.Handler(connection)
	}

}