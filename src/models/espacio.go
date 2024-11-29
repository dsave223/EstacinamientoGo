package models
type Espacio struct {
    ID       int
    Ocupado  bool
    Vehiculo *Vehiculo
    Posicion Pos
}

func NewEspacio(id int, x, y int32) *Espacio {
    return &Espacio{
        ID:       id,
        Ocupado:  false,
        Vehiculo: nil,
        Posicion: Pos{X: x, Y: y},
    }
}

func (e *Espacio) Ocupar(vehiculo *Vehiculo) {
    e.Ocupado = true
    e.Vehiculo = vehiculo
}

func (e *Espacio) Liberar() {
    e.Ocupado = false
    e.Vehiculo = nil
}
