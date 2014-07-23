package observable

type Observable interface {
    Register(interface{})
    Delete(interface{})
    Notify(interface{})
}

