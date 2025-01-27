package api_versions

import (
	"encoding/binary"
	"unsafe"
)

type Response struct {
	// MessageSize   int32
	CorrelationId int32
	ErrorCode     int16
	ApiKeys ApiKeys
	ThrottleTimeMs int32
}

type ApiKeys struct {
	ApiKey int16
	MinVersion int16
	MaxVersion int16
}

func (r Response) ToByte() (b []byte) {
	buf := make([]byte, unsafe.Sizeof(r))

	binary.BigEndian.PutUint32(buf, uint32(r.CorrelationId))
	binary.BigEndian.PutUint16(buf[4:], uint16(r.ErrorCode))
	binary.BigEndian.PutUint16(buf[6:], uint16(r.ApiKeys.ApiKey))
	binary.BigEndian.PutUint16(buf[8:], uint16(r.ApiKeys.MinVersion))
	binary.BigEndian.PutUint16(buf[10:], uint16(r.ApiKeys.MaxVersion))
	binary.BigEndian.PutUint32(buf[12:], uint32(r.ThrottleTimeMs))

	return buf
}