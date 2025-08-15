package dispatcher

import "justQit/dispatcher/inmemory"

var reqCounter uint64 = 0;

func NewDispatcher(dispatcherType string) Dispatcher {
    switch dispatcherType {
    case "in-memory":
        return &inmemory.InMemoryDispatcher{}
    default:
        return nil
    }
}