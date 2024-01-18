package serve

import (
	"fmt"
	mpiserver "go-distributed/mpi_server"
	"log"
	"net"
	"net/http"
)

/**
 *	The server for dashboard will run on localhost:5000
 *	The api gates are as follows:
 *
 */

var mpi = mpiserver.NewMPI(0, []net.Conn{})

func ServeDashboard() {
	fmt.Printf("Serve dashboard started\n")
	
	http.HandleFunc("/run", runExperiment)
	http.HandleFunc("/add-slaves", addSlaves)
	http.HandleFunc("/upload", handleFileUpload)
	http.HandleFunc("/setup-workers", setupWorkers)
	
	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}