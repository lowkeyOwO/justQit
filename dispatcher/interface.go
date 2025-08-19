package dispatcher

import (
	"justQit/types"
	"net/http"
)

type Dispatcher interface {
	Initialize(types.DispatcherConfig)
	Enqueue(http.ResponseWriter, *http.Request)
	JobASAP(http.ResponseWriter, *http.Request)
	Ack(http.ResponseWriter, *http.Request)
}

