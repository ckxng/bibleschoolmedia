// provides a simple type for representing
// characters who may be used throughout the application as narrarators.
//
// Characters have a name and a URL to an avatar image.
package character

import (
    "encoding/json"
    "math/rand"
)

// creates a new character with a randomized id
func NewCharacter(name string, avatarUrl string) Character {
    return Character {
        id: rand.Int(),
        name: name,
        avatarUrl: avatarUrl,
    }
}

type Character struct {
    id          int     `json:"id"`
    name        string  `json:"name"`
    avatarUrl   string  `json:"avatarUrl"`
}

// gets the Id
func (c *Character) Id() int {
    return c.id
}

// gets the Name
func (c *Character) Name() string {
    return c.name
}

// gets the Avatar URL
func (c *Character) AvatarUrl() string {
    return c.avatarUrl
}

// JSON marshals to {id:int, name:string, avatarURL:string}
func (c *Character) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Id          int                         `json:"id"`
        Name        string                      `json:"name"`
        AvatarUrl   string                      `json:"avatarUrl"`
    }{
        Id:         c.id,
        Name:       c.name,
        AvatarUrl:  c.avatarUrl,
    })
}

type Characters []Character