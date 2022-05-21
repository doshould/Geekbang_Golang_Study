package length_field_based_frame_decoder

import (
	"fmt"
	"log"
	"net"
	"week9/protocol"
)

func TcpServerFrameDecoder(coon net.Conn) {
	fmt.Println("Server: length field based frame decoder")
	tmpBuffer := make([]byte, 0)
	readCh := make(chan []byte, 4096)
	go reader(readCh)

	buf := make([]byte, 1024)
	for {
		n, err := coon.Read(buf)
		if err != nil {
			Logger(coon.RemoteAddr().String(), ": ", err)
			return
		}
		tmpBuffer = protocol.DePack(append(tmpBuffer, buf[:n]...))
		readCh <- tmpBuffer
	}
}

//get channel data
func reader(readCh chan []byte) {
	for {
		select {
		case data := <-readCh:
			Logger("Channel:", string(data)) //print channel data
		}
	}
}

func Logger(v ...interface{}) {
	log.Println(v...)
}
