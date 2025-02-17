package api_versions

import (
	"encoding/binary"
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/model"
)

/**
DescribeTopicPartitions Response (Version: 0) => throttle_time_ms [topics] next_cursor TAG_BUFFER
  throttle_time_ms => INT32
  topics => error_code name topic_id is_internal [partitions] topic_authorized_operations TAG_BUFFER
    error_code => INT16
    name => COMPACT_NULLABLE_STRING
    topic_id => UUID
    is_internal => BOOLEAN
    partitions => error_code partition_index leader_id leader_epoch [replica_nodes] [isr_nodes] [eligible_leader_replicas] [last_known_elr] [offline_replicas] TAG_BUFFER
      error_code => INT16
      partition_index => INT32
      leader_id => INT32
      leader_epoch => INT32
      replica_nodes => INT32
      isr_nodes => INT32
      eligible_leader_replicas => INT32
      last_known_elr => INT32
      offline_replicas => INT32
    topic_authorized_operations => INT32
  next_cursor => topic_name partition_index TAG_BUFFER
    topic_name => COMPACT_STRING
    partition_index => INT32
**/

type Response struct {
	MessageSize int32
	Header      Header
	Body        Body
}

type Header struct {
	CorrelationId int32
}

type Body struct {
	ThrottleTimeMs int32
	Topic Topic
	NextCursor NextCursor
}

type Topic struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

type NextCursor struct {
	TopicName COMPACT_STRING
	PartitionIndex int32
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

func Handle(request model.Request) ([]byte, error) {
	fmt.Println("correlationId:", request.Header.CorrelationId, " ", request.Header.RequestApiVersion)
	var apiVersionResponse Response
	if request.Header.RequestApiVersion != 0 {
		
	} else {
		
	}
	fmt.Println("Handle ApiVersions api")
	return apiVersionResponse.ToByte(), nil
}
