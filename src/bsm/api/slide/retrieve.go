package slide

import (
    "github.com/gorilla/mux"
    "fmt"
    "net/http"
    "strconv"
    "bsm/api/character"
)

func Retrieve(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    id, ok := mux.Vars(r)["id"]
    if !ok || id == "" {
        return nil, fmt.Errorf("Missing ID")
    }
    
    iId, err := strconv.ParseInt(id,0,0)
    if err != nil {
        return nil, fmt.Errorf("Non-numeric ID: %s", id)
    }
    
    switch int(iId) {
        case 0:
            ts := NewTitleSlide(
                "Title Slide", 
                "Title Slide", 
                "A slide with a title",
            )
            ts.id = int(iId)
            return ts, nil
        case 1:
            is := NewImageSlide(
                "Image Slide",
                "Image Slide",
                "https://placehold.it/300?text=image",
                "An image caption",
            )
            is.id = int(iId)
            return is, nil
        case 2: 
            ns := NewNarrationSlide(
                "Narration Slide",
                character.NewCharacter(
                    "Narrarator",
                    "https://placehold.it/64?text=narr",
                ),
                "Lorem ipsum dolor sit amet.",
            )
            ns.id = int(iId)
            return ns, nil
    }
    
    return nil, fmt.Errorf("Invalid slide")
}