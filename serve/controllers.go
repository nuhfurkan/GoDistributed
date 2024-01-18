package serve

import (
	"encoding/json"
	"fmt"
	mpiserver "go-distributed/mpi_server"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDirectory = "./uploads/"

func addSlaves(w http.ResponseWriter, r *http.Request) {
	var new_connections []mpiserver.Connection
	err := json.NewDecoder(r.Body).Decode(&new_connections)
	if err != nil {
		log.Println("Error decoding JSON:", err)
	}

	for new_conn := range new_connections {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", new_connections[new_conn].Ip, new_connections[new_conn].Port))
		if err != nil {
			fmt.Println("Error connecting to the server:", err)
			return
		}
		mpi.Connections = append(mpi.Connections, conn)
	}

	
	fmt.Println("These connections added: ")
	for conn := range new_connections {
		fmt.Println("Connection", conn)
		fmt.Println(new_connections[conn].Ip)
		fmt.Println(new_connections[conn].Port)
	}
	
}

func setupWorkers(w http.ResponseWriter, r *http.Request) {
	var configs SetupConfig
	err := json.NewDecoder(r.Body).Decode(&configs)
	if err != nil {
		log.Println("Error decoding JSON:", err)
	}


	for conn := range mpi.Connections {
		mpi.SendFile(mpi.Connections[conn], configs.Filename, 1024)
	} 
}

func runExperiment(w http.ResponseWriter, r *http.Request) {
	var experiment Experiment
	err := json.NewDecoder(r.Body).Decode(&experiment)
	if err != nil {
		log.Println("Error decoding JSON:", err)
	}

	if len(mpi.Connections) == 0 {
		w.WriteHeader(http.StatusForbidden)
	    w.Write([]byte("No live slave found!"))
		return
	}

	createExperiment(experiment)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %s", err), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving the file: %s", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the uploads directory if it doesn't exist
	if _, err := os.Stat(uploadDirectory); os.IsNotExist(err) {
		os.Mkdir(uploadDirectory, os.ModePerm)
	}

	// Create a unique filename
	filename := handler.Filename
	if filename != "" {
		filename = filepath.Join(uploadDirectory, filename)
	} else {
		http.Error(w, "Empty filename received", http.StatusBadRequest)
		return
	}

	// Create the file on the server
	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating the file: %s", err), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Copy the file data into the created file
	_, err = io.Copy(outputFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error copying file data: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s uploaded successfully.", handler.Filename)
}