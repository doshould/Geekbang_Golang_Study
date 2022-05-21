package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	PackageLengthBytes = 4
	HeaderLengthBytes  = 2
	VersionBytes       = 2
	OperationBytes     = 4
	SequenceIDBytes    = 4
	HeaderLength       = PackageLengthBytes + HeaderLengthBytes + VersionBytes + OperationBytes + SequenceIDBytes
)

func DePack(buffer []byte) []byte {
	length := len(buffer)
	data := make([]byte, 32)
	var i int
	for i = 0; i < length; i++ {
		if length < i+HeaderLength {
			break
		}
		msgLength := ByteToInt(buffer[i : i+PackageLengthBytes])
		if length < i+HeaderLength+msgLength {
			break
		}
		site := i + PackageLengthBytes
		headerLength := ByteToInt(buffer[site : site+HeaderLengthBytes])
		site += HeaderLengthBytes

		protocolVersion := ByteToInt16(buffer[site : site+VersionBytes])
		site += VersionBytes

		operation := ByteToInt(buffer[site : site+OperationBytes])
		site += OperationBytes

		SequenceID := ByteToInt(buffer[site : site+OperationBytes])
		site += OperationBytes

		fmt.Printf("packageLength: %d, headerLength: %d , protocolVersion: %d, operation: %d, sequenceID: %d \n", msgLength, headerLength, protocolVersion, operation, SequenceID)
		data = buffer[i+headerLength : i+headerLength+msgLength]
	}
	if i == length {
		return make([]byte, 0)
	}
	return data
}

func Packet(msg []byte) []byte {
	body := append(Int32ToBytes(len(msg)), Int16ToBytes(0)...)
	body = append(body, Int16ToBytes(8)...)
	body = append(body, Int32ToBytes(99)...)
	body = append(body, Int32ToBytes(10)...)
	body = append(body, msg...)

	return body
}

func ByteToInt(n []byte) int {
	byteBuffer := bytes.NewBuffer(n)
	var x int32
	err := binary.Read(byteBuffer, binary.BigEndian, &x)
	if err != nil {
		return 0
	}
	return int(x)
}

func ByteToInt16(n []byte) int {
	byteBuffer := bytes.NewBuffer(n)
	var x int16
	err := binary.Read(byteBuffer, binary.BigEndian, &x)
	if err != nil {
		return 0
	}
	return int(x)
}

func Int32ToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

func Int16ToBytes(n int) []byte {
	x := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}
