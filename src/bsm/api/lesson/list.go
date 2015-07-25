package lesson

import (
    "fmt"
    "net/http"
    "appengine"
    "appengine/datastore"
)

// returns a LessonIndex of all available lessons
func List(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    gc := appengine.NewContext(r)
    c, err := appengine.Namespace(gc, "bsm")
    if err != nil {
        return nil, fmt.Errorf("unable to change to namespace: %s", err)
    }

    q := datastore.NewQuery("Lesson").Filter("Active =", true)
    var lessons []DatastoreLesson
    keys, err := q.GetAll(c, &lessons)
    if err != nil {
        return nil, fmt.Errorf("unable to retrieve lessons: %s", err)
    }
    
    index := make(LessonIndex, len(lessons))
    for i:=0; i<len(lessons); i++ {
        index[i] = LessonListing {
            Id: keys[i].IntID(),
            Name: lessons[i].Title,
        }
    }
    
    return index, nil
}