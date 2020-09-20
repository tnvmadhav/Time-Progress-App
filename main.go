package main

import (
	"fmt"
	"time"
	"math"
	"net/http"
	"os"
)

// Long running data producer
func Updated(updateChan chan float64) float64 {
	for {
		updateChan <- IsProgressed()
		time.Sleep(time.Hour * 8)
	}
}

// Return the right data
func IsProgressed() float64 {
	for {
		answer := GetTimeProgress()
		fmt.Println(answer)
		latency := answer - math.Floor(answer)
		// WHen to return the data...
		if latency > 0 && latency < 0.000004 {
			return answer
		}
		time.Sleep(time.Second * 1)
	}
}

// Just give me the damn number!
func GetTimeProgress() float64 {
	time_year_start := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC) //Time progressed in Hours
	time_year_end := time.Date(time.Now().Year() + 1 , 1, 1, 0, 0, 0, 0, time.UTC) // Get total time in hours
	return (float64(time.Since(time_year_start).Round(time.Second))/ float64(time_year_end.Sub(time_year_start).Round(time.Second))) * 100 // Year Progress
}


func main() {
	// Start log...
	fmt.Println("Time Progress App")
	http.HandleFunc("/", test)
	// Web listen for status checks
	go http.ListenAndServe(getPort(), nil)
	c := make(chan float64)
	// Start a routine in parallel with following statements
	go Updated(c)
	// Get the result from the routine
	answer := <- c
	// Tweet the Result
	fmt.Println(tweet(fmt.Sprintf("The Year has progressed: %d%", int(answer))))
}

// Deployment server's exposed port
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":5000"
}

// For remote checks
func test(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Pong")
}