package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	X      float64
	Y      float64
	Sprite *ebiten.Image
}

func NewPlayer() *Player {

	AddAction("right", KeyInput{Key: ebiten.KeyRight})
	AddAction("left", KeyInput{Key: ebiten.KeyLeft})

	p := Player{}
	p.X = 32
	p.Y = 32
	p.Sprite = GetImage("Player.png")
	return &p

}

func (p *Player) Update() {
	moveSpd := 1
	p.X += float64(moveSpd * (IsActionDownI("right") - IsActionDownI("left")))
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Sprite, &options)
}
