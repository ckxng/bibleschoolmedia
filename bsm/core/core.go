package core

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
    "bsm/api/lesson"
)

func Init() {
    router := mux.NewRouter().StrictSlash(false)
    router.HandleFunc("/hello", handler)
    router.HandleFunc("/api/v1/lesson", JSONHandlerFunc(lesson.List()))
    router.HandleFunc("/api/v1/lesson/1", JSONHandlerFunc(lesson.Retrieve(1)))
    // add controller routes
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

    stack := negroni.New()
    // add middleware
    stack.UseHandler(router)

    http.Handle("/", stack)
}

func JSONHandlerFunc(data interface{}) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        payload := struct{
            Data    interface{}         `json:"data"`
            Err     string              `json:"err"`
        } {
            Data: data,
            Err: "",
        }
        
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

