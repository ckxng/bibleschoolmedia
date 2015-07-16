package slide

import (
    "fmt"
    "encoding/json"
    "math/rand"
)

func NewImageSlide(name string, title string, url string, caption string) ImageSlide {
    is := ImageSlide {
        id: rand.Int(),
        name: name,
        myType: fmt.Sprintf("%T", ImageSlide{}),
        data: make(map[string]interface{}),
    }
    is.data["title"] = title
    is.data["url"] = url
    is.data["caption"] = caption
    return is
}

type ImageSlide struct {
    id          int                         `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
}

func (is ImageSlide) Id() int {
    return is.id
}
 
func (is ImageSlide) Name() string {
    return is.name
}

func (is ImageSlide) Type() string {
    return is.myType
}

func (is ImageSlide) Title() (string, bool) {
    if val, ok := is.data["title"]; ok {
        if title, ok := val.(string); ok {
            return title, true
        }
    }
    return "", false
}

func (is ImageSlide) Url() (string, bool) {
    if val, ok := is.data["url"]; ok {
        if url, ok := val.(string); ok {
            return url, true
        }
    }
    return "", false
}

func (is ImageSlide) Caption() (string, bool) {
    if val, ok := is.data["caption"]; ok {
        if caption, ok := val.(string); ok {
            return caption, true
        }
    }
    return "", false
}

func (is ImageSlide) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Id          int                         `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{
        Id:         is.id,
        Name:       is.name,
        MyType:     is.myType,
        Data:       is.data,
    })
}