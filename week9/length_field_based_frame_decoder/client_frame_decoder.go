package length_field_based_frame_decoder

import (
	"fmt"
	"math/rand"
	"net"
	"week9/protocol"
)

func TcpClientFrameDecoder(coon net.Conn) {
	fmt.Println("Client: length field based frame decoder")
	for i := 0; i < 10; i++ {
		name := randStringRunes(10)
		sendMsg := "{\"Name\":\"" + name + "20220521\",\"Meta\":\"Gp\",\"Content\":\"message\"}"
		_, err := coon.Write(protocol.Packet([]byte(sendMsg)))
		if err != nil {
			fmt.Println(err, "idx: ", i)
			return
		}
		fmt.Println(sendMsg)
	}
	Logger("Send Over")
}

var lr = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	bt := make([]rune, n)
	for i := range bt {
		bt[i] = lr[rand.Intn(len(lr))]
	}
	return string(bt)
}
