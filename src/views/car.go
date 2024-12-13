package views

import (
	"log"
	"math"
	"simulador/src/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// CarView representa la vista de un vehículo.
type CarView struct {
	Car   *models.Car
	image *ebiten.Image
}

// NewCarView crea una nueva instancia de CarView vinculada a un modelo lógico Car.
func NewCarView(car *models.Car) *CarView {
	img, _, err := ebitenutil.NewImageFromFile("assets/cars/car3.png") // Asegúrate de tener esta imagen.
	if err != nil {
		log.Fatalf("Error al cargar la imagen del auto: %v", err)
	}
	return &CarView{
		Car:   car,
		image: img,
	}
}

// Draw dibuja el auto en la pantalla.
func (c *CarView) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Usamos los getters para acceder a las propiedades del coche
	op.GeoM.Scale(c.Car.GetScale(), c.Car.GetScale())    // Escala
	op.GeoM.Rotate(c.Car.GetRotation() * (math.Pi / 180)) // Rotación
	op.GeoM.Translate(c.Car.GetPosX(), c.Car.GetPosY())        // Posición

	// Dibujamos la imagen del coche
	screen.DrawImage(c.image, op)
}
