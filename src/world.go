package main

import "github.com/hajimehoshi/ebiten"

type World struct {
	PlayerObj *Player
}

func NewWorld() *World {

	w := World{}
	w.PlayerObj = NewPlayer()
	return &w

}

func (w *World) Update(screen *ebiten.Image) {

	w.PlayerObj.Update()
	w.PlayerObj.Draw(screen)

}
