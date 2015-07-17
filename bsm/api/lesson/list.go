package lesson

import (
    "net/http"
    "math/rand"
)

func List(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    return LessonIndex {
        LessonListing {
            Id: rand.Int(),
            Name: "A randomly generated listing",
        },
        LessonListing {
            Id: rand.Int(),
            Name: "Another randomly generated listing",
        },
        LessonListing {
            Id: rand.Int(),
            Name: "An additional randomly generated listing",
        },
        LessonListing {
            Id: rand.Int(),
            Name: "A further randomly generated listing",
        },
        LessonListing {
            Id: rand.Int(),
            Name: "Yet another randomly generated listing",
        },
    }, nil
}