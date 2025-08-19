package dispatcher

import (
	"fmt"
	"justQit/types"
	"net/http"
	"strconv"
)


func Initialize(port int16, config types.DispatcherConfig, dispatcherType string) {

	dispatchHandler := NewDispatcher(dispatcherType)
	dispatchHandler.Initialize(config)

	// External Methods 
	http.HandleFunc("/dispatcher/enqueue", dispatchHandler.Enqueue) 

	// Internal Methods 
	http.HandleFunc("/dispatcher/ack", dispatchHandler.Ack)
	http.HandleFunc("/dispatcher/jobasap", dispatchHandler.JobASAP)

	addr := ":" + strconv.Itoa(int(port))
	fmt.Println(`
     __                __  ________  .__  __   
    |__|__ __  _______/  |_\_____  \ |__|/  |_ 
    |  |  |  \/  ___/\   __\/  / \  \|  \   __\
    |  |  |  /\___ \  |  | /   \_/.  \  ||  |  
/\__|  |____//____  > |__| \_____\ \_/__||__|  
\______|          \/              \__>` + "\nReady to queue on port" + addr)
	http.ListenAndServe(addr, nil)
}
