package Principal

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//pacimages "github.com/kgosse/pacmanresources/images"
)

var logo *ebiten.Image
var mplusNormalFont font.Face

type menu struct{}

//Create a new Scene
func newMenu() *menu {
	var err error
	s := &menu{} //create the structure pointer scene
	logo, _, err = ebitenutil.NewImageFromFile("pacman_logo.jpg", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    15,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return s //return the pointer structure scene
}

//it works to show things in the screen
func (m *menu) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() { // when IsDrawingSkipped is true, the rendered result is not adopted.
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	//_, h := logo.Size()
	//Obtener tamño de la ventana: screen.Bounds().Size().Y
	op.GeoM.Translate(38, 10)
	op.GeoM.Scale(0.8, 0.8)
	screen.DrawImage(logo, op)

	str := "PRESS ENTER TO START"
	//str2 := "By Ulises Bojorquez and Santiago Yeomans"
	x := 75
	y := (screen.Bounds().Size().Y / 2) + 65
	text.Draw(screen, str, mplusNormalFont, x, y, color.White)
	//text.Draw(screen, str2, mplusNormalFont, 10, y+30, color.White)
	return nil
}
