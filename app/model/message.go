package model

type Message struct {
	MessageSize int32
	Header       HeaderV0
	Body         Body
}

type Request struct {
	MessageSize int32
	Header       HeaderV2
	Body         Body
}

type HeaderV0 struct {
	CorrelationId int32
}

type HeaderV2 struct {
	RequestApiKey     int16
	RequestApiVersion int16
	CorrelationId      int32
	ClientId           string
}

type MockSlice struct {
	Array uintptr
	Len   int
	Cap   int
}

type Body struct {
}
