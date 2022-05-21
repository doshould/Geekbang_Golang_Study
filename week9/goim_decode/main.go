package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := encode("Hello, Golang!")
	decode(data)
}

/*
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
*/

func decode(data []byte) {
	if len(data) <= 16 {
		fmt.Println("data len < 16")
		return
	}

	_packSize := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("_packSize:%v\n", _packSize)

	_headerSize := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("_headerSize:%v\n", _headerSize)

	_verSize := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("_verSize:%v\n", _verSize)

	_opSize := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("_opSize:%v\n", _opSize)

	_seqSize := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("_seqSize:%v\n", _seqSize)

	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
}

func encode(body string) []byte {
	_headerSize := 16
	_packSize := len(body) + _headerSize
	res := make([]byte, _packSize)

	binary.BigEndian.PutUint32(res[:4], uint32(_packSize))
	binary.BigEndian.PutUint16(res[4:6], uint16(_headerSize))
	_verSize := 5
	binary.BigEndian.PutUint16(res[6:8], uint16(_verSize))
	_opSize := 6
	binary.BigEndian.PutUint32(res[8:12], uint32(_opSize))
	_seqSize := 7
	binary.BigEndian.PutUint32(res[12:16], uint32(_seqSize))

	bodyByte := []byte(body)
	copy(res[16:], bodyByte)
	return res
}
