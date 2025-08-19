package types

type InMemoryDispatcher struct {
	config    DispatcherConfig
	jobqueues []chan string
	jobmap    map[string]string
}
