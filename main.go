package main

import (
	"fmt"
	"time"
	"math"
	"net/http"
	"os"
)

func Updated(updateChan chan float64) float64 {
	for {
		updateChan <- IsProgressed()
		time.Sleep(time.Hour * 8)
	}
}

func IsProgressed() float64 {
	for {
		answer := GetTimeProgress()
		fmt.Println(answer)
		latency := answer - math.Floor(answer)
		if latency > 0 && latency < 0.000004 {
			return answer
		}
		time.Sleep(time.Second * 1)
	}
}



func GetTimeProgress() float64 {
	time_year_start := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC) //Time progressed in Hours
	time_year_end := time.Date(time.Now().Year() + 1 , 1, 1, 0, 0, 0, 0, time.UTC) // Get total time in hours
	return (float64(time.Since(time_year_start).Round(time.Second))/ float64(time_year_end.Sub(time_year_start).Round(time.Second))) * 100 // Year Progress
}

func main() {
	fmt.Println("Year Progress App")
	http.HandleFunc("/", test)
	go http.ListenAndServe(getPort(), nil)
	c := make(chan float64)
	go Updated(c)
	answer := <- c
	fmt.Println(tweet(fmt.Sprintf("The Year has progressed: %d%", answer)))
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func test(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "this is a test")
}