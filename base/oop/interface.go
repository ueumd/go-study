package main

import "fmt"

/*
*
在 Go 语言中，类对接口的实现和子类对父类的继承一样，并没有提供类似 implement 这种关键字显式声明该类实现了哪个接口
一个类只要实现了某个接口要求的所有方法，我们就说这个类实现了该接口。
*/
type IServer interface {
	Stop()
	Start()
	Serve()
}

type Server struct {
	Name string
	IP   string
	Port int
}

// 实现了IServer接口所有的方法
func (s *Server) Stop() {
	fmt.Println("[STOP] server , name ", s.Name)
}

func (s *Server) Start() {
	fmt.Println("[START] server , name ", s.Name)
}

func (s *Server) Serve() {
	fmt.Println("[Serve] server , name ", s.Name)
}

// 返回IServer类型 进行接口约束，必须实现接口中所有的方法  （推荐）
// 如果IServer其中1项方法未实现，则报错
func NewServer() IServer {
	server := &Server{
		Name: "OOO",
		IP:   "0.0.0.0",
		Port: 5173,
	}
	return server
}

// 构造函数 返回*Server 结构体指针类型 （推荐）
func NewServer2() *Server {
	server := &Server{
		Name: "OOO",
		IP:   "0.0.0.0",
		Port: 5173,
	}
	return server
}

// 不推荐使用此方式初始化
//
//	因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func NewServer3() Server {
	server := Server{
		Name: "OOO",
		IP:   "0.0.0.0",
		Port: 5173,
	}
	return server
}

func main() {
	server := &Server{
		Name: "OOO",
		IP:   "0.0.0.0",
		Port: 5173,
	}
	server.Start()
}
