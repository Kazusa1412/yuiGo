package main

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

func (this *Server) Start() {

}