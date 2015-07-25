package slide

import (
    "fmt"
    "net/http"
)

// returns an error.  There are no current plans to impliment this function.
func Retrieve(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    return nil, fmt.Errorf("Not implimented")
}