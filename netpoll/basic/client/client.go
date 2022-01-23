package main

import (
	"fmt"
	"github.com/cloudwego/netpoll"
	"time"
)

func main() {
	// Dial a connection with Dialer.
	dialer := netpoll.NewDialer()
	conn, err := dialer.DialConnection("tcp", "0.0.0.0:8888", 10 * time.Second)
	if err != nil {
		panic("dial netpoll connection failed")
	}
	reader, writer := conn.Reader(), conn.Writer()

	// writing
	var write_data = []byte("hello from client") 
	// make the write data
	alloc, _ := writer.Malloc(len(write_data))
	copy(alloc, write_data) // write data
	writer.Flush()

	// reading
	buf, _ := reader.Next(17)
	// parse the read data
	fmt.Println("data from server: " + string(buf))
	reader.Release()
}
