package lesson

import (
    "math/rand"
    "bsm/api/slide"
    "bsm/api/character"
)

func Retrieve(id int) Lesson {
    return Lesson {
        Id: rand.Int(),
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
    }
}