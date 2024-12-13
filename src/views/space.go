package views

import (
	"log"
	"math"
	"simulador/src/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// SpaceView representa la vista de un espacio de estacionamiento.
type SpaceView struct {
	space *models.Space
	image *ebiten.Image
}

// NewSpaceView crea una nueva instancia de SpaceView vinculada a un modelo lógico Space.
func NewSpaceView(space *models.Space) *SpaceView {
	// Cargar la imagen del espacio desde assets.
	img, _, err := ebitenutil.NewImageFromFile("assets/scene/parkingSing.png")
	if err != nil {
		log.Fatalf("Error al cargar la imagen del espacio: %v", err)
	}

	return &SpaceView{
		space: space,
		image: img,
	}
}

// Draw dibuja el espacio en la pantalla.
func (s *SpaceView) Draw(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}

    // Scale
    op.GeoM.Scale(s.space.GetScale(), s.space.GetScale())

    // Rotate
    op.GeoM.Rotate(s.space.GetRotation() * (math.Pi / 180))

    // Translate
    op.GeoM.Translate(s.space.GetX(), s.space.GetY())

    screen.DrawImage(s.image, op)
}
