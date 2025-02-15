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
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

func (r Response) ToByte() (b []byte) {
	buf := make([]byte, 26)
	binary.BigEndian.PutUint32(buf, uint32(r.Header.CorrelationId))
	binary.BigEndian.PutUint16(buf[4:], uint16(r.Body.ErrorCode))
	apiKeyCount := len(r.Body.ApiKeys)
	buf[6] = byte(apiKeyCount + 1)
	fmt.Println("apiKeyCount: ", apiKeyCount)
	for i := 0; i < apiKeyCount; i++ {
		binary.BigEndian.PutUint16(buf[7+7*i:], uint16(r.Body.ApiKeys[i].ApiKey))
		binary.BigEndian.PutUint16(buf[9+7*i:], uint16(r.Body.ApiKeys[i].MinVersion))
		binary.BigEndian.PutUint16(buf[11+7*i:], uint16(r.Body.ApiKeys[i].MaxVersion))
		buf[13+7*i] = 0
	}
	fmt.Println("ThrottleTimeMs: ", r.Body.ThrottleTimeMs)
	binary.BigEndian.PutUint32(buf[14+7*(apiKeyCount-1):], uint32(r.Body.ThrottleTimeMs))
	buf[18+7*(apiKeyCount-1)] = 0

	fmt.Println("buf:", buf)
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(len(buf)))
	b = append(b, buf...)
	fmt.Println("b:", b)
	return
}
