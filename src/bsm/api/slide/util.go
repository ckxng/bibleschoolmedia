package slide

import (
    "fmt"
    "bsm/api/character"
    "encoding/json"
)

// safely extract a string from sm, if sm[k] holds a string.
// the second return value is true on success
func getMappedString(sm map[string]interface{}, k string) (string, bool) {
    if val, ok := sm[k]; ok {
        if sVal, ok := val.(string); ok {
            return sVal, true
        }
    }
    return "", false
}

// safely extract a character.Character from sm, if sm[k] holds a 
// character.Character.  the second return value is true on success
func getMappedCharacter(sm map[string]interface{}, k string) (character.Character, bool) {
    if val, ok := sm[k]; ok {
        if c, ok := val.(character.Character); ok {
            return c, true
        }
    }
    return character.Character{}, false
}

// json.Marshal to {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{} {narrator:bsm.api.character.Character, 
// text:string}} 
func marshalSlideJSON(id int64, name string, myType string, data map[string]interface{}) ([]byte, error) {
    return json.Marshal(struct{
        Id          int64                       `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{
        Id:         id,
        Name:       name,
        MyType:     myType,
        Data:       data,
    })
}

// json.Unmarshal from {id:int, name:string, myType:fmt.Sprintf("%T"), 
// data:map[string]interface{}} 
func unmarshalSlideJSON(data []byte) (int64, string, string, map[string]interface{}, error) {
    um := struct{
        Id          int64                       `json:"id"`
        Name        string                      `json:"name"`
        MyType      string                      `json:"myType"`
        Data        map[string]interface{}      `json:"data"`
    }{}
    err := json.Unmarshal(data, &um)
    if err != nil {
        return 0, "", "", map[string]interface{}{}, fmt.Errorf("Unable to unmarshal data: %s", err)
    }
    
    return um.Id, um.Name, um.MyType, um.Data, nil
}
