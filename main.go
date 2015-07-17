// +build !appengine
	    
package main

import (
    "net/http"
    "bsm"
)

func main() {
    bsm.Init()
    http.ListenAndServe("0.0.0.0:8080", nil)        
}
