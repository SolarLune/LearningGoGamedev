package main

import "github.com/hajimehoshi/ebiten"

type World struct {
	PlayerObj *Player
	Level     int
}

func NewWorld() *World {

	w := World{}
	w.PlayerObj = NewPlayer()
	return &w

}

func (w *World) DrawTiles(screen *ebiten.Image) {

	drawOptions := ebiten.DrawImageOptions{}

	screen.DrawImage(GetImage("Tileset.png"), &drawOptions)

}

func (w *World) Update(screen *ebiten.Image) {

	w.DrawTiles(screen)

	w.PlayerObj.Update()
	w.PlayerObj.Draw(screen)

}
