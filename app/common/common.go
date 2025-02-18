package common

import "github.com/codecrafters-io/kafka-starter-go/app/util"

const (
	ApiVersions             = 18
	DescribeTopicPartitions = 75
)

const (
	SUCCESS                     int16 = 0
	UNKNOWN_TOPIC_OR_PARTITION  int16 = 3
	INVALID_REQUEST_API_VERSION int16 = 35
)

type COMPACT_STRING string
type COMPACT_NULLABLE_STRING string
type UUID string
type BOOLEAN bool

func (s COMPACT_STRING) ToByte() []byte {
	return []byte{}
}

type UNSIGNED_VARINT int32

func (i UNSIGNED_VARINT) ToByte() []byte {
	result := []byte{}
	result = append(result, byte(i&0b1111111))
	i = i >> 7
	for i > 0 {
		tmp := i&0b1111111 | 0b10000000
		result = append(result, byte(tmp))
		i = i >> 7
	}
	return util.Reverse(result)
}