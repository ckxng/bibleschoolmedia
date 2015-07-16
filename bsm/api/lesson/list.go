package lesson

import (
    "math/rand"
)

func List() LessonIndex {
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
    }
}