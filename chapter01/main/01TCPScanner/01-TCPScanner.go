package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	//IP = "154.38.91.254"
	IP = "127.0.0.1"
	//POSTS = 65535
	POSTS = 65535
)

var wg sync.WaitGroup

func TcpScannerBasic() {
	for i := 1; i <= POSTS; i++ {
		address := fmt.Sprintf("%s:%d", IP, i)
		// return conn ignore
		_, err := net.DialTimeout("tcp", address, time.Millisecond*100) //timeout 0.1s

		if err == nil {
			fmt.Printf("%s:%d open\n", IP, i)
		}
	}
}
func TcpScannerQuick() {
	for i := 1; i <= POSTS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", IP, i)
			// return conn ignore
			conn, err := net.DialTimeout("tcp", address, time.Millisecond*100) //timeout 0.5s

			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s:%d open\n", IP, i)
		}(i)
	}
}

func main() {
	// TcpScannerBasic()
	TcpScannerQuick()
	wg.Wait()
}
