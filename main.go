package main

import (
	"log"
	"simulador/src/models"
	"simulador/src/scenes"
	"simulador/src/views"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Inicializar el estacionamiento con autos y espacios
	parkingLot := models.NewParkingLot()

	// Inicializar las vistas de los autos y los espacios
	carViews := []*views.CarView{}
	spaceViews := []*views.SpaceView{}

	// Crear vistas correspondientes a cada auto
	for _, car := range parkingLot.Cars {
		carViews = append(carViews, views.NewCarView(car))
	}

	// Crear vistas correspondientes a cada espacio
	for _, space := range parkingLot.Spaces {
		spaceViews = append(spaceViews, views.NewSpaceView(space))
	}

	// Crear la escena del estacionamiento con los componentes
	scene := scenes.NewSceneParking(parkingLot, carViews, spaceViews)

	// Configurar la ventana de Ebiten
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Simulador de Estacionamiento")

	// Inicializar el juego
	if err := ebiten.RunGame(scene); err != nil {
		log.Fatal(err)
	}
}