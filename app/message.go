package main

type Message struct {
	message_size int32
	header       HeaderV0
	body         Body
}

type Request struct {
	message_size int32
	header       HeaderV2
	body         Body
}

type HeaderV0 struct {
	correlation_id int32
}

type HeaderV2 struct {
	request_api_key     int16
	request_api_version int16
	correlation_id      int32
	client_id           string
}

type MockSlice struct {
	array uintptr
	len   int
	cap   int
}

type Body struct {
}
