package slide

import (
    "bsm/api/character"
)

func getMappedString(sm map[string]interface{}, k string) (string, bool) {
    if val, ok := sm[k]; ok {
        if sVal, ok := val.(string); ok {
            return sVal, true
        }
    }
    return "", false
}

func getMappedCharacter(sm map[string]interface{}, k string) (character.Character, bool) {
    if val, ok := sm[k]; ok {
        if c, ok := val.(character.Character); ok {
            return c, true
        }
    }
    return character.Character{}, false
}