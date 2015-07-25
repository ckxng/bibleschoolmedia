package slide

import (
    "fmt"
    "encoding/json"
)

// create a new image slide with an id of 1
func NewImageSlide(name string, title string, url string, caption string) ImageSlide {
    is := ImageSlide {
        id: 1,
        name: name,
        myType: fmt.Sprintf("%T", ImageSlide{}),
        data: make(map[string]interface{}),
    }
    is.data["title"] = title
    is.data["url"] = url
    is.data["caption"] = caption
    return is
}

// a Slide type that holds a URL to a title, a URL to an image, and a caption.
type ImageSlide struct {
    id          int64                       `json:"id"`
    name        string                      `json:"name"`
    myType      string                      `json:"myType"`
    data        map[string]interface{}      `json:"data"`
}

// get the Id
func (is ImageSlide) Id() int64 {
    return is.id
}
 
// get the Name
func (is ImageSlide) Name() string {
    return is.name
}

// get the Type
func (is ImageSlide) Type() string {
    return is.myType
}

// get the Title
func (is ImageSlide) Title() (string, bool) {
    return getMappedString(is.data, "title")
}

// get the Url to the image
func (is ImageSlide) Url() (string, bool) {
    return getMappedString(is.data, "url")
}

// get the Caption
func (is ImageSlide) Caption() (string, bool) {
    return getMappedString(is.data, "caption")
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{} {title:string, url:string, caption:string}} 
func (is ImageSlide) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Id          int64                       `json:"id"`
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

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func (is *ImageSlide) UnmarshalJSON(data []byte) error {
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
    
    is.id = um.Id
    is.name = um.Name
    is.myType = um.MyType
    is.data = um.Data
    
    return nil
}
