package character

import (
    "encoding/json"
    "math/rand"
)

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

func (c Character) Id() int {
    return c.id
}

func (c Character) Name() string {
    return c.name
}

func (c Character) AvatarUrl() string {
    return c.avatarUrl
}

func (c Character) MarshalJSON() ([]byte, error) {
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