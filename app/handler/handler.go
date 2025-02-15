package handler

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/common"
	"github.com/codecrafters-io/kafka-starter-go/app/model"
	"github.com/codecrafters-io/kafka-starter-go/app/model/api_versions"
)

func Handle(request model.Request) ([]byte, error) {
	fmt.Println("API VERSION KEY:", request.Header.RequestApiKey)
	switch request.Header.RequestApiKey {
	case common.ApiVersions:
		return api_versions.Handle(request)
	}
	return []byte{}, errors.New("Unrecognized API")
}
