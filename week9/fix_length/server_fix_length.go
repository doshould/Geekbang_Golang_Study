package fix_length

import (
	"fmt"
	"net"
)

func TcpServerFixLength(server net.Conn) {
	fmt.Println("Server: fix length")
	for {
		buf := make([]byte, 1024)
		_, err := server.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Client data:", string(buf))
	}
}
