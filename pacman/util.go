package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
)

func isWall(e elem) bool {
	if w0 <= e && e <= w24 {
		return true
	}
	return false
}

func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	handleError(err)
	ebImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handleError(err)
	return ebImg
}

func loadImages(images [][]byte) []*ebiten.Image {
	var res []*ebiten.Image
	size := len(images)
	for i := 0; i < size; i++ {
		res = append(res, loadImage(images[i]))
	}
	return res
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
