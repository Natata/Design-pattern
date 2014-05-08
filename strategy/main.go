package main

import (
    "Design-pattern/strategy/Hunter"
    "Design-pattern/strategy/Hatsu"
)

func main(){
    gon := &hunter.Hunter{new(hatsu.Enhancer)}
    gon.Hatsu();
    gon.Ten();
}
