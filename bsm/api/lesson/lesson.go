package lesson

import (
    "bsm/api/slide"
)

type LessonListing struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
}

type LessonIndex []LessonListing

type Lesson struct {
    Id      int             `json:"id"`
    Name    string          `json:"name"`
    Deck    slide.Slides    `json:"slides"`
}

type Lessons []Lesson