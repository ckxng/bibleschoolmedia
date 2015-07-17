package slide

import (
    "net/http"
    "fmt"
)

// returns an error.  There are no current plans to impliment this function.
func List(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    return nil, fmt.Errorf("Not implimented")
}