package main

import (
    "fmt"
    "log"
    "net/http"
    //"encoding/json"
)

func test(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Test Successful !!")
    fmt.Println("Endpoint Hit: test")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/test", test)
    log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {
    handleRequests()
}