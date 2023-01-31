package main

import (
	"fmt"
	"log"
	"os"
)

type MyReader struct {
}

func (MyReader *MyReader) Read(b []byte) (int, error) {
	fmt.Print("in < ")
	//从标准输入读取数据
	return os.Stdin.Read(b)
}

type MyWrite struct {
}

func (MyWrite *MyWrite) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	//将数据写入标准输出
	return os.Stdout.Write(b)
}

var (
	reader MyReader
	write  MyWrite
)

func main() {
	//缓冲区
	input := make([]byte, 4096)

	//从terminal读取
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data")
	}
	fmt.Printf("read %d bytes from stdin\n", s)

	//写入
	s, err = write.Write(input)
	if err != nil {
		log.Fatalln("unable to write data")
	}
	fmt.Printf("write %d bytes to stout\n", s)

}
