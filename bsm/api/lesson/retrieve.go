package lesson

import (
    "github.com/gorilla/mux"
    "fmt"
    "net/http"
    "strconv"
    "bsm/api/slide"
    "bsm/api/character"
)

func Retrieve(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    id, ok := mux.Vars(r)["id"]
    if !ok || id == "" {
        return Lesson{}, fmt.Errorf("Missing ID")
    }
    
    iId, err := strconv.ParseInt(id,0,0)
    if err != nil {
        return Lesson{}, fmt.Errorf("Non-numeric ID: %s", id)
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