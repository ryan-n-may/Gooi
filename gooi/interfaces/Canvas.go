package Interfaces 



type Canvas_Interface interface {
	GetEventHandler() EventHandler_Interface
	SetEventHandler(EventHandler_Interface)

	SetWindow(Window_Interface)
	GetWindow() Window_Interface

	GetWidth() float32
	GetHeight() float32

	GetPrograms() uint32

	AddDisplayable(Displayable)
	GetDisplayables() []Displayable
	SetDisplayables([]Displayable)

	SetBackgroundColour([3]float32)

	CountComponents() int
	
	CompileCanvasShader(string, uint32) uint32

	Draw()
	RefreshCanvas()
}