package core

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
    "bsm/api/lesson"
)

type JSONHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

func Init() {
    router := mux.NewRouter().StrictSlash(false)
    router.HandleFunc("/hello", handler)
    router.HandleFunc("/api/v1/lesson", JSONDecorator(lesson.List))
    router.HandleFunc("/api/v1/lesson/{id:[0-9]+}", JSONDecorator(lesson.Retrieve))
    // add controller routes
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

    stack := negroni.New()
    // add middleware
    stack.UseHandler(router)

    http.Handle("/", stack)
}

func JSONDecorator(handler JSONHandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data, err := handler(w,r)
        payload := struct{
            Data    interface{}         `json:"data"`
            Err     error               `json:"err"`
        } {
            Data: data,
            Err: err,
        }
        
        w.Header().Set("Content-Type", "application/json")
        if bPayload, err := json.MarshalIndent(payload, "", "  "); err == nil {
            fmt.Fprint(w, string(bPayload))
        } else {
            fmt.Fprintf(w, "{\"data\":\"\",\"err\":\"Marshal error: %s\"}", err)
        }
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

