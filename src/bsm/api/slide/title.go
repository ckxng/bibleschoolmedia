package slide

import (
    "fmt"
    "encoding/json"
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
    return json.Marshal(struct{
        Id          int64                       `json:"id"`
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

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (ts *TitleSlide) UnmarshalJSON(data []byte) error {
    um := struct{
        Id          int64                       `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{}
    err := json.Unmarshal(data, &um)
    if err != nil {
        return fmt.Errorf("Unable to unmarshal data: %s", err)
    }
    
    ts.id = um.Id
    ts.name = um.Name
    ts.myType = um.MyType
    ts.data = um.Data
    
    return nil
}
