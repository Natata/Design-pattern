package main

import (
"fmt"
"strconv"
)

type subject struct {
    id       string
    observer map[string]*subscriber
}

func (s subject) register(o *subscriber) {
    s.observer[o.id] = o
}

func (s subject) update() {
    for k, _ := range s.observer {
    s.observer[k].update()
}
}

func (s subject) list() {
    for k, _ := range s.observer {
    fmt.Println(s.observer[k])
}
}

type subscriber struct {
    id    string
    value int
}

func (s *subscriber) update() {
    s.value += 1
}

func main() {
    sub := subject{id: "ker", observer: make(map[string]*subscriber)}

    for i := 0; i < 5; i++ {
        a := &subscriber{id: strconv.Itoa(i), value: i * 4}
        sub.register(a)
    }
    sub.list()
    sub.update()
    sub.list()
}
