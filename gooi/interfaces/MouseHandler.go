package Interfaces

import (
	glfw "github.com/go-gl/glfw/v3.2/glfw"
)

type MouseHandler_Interface interface {
	RegisterClickableToHandler(Clickable)
	CheckClick(float32, float32, int, glfw.ModifierKey) 
	GetClickData(*glfw.Window) (float32, float32, int)
	GetClickables() []Clickable
	SetName(string)
	GetName() string
}
