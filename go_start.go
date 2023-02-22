package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func listenGoogle(w http.ResponseWriter, req *http.Request) {

    resp, err := http.Get("http://www.google.com")
    if err != nil {
        log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    sb := string(body)
    fmt.Fprintf(w, sb)
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/listenGoogle", listenGoogle)

    http.ListenAndServe(":8090", nil)
}