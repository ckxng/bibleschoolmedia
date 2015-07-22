package lesson

import (
    "github.com/gorilla/mux"
    "fmt"
    "net/http"
    "strconv"
    "bsm/api/slide"
    "appengine"
    "appengine/datastore"
    "encoding/json"
)

// returns a fully specified Lesson including the deck of Slides
//
// Error: "Missing ID" if the http.Request does not have an "id" Var
//
// Error: "Non-numeric ID" if the ID is not fully numeric
func Retrieve(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    id, ok := mux.Vars(r)["id"]
    if !ok || id == "" {
        return nil, fmt.Errorf("Missing ID")
    }
    
    iId, err := strconv.ParseInt(id,0,0)
    if err != nil {
        return nil, fmt.Errorf("Non-numeric ID: %s", id)
    }
    
    gc := appengine.NewContext(r)
    c, err := appengine.Namespace(gc, "bsm")
    if err != nil {
        return nil, fmt.Errorf("unable to change to namespace: %s", err)
    }

    key := datastore.NewKey(c, "Lesson", "", iId, nil)
    lesson := DatastoreLesson{}
    err = datastore.Get(c, key, &lesson)
    if err != nil {
        return nil, fmt.Errorf("unable to retrieve lessons: %s", err)
    }
    
    deck := slide.UntypedSlides{}
    err = json.Unmarshal([]byte(lesson.Deck), &deck)
    if err != nil {
        return Lesson {
            Id: int(key.IntID()),
            Name: lesson.Title,
        }, fmt.Errorf("Unable to read slides: %s [%s]", err, lesson.Deck)
    }
    
    slides := make(slide.Slides, len(deck))
    for i:=0; i<len(deck); i++ {
        if ts, err := deck[i].Upgrade(); err == nil {
            slides[i] = ts
        } else {
            return Lesson {
                Id: int(key.IntID()),
                Name: lesson.Title,
                Deck: slides,
            }, fmt.Errorf("Failed to upgrade slide of type: %s", deck[i].Type())
        }
    }
    
    return Lesson {
        Id: int(key.IntID()),
        Name: lesson.Title,
        Deck: slides,
    }, nil
}