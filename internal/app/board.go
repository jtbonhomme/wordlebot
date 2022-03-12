package app

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jtbonhomme/wordlebot/internal/fonts"
)

const (
	tileHeight    int = 70
	tileWidth     int = 30
	bigTileHeight int = 90
	bigTileWidth  int = 50
	dpi               = 72
)

var (
	smallFont  font.Face
	normalFont font.Face
	bigFont    font.Face
	textColor  = color.RGBA{0xaa, 0xa8, 0xaf, 0xff}
)

// Board represents the game board.
type Board struct {
	currentWord string
	guessedWord string
}

// NewBoard generates a new Board with giving a size.
func NewBoard() (*Board, error) {
	tt, err := opentype.Parse(fonts.CircularMedium_ttf)
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
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	bigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	b := &Board{}
	return b, nil
}

var Keyb = [][]string{
	{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
	{"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"},
	{"U", "V", "W", "X", "Y", "Z"},
}

// Update updates the board state.
func (b *Board) Update(i *Input) error {
	if i.IsSettled() {
		// get key pressed
		x, y := i.LastPos()
		index := int(math.Ceil(float64(x)/40.0)) - 1
		if index < 0 || index > 9 {
			return fmt.Errorf("index error %d", index)
		}
		if y > 286 && y < 356 {
			if len(b.currentWord) < 5 {
				b.currentWord += Keyb[0][index]
			}
		} else if y > 355 && y < 431 {
			if len(b.currentWord) < 5 {
				b.currentWord += Keyb[1][index]
			}
		} else if y > 430 && y < 504 && x < 239 {
			if len(b.currentWord) < 5 {
				b.currentWord += Keyb[2][index]
			}
		} else if y > 430 && y < 504 && x > 238 && x < 315 {
			log.Printf("Press Enter")
		} else if y > 430 && y < 504 && x > 314 && x < 392 {
			if len(b.currentWord) > 0 {
				b.currentWord = b.currentWord[:len(b.currentWord)-1]
			}
		}

	}
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
		lx := 66 + i*w + (bigTileHeight-w)/2
		ly := 500 + (2*bigTileWidth-h)/2

		drawText(boardImage, f, lx, ly, str)
	}
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	// Draw keyboard
	drawKeyboard(boardImage)
	draw123(boardImage)
	drawEnter(boardImage)
	drawDel(boardImage)
	drawCurrentWord(boardImage, b.currentWord)
	drawGuessedWord(boardImage, b.guessedWord)
}

// SetGuessedWord set guessed word
func (b *Board) SetGuessedWord(w string) {
	b.guessedWord = w
}
