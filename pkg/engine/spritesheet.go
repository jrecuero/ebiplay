package engine

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	Image        *ebiten.Image
	Rows         int
	Columns      int
	Width        int
	Height       int
	FrameTypes   []string
	FrameType    string
	FrameMap     map[string][]int
	FrameIndex   int
	FrameSpeed   int
	FrameCounter int
}

func NewSpriteSheet(image *ebiten.Image, rows, columns, width, height int) *SpriteSheet {
	return &SpriteSheet{
		Image:   image,
		Rows:    rows,
		Columns: columns,
		Width:   width,
		Height:  height,
	}
}

func (s *SpriteSheet) GetFrameFor(frameType string) *ebiten.Image {
	if !s.IsValidFrameType(frameType) {
		log.Fatalf("invalid frame type %s", frameType)
	}
	// Update the frame type if it is a different one and reset all counters
	// and indexes.
	s.UpdateFrameType(frameType)
	framemap := s.FrameMap[frameType]
	if s.FrameCounter = (s.FrameCounter + 1) % s.FrameSpeed; s.FrameCounter == 0 {
		s.FrameIndex = (s.FrameIndex + 1) % 4
	}
	i := s.FrameIndex * 2
	xIndex, yIndex := framemap[i], framemap[i+1]
	return s.Image.SubImage(image.Rect(xIndex, yIndex, xIndex+s.Width, yIndex+s.Height)).(*ebiten.Image)
}

func (s *SpriteSheet) GetSpriteFromRowAndCol(row, col int) *ebiten.Image {
	x := col * s.Width
	y := row * s.Height
	return s.Image.SubImage(image.Rect(x, y, x+s.Width, y+s.Height)).(*ebiten.Image)
}

func (s *SpriteSheet) IsValidFrameType(frameType string) bool {
	for _, ft := range s.FrameTypes {
		if ft == frameType {
			return true
		}
	}
	return false
}

func (s *SpriteSheet) UpdateFrameType(frameType string) {
	if frameType != s.FrameType {
		s.FrameType = frameType
		s.FrameIndex = 0
		s.FrameCounter = 0
	}
}
