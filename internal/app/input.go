package app

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

// Dir represents a direction.
type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type touchState int

const (
	touchStateNone touchState = iota
	touchStatePressing
	touchStateSettled
	touchStateInvalid
)

// Input represents the current key states.
type Input struct {
	mouseState       mouseState
	mouseCurrentPosX int
	mouseCurrentPosY int

	touches          []ebiten.TouchID
	touchState       touchState
	touchID          ebiten.TouchID
	touchCurrentPosX int
	touchCurrentPosY int

	lastPosX int
	lastPosY int
}

// NewInput generates a new Input object.
func NewInput() *Input {
	return &Input{}
}

// ToString displays input object as a String
func (i *Input) ToString() string {
	return fmt.Sprintf("mouse state: %d (%d, %d)\ntouch state: %d (%d, %d)", i.mouseState, i.mouseCurrentPosX, i.mouseCurrentPosY, i.touchState, i.touchCurrentPosX, i.touchCurrentPosY)
}

// LastPos return last position
func (i *Input) LastPos() (int, int) {
	return i.lastPosX, i.lastPosY
}

// Update updates the current input states.
func (i *Input) Update() {
	switch i.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i.mouseCurrentPosX = x
			i.mouseCurrentPosY = y
			i.lastPosX = x
			i.lastPosY = y
			i.mouseState = mouseStatePressing
		}
	case mouseStatePressing:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			i.mouseState = mouseStateSettled
		}
	case mouseStateSettled:
		i.mouseState = mouseStateNone
	}

	i.touches = ebiten.AppendTouchIDs(i.touches[:0])
	switch i.touchState {
	case touchStateNone:
		if len(i.touches) == 1 {
			i.touchID = i.touches[0]
			x, y := ebiten.TouchPosition(i.touches[0])
			i.touchCurrentPosX = x
			i.touchCurrentPosY = y
			i.lastPosX = x
			i.lastPosY = y
			i.touchState = touchStatePressing
		}
	case touchStatePressing:
		if len(i.touches) >= 2 {
			break
		}
		if len(i.touches) == 1 {
			if i.touches[0] != i.touchID {
				i.touchState = touchStateInvalid
			} else {
				x, y := ebiten.TouchPosition(i.touches[0])
				i.touchCurrentPosX = x
				i.touchCurrentPosY = y
			}
			break
		}
		if len(i.touches) == 0 {
			i.touchState = touchStateSettled
		}
	case touchStateSettled:
		i.touchState = touchStateNone
	case touchStateInvalid:
		if len(i.touches) == 0 {
			i.touchState = touchStateNone
		}
	}
}

// IsSettled return true if touchState or mouseState is settled
func (i *Input) IsSettled() bool {
	if i.mouseState == mouseStateSettled || i.touchState == touchStateSettled {
		return true
	}
	return false
}

// Draw draws the input to the given boardImage.
func (i *Input) Draw(boardImage *ebiten.Image) {
	f := normalFont
	str := i.ToString()
	lx := 40
	ly := 650

	drawText(boardImage, f, lx, ly, str)

}
