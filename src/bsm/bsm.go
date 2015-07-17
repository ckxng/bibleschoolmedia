// initializes the application and provides Init(), which
// can be called by either appengine or main to configure the web server
package bsm

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
    "bsm/api/lesson"
    "bsm/api/slide"
)

// a specialized version of http.HandlerFunc that returns (interface{}, error)
// and is intended to be wrapped by JSONDecorator
type JSONHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

// initializes the web server and configure routes
func Init() {
    router := mux.NewRouter().StrictSlash(false)
    router.HandleFunc("/hello", handler)
    router.HandleFunc("/api/v1/lesson", JSONDecorator(lesson.List))
    router.HandleFunc("/api/v1/lesson/{id:[0-9]+}", JSONDecorator(lesson.Retrieve))
    router.HandleFunc("/api/v1/slide/{id:[0-9]+}", JSONDecorator(slide.Retrieve))
    // add controller routes
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

    stack := negroni.New()
    // add middleware
    stack.UseHandler(router)

    http.Handle("/", stack)
}

// wraps JSONHanderFunc functions and provide JSON Content-type, marshaling,
// and HTTP error code functionality
func JSONDecorator(handler JSONHandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data, err := handler(w,r)

        payload := struct{
            Data    interface{}         `json:"data"`
            Err     interface{}         `json:"err"`
        } {
            Data: data,
            Err: err,
        }
        
        if err != nil {
            payload.Err = err.Error()
        }
        
        w.Header().Set("Content-Type", "application/json")
        if bPayload, merr := json.MarshalIndent(payload, "", "  "); merr == nil {
            if err != nil {
                w.WriteHeader(500)
            }
            fmt.Fprint(w, string(bPayload))
        } else {
            w.WriteHeader(500)
            fmt.Fprintf(w, "{\"data\":\"\",\"err\":\"Marshal error: %s\"}", err)
        }
    }
}

// a simple test function independant of all other app capabilities
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

