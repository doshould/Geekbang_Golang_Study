package fix_length

import (
	"fmt"
	"net"
)

func TcpClientFixLength(server net.Conn) {
	fmt.Println("Client: fix length")
	sendByte := make([]byte, 1024)
	sendMessage := "{\"msg01\":1, \"msg02\":2}"
	for i := 0; i < 1000; i++ {
		tmp := []byte(sendMessage)
		for j := 0; j < len(tmp) && j < 1024; j++ {
			sendByte[j] = tmp[j]
		}
		_, err := server.Write(sendByte)
		if err != nil {
			panic(err)
		}
		fmt.Println("Send Over")
	}
}
