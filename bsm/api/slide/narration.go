package slide

import (
    "fmt"
    "encoding/json"
    "math/rand"
    "bsm/api/character"
)

func NewNarrationSlide(name string, narrator character.Character, text string) NarrationSlide {
    ns := NarrationSlide {
        id: rand.Int(),
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

func (ns NarrationSlide) Id() int {
    return ns.id
}
 
func (ns NarrationSlide) Name() string {
    return ns.name
}

func (ns NarrationSlide) Type() string {
    return ns.myType
}

func (ns NarrationSlide) Narrarator() (character.Character, bool) {
    if val, ok := ns.data["narrator"]; ok {
        if narrator, ok := val.(character.Character); ok {
            return narrator, true
        }
    }
    return character.Character{}, false
}

func (ns NarrationSlide) Text() (string, bool) {
    if val, ok := ns.data["text"]; ok {
        if text, ok := val.(string); ok {
            return text, true
        }
    }
    return "", false
}

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