package models

import (
    "fmt"
    "sync"
)

type Estacionamiento struct {
    Capacidad       int
    Espacios        []*Espacio
    EntradaSalida   sync.Mutex
    CondicionEspera *sync.Cond
    Observers       []Observer
}

func NewEstacionamiento(capacidad int) *Estacionamiento {
    mu := &sync.Mutex{}
    
    estacionamiento := &Estacionamiento{
        Capacidad:       capacidad,
        Espacios:        make([]*Espacio, capacidad),
        CondicionEspera: sync.NewCond(mu),
        Observers:       []Observer{},
    }

    // Configurar posiciones de espacios (ejemplo de disposición)
    for i := 0; i < capacidad; i++ {
        x := 50 + int32((i%5) * 120)
        y := 100 + int32((i/5) * 70)
        estacionamiento.Espacios[i] = NewEspacio(i+1, x, y)
    }

    return estacionamiento
}

func (e *Estacionamiento) Entrar(vehiculo *Vehiculo) *Espacio {
    e.EntradaSalida.Lock()
    defer e.EntradaSalida.Unlock()

    // Esperar mientras no haya espacios disponibles
    for e.EspaciosOcupados() >= e.Capacidad {
        e.NotificarObservadores(fmt.Sprintf("Vehículo %d esperando espacio", vehiculo.ID))
        e.CondicionEspera.Wait()
    }

    // Encontrar y ocupar un espacio disponible
    for _, espacio := range e.Espacios {
        if !espacio.Ocupado {
            espacio.Ocupar(vehiculo)
            e.NotificarObservadores(fmt.Sprintf("Vehículo %d entra en espacio %d", vehiculo.ID, espacio.ID))
            
            // Configurar posición del vehículo en el espacio
            vehiculo.SetPosicion(espacio.Posicion)
            
            return espacio
        }
    }

    return nil
}

func (e *Estacionamiento) Salir(espacio *Espacio) {
    e.EntradaSalida.Lock()
    defer e.EntradaSalida.Unlock()

    espacio.Liberar()
    e.NotificarObservadores(fmt.Sprintf("Vehículo sale del espacio %d", espacio.ID))

    // Despertar vehículos en espera
    e.CondicionEspera.Broadcast()
}

func (e *Estacionamiento) EspaciosOcupados() int {
    ocupados := 0
    for _, espacio := range e.Espacios {
        if espacio.Ocupado {
            ocupados++
        }
    }
    return ocupados
}

func (e *Estacionamiento) Register(observer Observer) {
    e.Observers = append(e.Observers, observer)
}

func (e *Estacionamiento) Unregister(observer Observer) {
    for i, obs := range e.Observers {
        if obs == observer {
            e.Observers = append(e.Observers[:i], e.Observers[i+1:]...)
            break
        }
    }
}

func (e *Estacionamiento) NotifyAll() {
    // Implementación opcional de notificación general
}

func (e *Estacionamiento) NotificarObservadores(evento string) {
    for _, obs := range e.Observers {
        obs.Update(Pos{}) // Adaptación al Observer original
    }
}