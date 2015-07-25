package slide

import (
    "fmt"
)

// a struct for simply passing slide data
type UntypedSlide struct {
    id          int64                       `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
}

type UntypedSlides []UntypedSlide


// get the Id
func (us UntypedSlide) Id() int64 {
    return us.id
}
 
// get the Name
func (us UntypedSlide) Name() string {
    fmt.Printf("my name is: %s\n", us.name)
    return us.name
}

// get the Type
func (us UntypedSlide) Type() string {
    return us.myType
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (us UntypedSlide) MarshalJSON() ([]byte, error) {
    return marshalSlideJSON(us.id, us.name, us.myType, us.data)
}

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (us *UntypedSlide) UnmarshalJSON(data []byte) error {
    var err error
    us.id, us.name, us.myType, us.data, err = unmarshalSlideJSON(data)
    return err
}

// upgrades UntypedSlide to a more full-featured type that impliments the 
// Slide interface
func (us UntypedSlide) Upgrade() (Slide, error) {
    switch us.Type() {
        case "slide.ImageSlide":
            return ImageSlide {
                id: us.id,
                name: us.name,
                myType: us.myType,
                data: us.data,
            }, nil
        case "slide.TitleSlide":
            return TitleSlide {
                id: us.id,
                name: us.name,
                myType: us.myType,
                data: us.data,
            }, nil
        case "slide.NarrationSlide":
            return NarrationSlide {
                id: us.id,
                name: us.name,
                myType: us.myType,
                data: us.data,
            }, nil
    }
    
    return nil, fmt.Errorf("Unable to detect slide type: %s", us.Type())
}
