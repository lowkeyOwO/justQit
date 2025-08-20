package dispatcher

import "justQit/dispatcher/inmemory"


func NewDispatcher(dispatcherType string) Dispatcher {
    switch dispatcherType {
    case "in-memory":
        return &inmemory.InMemoryDispatcher{}
    default:
        return nil
    }
}