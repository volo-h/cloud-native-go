package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/api/echo", EchoHandleFunc)
  http.HandleFunc("/api/hello", HelloHandleFunc)

  http.HandleFunc("/api/books", BooksHandleFunc)
  http.HandleFunc("/api/books/", BookHandleFunc)

  http.ListenAndServe(port(), nil)
}

func port() string {
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8080"
  }
  return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Welcome")
}
