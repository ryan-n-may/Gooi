package components 
/**
 * Updated 11/11/2023.
 * Button Component.
 * Implements Drawable -> Component -> Clickable.
 **/
import (
	intf 		"gooi/interfaces"
	cons 		"gooi/base/constants"
	foundations "gooi/base/components/foundation"
	colours 	"gooi/base/colours"
	"fmt"
)
type Rectangle_Struct struct {
	canvas 			intf.Canvas_Interface
	masterStruct 	intf.Displayable

	name 			string
	
	colour 			[3]float32

	posX, posY, posZ float32
	radius float32

	openGLWindowWidth, openGLWindowHeight float32
	masterWidth, masterHeight float32
	slaveWidth, slaveHeight float32

	drawable *foundations.Drawable

	fillStyle int
	positionStyle int
}
func NewRectangle(
	canvas 				intf.Canvas_Interface, 
	masterStruct		intf.Displayable,
	name 				string,
	width, height 		float32,
	pos_x, pos_y, pos_z	float32,
	radius				float32, 
	colour  			[3]float32,
	fill_style   		int,
	position_style 		int,
) *Rectangle_Struct {
	var r = Rectangle_Struct{
		canvas, 
		masterStruct,
		name,
		colour, 
		pos_x, pos_y, pos_z,
		radius,
		canvas.GetWidth(), canvas.GetHeight(),
		masterStruct.GetWidth(), masterStruct.GetHeight(),
		0,0,
		foundations.NewDrawable(
			canvas, 
			masterStruct,
			canvas.GetWidth(), canvas.GetHeight(),
		),
		fill_style,
		position_style,
	}

	if fill_style == cons.FILL_MASTER_DIMENSIONS {
		r.slaveWidth = masterStruct.GetWidth()
		r.slaveHeight = masterStruct.GetHeight()
	} else {
		r.slaveHeight = height
		r.slaveWidth = width
		r.posX = pos_x
		r.posY = pos_y
		r.posZ = pos_z
	}

	if position_style == cons.MATCH_MASTER_POSITION {
		r.posX, r.posY, r.posZ = masterStruct.GetPos()
	} else {
		r.posX = pos_x
		r.posY = pos_y
		r.posZ = pos_z
	}
	
	r.GeneratePolygons()
	
	return &r
}
// GeneratePolygons()
// Generates the VAO array of the polygons used to draw the button. 
// Stores the VAO in intf.Drawing_Struct alongisde the drawing mode (gl.TRIANGLE or gl.TRIANGLE_FAN)
func (r *Rectangle_Struct) GeneratePolygons(){
	if r.fillStyle == cons.FILL_MASTER_DIMENSIONS {
		r.slaveWidth = r.masterStruct.GetWidth()
		fmt.Printf("Slave width is = %v\n", r.slaveWidth)
		r.slaveHeight = r.masterStruct.GetHeight()
		r.posX, r.posY, r.posZ = r.masterStruct.GetPos()
		fmt.Println("\t\tREDRAWING USING NEW SIZE!!!!")
	}
	if r.colour != colours.NONE {
		r.drawable.ClearPolygons()
		r.drawable.CreateRoundedRectangle(r.colour, r.slaveWidth, r.slaveHeight, r.posX, r.posY, r.posZ, r.radius)
	}
}
// Draw()
// This method draws the VAO array to gl using the canvas program.
func (r *Rectangle_Struct) Draw() {
	if r.colour != colours.NONE {
		r.drawable.Draw()
	}
}
func (r *Rectangle_Struct) Redraw() {
	if r.colour != colours.NONE {
		r.GeneratePolygons()
		r.drawable.Draw()
	}
}

func (r *Rectangle_Struct) GetWidth() float32 {
	return r.slaveWidth
}
func (r *Rectangle_Struct) GetHeight() float32 {
	return r.slaveHeight
}

// SetPos(float32, float32) & GetPos() float32, float32
// Sets the position on the window (absolute)
// Origin starting in the bottom left corner of the window.
// Updates clickable bounds and re-draws.
func (r *Rectangle_Struct) SetPos(x, y, z float32) { 
	r.posX = x
	r.posY = y
	r.posZ = z
	r.GeneratePolygons()
	r.canvas.RefreshCanvas()	
}
func (r *Rectangle_Struct) GetPos() (float32, float32, float32) { return r.posX, r.posY, r.posZ }
/**
 * Other Setter and Getter Methods
 **/
func (r *Rectangle_Struct) SetName(name string){ r.name = name }
func (r *Rectangle_Struct) GetName() string { return r.name }

func (r *Rectangle_Struct) GetMasterStruct() intf.Displayable { return r.masterStruct }
func (r *Rectangle_Struct) SetMasterStruct(master intf.Displayable) { r.masterStruct = master }