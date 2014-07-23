package observable

import "fmt"

type Subject interface {
    Register(interface{})
    Delete(interface{})
    Notify(interface{})
    SetChange()
}

type Observer interface {
    Update()
}




