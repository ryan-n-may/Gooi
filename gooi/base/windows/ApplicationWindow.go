package windows
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	comp 	"gooi/base/components"
	intf 	"gooi/interfaces"
	fmt  	"fmt"
	runtime "runtime"
	colours "gooi/base/colours"
	log 	"log"
	time 	"time"
)
var (
	resizing bool
)

type ApplicationWindow_Struct struct {
	WindowCanvas *comp.Canvas_Struct
	MouseHandler intf.MouseHandler_Interface
	Window *glfw.Window
	Width, Height *float32
	Initial_Width, Initial_Height float32
	Title string
	Resizable int
	CloseMethod func()
	OpenMethod func() 
	err error

	Theme *colours.Theme
}
func NewWindow(title string, width, height float32) *ApplicationWindow_Struct {
	log.Println("new [Window] struct created.")

	var window_width = width
	var window_height = height

	var window = ApplicationWindow_Struct{
		WindowCanvas: nil,
		Window: nil,
		Width: &window_width,
		Height: &window_height, 
		Initial_Width: window_width,
		Initial_Height: window_height,
		Title: title,
		Resizable: 1,
		CloseMethod: void,
		OpenMethod: void,
		err: nil,
	}
	return &window
}

func (A *ApplicationWindow_Struct) SetTheme(theme *colours.Theme) {
	log.Println("new [Theme] struct applied to [Window].")
	A.Theme = theme
	A.SetBackgroundColour(A.Theme.WindowBackground)
}

func (A *ApplicationWindow_Struct) GetTheme() *colours.Theme {
	return A.Theme
}

func void() {}
func (A *ApplicationWindow_Struct) SetWindow(w *glfw.Window) { A.Window = w }
func (A *ApplicationWindow_Struct) GetWindow() *glfw.Window { return A.Window }
func (A *ApplicationWindow_Struct) SetBackgroundColour(colour [3]float32){ A.GetWindowCanvas().SetBackgroundColour(colour) }
func (A *ApplicationWindow_Struct) SetWindowHeight(x float32){ *A.Height = x }
func (A *ApplicationWindow_Struct) GetWindowHeight() *float32 { 
	if A.Window != nil{
		var width, height = A.Window.GetSize()
		*A.Width = float32(width)
		*A.Height = float32(height)
	}
	return A.Height 
}
func (A *ApplicationWindow_Struct) SetWindowWidth(x float32){ *A.Width = x }
func (A *ApplicationWindow_Struct) GetWindowWidth() *float32 { 
	if A.Window != nil{
		var width, height = A.Window.GetSize()
		*A.Width = float32(width)
		*A.Height = float32(height)
	}
	return A.Width 
}
func (A *ApplicationWindow_Struct) SetResizable(r int) { A.Resizable = r }
func (A *ApplicationWindow_Struct) GetResizable() int { return A.Resizable }
func (A *ApplicationWindow_Struct) SetWindowTitle(t string) { A.Title = t }
func (A *ApplicationWindow_Struct) GetWindowTitle() string { return A.Title }
func (A *ApplicationWindow_Struct) SetCloseMethod(f func()) { A.CloseMethod = f }
func (A *ApplicationWindow_Struct) GetCloseMethod() func() { return A.CloseMethod }
func (A *ApplicationWindow_Struct) SetOpenMethod(f func()) { A.OpenMethod = f }
func (A *ApplicationWindow_Struct) GetOpenMethod() func() { return A.OpenMethod }
func (A *ApplicationWindow_Struct) CloseWindow() {
	A.CloseMethod()
	glfw.Terminate()
}
func (A *ApplicationWindow_Struct) OpenWindow() {
	log.Println("[Window] is being opened.")
	runtime.LockOSThread()
	A.err = glfw.Init()
	if A.err != nil {
		panic(A.err)
	}
	glfw.WindowHint(glfw.Resizable, A.Resizable)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
	A.Window, A.err = glfw.CreateWindow(int(*A.Width), int(*A.Height), A.Title, nil, nil) 
	if A.err != nil {
		panic(A.err)
	}
	A.Window.MakeContextCurrent()
	A.err = gl.Init()
	if A.err != nil {
		panic(A.err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println(fmt.Sprintf("OpenGL version: \" %s \"\n", version))
	var canvas = comp.NewCanvas(A)
	log.Println("[Canvas] is assigned to [Window] struct.")
	A.WindowCanvas = canvas
	A.OpenMethod()
}
func (A *ApplicationWindow_Struct) GetWindowCanvas() *comp.Canvas_Struct {
	return A.WindowCanvas
}
func (A *ApplicationWindow_Struct) SetWindowCanvas(c *comp.Canvas_Struct) {
	A.WindowCanvas = c
}
func (A *ApplicationWindow_Struct) SetMouseHandler(ml intf.MouseHandler_Interface){
	A.MouseHandler = ml
}
func (A *ApplicationWindow_Struct) GetMouseHandler() intf.MouseHandler_Interface{
	return A.MouseHandler
}
func (A *ApplicationWindow_Struct) RunWindow() {
	log.Println("[Window] is running until termination.")
	defer glfw.Terminate()
	A.GetWindow().SetMouseButtonCallback(A.mouseButtonCallback)
	A.GetWindow().SetSizeCallback(A.resizeCallback)
	log.Println("[Window] loop refreshes canvas, polls events, and executes event queue.")
	for !A.Window.ShouldClose(){
		A.GetWindowCanvas().RefreshCanvas()
		glfw.PollEvents()
		A.GetWindowCanvas().GetEventHandler().ExecuteNextEvent()
	}
	A.CloseWindow()
	log.Println("Closing [Window].")
}
func (A* ApplicationWindow_Struct) mouseButtonCallback(
	window *glfw.Window, 
	button glfw.MouseButton, 
	action glfw.Action,
	mods glfw.ModifierKey) {
		log.Println("Mouse clicked in [Window].")
		var mouse_x, mouse_y, pressed = A.GetMouseHandler().GetClickData(A.Window)
		A.GetMouseHandler().CheckClick(mouse_x, mouse_y, pressed, mods)
		A.GetWindow().Show()
		log.Println("Refreshing [Canvas].")
		A.GetWindowCanvas().RefreshCanvas()
}

func (A* ApplicationWindow_Struct) resizeCallback(window *glfw.Window, width int, height int){
	if !resizing {
		resizing = true
		go A.resizeTimeout()
	}
}

func (A* ApplicationWindow_Struct) resizeTimeout(){
	time.Sleep(250 * time.Millisecond)
	resizing = false
	A.resize()
}

func (A* ApplicationWindow_Struct) resize() {
	var window_w, window_h = A.Window.GetSize()
	*A.Width = float32(window_w)
	*A.Height = float32(window_h)
	A.GetWindowCanvas().Redraw()
	A.GetWindowCanvas().RefreshCanvas()
}