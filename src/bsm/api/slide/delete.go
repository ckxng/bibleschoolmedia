package slide

import (
    "fmt"
    "net/http"
)

// returns an error.  There are no current plans to impliment this function.
func Delete(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    return nil, fmt.Errorf("Not implimented")
}