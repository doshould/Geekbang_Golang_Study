package delimiter_based

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func TcpClientDelimiter(coon net.Conn) {
	var buffer bytes.Buffer
	fmt.Println("Client: delimiter based")
	sendMsg := "{\"msg01\":1, \"msg02\":2}\n"
	for i := 0; i < 10; i++ {
		buffer.WriteString(sendMsg)
		s, err := coon.Write([]byte(buffer.String()))
		Logger(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		Logger("Send Once Over")
		//time.Sleep(2 * time.Second)
	}
}

func Logger(v ...interface{}) {
	log.Println(v...)
}
