package api_versions

import (
	"encoding/binary"
	"fmt"
)

type Response struct {
	MessageSize int32
	Header      Header
	Body        Body
}

type Header struct {
	CorrelationId int32
}

type Body struct {
	ErrorCode      int16
	ApiKeys        []ApiKey
	ThrottleTimeMs int32
}

type ApiKey struct {
	ApiKey    int16
	MinVersion int16
	MaxVersion int16
}

func (r Response) ToByte() (b []byte) {
	buf := make([]byte, 19)
	binary.BigEndian.PutUint32(buf, uint32(r.Header.CorrelationId))
	binary.BigEndian.PutUint16(buf[4:], uint16(r.Body.ErrorCode))
	buf[6] = byte(len(r.Body.ApiKeys))
	binary.BigEndian.PutUint16(buf[7:], uint16(r.Body.ApiKeys[0].ApiKey))
	binary.BigEndian.PutUint16(buf[9:], uint16(r.Body.ApiKeys[0].MinVersion))
	binary.BigEndian.PutUint16(buf[11:], uint16(r.Body.ApiKeys[0].MaxVersion))
	buf[13] = 0
	binary.BigEndian.PutUint16(buf[14:], uint16(r.Body.ThrottleTimeMs))
	buf[18] = 0
	fmt.Println("buf:", buf)
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(len(buf)))
	b = append(b, buf...)
	fmt.Println("b:", b)
	return
}
