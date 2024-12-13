package models

type Pos struct {
    X, Y     int
}

type Observer interface {
    Update(pos Pos)
}

type Subject interface {
    Register(observer Observer)
    Unregister(observer Observer)
    NotifyAll()
}