package app

import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jtbonhomme/wordlebot/internal/fonts"
)

const (
	tileHeight int = 70
	tileWidth  int = 30
	dpi            = 72
)

var (
	smallFont  font.Face
	normalFont font.Face
	bigFont    font.Face
	textColor  = color.RGBA{0xaa, 0xa8, 0xaf, 0xff}
)

// Board represents the game board.
type Board struct {
}

// NewBoard generates a new Board with giving a size.
func NewBoard() (*Board, error) {
	tt, err := opentype.Parse(fonts.OutfitRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	smallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	bigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    30,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	b := &Board{}
	return b, nil
}

// Update updates the board state.
func (b *Board) Update(input *Input) error {
	return nil
}

func drawText(boardImage *ebiten.Image, f font.Face, x, y int, str string) {
	bound, _ := font.BoundString(f, str)
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	y += h
	text.Draw(boardImage, str, f, x, y, textColor)
}

func drawKeyboard(boardImage *ebiten.Image) {
	f := normalFont
	var x, y int
	A := "A"
	for i := 0; i < 26; i++ {
		if i%10 == 0 {
			x = 0
			y++
			if y == 3 {
				x += 2
			}
		}
		str := string(byte(A[0]) + byte(i))
		bound, _ := font.BoundString(f, str)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		lx := x*38 + (tileHeight-w)/2 - 2
		ly := 280 + (y-1)*76 + (tileWidth-h)/2 + h

		drawText(boardImage, f, lx, ly, str)
		x++
	}
}

func draw123(boardImage *ebiten.Image) {
	f := normalFont
	var x, y int
	A := "1"
	for i := 0; i < 3; i++ {
		str := string(byte(A[0]) + byte(i))
		bound, _ := font.BoundString(f, str)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		lx := 135 + x*38 + (tileHeight-w)/2 - 2
		ly := 280 + (y-1)*76 + (tileWidth-h)/2 + h

		drawText(boardImage, f, lx, ly, str)
		x++
	}
}

func drawEnter(boardImage *ebiten.Image) {
	f := normalFont
	str := "Enter"
	bound, _ := font.BoundString(f, str)
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	lx := 16 + (tileHeight-w)/2
	ly := 436 + (2*tileWidth-h)/2

	drawText(boardImage, f, lx, ly, str)
}

func drawDel(boardImage *ebiten.Image) {
	f := normalFont
	str := "Del"
	bound, _ := font.BoundString(f, str)
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	lx := 318 + (tileHeight-w)/2
	ly := 436 + (2*tileWidth-h)/2

	drawText(boardImage, f, lx, ly, str)
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	// Draw keyboard
	drawKeyboard(boardImage)
	draw123(boardImage)
	drawEnter(boardImage)
	drawDel(boardImage)
}
