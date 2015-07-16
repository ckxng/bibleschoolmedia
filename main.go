// +build !appengine
	    
package main

import (
    "net/http"
    "bsm/core"
)

func main() {
    core.Init()
    http.ListenAndServe("0.0.0.0:8080", nil)        
}
