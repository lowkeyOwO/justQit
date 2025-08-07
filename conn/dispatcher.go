package conn

import (
	"fmt"
	"justQit/types"
	"net/http"
	"strconv"
)

func Initialize(port int16, config types.DispatcherConfig) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from handler!")
	})
	addr := ":" + strconv.Itoa(int(port))
	fmt.Println("Listening on port " + addr)
	http.ListenAndServe(addr, nil)
}
