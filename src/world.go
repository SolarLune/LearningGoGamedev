package main

import (
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

type World struct {
	PlayerObj *Player
	Level     int
}

func NewWorld() *World {

	w := World{}
	w.Level = 1
	w.PlayerObj = NewPlayer()
	return &w

}

func (w *World) DrawTiles(screen *ebiten.Image) {

	levelMap := GetImage("levels/Level" + strconv.Itoa(w.Level) + ".png")

	size := levelMap.Bounds()

	drawOptions := ebiten.DrawImageOptions{}

	for y := 0; y < size.Dy(); y++ {

		for x := 0; x < size.Dx(); x++ {

			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(float64(x*16), float64(y*16))

			cr, cg, cb, _ := levelMap.At(x, y).RGBA()

			if cr == 0 && cg == 0 && cb == 0 {
				r := image.Rect(16, 0, 32, 16)
				drawOptions.SourceRect = &r
				screen.DrawImage(GetImage("Tileset.png"), &drawOptions)
			}

		}

	}

}

func (w *World) Update(screen *ebiten.Image) {

	w.DrawTiles(screen)

	w.PlayerObj.Update()
	w.PlayerObj.Draw(screen)

}
