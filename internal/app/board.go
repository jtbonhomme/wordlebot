package app

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Board represents the game board.
type Board struct {
}

// NewBoard generates a new Board with giving a size.
func NewBoard() (*Board, error) {
	b := &Board{}
	return b, nil
}

// Update updates the board state.
func (b *Board) Update(input *Input) error {
	return nil
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
}
