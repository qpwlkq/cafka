package handler

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/model"
	"github.com/codecrafters-io/kafka-starter-go/app/model/api_versions"
)

const (
	ApiVersions = 18
)

func Handle(request model.Request) ([]byte, error) {
	fmt.Println("API VERSION KEY:", request.Header.RequestApiKey)
	switch request.Header.RequestApiKey {
	case ApiVersions:
		apiVersionResponse := api_versions.Response{
			MessageSize:   0,
			CorrelationId: request.Header.CorrelationId,
			ErrorCode:     35,
		}
		fmt.Println("Handle ApiVersions api")
		return apiVersionResponse.ToByte(), nil
	}
	return []byte{}, errors.New("Unrecognized API")
}
