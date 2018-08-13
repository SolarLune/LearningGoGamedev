package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var _assets map[string]interface{} = make(map[string]interface{}, 0)

func GetImage(imageName string) *ebiten.Image {

	_, exists := _assets[imageName]

	if !exists {
		_assets[imageName], _, _ = ebitenutil.NewImageFromFile("assets/graphics/"+imageName, ebiten.FilterLinear)
	}

	img := _assets[imageName].(*ebiten.Image)
	return img

}
