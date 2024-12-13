package scenes

import (
	"simulador/src/models"
	"simulador/src/views"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneParking struct {
	parkingView *views.ParkingLotView
	parkingLot *models.ParkingLot
	CarViews   []*views.CarView
	SpaceViews []*views.SpaceView
	RoadView   *views.RoadView
}

// NewSceneParking crea una nueva escena de estacionamiento a partir del modelo lógico.
func NewSceneParking(parkingLot *models.ParkingLot, carViews []*views.CarView, spaceViews []*views.SpaceView) *SceneParking {
	roadView := views.NewRoadView()
	parkingV := views.NewParkingLotView(parkingLot)

	return &SceneParking{
		parkingLot: parkingLot,
		parkingView: parkingV,
		CarViews:   carViews,
		SpaceViews: spaceViews,
		RoadView:   roadView,
	}
}

func (s *SceneParking) Update() error {
	s.parkingLot.Update()
	return nil
}

// Draw dibuja todas las vistas en la pantalla.
func (s *SceneParking) Draw(screen *ebiten.Image) {
	s.parkingView.Draw(screen)

	s.RoadView.Draw(screen)

	for _, spaceView := range s.SpaceViews {
		spaceView.Draw(screen)
	}

	for _, carView := range s.CarViews {
		carView.Draw(screen)
	}
}

// Layout retorna el tamaño lógico de la escena.
func (s *SceneParking) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}