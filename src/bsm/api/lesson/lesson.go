// provides the types and api methods required to serve lesson data throughout
// the application
package lesson

import (
    "bsm/api/slide"
)

// a simple listing of lessons consisting only of an Id and Name
type LessonListing struct {
    Id      int64   `json:"id"`
    Name    string  `json:"name"`
}

type LessonIndex []LessonListing

// the data type stored by Google Datastore
type DatastoreLesson struct {
    ID      string
    Title   string
    Deck    string
    Active  bool
}

// a fully specified lesson including the slide deck
type Lesson struct {
    Id      int64           `json:"id"`
    Name    string          `json:"name"`
    Deck    slide.Slides    `json:"slides"`
}

type Lessons []Lesson