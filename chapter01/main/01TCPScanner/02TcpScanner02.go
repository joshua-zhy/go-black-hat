package main

import (
	"fmt"
	"net"
	"time"
)

const (
	IP2 = "127.0.0.1"
	//POSTS2 = 65535

	// IP2    = "154.38.91.254"
	POSTS2 = 65535
)

// goroutine work pool
func workPool(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", IP2, p)
		// return conn ignore
		conn, err := net.DialTimeout("tcp", address, time.Millisecond*200) //timeout 0.2s
		if err != nil {
			results <- -1 //closed port
			continue
		}
		conn.Close()
		results <- p //open port
	}
}

func TCPScanner() {
	ports := make(chan int, 100)
	results := make(chan int)

	defer close(ports)
	defer close(results)

	//channel是安全的 没有数据会阻塞
	for i := 0; i < cap(ports); i++ {
		go workPool(ports, results)
	}

	//存入数据
	go func() {
		for i := 0; i < POSTS2; i++ {
			ports <- i
		}
	}()

	//读取数据
	for p := range results { //for range chan 要求chan 是关闭的 不然会一直读取数据 造成死锁
		if p != -1 {
			fmt.Printf("%s:%d open\n", IP2, p)
		}
	}
}

func main() {
	TCPScanner()
}
