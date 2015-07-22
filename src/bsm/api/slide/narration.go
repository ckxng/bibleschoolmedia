package slide

import (
    "fmt"
    "encoding/json"
    "bsm/api/character"
)

// create a new narration slide with an id of 2
func NewNarrationSlide(name string, narrator character.Character, text string) NarrationSlide {
    ns := NarrationSlide {
        id: 2,
        name: name,
        myType: fmt.Sprintf("%T", NarrationSlide{}),
        data: make(map[string]interface{}),
    }
    ns.data["narrator"] = narrator
    ns.data["text"] = text
    return ns
}

type NarrationSlide struct {
    id          int                         `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
}

// get the Id
func (ns NarrationSlide) Id() int {
    return ns.id
}
 
// get the Name
func (ns NarrationSlide) Name() string {
    return ns.name
}

// get the Type
func (ns NarrationSlide) Type() string {
    return ns.myType
}

// get the Narrator
func (ns NarrationSlide) Narrarator() (character.Character, bool) {
    if val, ok := ns.data["narrator"]; ok {
        if narrator, ok := val.(character.Character); ok {
            return narrator, true
        }
    }
    return character.Character{}, false
}

// get the Text
func (ns NarrationSlide) Text() (string, bool) {
    if val, ok := ns.data["text"]; ok {
        if text, ok := val.(string); ok {
            return text, true
        }
    }
    return "", false
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{} {narrator:bsm.api.character.Character, 
// text:string}} 
func (ns NarrationSlide) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Id          int                         `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{
        Id:         ns.id,
        Name:       ns.name,
        MyType:     ns.myType,
        Data:       ns.data,
    })
}