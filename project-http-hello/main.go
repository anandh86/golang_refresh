package main

import (
  "net/http"
  "log"
  "fmt"
)

var (
  calls []string
  stats map[string]int
)

func main() {
	// Your solution goes here. Good luck!
  stats = make(map[string]int)

  http.HandleFunc("/hello", funcHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func funcHandler(w http.ResponseWriter, r *http.Request)  {

  name := r.URL.Query().Get("name")

  if name == "" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  fmt.Fprint(w, "Hello, ", name)

  calls = append(calls,name)

  if _, ok := stats[name]; !ok {
    stats[name] = 1

  } else {
    stats[name]++
  }

  fmt.Printf("calls: %#v\n", calls)
  fmt.Printf("stats: %#v\n\n", stats)
}
