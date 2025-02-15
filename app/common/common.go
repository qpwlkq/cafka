package common

const (
	ApiVersions             = 18
	DescribeTopicPartitions = 75
)

const (
	SUCCESS                     int16 = 0
	UNKNOWN_TOPIC_OR_PARTITION  int16 = 3
	INVALID_REQUEST_API_VERSION int16 = 35
)