package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"strconv"
)

type SortPayload struct {
	ToSort [][]int `json:"to_sort"`
}

type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}
func parseRequest(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// sendJSONResponse sends a JSON response to the client
func sendJSONResponse(w http.ResponseWriter, sortedArrays [][]int, duration time.Duration) {
	response := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       duration.Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(response)
}
func main() {

	mux := http.NewServeMux()
	
	// Define the handlers for the endpoints
	mux.HandleFunc("/",mainpageHandler)
	mux.HandleFunc("/process-single", processSingleHandler)
	mux.HandleFunc("/process-concurrent", processConcurrentHandler)

	// Get the port from environment variables or use a default value
	port := getPort()

	// Create a custom logger
	logger := log.New(os.Stdout, "[GoServer] ", log.LstdFlags|log.Lmicroseconds)

	// Create a server with timeouts and the custom logger
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      loggingMiddleware(mux, logger),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Handle graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		<-sigint

		logger.Println("Shutting down server...")

		// Allow 30 seconds for existing connections to finish processing
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Printf("Error during server shutdown: %v\n", err)
		}

		logger.Println("Server gracefully stopped")

		// Exit the program
		os.Exit(0)
	}()

	logger.Printf("Server is listening on :%d...\n", port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Error starting server: %v\n", err)
	}
}

func mainpageHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Server is running Successfully")
	log.Println("Server is running Successfully")
}

// processSingleHandler handles the /process-single endpoint for sequential processing
func processSingleHandler(w http.ResponseWriter, r *http.Request) {
	var payload SortPayload
	if err := parseRequest(r, &payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	// Sequentially sort each sub-array using Quick Sort
	for i, subArray := range payload.ToSort {
		quickSort(subArray)
		payload.ToSort[i] = subArray
	}

	duration := time.Since(startTime)

	// Send the sorted result as JSON response
	sendJSONResponse(w, payload.ToSort, duration)

}

// processConcurrentHandler handles the /process-concurrent endpoint for concurrent processing
func processConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request to /process-concurrent with method %s\n", r.Method)
	var payload SortPayload
	if err := parseRequest(r, &payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	// Simulate concurrent processing using goroutines and channels
	var wg sync.WaitGroup
	resultCh := make(chan []int, len(payload.ToSort))

	for _, subArray := range payload.ToSort {
		wg.Add(1)
		go func(arr []int) {
			defer wg.Done()
			quickSort(arr)
			resultCh <- arr
		}(subArray)
	}

	// Close the channel when all goroutines finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect results from the channel
	var sortedArrays [][]int
	for sortedArray := range resultCh {
		sortedArrays = append(sortedArrays, sortedArray)
	}

	duration := time.Since(startTime)

	// Send the sorted result as JSON response
	sendJSONResponse(w, sortedArrays, duration)
}

func quickSort(arr []int)error {
	if len(arr) <= 1 {
		return nil
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quickSort(arr[:right+1])
	quickSort(arr[left:])
	return nil
}

func getPort() int {
	port := 8000
    if envPort, ok := os.LookupEnv("PORT"); ok {
        if p, err := strconv.Atoi(envPort); err == nil {
            port = p
        } else {
            log.Printf("Error converting PORT to integer: %v\n", err)
        }
    }
    return port
}

// loggingMiddleware is a middleware that logs incoming requests
func loggingMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			logger.Printf("[%s] %s %s %v\n", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
		}()

		next.ServeHTTP(w, r)
	})
}
