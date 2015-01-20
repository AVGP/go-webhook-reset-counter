package main

import(
    "fmt"
    "log"
    "flag"
    "time"
    "net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)
type resetFunc func()

func resetAfter(n time.Duration, onReset resetFunc) {
    for _ = range time.Tick(n * time.Second) {
        log.Printf("Resetting...")
        onReset()
    }
}

func main() {
    var is_open bool = false
    var actuated int = 0

    addr := flag.String("addr", ":8080", "The address on which to listen")
    resetTimeout := flag.Int("reset", 60, "The seconds after which the state is reset")
    flag.Parse()

    if *resetTimeout > 0 {
        go resetAfter(time.Duration(*resetTimeout), func() {
            is_open = false
            actuated = 0
        })
    }

    // Read the value
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        fmt.Fprintf(w, "%v", is_open)
    })

    // Statistics
    http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        fmt.Fprintf(w, "%v", actuated)
    })

    // Set the value
    http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        is_open = true
        actuated++
        fmt.Fprintf(w, "%v", is_open)
    })

    log.Printf("Starting server on address %s", *addr)
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatalf("Error setting up server: %v", err)
    }
}