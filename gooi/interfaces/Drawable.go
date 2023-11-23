package Interfaces

type Drawable_Interface interface {
	Draw()
	Redraw()

	SetPos(float32, float32)
	GetPos() (float32, float32)

	SetPosZ(float32)
	GetPosZ() float32

	GetBounds() (float32, float32)
}
