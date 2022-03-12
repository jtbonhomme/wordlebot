package app

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jtbonhomme/wordlebot/internal/keyboard"
)

const (
	ScreenWidth  = 400
	ScreenHeight = 720
)

var backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}

// Game represents a game state.
type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
	}
	var err error
	g.board, err = NewBoard()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(keyboard.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	g.boardImage = ebiten.NewImageFromImage(img)

	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()
	if err := g.board.Update(g.input); err != nil {
		return err
	}
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(ScreenWidth, ScreenHeight)
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
