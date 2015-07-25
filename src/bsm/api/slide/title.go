package slide

import (
    "fmt"
)

// create a new title slide with an id of 0
func NewTitleSlide(name string, title string, subtitle string) TitleSlide {
    ts := TitleSlide {
        id: 0,
        name: name,
        myType: fmt.Sprintf("%T", TitleSlide{}),
        data: make(map[string]interface{}),
    }
    ts.data["title"] = title
    ts.data["subtitle"] = subtitle
    return ts
}

type TitleSlide struct {
    id          int64                       `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
    Test        string                      `json:"test"`
}

// get the Id
func (ts TitleSlide) Id() int64 {
    return ts.id
}

// get the Name
func (ts TitleSlide) Name() string {
    return ts.name
}

// get the Type
func (ts TitleSlide) Type() string {
    return ts.myType
}

// get the Title
func (ts TitleSlide) Title() (string, bool) {
    return getMappedString(ts.data, "title")
}

// get the Subtitle
func (ts TitleSlide) Subtitle() (string, bool) {
    return getMappedString(ts.data, "subtitle")
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{} {title:string, subtitle:string}} 
func (ts TitleSlide) MarshalJSON() ([]byte, error) {
    return marshalSlideJSON(ts.id, ts.name, ts.myType, ts.data)
}

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (ts *TitleSlide) UnmarshalJSON(data []byte) error {
    var err error
    ts.id, ts.name, ts.myType, ts.data, err = unmarshalSlideJSON(data)
    return err
}
