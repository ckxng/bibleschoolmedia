package core

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
)

func Init() {
    router := mux.NewRouter().StrictSlash(false)
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
    router.HandleFunc("/", handler)
    // add controller routes

    stack := negroni.New()
    // add middleware
    stack.UseHandler(router)

    http.Handle("/", stack)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

