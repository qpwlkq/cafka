package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"unsafe"

	"github.com/codecrafters-io/kafka-starter-go/app/handler"
	"github.com/codecrafters-io/kafka-starter-go/app/model"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	buf := make([]byte, 40)
	conn.Read(buf)
	fmt.Println("input", buf)
	request := ByteToRequestInBigEndian(buf)

	responseBytes, err := handler.Handle(request)
	if err == nil {
		conn.Write(responseBytes)
		fmt.Println("return")
		return
	}

	// print: 00000023001200046f7fc66100096b61666b612d636c69000a6b61666b612d636c6904302e310000
	// fmt.Println(hex.EncodeToString(buf))

	// print: 2300000000000000001261c661c67f6f000000000000000000000000000000000000000000000000
	// !!!!!!! 计算机存储使用小端序, 但是 tcp网络传输使用大端序!
	// fmt.Println(hex.EncodeToString(RequestToByte(&request)))

	response := model.Message{
		Header: model.HeaderV0{
			CorrelationId: request.Header.CorrelationId,
		},
	}
	response_byte := MessageToByteInBigEndian(&response)
	conn.Write(response_byte)

	fmt.Println(hex.EncodeToString(response_byte))
}

// []byte => request by pointer conversion
func ByteToRequest(b []byte) (request model.Request) {
	request = *(*(**model.Request)(unsafe.Pointer(&b)))
	return
}

// []byte => request in bigEndian
func ByteToRequestInBigEndian(b []byte) (request model.Request) {
	request.MessageSize = int32(binary.BigEndian.Uint32(b[0:4]))
	request.Header.RequestApiKey = int16(binary.BigEndian.Uint16(b[4:6]))
	request.Header.RequestApiVersion = int16(binary.BigEndian.Uint16(b[6:8]))
	request.Header.CorrelationId = int32(binary.BigEndian.Uint32(b[8:13]))
	return
}

// message => []byte in bigEndian
func MessageToByteInBigEndian(message *model.Message) (b []byte) {
	b = make([]byte, unsafe.Sizeof(*message))
	correlation_id_size := unsafe.Sizeof(message.Header.CorrelationId)
	fmt.Println(correlation_id_size)     // 4
	fmt.Println(unsafe.Sizeof(*message)) // 12 ????
	buf := make([]byte, correlation_id_size)
	binary.BigEndian.PutUint32(buf, uint32(message.Header.CorrelationId))
	for i := 0; i < 4; i++ {
		b[4+i] = buf[i]
	}
	return
}

// message => byte by pointer conversion
func MessageToByte(message_ptr *model.Message) (b []byte) {
	len := unsafe.Sizeof(*message_ptr)
	mock_slice := &model.MockSlice{
		Array: uintptr(unsafe.Pointer(message_ptr)),
		Cap:   int(len),
		Len:   int(len),
	}
	b = *(*[]byte)(unsafe.Pointer(mock_slice))
	return
}

// request => byte by pointer conversion
func RequestToByte(request_ptr *model.Request) (b []byte) {
	len := unsafe.Sizeof(*request_ptr)
	mock_slice := &model.MockSlice{
		Array: uintptr(unsafe.Pointer(request_ptr)),
		Cap:   int(len),
		Len:   int(len),
	}
	b = *(*[]byte)(unsafe.Pointer(mock_slice))
	return
}
