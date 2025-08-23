package inmemory

import (
	"justQit/types"
	"justQit/utils"
)

type InMemoryDispatcher struct {
	config      *types.DispatcherConfig
	jobQueues   []chan string
	jobMap      *utils.SafeMap[string, *types.DBSchema]
	dispatchMap *utils.SafeMap[string, *types.DBSchema]
}

func (inmemory *InMemoryDispatcher) Initialize(config types.DispatcherConfig) {
	inmemory.config = &config
	inmemory.jobQueues = make([]chan string, len(inmemory.config.QueueSize))
	for i := range len(inmemory.jobQueues) {
		queueSize := inmemory.config.QueueSize[i]
		inmemory.jobQueues[i] = make(chan string, queueSize)
	}

	inmemory.jobMap = utils.NewSafeMap[string, *types.DBSchema]()
	inmemory.dispatchMap = utils.NewSafeMap[string, *types.DBSchema]()

	go inmemory.jobRequestHandler()
}
