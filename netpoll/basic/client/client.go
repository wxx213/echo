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
	// reading
	buf, _ := reader.Next(0)
	// parse the read data
	fmt.Println(buf)
	reader.Release()

	// writing
	var write_data []byte
	// make the write data
	alloc, _ := writer.Malloc(len(write_data))
	copy(alloc, write_data) // write data
	writer.Flush()
}
