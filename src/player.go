package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	X      float64
	Y      float64
	Sprite *ebiten.Image
}

func (p *Player) Update() {
	p.X += 1
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Sprite, &options)
}

func NewPlayer() *Player {

	p := Player{}
	p.X = 32
	p.Y = 32
	p.Sprite = GetImage("Player.png")
	fmt.Println(p.Sprite)
	return &p

}
