package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"

    "github.com/naoina/denco"
)

func Index(w http.ResponseWriter, r *http.Request, params denco.Params) {
    fmt.Fprintf(w, "Welcome to Denco!\n")
}

func User(w http.ResponseWriter, r *http.Request, params denco.Params) {
  name := params.Get("name")
  log.Printf("form:%s!!\n", r.FormValue("auth"))
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
      fmt.Fprintf(w, "Hello %s. Error happened, error:%#v!\n", name, err) 
    } else {
      fmt.Fprintf(w, "Hello %s. data: %s!\n", name, data)
    }
}

func main() {
    mux := denco.NewMux()
    handler, err := mux.Build([]denco.Handler{
        mux.GET("/", Index),
        mux.GET("/user/:name", User),
        mux.POST("/user/:name", User),
    })
    if err != nil {
        panic(err)
    }
    log.Fatal(http.ListenAndServe("localhost:8080", handler))
}