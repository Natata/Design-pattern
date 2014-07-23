package observer

import "Design-pattern/observe/observable"

type Observer interface {
    Update()
}

func GetNewObserver(subject observable.Observable, values ...interface{}) Observer {
    for 
    obs := new(Observer)
}
