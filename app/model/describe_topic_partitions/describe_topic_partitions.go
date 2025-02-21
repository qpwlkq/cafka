package api_versions

import (
	"encoding/binary"
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/common"
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
	Topics []Topic
	NextCursor NextCursor
}

type Topic struct {
	ErrorCode int16
	Name common.COMPACT_NULLABLE_STRING 
	TopicId common.UUID
	IsInternal common.BOOLEAN
	Partitions []Partition
	TopicAuthorizedOperations int32
}

type Partition struct {
	ErrorCode int16
	PartitionIndex int32
	LeaderId int32
	LeaderEpoch int32
	ReplicaNodes int32
	IsrNodes int32
	EligibleLeaderReplicas int32
	LastKnownElr int32
	OfflineReplicas int32
}

type NextCursor struct {
	TopicName common.COMPACT_STRING
	PartitionIndex int32
}

func (r Response) ToByte() (b []byte) {
	buf := make([]byte, 20)
	binary.BigEndian.PutUint32(buf, uint32(r.Header.CorrelationId))
	binary.BigEndian.PutUint32(buf[4:], uint32(r.Body.ThrottleTimeMs))

	buf[8] = byte(len(r.Body.Topics) + 1)
	for i := 0; i < len(r.Body.Topics); i++ {
		buf = append(buf, r.Body.Topics[i].ToByte()...)
	}
	return
}

func (t Topic) ToByte() (b []byte) {
	buf := make([]byte, 10)
	binary.BigEndian.PutUint32(buf, uint32(t.ErrorCode))
}

func Handle(request model.Request) ([]byte, error) {
	fmt.Println("correlationId:", request.Header.CorrelationId, " ", request.Header.RequestApiVersion)
	var apiVersionResponse Response
	if request.Header.RequestApiVersion != 0 {
		apiVersionResponse = // The `Response` struct in the code defines the structure of the response that
		// will be sent back by the API endpoint handling the request. It contains three
		// main fields:
		Response{
			Header: Header{
				CorrelationId: request.Header.CorrelationId,
			},
			Body: Body{
				ThrottleTimeMs: 666,
				Topics: []Topic{
					{
						ErrorCode: common.UNKNOWN_TOPIC_OR_PARTITION,
						Name: "topic 123",
						TopicId: "00000000-0000-0000-0000-000000000000",
						Partitions: []Partition{},
					},
				},
				NextCursor: NextCursor{},
			},
		}
	} else {
		
	}
	fmt.Println("Handle ApiVersions api")
	return apiVersionResponse.ToByte(), nil
}
