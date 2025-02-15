package handler

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/model"
	"github.com/codecrafters-io/kafka-starter-go/app/model/api_versions"
)

const (
	ApiVersions = 18
	DescribeTopicPartitions = 75
)

func Handle(request model.Request) ([]byte, error) {
	fmt.Println("API VERSION KEY:", request.Header.RequestApiKey)
	switch request.Header.RequestApiKey {
	case ApiVersions:
		fmt.Println("correlationId:", request.Header.CorrelationId, " ", request.Header.RequestApiVersion)
		var apiVersionResponse api_versions.Response
		if request.Header.RequestApiVersion < 4 {
			apiVersionResponse = api_versions.Response{
				Header: api_versions.Header{
					CorrelationId: request.Header.CorrelationId,
				},
				Body: api_versions.Body{
					ErrorCode: 35,
					ApiKeys: []api_versions.ApiKey{
						{
							ApiKey:     ApiVersions,
							MinVersion: 3,
							MaxVersion: 4,
						},
						{
							ApiKey: DescribeTopicPartitions,
							MinVersion: 0,
							MaxVersion: 0,
						},
					},
					ThrottleTimeMs: 0,
				},
			}
		} else {
			apiVersionResponse = api_versions.Response{
				Header: api_versions.Header{
					CorrelationId: request.Header.CorrelationId,
				},
				Body: api_versions.Body{
					ErrorCode: 66,
					ApiKeys: []api_versions.ApiKey{
						{
							ApiKey:     18,
							MinVersion: 3,
							MaxVersion: 4,
						},
						{
							ApiKey: DescribeTopicPartitions,
							MinVersion: 0,
							MaxVersion: 0,
						},
					},
					ThrottleTimeMs: 0,
				},
			}
		}
		fmt.Println("Handle ApiVersions api")
		return apiVersionResponse.ToByte(), nil
	}
	return []byte{}, errors.New("Unrecognized API")
}
