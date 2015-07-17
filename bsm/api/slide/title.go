package slide

import (
    "fmt"
    "encoding/json"
    "math/rand"
)

// create a new title slide with a randomly generated id
func NewTitleSlide(name string, title string, subtitle string) TitleSlide {
    ts := TitleSlide {
        id: rand.Int(),
        name: name,
        myType: fmt.Sprintf("%T", TitleSlide{}),
        data: make(map[string]interface{}),
    }
    ts.data["title"] = title
    ts.data["subtitle"] = subtitle
    return ts
}

type TitleSlide struct {
    id          int                         `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
    Test        string                      `json:"test"`
}

// get the Id
func (ts TitleSlide) Id() int {
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
    if val, ok := ts.data["title"]; ok {
        if title, ok := val.(string); ok {
            return title, true
        }
    }
    return "", false
}

// get the Subtitle
func (ts TitleSlide) Subtitle() (string, bool) {
    if val, ok := ts.data["subtitle"]; ok {
        if subtitle, ok := val.(string); ok {
            return subtitle, true
        }
    }
    return "", false
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{} {title:string, subtitle:string}} 
func (ts TitleSlide) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Id          int                         `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{
        Id:         ts.id,
        Name:       ts.name,
        MyType:     ts.myType,
        Data:       ts.data,
    })
}