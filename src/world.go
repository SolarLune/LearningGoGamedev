package main

import (
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

const (
	LevelEmpty = iota
	LevelSolid
	LevelPlayerStart
)

type World struct {
	PlayerObj *Player
	Level     int
}

func NewWorld() *World {

	w := World{}
	w.Level = 1
	w.ResetWorld()
	return &w

}

func (w *World) ResetWorld() {

	lvl := "levels/Level" + strconv.Itoa(w.Level) + ".png"

	assets = make(map[string]interface{}, 0)

	levelMap := GetImage(lvl)

	w.PlayerObj = NewPlayer()

	for y := 0; y < levelMap.Bounds().Dy(); y++ {

		for x := 0; x < levelMap.Bounds().Dx(); x++ {

			wx := x * 16
			wy := y * 16

			if w.CodeAt(x, y) == LevelPlayerStart {
				w.PlayerObj.X = float64(wx)
				w.PlayerObj.Y = float64(wy)
			}

		}

	}

}

func (w *World) CodeAt(x, y int) int {

	levelMap := GetImage("levels/Level" + strconv.Itoa(w.Level) + ".png")

	cr, cg, cb, _ := levelMap.At(x, y).RGBA()

	if cr == 0 && cg == 0 && cb == 0 {
		return LevelSolid
	} else if cr == 0 && cg == 0 && cb == 65535 {
		return LevelPlayerStart
	}

	return LevelEmpty

}

func (w *World) DrawTiles(screen *ebiten.Image) {

	levelMap := GetImage("levels/Level" + strconv.Itoa(w.Level) + ".png")

	drawOptions := ebiten.DrawImageOptions{}

	for y := 0; y < levelMap.Bounds().Dy()+1; y++ {

		for x := 0; x < levelMap.Bounds().Dx(); x++ {

			wx := x * 16
			wy := y * 16

			r := image.Rectangle{}

			if w.CodeAt(x, y) == LevelSolid {
				r = image.Rect(0, 0, 16, 16)

				if w.CodeAt(x, y-1) == LevelSolid {
					r = image.Rect(0, 16, 16, 32)
				}

				drawOptions.SourceRect = &r
			}

			if !r.Empty() {
				drawOptions.GeoM.Reset()
				drawOptions.GeoM.Translate(float64(wx), float64(wy))
				screen.DrawImage(GetImage("Tileset.png"), &drawOptions)
			}

		}

	}

}

func (w *World) Update(screen *ebiten.Image) {

	w.DrawTiles(screen)

	w.PlayerObj.Update()
	w.PlayerObj.Draw(screen)

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		w.ResetWorld()
	}

}
