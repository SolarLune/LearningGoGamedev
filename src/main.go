package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var screenW = 320
var screenH = 180

var world = NewWorld()

func update(screen *ebiten.Image) error {

	var finishedState error

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		finishedState = errors.New("Exit")
	}

	world.Update(screen)

	return finishedState
}

func main() {

	if err := ebiten.Run(update, screenW, screenH, 4, "Learning Go Gamedev"); err != nil {
		log.Fatal(err)
	}

}
