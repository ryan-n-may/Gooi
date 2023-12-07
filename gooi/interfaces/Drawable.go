package Interfaces

import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
)


type Drawable interface {
	SetXYZ([]float32) 
	GetXYZ() []float32

	SetRGB([]float32) 
	GetRGB() []float32

	SetDIM([]float32) 
	GetDIM() []float32

	SetVAO([]Drawing) 
	GetVAO() []Drawing	

	GetWidth() float32
	GetHeight() float32
}

type Drawing struct {
	VAO 		uint32
	DrawMode 	uint32
}

type Displayable interface {
	SetPos(float32, float32, float32)
	GetPos() (float32, float32, float32)

	GetWidth() float32
	GetHeight() float32

	GetMasterStruct() Displayable
	SetMasterStruct(Displayable)

	Draw()
	Redraw()
}

type Clickable interface {
	GetClickableBounds() (float32, float32, float32, float32, float32)
	Click(*bool, int, float32, float32, glfw.ModifierKey) 
}
