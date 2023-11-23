package Interfaces 



type Canvas_Interface interface {
	GetEventHandler() EventHandler_Interface
	SetEventHandler(EventHandler_Interface)

	SetWindow(Window_Interface)
	GetWindow() Window_Interface

	GetPrograms() uint32

	AddComponent(Drawable_Interface)
	SetComponents([]Drawable_Interface)

	SetBackgroundColour([3]float32)

	CountComponents() int
	
	CompileCanvasShader(string, uint32) uint32

	Draw()
	RefreshCanvas()
}