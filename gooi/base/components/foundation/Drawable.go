package foundation

import (
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	
	base 	"gooi/base"
	intf 	"gooi/interfaces"
)

type Drawable struct {
	canvas 		 	intf.Canvas_Interface

	xyz 			[]float32
	rgb 			[]float32 
	dim 			[]float32
	vao 			[]intf.Drawing

	masterStruct intf.Displayable
	masterWidth float32
	masterHeight float32

	slaveWidth float32
	slaveHeight	float32

	openGLWindowHeight float32
	openGLWindowWidth float32
}

func NewDrawable(canvas intf.Canvas_Interface, masterStruct intf.Displayable, openGLWindowWidth, openGLWindowHeight float32) *Drawable{
	var drawable = Drawable{}
	drawable.vao = make([]intf.Drawing, 0, 0)
	drawable.xyz = make([]float32, 0, 0)
	drawable.rgb = make([]float32, 0, 0)
	drawable.dim = make([]float32, 0, 0)

	drawable.canvas = canvas
	drawable.masterStruct = masterStruct
	drawable.openGLWindowWidth = openGLWindowWidth
	drawable.openGLWindowHeight = openGLWindowHeight
	drawable.masterWidth = drawable.masterStruct.GetWidth()
	drawable.masterHeight = drawable.masterStruct.GetHeight()
	drawable.slaveWidth = 0.0
	drawable.slaveHeight = 0.0
	return &drawable
}

func (drawable *Drawable) SetXYZ(xyz []float32) { drawable.xyz = xyz }
func (drawable *Drawable) GetXYZ() []float32 { return drawable.xyz }

func (drawable *Drawable) SetRGB(rgb []float32) { drawable.rgb = rgb }
func (drawable *Drawable) GetRGB() []float32 { return drawable.xyz }

func (drawable *Drawable) SetDIM(dim []float32) { drawable.dim = dim } 
func (drawable *Drawable) GetDIM() []float32 { return drawable.dim }

func (drawable *Drawable) SetVAO(vao []intf.Drawing) { drawable.vao = vao }
func (drawable *Drawable) GetVAO() []intf.Drawing	{ return drawable.vao }

func (drawable *Drawable) GetWidth() float32 { return drawable.slaveWidth }
func (drawable *Drawable) GetHeight() float32 { return drawable.slaveHeight }

func (drawable *Drawable) ClearPolygons(){
	drawable.vao = make([]intf.Drawing, 0, 0)
	drawable.xyz = make([]float32, 0, 0)
	drawable.rgb = make([]float32, 0, 0)
	drawable.dim = make([]float32, 0, 0)
	return
}

func (drawable *Drawable) CreateRectangle(colour [3]float32, width, height, pos_x, pos_y, pos_z float32){
	var newDrawing = intf.Drawing{ 
		base.GenerateRectangle(
			drawable, 
			colour, 
			width, height, 
			pos_x, pos_y, pos_z,
			drawable.openGLWindowHeight, drawable.openGLWindowWidth), 
		gl.TRIANGLES,
	}
	if width > drawable.slaveWidth {
		drawable.slaveWidth = width
	} 
	if height > drawable.slaveHeight {
		drawable.slaveHeight = height 
	}
	drawable.vao = append(drawable.vao, newDrawing)
	return
}

func (drawable *Drawable) Draw(){
	gl.UseProgram(drawable.canvas.GetPrograms())
	for _, v := range drawable.vao{
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(drawable.xyz)/3))
	}
}

func (drawable *Drawable) CreateRoundedRectangle(colour [3]float32, width, height, pos_x, pos_y, pos_z, radius float32){
	// Border rectangles 
		drawable.CreateRectangle(colour, width-radius, height, pos_x+radius/2, pos_y, pos_z)
		drawable.CreateRectangle(colour, width, height-radius, pos_x, pos_y+radius/2, pos_z)
		// Border corner circles
		drawable.CreateCircle(colour, radius/2, pos_x+radius/2, pos_y+radius/2, pos_z)
		drawable.CreateCircle(colour, radius/2, pos_x+width-radius/2, pos_y+radius/2, pos_z)
		drawable.CreateCircle(colour, radius/2, pos_x+radius/2, pos_y+height-radius/2, pos_z)
		drawable.CreateCircle(colour, radius/2, pos_x+width-radius/2, pos_y+height-radius/2, pos_z)
		return
}

func (drawable *Drawable) CreateCircle(colour [3]float32, radius, pos_x, pos_y, pos_z float32) {
	var newDrawing = intf.Drawing{ 
		base.GenerateCircle(
			drawable, 
			colour, 
			pos_x, pos_y, pos_z, 
			radius,
			drawable.openGLWindowWidth, drawable.openGLWindowHeight, 
			64), 
		gl.TRIANGLE_FAN }
	if radius*2 > drawable.slaveWidth {
		drawable.slaveWidth = radius*2
	} 
	if radius*2 > drawable.slaveHeight {
		drawable.slaveHeight = radius*2 
	}
	drawable.vao = append(drawable.vao, newDrawing)
	return
}