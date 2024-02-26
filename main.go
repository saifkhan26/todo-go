package main

import (
    "fmt"
    "net/http"
)

func main() {
    i := 1
    fmt.Println("Out", i)
    for i <= 3{
        fmt.Println("Hello From the loop", i)
        i = i + 1
    }
    for j := 0; j <= 3; j++{
        fmt.Println("Hello from traditional Loop")
    }
    for i := range 3{

    }
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
    })

    http.ListenAndServe(":8080", nil)
}
