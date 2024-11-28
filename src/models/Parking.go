package models

import (
	"sync"
)

type Parking struct {
	Espacios []*Espacio
	Mutex    sync.Mutex
}

// NewParking crea un nuevo estacionamiento con el número de espacios especificado.
func NewParking(numEspacios int) *Parking {
	espacios := make([]*Espacio, numEspacios)
	for i := 0; i < numEspacios; i++ {
		espacios[i] = NewEspacio(i)
	}
	return &Parking{
		Espacios: espacios,
	}
}

// VerificarDisponibilidad retorna true si hay al menos un espacio disponible.
func (p *Parking) VerificarDisponibilidad() bool {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	for _, espacio := range p.Espacios {
		if !espacio.Ocupado {
			return true
		}
	}
	return false
}

// OcuparEspacio asigna un vehículo al primer espacio disponible.
func (p *Parking) OcuparEspacio(vehiculo *Vehiculo) *Espacio {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	for _, espacio := range p.Espacios {
		if !espacio.Ocupado {
			espacio.Ocupar(vehiculo)
			return espacio
		}
	}
	return nil // No hay espacios disponibles
}

// LiberarEspacio libera un espacio ocupado.
func (p *Parking) LiberarEspacio(idEspacio int) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if idEspacio >= 0 && idEspacio < len(p.Espacios) {
		p.Espacios[idEspacio].Liberar()
	}
}
