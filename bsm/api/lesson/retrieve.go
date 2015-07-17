package lesson

import (
    "github.com/gorilla/mux"
    "fmt"
    "net/http"
    "strconv"
    "bsm/api/slide"
    "bsm/api/character"
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
    
    return Lesson {
        Id: int(iId),
        Name: "Example Lesson",
        Deck: slide.Slides {
            slide.NewTitleSlide(
                "Title Slide", 
                "Title Slide", 
                "A slide with a title",
            ),
            slide.NewImageSlide(
                "Image Slide",
                "Image Slide",
                "https://placehold.it/300?text=image",
                "An image caption",
            ),
            slide.NewNarrationSlide(
                "Narration Slide",
                character.NewCharacter(
                    "Narrarator",
                    "https://placehold.it/64?text=narr",
                ),
                "Lorem ipsum dolor sit amet.",
            ),
        },
    }, nil
}