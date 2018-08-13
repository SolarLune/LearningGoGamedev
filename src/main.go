package main

import (
	"errors"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var screenW = 320
var screenH = 180

var world = NewWorld()

func update(screen *ebiten.Image) error {

	var finishedState error

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		finishedState = errors.New("Exit")
	}

	ebitenutil.DrawRect(screen, 0, 0, float64(screenW), float64(screenH), color.White)
	world.Update(screen)

	return finishedState
}

func main() {

	if err := ebiten.Run(update, screenW, screenH, 4, "Learning Go Gamedev"); err != nil {
		log.Fatal(err)
	}

}
