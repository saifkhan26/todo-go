package main

import (
	"fmt"
	"reflect"
	//    "net/http"
)

func main() {
    whatIsTheType := func(v interface{}){
        println(reflect.TypeOf(v))
        switch v.(type){
        case bool:
        fmt.Println("This is a boolean")
        case int:
        fmt.Println("This is a Int")
        default:
        fmt.Println("This type is not recognizable")
    }
    }
    whatIsTheType(1)
    whatIsTheType(true)
    whatIsTheType("Hello")
//    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//        fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
//    })
//
//    http.ListenAndServe(":8080", nil)
}
