package types

type AckPayload struct {
	JobId string
	Ack bool
	Log string
}