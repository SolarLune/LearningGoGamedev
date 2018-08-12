package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	return nil
}

func main() {

	if err := ebiten.Run(update, 320, 240, 2, "Test"); err != nil {
		log.Fatal(err)
	}

}
