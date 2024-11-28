package scenes

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
)

type Scene struct {
    scene     fyne.Window
    container *fyne.Container
}

func NewScene(scene fyne.Window) *Scene {
    return &Scene{scene: scene, container: nil}
}

func (s *Scene) Init() {
    rect := canvas.NewImageFromURI(storage.NewFileURI("./assets/scene/frame.png"))
    rect.Resize(fyne.NewSize(1280, 580))
    rect.Move(fyne.NewPos(0, 0))

    // Crear contenedor con la imagen de fondo
    s.container = container.NewWithoutLayout(rect)
    s.scene.SetContent(s.container)
}

func (s *Scene) AddWidget(widget fyne.Widget) {
    s.container.Add(widget)
    s.container.Refresh()
}

func (s *Scene) AddContainer(cont *fyne.Container) {
    s.container.Add(cont)
    s.container.Refresh()
}

func (s *Scene) AddImage(image *canvas.Image) {
    s.container.Add(image)
    s.container.Refresh()
}

func (s *Scene) GetContainer() *fyne.Container {
    return s.container
} 