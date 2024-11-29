package models

import (
    "math/rand"
    "time"
)

type Vehiculo struct {
    ID             int
    TiempoEstancia time.Duration
    Posicion       Pos
    ImagenPath     string
    Observers      []Observer
}

func NewVehiculo(id int) *Vehiculo {
    imagenes := []string{
        "./assets/vehiculos/vehiculo1.png",
        "./assets/vehiculos/vehiculo2.png",
        "./assets/vehiculos/vehiculo3.png",
    }
    
    return &Vehiculo{
        ID:         id,
        ImagenPath: imagenes[rand.Intn(len(imagenes))],
        Posicion:   Pos{X: 0, Y: 0},
        Observers:  []Observer{},
    }
}

func (v *Vehiculo) GenerarTiempoEstancia() time.Duration {
    return time.Duration(rand.Intn(2000)+3000) * time.Millisecond
}

func (v *Vehiculo) Register(observer Observer) {
    v.Observers = append(v.Observers, observer)
}

func (v *Vehiculo) Unregister(observer Observer) {
    for i, obs := range v.Observers {
        if obs == observer {
            v.Observers = append(v.Observers[:i], v.Observers[i+1:]...)
            break
        }
    }
}

func (v *Vehiculo) NotifyAll() {
    for _, observer := range v.Observers {
        observer.Update(v.Posicion)
    }
}

func (v *Vehiculo) SetPosicion(pos Pos) {
    v.Posicion = pos
    v.NotifyAll()
}