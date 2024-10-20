package main

import (
  "fmt"
  "log"
  "net/http"
  )

  // Form function to handle /form path. 
func formHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(), err != nil {
    fmt.Fprintf(w, "ParseForm() err: %v ", err)
    return
  }
  fmt.Fprintf("Post request Successful.")
  name := r.FormValue("name")
  address := r.FormValue("address")
  fmt.Printf(w, "Name = %s \n", name)
  fmt.Printf(w, "Address = %s \n", address)
}

  // hello function to handle /hello path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.path != "/hello" {
  http.Error(w, "404 nnot found", http.StatusNotFound)
  return
  }
if r.Method != "GET" {
http.Error(w, "method is not supported.", http.StatusNotFound)
  return
  }
  fmt.Fprintf(w, "hello!")
}

  // main function
func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/form", formHandler)
  http.HandleFunc("/hello", helloHandler)
  fmt.Printf("Starting server  at port  8080\n")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
