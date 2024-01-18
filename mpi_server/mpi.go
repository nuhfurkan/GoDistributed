package mpiserver

import (
	"bufio"
	"fmt"
	"go-distributed/representations"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type MPI struct {
	Connections []net.Conn
	rank		int
}

func NewMPI(rank int, connections []net.Conn) *MPI {
	return &MPI{
		Connections: 	connections,
		rank:			rank,
	}
}

type Connection struct {
	Ip   string		`json:"ip"`
	Port int		`json:"port"`
}

type WorkChain struct {
    data	chan	representations.Representation
    quit	chan	bool
    stopped	bool
}

type RetrievedResult struct {
	Data 	representations.Representation
	Score	float64
}

func processIncomingData(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		data := scanner.Text()

		fmt.Println(data)
	
		scriptPath, err := os.Executable()
    	if err != nil {
        	fmt.Println("Error getting executable path:", err)
        	return err
    	}

    	scriptPath = scriptPath[:len(scriptPath)-len("./godistributed")] + "/saves/start.sh"

    	cmd := exec.Command(scriptPath, data)

    	output, err := cmd.CombinedOutput()
    	if err != nil {
        	fmt.Println("Error:", err)
    	} else {
			fmt.Println("Output:", string(output))
			response, err := strconv.ParseFloat(string(output), 64)
			if err != nil {
				fmt.Println("Error converting float64", err)
				return err
			} else {
				fmt.Println(response)
				_, err = fmt.Fprintf(conn, "%f\n", response)
				if err != nil {
					fmt.Println("Error sending response:", err)
					return err
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return err
	}

	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("[%s] connected\n", conn.RemoteAddr())

	if err := processIncomingData(conn); err != nil {
		fmt.Println("Error processing data:", err)
	}

	fmt.Printf("[%s] disconnected\n", conn.RemoteAddr())
}

func scanOnce(myconn net.Conn, results chan<- RetrievedResult, data representations.Representation, internal_wg *sync.WaitGroup) {
	scanner := bufio.NewScanner(myconn)
	scanner.Scan()
	my_str := scanner.Text()
	fmt.Println(my_str)
	s, err := strconv.ParseFloat(my_str, 64)
	if err != nil {
		fmt.Println("Can't convert this to an float64!")
	} else {
		results <- RetrievedResult{
			Data: data,
			Score: s,
		}
		fmt.Println("recived:", s, "inside:", myconn.RemoteAddr())
		internal_wg.Done()
	}
}

func worker(conn net.Conn, jobs WorkChain, results chan<- RetrievedResult, wg *sync.WaitGroup) {
	defer wg.Done()

	var internal_wg sync.WaitGroup
	var send_wg sync.WaitGroup
	send_wg.Add(1)

	// Goroutine for sending jobs
	go func(s_wg *sync.WaitGroup) {
		defer s_wg.Done()
		for job := range jobs.data {
			internal_wg.Add(1)
			go scanOnce(conn, results, job, &internal_wg)
			sendJob(conn, job)
			
			internal_wg.Wait()
			fmt.Println("send")
		}
	}(&send_wg)

	send_wg.Wait()
}

func receiveFile(conn net.Conn, saveDirectory string, bufferSize int) error {
	// Read file name and size from the connection
	fileInfoBuffer := make([]byte, bufferSize)
	n, err := conn.Read(fileInfoBuffer)
	if err != nil {
		return err
	}
	fileInfo := string(fileInfoBuffer[:n])
	fmt.Println(fileInfo)
	fileName, fileSizeStr := parseFileInfo(fileInfo)

	

	// Convert the file size string to int
	fileSize, err := strconv.ParseInt(fileSizeStr, 10, 64)
	if err != nil {
		return err
	}

	// Create a file to save the received data
	filePath := filepath.Join(saveDirectory, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the file data from the connection to the file
	_, err = io.CopyN(file, conn, fileSize)
	if err != nil {
		return err
	}

	fmt.Printf("File '%s' received and saved to %s\n", fileName, filePath)
	cmd := exec.Command("unzip", "./saves/"+fileName, "-d", "./saves/")

	// The `Output` method executes the command and
	// collects the output, returning its value
	_, err = cmd.Output()
	if err != nil {
	  // if there was any error, print it here
	  fmt.Println("could not run command: ", err)
	  fmt.Println("In file unzip")
		exec.Command("rm", "./saves/"+fileName).Run()
	} else {
		// if you want to see the out uncomment and define out up
		// fmt.Println(out)
	}

	return nil
}

func parseFileInfo(fileInfo string) (string, string) {
	parts := strings.Split(fileInfo, ",")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	fmt.Println(parts)
	return "", ""
}


func readString(conn net.Conn) (string, error) {
	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func (mpi *MPI) SendFile(conn net.Conn, filePath string, bufferSize int) error {
	// Open the file to be sent
	file, err := os.Open("uploads/"+filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file info for sending
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}


	fmt.Println(fileInfo.Size())

	// Send file name and size to the client
	fileInfoStr := fmt.Sprintf("%s,%d", fileInfo.Name(), fileInfo.Size())
	_, err = conn.Write([]byte(fileInfoStr))
	if err != nil {
		return err
	}

	// Copy the file data to the connection
	_, err = io.CopyBuffer(conn, file, make([]byte, bufferSize))
	if err != nil {
		return err
	}

	fmt.Printf("File '%s' sent successfully\n", fileInfo.Name())
	return nil
}


func sendJob(conn net.Conn, job representations.Representation) {
	// Use a mutex for synchronization to avoid concurrent writes
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Fprintln(conn, 
		strings.Trim(fmt.Sprint(
			job.Get()["value"],
		), "[]"),
	)
	//fmt.Fprintf(conn, "%v\n", encodedData)
}

func waitForSetup(conn net.Conn, wg *sync.WaitGroup) error {
	defer wg.Done()
	err := receiveFile(conn, "saves", 1024)
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return err
	}
	return nil
}

func (mpi *MPI) StartWorker(ip string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		var wg sync.WaitGroup
		wg.Add(1)
		err = waitForSetup(conn, &wg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		wg.Wait()

		go handleConnection(conn)
	}
}

func (mpi *MPI) StartMaster(job_stack SafeStack) []RetrievedResult {
	jobs := WorkChain{
		data:    make(chan representations.Representation, job_stack.Size()),
		quit:    make(chan bool),
		stopped: false,
	}
	results := make(chan RetrievedResult, job_stack.Size())
	var wg sync.WaitGroup

	// Start multiple workers
	for _, conn := range mpi.Connections {
		wg.Add(1)
		go worker(conn, jobs, results, &wg)
	}

	// Send jobs to the workers

	for !job_stack.IsEmpty() {
		job, err := job_stack.Pop()
		if err == nil {
			if assert_job, ok := job.(representations.Representation); ok {
				jobs.data <- assert_job
			} else {
				fmt.Println("could not assert type")
			}
		} else {
			log.Println(err)
		}
	} 

	close(jobs.data)

	// Wait for all jobs to be done
	wg.Wait()

	// Close the results channel to signal that no more results will be sent
	close(results)

	// Collect and print the results
	var max_score1 RetrievedResult
	var max_score2 RetrievedResult  
	for result := range results {
		if result.Score > max_score1.Score {
			max_score2 = max_score1
			max_score1 = result
		} else if result.Score > max_score2.Score {
			max_score2 = result
		}
		fmt.Println("Received result", "\n", result.Data.Show(), "\n", result.Score, "\n", "-------------------")
	}

	return []RetrievedResult{
		max_score1,
		max_score2,
	}
}