package main

import "fmt"

type Human struct {
    Weapon func()
}

func Bomb() {
    fmt.Println("Bomb!")
}

func AK47() {
    fmt.Println("Da da da...")
}

func main(){

    solder := Human{Bomb}
    solder.Weapon()
    solder.Weapon = AK47;
    solder.Weapon()

}
