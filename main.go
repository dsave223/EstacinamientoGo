package main

import (
    "fmt"
    "simulador/src/models"
    "simulador/src/scenes"
    "simulador/src/views"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
)

type SimuladorState struct {
    autosActivos []*models.Auto
    pausaGlobal  bool
}