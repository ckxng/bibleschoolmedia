package slide

import (
    "fmt"
    "encoding/json"
)

// a struct for simply passing slide data
type UntypedSlide struct {
    id          int                         `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
}

type UntypedSlides []UntypedSlide


// get the Id
func (us UntypedSlide) Id() int {
    return us.id
}
 
// get the Name
func (us UntypedSlide) Name() string {
    return us.name
}

// get the Type
func (us UntypedSlide) Type() string {
    return us.myType
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (us UntypedSlide) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Id          int                         `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{
        Id:         us.id,
        Name:       us.name,
        MyType:     us.myType,
        Data:       us.data,
    })
}

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (us UntypedSlide) UnmarshalJSON(data []byte) error {
    // BUG(cking) Unmarshal is not yet implimented.
    return fmt.Errorf("UntypedSlide.UnmarshalJSON Unimplimented")
}