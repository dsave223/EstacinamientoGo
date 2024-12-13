package views

import (
	"image/color"
	"log"
	"math"
	"simulador/src/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ParkingLotView representa la vista del estacionamiento.
type ParkingLotView struct {
	background  *ebiten.Image
	parkingLot  *models.ParkingLot
	carView []*CarView
}

// NewParkingLotView crea una nueva vista del estacionamiento.
func NewParkingLotView(parkingLot *models.ParkingLot) *ParkingLotView {
	bg, _, err := ebitenutil.NewImageFromFile("assets/scene/parking.png")
	if err != nil {
		log.Fatalf("Error al cargar la imagen del estacionamiento: %v", err)
	}

	var carViews []*CarView
	for _, car := range parkingLot.Cars {
		carViews = append(carViews, NewCarView(car))
	}

	return &ParkingLotView{
		background: bg,
		parkingLot: parkingLot,
		carView:   carViews,
	}
}

// Draw dibuja la vista del estacionamiento en la pantalla.
func (p *ParkingLotView) Draw(screen *ebiten.Image) {
	p.drawWalls(screen)

	op := &ebiten.DrawImageOptions{}

	b := p.background.Bounds()
	width := b.Dx()
	height := b.Dy()

	scaleX := float64(800) / float64(width)
	scaleY := float64(600) / float64(height)

	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	additionalScale := 0.7

	op.GeoM.Scale(scale*additionalScale, scale*additionalScale)

	op.GeoM.Rotate(90 * (math.Pi / 180))

	op.GeoM.Translate(780, 20)

	screen.DrawImage(p.background, op)

	for _, carView := range p.carView {
		carView.Draw(screen)
	}
}

// drawWalls dibuja las paredes del estacionamiento.
func (p *ParkingLotView) drawWalls(screen *ebiten.Image) {
	fillColor := color.RGBA{R: 204, G: 204, B: 204, A: 255}

	x := 370.0
	y := 13.0
	width := 600.0 * 0.7
	height := 820.0 * 0.7

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), fillColor, false)
}
