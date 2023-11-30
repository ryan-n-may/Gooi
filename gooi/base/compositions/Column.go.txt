package compositions
import (
	intf "gooi/interfaces"
	cons "gooi/base/constants"	
	log  "log"
	"fmt"
)
type Column_Struct struct {
	ColumnName string
	Master_Pos_x float32
	Master_Pos_y float32
	Master_Height float32
	Master_Width float32

	Sub_Width *float32
	Sub_Height *float32

	Padding float32
	Alignment int
	Drawables []intf.Drawable_Interface
}
func (col *Column_Struct) SetSubWidth(w *float32){
	col.Sub_Width = w
}
func (col *Column_Struct) SetSubHeight(h *float32){
	col.Sub_Height = h
}
func NewColumnComposition(name string, pos_x, pos_y, padding float32, alignment int) (*Column_Struct) {
		log.Println("new [Column].")
		var col = Column_Struct{}
		var zero float32 = 0 
		col.SetSubWidth(&zero)
		col.SetSubHeight(&zero)
		col.ColumnName = name
		col.Master_Pos_x = pos_x
		col.Master_Pos_y = pos_y
		col.Padding = padding
		col.Master_Width = 0
		col.Master_Height = 0
		col.Alignment = alignment
		col.Drawables = make([]intf.Drawable_Interface, 0)
		return &col
}
func (col *Column_Struct) AddDrawable(drawable intf.Drawable_Interface){
	drawable.SetSubWidth(col.Sub_Width)
	var _, height = drawable.GetBounds()
	drawable.SetSubHeight(&height)
	col.Drawables = append(col.Drawables, drawable)
	col.MoveComponents()
}
func (col *Column_Struct) MoveComponents(){
	//log.Println("moving [Column] components.")
	var alignment float32 = 0
	if col.Alignment == cons.ALIGN_CENTRE_COLUMN {
		alignment = 0.5
	} else if col.Alignment == cons.ALIGN_LEFT {
		alignment = 0
	} else if col.Alignment == cons.ALIGN_RIGHT {
		alignment = 1
	} else {
		alignment = 0.5
	}
	var current_x = col.Master_Pos_x
	var current_y = col.Master_Pos_y
	col.Master_Height = 0

	for i := 0 ; i < len(col.Drawables); i++ {
		var current_width, _ = col.Drawables[i].GetBounds()
		if current_width > col.Master_Width {
			col.Master_Width = current_width
		}
	}

	var width, height = col.Drawables[0].GetBounds()
	if width > col.Master_Width {
		col.Master_Width = width
	}
	col.Master_Height += height 
	col.Drawables[0].SetPos(current_x - width*alignment + col.Master_Width*alignment, current_y)
	for i := 1 ; i < len(col.Drawables); i++ {
		var _, prev_height = col.Drawables[i-1].GetBounds()
		var current_width, current_height = col.Drawables[i].GetBounds()
		col.Master_Height += (current_height + col.Padding)
		current_y = current_y + prev_height + col.Padding
		col.Drawables[i].SetPos(current_x - current_width*alignment + col.Master_Width*alignment, current_y)
		if current_width > col.Master_Width {
			col.Master_Width = current_width
		}
	}
	/*for i := 0 ; i < len(col.Drawables); i++ {
		var current_width, _ = col.Drawables[i].GetBounds()
		var current_x, current_y = col.Drawables[i].GetPos()
		col.Drawables[i].SetPos(current_x - current_width*alignment + col.Master_Width*alignment, current_y)
	}*/
}
func (col *Column_Struct) Draw(){
	for i := 0 ; i < len(col.Drawables); i++ {
		col.Drawables[i].Draw()
	}
}
func (col *Column_Struct) Redraw(){
	col.MoveComponents()
	for i := 0 ; i < len(col.Drawables); i++ {
		col.Drawables[i].Redraw()
	}
}
func (col *Column_Struct) SetPos(x, y float32){
	col.Master_Pos_x = x
	col.Master_Pos_y = y
	col.Redraw()
}
func (col *Column_Struct) GetPos() (float32, float32){
	fmt.Printf("Column master x y, %v, %v\n", col.Master_Pos_x, col.Master_Pos_y)
	return col.Master_Pos_x, col.Master_Pos_y
}
func (col *Column_Struct) GetBounds() (float32, float32){
	return col.Master_Width, col.Master_Height
}

func (col *Column_Struct) SetPosZ(z float32) {
	for i := 0 ; i < len(col.Drawables); i++ {
		col.Drawables[i].SetPosZ(z)
	}
}
func (col *Column_Struct) GetPosZ() float32 {
	return col.Drawables[0].GetPosZ()
}