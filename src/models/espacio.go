package models

type Espacio struct {
	ID       int
	Ocupado  bool
	Vehiculo *Vehiculo
}

// NewEspacio crea un nuevo espacio de estacionamiento.
func NewEspacio(id int) *Espacio {
	return &Espacio{
		ID:      id,
		Ocupado: false,
	}
}

// Ocupar marca el espacio como ocupado por un vehículo.
func (e *Espacio) Ocupar(vehiculo *Vehiculo) {
	e.Ocupado = true
	e.Vehiculo = vehiculo
}

// Liberar libera el espacio, dejándolo disponible.
func (e *Espacio) Liberar() {
	e.Ocupado = false
	e.Vehiculo = nil
}
