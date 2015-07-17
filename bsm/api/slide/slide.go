// provides the types and api methods required to represent slides throughout
// the application.
package slide

// an interface common to all slides.
// 
// This is the minimum amount of information required to list slides and
// manipulate them from a high level.  The slide must be upgraded to a more 
// specific type before the slide content is accessible.
//
// Structs that impliment the Slide interface are encouraged to json.Marshal to
// {id:int, name:string, myType:fmt.Sprintf("%T"), data:map[string]interface{}}
// for maximum compatibility.
type Slide interface {
    Id()            int
    Name()          string
    Type()          string
}

type Slides []Slide