package models

type Pos struct {
    X, Y     int32
    Rotation float64  // Agregamos el campo Rotation
}

type Observer interface {
    Update(pos Pos)
}

type Subject interface {
    Register(observer Observer)
    Unregister(observer Observer)
    NotifyAll()
}