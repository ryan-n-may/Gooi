package compositions
import (
	intf 	"gooi/interfaces"
	cons 	"gooi/base/constants"
	log 	"log"
)
type Row_Struct struct {
	RowName string
	Master_Pos_x float32
	Master_Pos_y float32
	Master_Height float32
	Master_Width float32
	Padding float32
	Alignment int
	Drawables []intf.Drawable_Interface
}
func NewRowComposition(name string, pos_x, pos_y, padding float32, alignment int) (*Row_Struct) {
		log.Println("new [Row].")
		var row = Row_Struct{}
		row.RowName = name
		row.Master_Pos_x = pos_x
		row.Master_Pos_y = pos_y
		row.Padding = padding
		row.Master_Width = 0
		row.Master_Height = 0
		row.Alignment = alignment
		row.Drawables = make([]intf.Drawable_Interface, 0)
		return &row
}
func (row *Row_Struct) AddDrawable(drawable intf.Drawable_Interface){
	row.Drawables = append(row.Drawables, drawable)
	row.MoveComponents()
}
func (row *Row_Struct) MoveComponents(){
	//log.Println("moving [Row] components.")
	var alignment float32 = 0
	if row.Alignment == cons.ALIGN_TOP {
		alignment = 1
	} else if row.Alignment == cons.ALIGN_BOTTOM {
		alignment = 0
	} else if row.Alignment == cons.ALIGN_CENTRE_ROW {
		alignment = 0.5
	} else {
		alignment = 0.5
	}
	var current_x = row.Master_Pos_x
	var current_y = row.Master_Pos_y
	row.Master_Width = 0
	var width, height = row.Drawables[0].GetBounds()
	if height > row.Master_Height {
		row.Master_Height = height
	}
	row.Master_Width += width 
	row.Drawables[0].SetPos(current_x, current_y - height*alignment)
	for i := 1 ; i < len(row.Drawables); i++ {
		var prev_width, _ = row.Drawables[i-1].GetBounds()
		var current_width, current_height = row.Drawables[i].GetBounds()
		row.Master_Width += (current_width + row.Padding)
		current_x = current_x + prev_width + row.Padding
		row.Drawables[i].SetPos(current_x, current_y - current_height*alignment)
		if current_height > row.Master_Height {
			row.Master_Height = current_height
		}
	}
	for i := 0 ; i < len(row.Drawables); i++ {
		var x, y = row.Drawables[i].GetPos()
		if row.Alignment == cons.ALIGN_TOP {
			row.Drawables[i].SetPos(x, y + row.Master_Height)
		} else if row.Alignment == cons.ALIGN_BOTTOM {
			row.Drawables[i].SetPos(x, y + 0)
		} else if row.Alignment == cons.ALIGN_CENTRE_ROW {
			row.Drawables[i].SetPos(x, y + row.Master_Height/2)
		} else {
			row.Drawables[i].SetPos(x, y + row.Master_Height/2)
		}
	}
}
func (row *Row_Struct) Draw(){
	for i := 0 ; i < len(row.Drawables); i++ {
		row.Drawables[i].Draw()
	}
}
func (row *Row_Struct) Redraw(){
	row.MoveComponents()
	for i := 0 ; i < len(row.Drawables); i++ {
		row.Drawables[i].Redraw()
	}
}
func (row *Row_Struct) SetPos(x, y float32){
	row.Master_Pos_x = x
	row.Master_Pos_y = y
	row.MoveComponents()
}
func (row *Row_Struct) GetPos() (float32, float32){
	return row.Master_Pos_x, row.Master_Pos_y
}
func (row *Row_Struct) GetBounds() (float32, float32){
	return row.Master_Width, row.Master_Height
}

func (row *Row_Struct) SetPosZ(z float32) {
	for i := 0 ; i < len(row.Drawables); i++ {
		row.Drawables[i].SetPosZ(z)
	}
}
func (row *Row_Struct) GetPosZ() float32 {
	return row.Drawables[0].GetPosZ()
}