package compositions
import (
	intf "gooi/interfaces"
	cons "gooi/base/constants"	
	log  "log"
	//fmt  "fmt"
)
type Stack_Struct struct {
	StackName string
	Master_Pos_x float32
	Master_Pos_y float32
	Master_Pos_z float32
	Master_Height *float32
	Master_Width *float32
	Alignment int
	Drawables []intf.Drawable_Interface
}
func NewStackComposition(name string, x, y float32, width, height *float32, alignment int) (*Stack_Struct) {
		log.Println("new [Stack].")
		var box = Stack_Struct{}
		box.StackName = name
		box.Master_Pos_x = x
		box.Master_Pos_y = y
		box.Master_Width = width
		box.Master_Height = height
		box.Alignment = alignment
		box.Drawables = make([]intf.Drawable_Interface, 0)
		return &box
}
func (box *Stack_Struct) AddDrawable(drawable intf.Drawable_Interface){
	box.Drawables = append(box.Drawables, drawable)
	box.MoveComponents()
}
func (box *Stack_Struct) MoveComponents(){
	for i, _ := range box.Drawables {
		var component_width, component_height = box.Drawables[i].GetBounds()
		if box.Alignment == cons.ALIGN_BOTTOM_LEFT {
			box.Drawables[i].SetPos(box.Master_Pos_x + 0, 
				box.Master_Pos_y)
		} else if box.Alignment == cons.ALIGN_BOTTOM_RIGHT {
			box.Drawables[i].SetPos(box.Master_Pos_x + (*box.Master_Width) - component_width, 
				box.Master_Pos_y)
		} else if box.Alignment == cons.ALIGN_TOP_LEFT {
			box.Drawables[i].SetPos(box.Master_Pos_x + 0, 
				box.Master_Pos_y + (*box.Master_Height) - component_height)
		} else if box.Alignment == cons.ALIGN_TOP_RIGHT {
			box.Drawables[i].SetPos(box.Master_Pos_x + (*box.Master_Width) - component_width, 
				box.Master_Pos_y + (*box.Master_Height)-component_height)
		} else if box.Alignment == cons.ALIGN_CENTRE {
			box.Drawables[i].SetPos(box.Master_Pos_x + (*box.Master_Width/2) - component_width/2, 
				box.Master_Pos_y + (*box.Master_Height)/2 - component_height/2)
		} else if box.Alignment == cons.ALIGN_CENTRE_LEFT {
			box.Drawables[i].SetPos(box.Master_Pos_x + 0, 
				box.Master_Pos_y + (*box.Master_Height)/2-component_height/2)
		} else if box.Alignment == cons.ALIGN_CENTRE_RIGHT {
			box.Drawables[i].SetPos(box.Master_Pos_x + (*box.Master_Width) - component_width, 
				box.Master_Pos_y + (*box.Master_Height)/2-component_height/2)
		} else if box.Alignment == cons.ALIGN_TOP_CENTRE {
			box.Drawables[i].SetPos(box.Master_Pos_x + (*box.Master_Width/2) - component_width/2, 
				box.Master_Pos_y + (*box.Master_Height)-component_height)
		} else if box.Alignment == cons.ALIGN_BOTTOM_CENTRE {
			box.Drawables[i].SetPos(box.Master_Pos_x + (*box.Master_Width/2) - component_width/2, 
				box.Master_Pos_y)
		} else {
			box.Drawables[i].SetPos(box.Master_Pos_x + *box.Master_Width/2 - component_width/2, 
				box.Master_Pos_y + (*box.Master_Height)/2-component_height/2)
		}
	}
}
func (box *Stack_Struct) Draw(){
	for i, _ := range box.Drawables {
		if box.Drawables[i] != nil {
			box.Drawables[i].Draw()
		}
	}
}
func (box *Stack_Struct) Redraw(){
	box.MoveComponents()
	for i, _ := range box.Drawables {
		if box.Drawables[i] != nil {
			box.Drawables[i].Draw()
		}
	}
}
func (box *Stack_Struct) SetPos(x, y float32){
	box.Master_Pos_x = x
	box.Master_Pos_y = y
	box.Redraw()
}

func (box *Stack_Struct) SetWidth(w float32) {
	box.Master_Width = &w
}

func (box *Stack_Struct) SetHeight(h float32) {
	box.Master_Height = &h
}

func (box *Stack_Struct) GetPos() (float32, float32){
	return box.Master_Pos_x, box.Master_Pos_y
}
func (box *Stack_Struct) GetBounds() (float32, float32){
	return *box.Master_Width, *box.Master_Height
}

func (box *Stack_Struct) SetPosZ(z float32) {
	for i, _ := range box.Drawables {
		if box.Drawables[i] != nil {
			box.Drawables[i].SetPosZ(z)
		}
	}
}
func (box *Stack_Struct) GetPosZ() float32 {
	return box.Drawables[0].GetPosZ()
}