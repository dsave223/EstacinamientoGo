package models

import (
	"time"
)

type Car struct {
	id         int
	posX, posY float64
	scale      float64
	status     bool
	observers  []Observer
	time       time.Duration
	speed      float64
	rotation   float64
}

func NewCar(id int, X, Y, scale, speed float64) *Car {
	return &Car{
		id:       id,
		posX:     X,
		posY:     Y,
		scale:    scale,
		status:   true,
		observers: []Observer{},
		time:     time.Duration(0),
		speed:    speed,
		rotation: 90,
	}
}

func (c *Car) Update() {
	if c.posY < 200 {
		c.posY += c.speed
	} else if c.posY < 250 {
		c.posY += c.speed
	} else if c.posX < 250 {
		c.posY = 268
		c.rotation = 360
		c.posX += c.speed
	} else if c.posY < 370 {
		c.posX = 327
		c.rotation = 90
		c.posY += c.speed
	} else if c.posX < 329 {
		c.posY = 380
		c.rotation = 360
		c.posX += c.speed
	}
}

// SetStatus actualiza el estado del auto.
func (c *Car) SetStatus(status bool) {
	c.status = status
}

// Register añade un observador a la lista.
func (c *Car) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}

// Unregister elimina un observador de la lista.
func (c *Car) Unregister(observer Observer) {
	for i, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll notifica a todos los observadores sobre una actualización.
func (c *Car) NotifyAll() {
	for _, observer := range c.observers {
		observer.Update(Pos{X: int(c.posX), Y: int(c.posY)})
	}
}

// Getters

// GetId retorna el ID del carro.
func (c *Car) GetId() int {
	return c.id
}

// GetPosX retorna la posición X del carro.
func (c *Car) GetPosX() float64 {
	return c.posX
}

// GetPosY retorna la posición Y del carro.
func (c *Car) GetPosY() float64 {
	return c.posY
}

// GetScale retorna la escala del carro.
func (c *Car) GetScale() float64 {
	return c.scale
}

// GetStatus retorna el estado del carro (true = activo).
func (c *Car) GetStatus() bool {
	return c.status
}

// GetRotation retorna la rotación actual del carro.
func (c *Car) GetRotation() float64 {
	return c.rotation
}

// GetSpeed retorna la velocidad del carro.
func (c *Car) GetSpeed() float64 {
	return c.speed
}

// GetTime retorna el tiempo asociado al carro (para simular movimiento en el tiempo).
func (c *Car) GetTime() time.Duration {
	return c.time
}