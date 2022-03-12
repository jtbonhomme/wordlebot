package app

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"
	"math"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jtbonhomme/wordlebot/internal/fonts"
	"github.com/jtbonhomme/wordlebot/internal/wordle"
	"github.com/jtbonhomme/wordlebot/internal/words"
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
	result      string
	guessedWord string
	wg          *wordle.Game
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
	b.wg = wordle.New(words.Words)
	return b, nil
}

var Keyb = [][]string{
	{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
	{"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"},
	{"U", "V", "W", "X", "Y", "Z"},
}
var Res = []string{"0", "1", "2"}

// Update updates the board state.
func (b *Board) Update(i *Input) error {
	if i.IsSettled() {
		// get key pressed
		x, y := i.LastPos()
		index := int(math.Ceil(float64(x)/40.0)) - 1
		if index < 0 || index > 9 {
			return fmt.Errorf("index error %d", index)
		}
		if y > 206 && y < 281 && x > 146 && x < 260 { // result choice
			if len(b.result) < 5 /*&& len(b.currentWord) > len(b.result)*/ {
				index2 := int(math.Ceil(float64(x-146)/40.0)) - 1
				b.result += Res[index2]
			}
		} else if y > 209 && y < 356 { // first letters row
			if len(b.currentWord) < 5 /*&& len(b.currentWord) == len(b.result) */ {
				b.currentWord += Keyb[0][index]
			}
		} else if y > 355 && y < 431 { // second letters row
			if len(b.currentWord) < 5 /*&& len(b.currentWord) == len(b.result)*/ {
				b.currentWord += Keyb[1][index]
			}
		} else if y > 430 && y < 504 && x < 239 { // third letters row
			if len(b.currentWord) < 5 /*&& len(b.currentWord) == len(b.result)*/ {
				b.currentWord += Keyb[2][index]
			}
		} else if y > 430 && y < 504 && x > 238 && x < 315 { // enter
			b.guessedWord = b.play()
			b.result = ""
			b.currentWord = ""
		} else if y > 430 && y < 504 && x > 314 && x < 392 { // delete
			if len(b.currentWord) > 0 {
				b.currentWord = b.currentWord[:len(b.currentWord)-1]
				if len(b.result) > len(b.currentWord) {
					b.result = b.result[:len(b.result)-1]
				}
			}
		}
	}
	return nil
}

// play looks for the next best word
func (b *Board) play() string {
	var result []int
	if len(b.currentWord) != 5 || len(b.result) != 5 {
		return ""
	}

	for _, c := range b.result {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			return ""
		}
		result = append(result, i)
	}

	b.wg.Filter(b.currentWord, result, true)
	b.wg.Commit()

	var maxEntropy float64
	var bestWord string
	if len(b.wg.Words()) == 0 {
		return ""
	}

	for _, w := range b.wg.Words() {
		b.SetGuessedWord(w)

		e, _, err := b.wg.Entropy(w, true)
		if err != nil {
			return ""
		}
		if e > maxEntropy {
			maxEntropy = e
			bestWord = w
		}
	}

	if len(b.wg.Words()) > 0 && bestWord == "" && maxEntropy == 0 {
		return ""
	}
	return bestWord
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	// Draw keyboard
	drawKeyboard(boardImage)
	draw123(boardImage)
	drawEnter(boardImage)
	drawDel(boardImage)
	drawCurrentWord(boardImage, b.currentWord)
	drawResult(boardImage, b.result)
	drawGuessedWord(boardImage, b.guessedWord)
}

// SetGuessedWord set guessed word
func (b *Board) SetGuessedWord(w string) {
	b.guessedWord = w
}
