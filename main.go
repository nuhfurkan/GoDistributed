package main

import (
	"flag"
	mpiserver "go-distributed/mpi_server"
	serve "go-distributed/serve"
)

func main() {
	var isSlave int
	var ip string
	var port int 

	flag.IntVar(&isSlave, "rank", 0, "Rank of the process")
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address of the process")
	flag.IntVar(&port, "port", 3000, "Port number for communication")

	flag.Parse()

	if isSlave == 1 {
		mpiserver.NewMPI(1, nil).StartWorker(ip, port)
	} else {
		serve.ServeDashboard()
	}
}