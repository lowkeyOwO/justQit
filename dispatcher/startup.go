package dispatcher

import (
	"fmt"
	"justQit/database"
	"justQit/types"
	"net/http"
	"strconv"
)

func StartUp(port int16, config types.DispatcherConfig, dispatcherType string) {
	if config.LogToDatabase {
		database.InitializeLogger(config.LogAfterXRequests)
	}

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
