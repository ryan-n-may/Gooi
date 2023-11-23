package Interfaces

import (
	glfw "github.com/go-gl/glfw/v3.2/glfw"
)

type MouseHandler_Interface interface {
	RegisterClickableToHandler(Clickable_Interface)
	CheckClick(int, int, int, glfw.ModifierKey) 
	GetClickData(*glfw.Window) (float64, float64, int)
	GetClickables() []Clickable_Interface
	SetName(string)
	GetName() string
}
