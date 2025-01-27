package api_versions

import (
	"encoding/binary"
	"unsafe"
)

type Response struct {
	MessageSize   int32
	CorrelationId int32
	ErrorCode     int16
	Body Body
}

type Body struct {
	
}

func (r Response) ToByte() (b []byte) {
	b = make([]byte, unsafe.Sizeof(r))
	buf := make([]byte, 8)

	// set CorrelationId
	binary.BigEndian.PutUint32(buf, uint32(r.CorrelationId))
	for i := 0; i < 4; i++ {
		b[4 + i] = buf[i]
	}

	// set ErrorCode
	binary.BigEndian.PutUint16(buf, uint16(r.ErrorCode))
	for i := 0; i < 2; i++ {
		b[8 + i] = buf[i]
	}
	return
}