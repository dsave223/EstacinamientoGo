package views

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RoadView struct {
	image1 *ebiten.Image
	image2 *ebiten.Image
	image3 *ebiten.Image
	image4 *ebiten.Image
}

func NewRoadView() *RoadView {
	img1, _, err := ebitenutil.NewImageFromFile("assets/scene/road1.png")
	if err != nil {
		log.Fatalf("Error al cargar la primera imagen de carretera: %v", err)
	}
	img2, _, err := ebitenutil.NewImageFromFile("assets/scene/road2.png")
	if err != nil {
		log.Fatalf("Error al cargar la segunda imagen de carretera: %v", err)
	}
	img3, _, err := ebitenutil.NewImageFromFile("assets/scene/road3.png")
	if err != nil {
		log.Fatalf("Error al cargar la segunda imagen de carretera: %v", err)
	}
	img4, _, err := ebitenutil.NewImageFromFile("assets/scene/road4.png")
	if err != nil {
		log.Fatalf("Error al cargar la segunda imagen de carretera: %v", err)
	}

	return &RoadView{
		image1: img1,
		image2: img2,
		image3: img3,
		image4: img4,
	}
}

// Draw dibuja las tres imágenes de carretera en la pantalla.
func (r *RoadView) Draw(screen *ebiten.Image) {

	//carretera 1
	road1 := &ebiten.DrawImageOptions{}
	road1.GeoM.Scale(0.13, 0.13)
	road1.GeoM.Translate(88.5, 98)
	screen.DrawImage(r.image1, road1)

	road2 := &ebiten.DrawImageOptions{}
	road2.GeoM.Scale(0.13, 0.13)
	road2.GeoM.Translate(88.5, 360)
	screen.DrawImage(r.image1, road2)

	road3 := &ebiten.DrawImageOptions{}
	road3.GeoM.Scale(0.13, 0.13)
	road3.GeoM.Translate(88.5, 490)
	screen.DrawImage(r.image1, road3)

	road4 := &ebiten.DrawImageOptions{}
	road4.GeoM.Scale(0.13, 0.13)
	road4.GeoM.Translate(88.5, 0)
	screen.DrawImage(r.image1, road4)

	//Enlace veicular
	ev1 := &ebiten.DrawImageOptions{}
	ev1.GeoM.Scale(0.13, -0.13)
	ev1.GeoM.Rotate(90 * (math.Pi / 180))
	ev1.GeoM.Translate(88, 231)
	screen.DrawImage(r.image2, ev1)

	ev2 := &ebiten.DrawImageOptions{}
	ev2.GeoM.Scale(0.13, 0.13)
	ev2.GeoM.Translate(220, 212)
	screen.DrawImage(r.image4, ev2)

	//Carreteras con curva
	cc1 := &ebiten.DrawImageOptions{}
	cc1.GeoM.Scale(-0.13, 0.13)
	cc1.GeoM.Rotate(90 * (math.Pi / 180))
	cc1.GeoM.Translate(382, 259)
	screen.DrawImage(r.image3, cc1)

	cc2 := &ebiten.DrawImageOptions{}
	cc2.GeoM.Scale(0.13, 0.13)
	cc2.GeoM.Rotate(90 * (math.Pi / 180))
	cc2.GeoM.Translate(382, 341)
	screen.DrawImage(r.image3, cc2)
}