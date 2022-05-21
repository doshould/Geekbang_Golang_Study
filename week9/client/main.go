package main

import (
	"fmt"
	"net"
	"os"
	"week9/delimiter_based"
	"week9/fix_length"
	"week9/length_field_based_frame_decoder"
)

func main() {
	addr := "127.0.0.1:10010"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	coon, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer coon.Close()
	fmt.Println("connected...")
	fix_length.TcpClientFixLength(coon)
	delimiter_based.TcpClientDelimiter(coon)
	length_field_based_frame_decoder.TcpClientFrameDecoder(coon)
}
