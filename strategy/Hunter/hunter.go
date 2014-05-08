package hunter

import (
    "Design-pattern/strategy/Hatsu"
    "fmt"
)

type Hunter struct{
    hatsu.Hatsuer
}

func (h Hunter) Ten() {
    fmt.Println("Ten!")
}

