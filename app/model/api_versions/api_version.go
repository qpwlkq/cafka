package api_versions

import (
	"encoding/binary"
	"fmt"
)

type Response struct {
	MessageSize int32
	Header Header
	ErrorCode     int16
	ApiKeys ApiKeys
	ThrottleTimeMs int32
}

type Header struct {
	CorrelationId int32
}

type ApiKeys struct {
	ApiKey int16
	MinVersion int16
	MaxVersion int16
}

func (r Response) ToByte() (b []byte) {
	buf := make([]byte, 19)
	binary.BigEndian.PutUint32(buf, uint32(r.Header.CorrelationId))
	binary.BigEndian.PutUint16(buf[4:], uint16(r.ErrorCode))
	buf[6] = 2
	binary.BigEndian.PutUint16(buf[7:], uint16(r.ApiKeys.ApiKey))
	binary.BigEndian.PutUint16(buf[9:], uint16(r.ApiKeys.MinVersion))
	binary.BigEndian.PutUint16(buf[11:], uint16(r.ApiKeys.MaxVersion))
	buf[13] = 0
	binary.BigEndian.PutUint16(buf[14:], uint16(r.ThrottleTimeMs))
	buf[18] = 0
	fmt.Println("buf:", buf)
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(len(buf)))
	b = append(b, buf...)
	fmt.Println("b:", b)
	return
}