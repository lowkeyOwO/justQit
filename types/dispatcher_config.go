package types

type CustomFieldMap struct {
	JobId string
	GroupId string
	Priority string
}

type DispatcherConfig struct {
	Priority []int
	QueueSize []int
	MaxWorkers int
	MaxDispatch int
	CustomFields bool
	FieldMap CustomFieldMap 
}