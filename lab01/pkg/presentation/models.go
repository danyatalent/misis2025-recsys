package presentation

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// ResultBox
//
// Структура для результирующего контейнера
type ResultBox struct {
	Container   *fyne.Container
	Rect        *canvas.Rectangle
	ResultLabel *widget.Label
	JSONView    *widget.Entry
}
