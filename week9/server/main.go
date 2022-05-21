package main

import (
	"fmt"
	"log"
	"net"
	"week9/delimiter_based"
	"week9/fix_length"
	"week9/length_field_based_frame_decoder"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			return
		}
	}(listen)
	Logger("waiting...")

	for {
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("Conn error", err)
			return
		}
		Logger(coon.RemoteAddr().String(), "tcp connected")
		go fix_length.TcpServerFixLength(coon)
		go delimiter_based.TcpServerDelimiter(coon)
		go length_field_based_frame_decoder.TcpServerFrameDecoder(coon)
	}
}

func Logger(v ...interface{}) {
	log.Println(v...)
}
