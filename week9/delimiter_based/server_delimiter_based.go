package delimiter_based

import (
	"bufio"
	"fmt"
	"net"
)

func TcpServerDelimiter(coon net.Conn) {
	fmt.Println("Server: delimiter based")
	reader := bufio.NewReader(coon)
	for {
		slice, err := reader.ReadSlice('\n')
		Logger(slice)
		if err != nil {
			Logger("delimiter based: ", err)
			return
		}
		fmt.Printf("slice: %s", slice)
	}
}
