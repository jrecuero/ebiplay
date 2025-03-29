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
	frameTypes   []string
	frameType    string
	frameMap     map[string][]int
	frameIndex   int
	frameSpeed   int
	frameCounter int
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

func NewSpriteSheetImage(image *ebiten.Image, width, height int) *SpriteSheet {
	return NewSpriteSheet(image, 1, 1, width, height)
}

func (s *SpriteSheet) GetFrameFor(frameType string) *ebiten.Image {
	if !s.IsValidFrameType(frameType) {
		log.Fatalf("invalid frame type %s", frameType)
	}
	// Update the frame type if it is a different one and reset all counters
	// and indexes.
	s.UpdateFrameType(frameType)
	framemap := s.frameMap[frameType]
	if s.frameCounter = (s.frameCounter + 1) % s.frameSpeed; s.frameCounter == 0 {
		s.frameIndex = (s.frameIndex + 1) % 4
	}
	i := s.frameIndex * 2
	xIndex, yIndex := framemap[i], framemap[i+1]
	return s.Image.SubImage(image.Rect(xIndex, yIndex, xIndex+s.Width, yIndex+s.Height)).(*ebiten.Image)
}

func (s *SpriteSheet) GetFrameType() string {
	return s.frameType
}

func (s *SpriteSheet) GetSpriteFromRowAndCol(row, col int) *ebiten.Image {
	x := col * s.Width
	y := row * s.Height
	return s.Image.SubImage(image.Rect(x, y, x+s.Width, y+s.Height)).(*ebiten.Image)
}

func (s *SpriteSheet) IsValidFrameType(frameType string) bool {
	for _, ft := range s.frameTypes {
		if ft == frameType {
			return true
		}
	}
	return false
}

func (s *SpriteSheet) SetFrameMap(m map[string][]int) *SpriteSheet {
	s.frameTypes = make([]string, 0, len(m))
	for key := range m {
		s.frameTypes = append(s.frameTypes, key)
	}
	s.frameMap = m
	s.frameType = s.frameTypes[0] // by default set to the first entry.
	return s
}

func (s *SpriteSheet) SetFrameSpeed(speed int) *SpriteSheet {
	s.frameSpeed = speed
	return s
}

func (s *SpriteSheet) SetFrameType(t string) *SpriteSheet {
	s.frameType = t
	return s
}

func (s *SpriteSheet) UpdateFrameType(frameType string) {
	if frameType != s.frameType {
		s.frameType = frameType
		s.frameIndex = 0
		s.frameCounter = 0
	}
}
