package models

// Space representa un espacio de estacionamiento lógico.
type Space struct {
	id int
    car *Car
	occupied bool
	x float64
	y float64
	rotation float64
	scale float64
}

// NewSpace crea una nueva instancia de Space.
func NewSpace(id int, x, y, scale, rotación float64) *Space {
	return &Space{
		id:       id,
        car: nil,
		occupied: false,
		x:        x,
		y:        y,
		rotation: rotación,
		scale:    scale,
	}
}

// Occupy ocupa el espacio con una matrícula.
func (s *Space) Occupy(car *Car) {
	s.occupied = true
	s.car = car
}

// Vacate libera el espacio.
func (s *Space) Vacate(car *Car) {
	s.occupied = false
	s.car = car
}

// GetID retorna el ID del espacio.
func (s *Space) GetID() int {
	return s.id
}

// GetCar retorna el carro en el espacio, si existe.
func (s *Space) GetCar() *Car {
	return s.car
}

// GetOccupied retorna si el espacio está ocupado o no.
func (s *Space) GetOccupied() bool {
	return s.occupied
}

// GetX retorna la coordenada X del espacio.
func (s *Space) GetX() float64 {
	return s.x
}

// GetY retorna la coordenada Y del espacio.
func (s *Space) GetY() float64 {
	return s.y
}

// GetRotation retorna la rotación del espacio.
func (s *Space) GetRotation() float64 {
	return s.rotation
}

// GetScale retorna la escala del espacio.
func (s *Space) GetScale() float64 {
	return s.scale
}
