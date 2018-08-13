package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var assets map[string]interface{} = make(map[string]interface{}, 0)

func GetImage(imageName string) *ebiten.Image {

	_, exists := assets[imageName]

	if !exists {

		var err error

		assets[imageName], _, err = ebitenutil.NewImageFromFile("assets/graphics/"+imageName, ebiten.FilterLinear)

		if err != nil {

			log.Fatal("ERROR: Couldn't load " + imageName + "!!!")

		}

	}

	img := assets[imageName].(*ebiten.Image)
	return img

}
