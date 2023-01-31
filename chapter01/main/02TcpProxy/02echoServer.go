package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo1(conn net.Conn) {
	defer conn.Close()
	//创建＿个缓冲区来存储接收到的数据
	b := make([]byte, 1024)
	for {
		//通过conn.Read接收数据到-个缓冲区
		size, err := conn.Read(b[0:])
		if err == io.EOF { //读到末尾了
			log.Println("client disconnected")
			break
		}
		if err != nil {
			log.Println("upexpected error")
			break
		}
		log.Printf("received %d bytes:%s\n,", size, string(b))

		//通过conn.write发送数据
		log.Println("writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("unable to write data")
		}

	}
}

func echo2(conn net.Conn) {
	defer conn.Close()
	// 见O）初始化新的带缓冲的Reader和
	// W∏ter。这些调用都以现有的ReadeT和Wnter作为参数（记住’Com类型实现了＿些必
	// 要的功能’使其能同时被视为Reader和W∏ter）

	reader := bufio.NewReader(conn) 
	s,err := reader.ReadString('\n');
	if err != nil {
		log.Println("unable to read data")
	}
	log.Println("writing data")

	write := bufio.NewWriter(conn) 
	if _,err := write.WriteString(s);err != nil {
		log.Fatalln("unable to write data")
	}
	write.Flush()
}

func main() {
	//在所有接口上绑定TCP端口8888
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("unable to bind port")
	}

	log.Println("listening on 0.0.0.0:8888")

	for {
		//等待连接。在已建立的连接上net.conn
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		log.Println("received connection")

		go echo2(conn)
	}
}
