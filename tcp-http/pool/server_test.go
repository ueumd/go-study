package pool

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"time"
)

// 简单定义两个结构用于客户端与服务端交互，传输协议用json示范

type Message struct {
	Uid string
	Val string
}

type Resp struct {
	Uid string
	Val string
	Ts  string
}

const TAG = "server: hello."

func ListenAndServer() {
	log.Print("Start Server...")

	// 启动监听
	listen, err := net.Listen("tcp", "0.0.0.0:5173")
	if err != nil {
		log.Fatal("Listen failed. msg: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept failed, err: %v", err)
			continue
		}

		go transfer(conn)
	}
}

func transfer(conn net.Conn) {
	defer func() {
		remoteAddr := conn.RemoteAddr().String()
		log.Print("discard remove add: ", remoteAddr)

		conn.Close()
	}()

	// 设置10秒关闭连接
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	for {
		var msg Message

		if err := json.NewDecoder(conn).Decode(&msg); err != nil && err != io.EOF {
			log.Printf("Decode from client err: %v", err)
			// todo... 仿照redis协议写入err前缀符号`-`，通知client错误处理
			return
		}

		if msg.Uid != "" || msg.Val != "" {
			var rsp Resp
			rsp.Uid = msg.Uid
			rsp.Val = TAG + msg.Val

			ser, _ := json.Marshal(msg)

			// 模拟服务端耗时
			time.Sleep(2 * time.Second)

			conn.Write(append(ser, '\n'))
		}
	}
}
