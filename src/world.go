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
	LevelData [][]int
}

func NewWorld() *World {

	w := World{}
	w.Level = 1
	w.LevelData = make([][]int, 0)
	w.ResetWorld()
	return &w

}

func (w *World) ResetWorld() {

	assets = make(map[string]interface{}, 0)

	levelMap := GetImage("levels/Level" + strconv.Itoa(w.Level) + ".png")

	w.PlayerObj = NewPlayer()

	// Populate the World's LevelData array with values

	for y := 0; y < levelMap.Bounds().Dy(); y++ {

		w.LevelData = append(w.LevelData, []int{})

		for x := 0; x < levelMap.Bounds().Dx(); x++ {

			wx := x * 16
			wy := y * 16

			cr, cg, cb, _ := levelMap.At(x, y).RGBA()

			cellCode := 0

			if cr == 0 && cg == 0 && cb == 0 {
				cellCode = LevelSolid
			} else if cr == 0 && cg == 0 && cb == 65535 {
				cellCode = LevelPlayerStart
			}

			if cellCode == LevelPlayerStart {
				w.PlayerObj.X = float64(wx)
				w.PlayerObj.Y = float64(wy)
			}

			w.LevelData[y] = append(w.LevelData[y], cellCode)

		}

	}

}

func (w *World) CodeAt(x, y int) int {

	if x < 0 || x >= len(w.LevelData[0]) || y < 0 || y >= len(w.LevelData) {
		return LevelSolid
	}

	return w.LevelData[y][x]

}

func (w *World) DrawTiles(screen *ebiten.Image) {

	drawOptions := ebiten.DrawImageOptions{}

	for y := 0; y < len(w.LevelData); y++ {

		for x := 0; x < len(w.LevelData[y]); x++ {

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
