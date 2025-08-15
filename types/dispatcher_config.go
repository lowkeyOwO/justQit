package types

type customFieldMap struct {
	JobId string
	Priority string
}

type DispatcherConfig struct {
	Priority []int
	QueueSize []int
	MaxWorkers int
	MaxDispatch int
	CustomFields bool
	FieldMap customFieldMap 
}