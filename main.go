package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"
	// "time"
	// "github.com/gosuri/uilive"
)

var paused bool = true

var Webclock string = `<div id="clock" class="flex flex-row items-center"
      hx-trigger="every 1s"
      hx-get="/webclock"
      hx-swap="this">
      %s
    </div>
`
var on string = `<div class="h-12 w-12 text-bold rounded-full bg-lime-300"></div>`
var off string = `<div class="h-12 w-12 text-bold rounded-full bg-green-700"></div>`
var delimiter string = `<div class="h-12 w-12 text-bold rounded-full bg-red-700"></div>`

func greet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func styles(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "./output.css")
}

func toggle(w http.ResponseWriter, r *http.Request) {
	if paused {
		webclock(w, r)
		paused = false
	} else {
		fmt.Fprintf(w, `<div id="clock"><h1 class="text-3xl font-bold text-lime-500">Paused.</h1></div>`)
		paused = true
	}
}

func webclock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Webclock, bclock())
}

func bclock() string {
	clock := time.Now()
	seconds := []byte(fmt.Sprintf("%b", clock.Second()))
	minutes := []byte(fmt.Sprintf("%b", clock.Minute()))
	hours := []byte(fmt.Sprintf("%b", clock.Hour()))
	d := []byte(":")
	bclock := slices.Concat(hours, d, minutes, d, seconds)
	output := ``
	for _, v := range bclock {
		if v == '1' {
			output += on
		} else if v == '0' {
			output += off
		} else {
			output += delimiter
		}
	}
	return output
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/output.css", styles)
	http.HandleFunc("/toggle", toggle)
	http.HandleFunc("/webclock", webclock)
	log.Println("Starting server on port :8080")
	http.ListenAndServe(":8080", nil)
}
