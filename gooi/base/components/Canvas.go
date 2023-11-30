package components
/**
 * Updated 11/11/2023.
 * Canvas Component.
 * Implements Canvas.
 * Contains drawables (Components, Compositions, and TextComponents).
 * Highest level component before being drawn to the ApplicationWindow.
 **/
import (
	gl 			"github.com/go-gl/gl/v4.1-core/gl"

	intf 		"gooi/interfaces"
	font 		"gooi/base/font"
	shaders 	"gooi/base/shaders"
	colours 	"gooi/base/colours"
	
	log 		"log"
	strings 	"strings"
)
type Canvas_Struct struct {
	err error
	EventHandler intf.EventHandler_Interface
	CanvasWindow intf.Window_Interface
	VertexShader uint32
	FragmentShader uint32
	Prog uint32
	Components []intf.Displayable
	Canvas_Font *font.Font_Struct
	BackgroundColour [3]float32

	masterStruct intf.Displayable
}
/**
 * Creats new Canvas_Struct 
 **/
func NewCanvas(window intf.Window_Interface) *Canvas_Struct {
	log.Println("Creating new [Canvas] struct.")
	var c = Canvas_Struct{}
	c.SetWindow(window)
	// Compile shaders for this canvas
	c.VertexShader = c.CompileCanvasShader(shaders.VertexShaderSource, gl.VERTEX_SHADER)
	c.FragmentShader = c.CompileCanvasShader(shaders.FragmentShaderSource, gl.FRAGMENT_SHADER)
	// Assign a new program to the canvas and link to gl
	c.Prog = gl.CreateProgram()
	gl.AttachShader(c.Prog, c.VertexShader)
	gl.AttachShader(c.Prog, c.FragmentShader)
	gl.LinkProgram(c.Prog)
	// Set background colour to default WHITE
	c.BackgroundColour = colours.WHITE
	gl.ClearColor(c.BackgroundColour[0], c.BackgroundColour[1], c.BackgroundColour[2], 1.0)
	// Create empty array of drawables
	c.Components = make([]intf.Displayable, 0)	
	c.masterStruct = &c
	return &c
}
// AddComponent(drawable)
// Adds a component to the components array.
func (c *Canvas_Struct) AddDisplayable(a intf.Displayable) { c.Components = append(c.Components, a) }
func (c *Canvas_Struct) GetDisplayables() []intf.Displayable { return c.Components }
func (c *Canvas_Struct) SetDisplayables(a []intf.Displayable) { c.Components = a }
// Count all drawables in canvas
func (c *Canvas_Struct) CountComponents() int { return len(c.Components) }
func (c *Canvas_Struct) RefreshCanvas() {
	gl.UseProgram(c.Prog)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	c.Draw()
	c.GetWindow().GetWindow().SwapBuffers()
}
func (c *Canvas_Struct) Redraw() {
	for _, component := range c.Components {
		gl.UseProgram(c.Prog)
		component.Redraw()
	}	
}
// Draw method for canvas
func (c *Canvas_Struct) Draw() {
	for _, component := range c.Components {
		gl.UseProgram(c.Prog)
		component.Draw()
	}	
}
// Compile the canvas shaders (from examples on go-gl)
func (c *Canvas_Struct) CompileCanvasShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
    csources, free := gl.Strs(source)
    gl.ShaderSource(shader, 1, csources, nil)
    free()
    gl.CompileShader(shader)
    var status int32
    gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
    if status == gl.FALSE {
        var logLength int32
        gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
        log := strings.Repeat("\x00", int(logLength+1))
        gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
        return 0
    }
    return shader
}
/** 
 * Setters and Getters 
 **/
func (c *Canvas_Struct) SetBackgroundColour(colour [3]float32){
	c.BackgroundColour = colour
	gl.ClearColor(c.BackgroundColour[0]/255, c.BackgroundColour[1]/255, c.BackgroundColour[2]/255, 1.0)
}
func (c *Canvas_Struct) GetPrograms() (uint32){ return c.Prog }
func (c *Canvas_Struct) GetEventHandler() intf.EventHandler_Interface{ return c.EventHandler }
func (c *Canvas_Struct) SetEventHandler(l intf.EventHandler_Interface){ c.EventHandler = l }
func (c *Canvas_Struct) SetWindow(w intf.Window_Interface){ c.CanvasWindow = w }
func (c *Canvas_Struct) GetWindow() intf.Window_Interface { return c.CanvasWindow }

func (c *Canvas_Struct) GetWidth() float32 { return *c.CanvasWindow.GetWindowWidth() }
func (c *Canvas_Struct) GetHeight() float32 { return *c.CanvasWindow.GetWindowHeight() }

func (c *Canvas_Struct) SetPos(x, y, z float32) { //nothing 
}
func (c *Canvas_Struct) GetPos() (float32, float32, float32) {
	return 0, 0, 0
}

func (c *Canvas_Struct) GetMasterStruct() intf.Displayable { return c }
func (c *Canvas_Struct) SetMasterStruct(master intf.Displayable) { c.masterStruct = master }