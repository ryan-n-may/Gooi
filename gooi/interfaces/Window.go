package Interfaces

import (
	glfw 		"github.com/go-gl/glfw/v3.2/glfw"
)

type Window_Interface interface {
	SetWindowHeight(float32)
	GetWindowHeight() *float32

	SetWindowWidth(float32)
	GetWindowWidth() *float32

	SetResizable(int)
	GetResizable() int

	SetWindowTitle(string)
	GetWindowTitle() string

	SetBackgroundColour([3]float32)

	CloseWindow()
	OpenWindow()

	SetCloseMethod(func())
	GetCloseMethod() func()

	SetOpenMethod(func())
	GetOpenMethod() func()

	RunWindow()

	SetWindow(*glfw.Window)
	GetWindow() *glfw.Window

	SetMouseHandler(MouseHandler_Interface)
	GetMouseHandler() MouseHandler_Interface
}
