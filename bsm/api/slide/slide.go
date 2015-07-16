package slide

type Slide interface {
    Id()            int
    Name()          string
    Type()          string
}

type Slides []Slide