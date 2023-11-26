package compositions
import (
	intf "gooi/interfaces"
	cons "gooi/base/constants"	
	log  "log"
	//fmt  "fmt"
)
type Box_Struct struct {
	ColumnName string
	Master_Pos_x float32
	Master_Pos_y float32
	Master_Pos_z float32
	Master_Height *float32
	Master_Width *float32

	Sub_Width *float32
	Sub_Height *float32

	Alignment int
	Drawables []intf.Drawable_Interface
}
func (box *Box_Struct) SetSubWidth(w *float32){
	box.Sub_Width = w
}
func (box *Box_Struct) SetSubHeight(h *float32){
	box.Sub_Height = h
}

func NewBoxComposition(name string, x, y float32, width, height *float32, alignment int) (*Box_Struct) {
		log.Println("new [Box].")
		var box = Box_Struct{}
		var zero float32 = 0
		box.SetSubWidth(&zero)
		box.SetSubHeight(&zero)
		box.ColumnName = name
		box.Master_Pos_x = x
		box.Master_Pos_y = y
		box.Master_Width = width
		box.Master_Height = height
		box.Alignment = alignment
		box.Drawables = make([]intf.Drawable_Interface, 1)
		return &box
}
func (box *Box_Struct) AddDrawable(drawable intf.Drawable_Interface){
	box.Drawables[0] = drawable
	box.Drawables[0].SetSubWidth(box.Master_Width)
	box.Drawables[0].SetSubHeight(box.Master_Height)
	box.MoveComponents()
}
func (box *Box_Struct) GetDrawables() intf.Drawable_Interface {
	return box.Drawables[0]
}
func (box *Box_Struct) MoveComponents(){
	
	//log.Println("moving [Box] components.")
	var component_width, component_height = box.Drawables[0].GetBounds()
	if box.Alignment == cons.ALIGN_BOTTOM_LEFT {
		box.Drawables[0].SetPos(box.Master_Pos_x + 0, 
			box.Master_Pos_y)
	} else if box.Alignment == cons.ALIGN_BOTTOM_RIGHT {
		box.Drawables[0].SetPos(box.Master_Pos_x + (*box.Master_Width) - component_width, 
			box.Master_Pos_y)
	} else if box.Alignment == cons.ALIGN_TOP_LEFT {
		box.Drawables[0].SetPos(box.Master_Pos_x + 0, 
			box.Master_Pos_y + (*box.Master_Height) - component_height)
	} else if box.Alignment == cons.ALIGN_TOP_RIGHT {
		box.Drawables[0].SetPos(box.Master_Pos_x + (*box.Master_Width) - component_width, 
			box.Master_Pos_y + (*box.Master_Height)-component_height)
	} else if box.Alignment == cons.ALIGN_CENTRE {
		box.Drawables[0].SetPos(box.Master_Pos_x + (*box.Master_Width/2) - component_width/2, 
			box.Master_Pos_y + (*box.Master_Height)/2 - component_height/2)
	} else if box.Alignment == cons.ALIGN_CENTRE_LEFT {
		box.Drawables[0].SetPos(box.Master_Pos_x + 0, 
			box.Master_Pos_y + (*box.Master_Height)/2-component_height/2)
	} else if box.Alignment == cons.ALIGN_CENTRE_RIGHT {
		box.Drawables[0].SetPos(box.Master_Pos_x + (*box.Master_Width) - component_width, 
			box.Master_Pos_y + (*box.Master_Height)/2-component_height/2)
	} else if box.Alignment == cons.ALIGN_TOP_CENTRE {
		box.Drawables[0].SetPos(box.Master_Pos_x + (*box.Master_Width/2) - component_width/2, 
			box.Master_Pos_y + (*box.Master_Height)-component_height)
	} else if box.Alignment == cons.ALIGN_BOTTOM_CENTRE {
		box.Drawables[0].SetPos(box.Master_Pos_x + (*box.Master_Width/2) - component_width/2, 
			box.Master_Pos_y)
	} else {
		box.Drawables[0].SetPos(box.Master_Pos_x + *box.Master_Width/2 - component_width/2, 
			box.Master_Pos_y + (*box.Master_Height)/2-component_height/2)
	}
}
func (box *Box_Struct) Draw(){
	if box.Drawables[0] != nil {
		box.Drawables[0].Draw()
	}
}
func (box *Box_Struct) Redraw(){
	box.MoveComponents()
	if box.Drawables[0] != nil {
		box.Drawables[0].Redraw()
	}
}
func (box *Box_Struct) SetPos(x, y float32){
	box.Master_Pos_x = x
	box.Master_Pos_y = y
	box.Redraw()
}

func (box *Box_Struct) SetWidth(w float32) {
	box.Master_Width = &w
}

func (box *Box_Struct) SetHeight(h float32) {
	box.Master_Height = &h
}

func (box *Box_Struct) GetWidth() *float32 {
	return box.Master_Width 
}

func (box *Box_Struct) GetHeight() *float32 {
	return box.Master_Height 
}

func (box *Box_Struct) GetPos() (float32, float32){
	return box.Master_Pos_x, box.Master_Pos_y
}
func (box *Box_Struct) GetBounds() (float32, float32){
	return *box.Master_Width, *box.Master_Height
}

func (box *Box_Struct) SetPosZ(z float32) {
	box.Drawables[0].SetPosZ(z)
}
func (box *Box_Struct) GetPosZ() float32 {
	return box.Drawables[0].GetPosZ()
}