package tool

import (
    "Design-pattern/strategy/Hatsu"
    "Design-pattern/strategy/Hunter"
)

func WaterDivination(appear string) *hunter.Hunter {
    switch appear {
        case "overflow":
            return &hunter.Hunter{new(hatsu.Enhancer)}
        default:
            return &hunter.Hunter{} 
    
    }
}
