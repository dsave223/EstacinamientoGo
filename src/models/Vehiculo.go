package models

import (
	"fmt"
	"time"
)

type Vehiculo struct {
    ID       int
	Pos      Pos
	Status   bool
	Observers []Observer
} 

func NewVehiculo(id int) *Vehiculo {
    return &Vehiculo{
		ID:       id,
		Pos:      Pos{X: 0, Y: 0},
		Status:   true,
		Observers: []Observer{},
	}
}

// Register añade un observador al vehículo
func (v *Vehiculo) Register(observer Observer) {
	v.Observers = append(v.Observers, observer)
	fmt.Println("nos")
}

// Unregister elimina un observador del vehículo
func (v *Vehiculo) Unregister(observer Observer) {
	for i, o := range v.Observers {
		if o == observer {
			v.Observers = append(v.Observers[:i], v.Observers[i+1:]...)
			break
		}
	}
}

// NotifyAll notifica a todos los observadores sobre una actualización
func (v *Vehiculo) NotifyAll() {
	for _, observer := range v.Observers {
		observer.Update(v.Pos)
	}
}

// SimularMovimiento simula el movimiento del vehículo y notifica cambios
func (v *Vehiculo) SimularMovimiento(x, y int32) {
	v.Pos = Pos{X: x, Y: y}
	v.NotifyAll()
	time.Sleep(1 * time.Second) // Simula un tiempo de movimiento
}


