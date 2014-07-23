package observable

import (
    custome "Design-pattern/observe/observer"
)

type Subject struct {
    observers map[custome.Id]custome.Observer
}

func (s *Subject) Register(c custome.Observer) {
    s.observers[c.Id] = c
}

func (s *Subject) Delete(c custome.Observer) {
    delete(s.observers, c.Id)
}

func (s *Subject) Notify() {
    for _, o := range s.observers {
        o.Update()
    }
}


