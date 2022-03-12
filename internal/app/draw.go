package app

import (
	"strings"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

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
	A := "0"
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
	lx := 245 + (tileHeight-w)/2
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

func drawCurrentWord(boardImage *ebiten.Image, word string) {
	f := bigFont
	if len(word) > 5 {
		return
	}
	word = strings.ToUpper(word)
	for i := 0; i < len(word); i++ {
		str := string(word[i])
		bound, _ := font.BoundString(f, str)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		lx := 66 + i*50 + (bigTileHeight-w)/2
		ly := 100 + (2*bigTileWidth-h)/2

		drawText(boardImage, f, lx, ly, str)
	}
}

func drawGuessedWord(boardImage *ebiten.Image, word string) {
	f := normalFont
	if len(word) > 5 {
		return
	}
	word = strings.ToUpper(word)
	for i := 0; i < len(word); i++ {
		str := string(word[i])
		bound, _ := font.BoundString(f, str)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		lx := 95 + i*30 + (bigTileHeight-w)/2
		ly := 500 + (2*bigTileWidth-h)/2

		drawText(boardImage, f, lx, ly, str)
	}
}

func drawResult(boardImage *ebiten.Image, word string) {
	f := normalFont
	if len(word) > 5 {
		return
	}
	word = strings.ToUpper(word)
	for i := 0; i < len(word); i++ {
		str := string(word[i])
		bound, _ := font.BoundString(f, str)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		lx := 95 + i*30 + (bigTileHeight-w)/2
		ly := 530 + (2*bigTileWidth-h)/2

		drawText(boardImage, f, lx, ly, str)
	}
}
