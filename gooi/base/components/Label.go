package components 
/**
 * Updated 11/11/2023.
 * Label -> Implements Drawable
 **/
import (
	intf 		"gooi/interfaces"
	foundations "gooi/base/components/foundation"
)
type Label_Struct struct {
	canvas 						intf.Canvas_Interface
	masterStruct 				intf.Displayable
	name 						string
	posX, posY, posZ 			float32
	masterWidth, masterHeight 	float32 
	slaveWidth, slaveHeight 	float32 
	openGLWindowWidth 			float32 
	openGLWindowHeight 			float32
	writing 					*foundations.Writing
}
func NewLabel(
	canvas intf.Canvas_Interface, 
	masterStruct intf.Displayable,
	name string,
	pos_x, pos_y, pos_z float32, 
	font_name string,
	font_path string,
	font_size int,
	) *Label_Struct {
	var l = Label_Struct{
		canvas, 
		masterStruct, 
		name, 
		pos_x, pos_y, pos_z, 
		masterStruct.GetWidth(), 
		masterStruct.GetHeight(),
		0, 0,
		canvas.GetWidth(), 
		canvas.GetHeight(),
		nil,
	}

	l.writing = foundations.NewWriting(
		canvas, 
		masterStruct, 
		name, 
		pos_x, pos_y, pos_z, 
		canvas.GetWidth(), canvas.GetHeight(),
		font_path, font_name, font_size,
	)
		
	l.writing.SetPosition(pos_x + l.writing.GetWidth()/2, pos_y + l.writing.GetHeight()/2, pos_z)

	l.slaveWidth = l.writing.GetWidth()
	l.slaveHeight = l.writing.GetHeight()

	return &l
}

func (l *Label_Struct) Draw() {	
	l.writing.Draw()
}
func (l *Label_Struct) Redraw() { 
	l.writing.Draw() // no polygons so this is the same as Draw here... 
}
/**
 * Setters and Getters 
 **/
func (l *Label_Struct) SetPos(x, y, z float32) { 
	l.posX = x + l.writing.GetWidth()/2
	l.posY = y + l.writing.GetHeight()/2 
	l.posZ = z
	l.writing.SetPosition(l.posX, l.posY, l.posZ)
}
func (l *Label_Struct) GetPos() (float32, float32, float32) { return l.posX, l.posY, l.posZ }

func (l *Label_Struct) SetDisplayText(s string) { 
	l.name = s 
	l.writing.SetLabel(l.name)
	l.slaveWidth = l.writing.GetWidth()
	l.slaveHeight = l.writing.GetHeight()
}
func (l *Label_Struct) GetDisplayText() string { return l.name }


func (l *Label_Struct) GetHeight() float32 { return l.slaveHeight }
func (l *Label_Struct) GetWidth() float32 { return l.slaveWidth }

func (l *Label_Struct) GetMasterStruct() intf.Displayable { return l.masterStruct }
func (l *Label_Struct) SetMasterStruct(master intf.Displayable) { l.masterStruct = master }