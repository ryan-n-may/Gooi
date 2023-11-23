package components 
/**
 * Updated 11/11/2023.
 * Button Component.
 * Implements Drawable -> Component -> Clickable.
 **/
import (
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	
	intf 	"gooi/interfaces"
	
	log 	"log"
	fmt 	"fmt"
)
type Rectangle_Struct struct {
	Name 			string

	XYZ 			[]float32
	RGB 			[]float32 
	DIM 			[]float32
	VAO 			[]intf.Drawing_Struct
	Canvas 		 	*Canvas_Struct
	
	Width, Height 	float32
	Pos_x, Pos_y 	float32
	Pos_z 			float32
	
	WindowHeight    		*float32
	WindowHeight_Initial 	float32
	WindowWidth   			*float32
	WindowWidth_Initial 	float32
	
	BackgroundColour [3]float32
}
/**
 * CreateButton
 * 	Creates a new button composable.
 **/
func CreateRectangle(
	canvas 			*Canvas_Struct, 
	width, height 	float32,
	pos_x, pos_y 	float32, 
	colour  		[3]float32,
) *Rectangle_Struct {
	log.Println("creating new thembed [Button] struct.")
	var r = Rectangle_Struct{}
	// Specified paramaters
	r.Canvas = canvas
	r.Width = width
	r.Height = height
	r.Pos_x = pos_x
	r.Pos_y = pos_y
	r.Pos_z = 0.0
	

	r.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
	r.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
	r.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
	r.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()
	
	r.BackgroundColour = colour
	
	r.GeneratePolygons()
	
	return &r
}
// GeneratePolygons()
// Generates the VAO array of the polygons used to draw the button. 
// Stores the VAO in intf.Drawing_Struct alongisde the drawing mode (gl.TRIANGLE or gl.TRIANGLE_FAN)
func (r *Rectangle_Struct) GeneratePolygons(){
	r.VAO = make([]intf.Drawing_Struct, 2)
	// Border rectangles 
	r.VAO[0] = intf.Drawing_Struct{ intf.GenerateRectangle(r, r.BackgroundColour, r.Width, r.Height, r.Pos_x, r.Pos_y, r.Pos_z, r.WindowHeight_Initial, r.WindowWidth_Initial), gl.TRIANGLES }
	r.VAO[1] = intf.Drawing_Struct{ intf.GenerateRectangle(r, r.BackgroundColour, r.Width, r.Height, r.Pos_x, r.Pos_y, r.Pos_z, r.WindowHeight_Initial, r.WindowWidth_Initial), gl.TRIANGLES }
}
// Draw()
// This method draws the VAO array to gl using the canvas program.
func (r *Rectangle_Struct) Draw() {
	// Obtain program that isnt FontProg
	gl.UseProgram(r.GetCanvas().GetPrograms())
	for _, v := range r.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(r.GetXYZ())/3))
	}
	// Draw button text (Uses FontProg program)
	// Modified glText implementation:
}
func (r *Rectangle_Struct) Redraw() {
	r.GeneratePolygons()
	gl.UseProgram(r.GetCanvas().GetPrograms())
	for _, v := range r.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(r.GetXYZ())/3))
	}
	// Draw button text (Uses FontProg program)
	// Modified glText implementation:
	//b.Button_Text.Text.Font.ResizeWindow(*b.WindowWidth, *b.WindowHeight)
}
// Move()
// This method moves the button and the clickable area of the button.
func (r *Rectangle_Struct) Move(delta_x, delta_y float32) {
	log.Println(fmt.Sprintf("moving [Button] by %v, %v.", delta_x, delta_y))
	r.Pos_y = r.Pos_y + delta_y
	r.Pos_x = r.Pos_x + delta_x
	r.GeneratePolygons()
} 
// SetPos(float32, float32) & GetPos() float32, float32
// Sets the position on the window (absolute)
// Origin starting in the bottom left corner of the window.
// Updates clickable bounds and re-draws.
func (r *Rectangle_Struct) SetPos(x, y float32) { 
	r.Pos_x = x
	r.Pos_y = y
	r.GeneratePolygons()
	r.GetCanvas().RefreshCanvas()	
}
func (r *Rectangle_Struct) GetPos() (float32, float32) { return r.Pos_x, r.Pos_y }
/**
 * Other Setter and Getter Methods
 **/
func (r *Rectangle_Struct) SetName(name string){ r.Name = name }
func (r *Rectangle_Struct) GetName() string { return r.Name }
func (r *Rectangle_Struct) SetDIM(dim []float32){ r.DIM = dim }
func (r *Rectangle_Struct) GetDIM() []float32 { return r.DIM }
func (r *Rectangle_Struct) SetXYZ(xyz []float32){ r.XYZ = xyz }
func (r *Rectangle_Struct) GetXYZ() []float32 { return r.XYZ }
func (r *Rectangle_Struct) SetRGB(rgb []float32){ r.RGB = rgb }
func (r *Rectangle_Struct) GetRGB() []float32 { return r.RGB }
func (r *Rectangle_Struct) SetVAO(vao []intf.Drawing_Struct){ r.VAO = vao }
func (r *Rectangle_Struct) GetVAO() []intf.Drawing_Struct { return r.VAO }
func (r *Rectangle_Struct) GetCanvas() intf.Canvas_Interface { return r.Canvas }
func (r *Rectangle_Struct) SetPosZ(z float32){
	r.Pos_z = z
}
func (r *Rectangle_Struct) GetPosZ() float32 {
	return r.Pos_z
}
