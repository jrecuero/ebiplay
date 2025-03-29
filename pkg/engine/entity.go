package engine

type IEntity interface {
	IBase
	GetHeight() int
	GetPos() (float64, float64)
	GetSize() (int, int)
	GetWidth() int
	GetX() float64
	GetY() float64
	SetHeight(int) IEntity
	SetPos(float64, float64) IEntity
	SetSize(int, int) IEntity
	SetWidth(int) IEntity
	SetX(float64) IEntity
	SetY(float64) IEntity
}

type Entity struct {
	*Base
	x      float64
	y      float64
	height int
	width  int
}

func NewEntity(name string, x, y float64, w, h int) *Entity {
	return &Entity{
		Base:   NewBase(name),
		x:      x,
		y:      y,
		width:  w,
		height: h,
	}
}

func (e *Entity) GetHeight() int {
	return e.height
}

func (e *Entity) GetPos() (float64, float64) {
	return e.x, e.y
}

func (e *Entity) GetSize() (int, int) {
	return e.width, e.height
}

func (e *Entity) GetWidth() int {
	return e.width
}

func (e *Entity) GetX() float64 {
	return e.x
}

func (e *Entity) GetY() float64 {
	return e.y
}

func (e *Entity) SetHeight(h int) IEntity {
	e.height = h
	return e
}

func (e *Entity) SetPos(x, y float64) IEntity {
	e.x, e.y = x, y
	return e
}

func (e *Entity) SetSize(w, h int) IEntity {
	e.width, e.height = w, h
	return e
}

func (e *Entity) SetWidth(w int) IEntity {
	e.width = w
	return e
}

func (e *Entity) SetX(x float64) IEntity {
	e.x = x
	return e
}

func (e *Entity) SetY(y float64) IEntity {
	e.y = y
	return e
}

var _ IEntity = (*Entity)(nil)
