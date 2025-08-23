package inmemory

import (
	"net/http"
	"fmt"
)

func (inmemory *InMemoryDispatcher) Ack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ack called")
}