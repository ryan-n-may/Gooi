package compositions
import (
	intf 	"gooi/interfaces"
	cons 	"gooi/base/constants"
	fmt 	"fmt"
	log 	"log"
)
type Vertical_Div_Struct struct {
	VerticalName 		string

	Master_Pos_x 		float32
	Master_Pos_y 		float32
	Master_Height 		*float32
	Master_Width 		*float32
	
	ColumnWidths 		[]float32
	Columns   			[]*Box_Struct

	Offset_Pos_x 		[]float32

	Sub_Width, Sub_Height *float32
}
func (b *Vertical_Div_Struct) SetSubWidth(w *float32){
	b.Sub_Width = w
}
func (b *Vertical_Div_Struct) SetSubHeight(h *float32){
	b.Sub_Height = h
}


func NewVerticalDivision(name string, widths []float32, posx, posy float32, width, height *float32) *Vertical_Div_Struct {
	log.Println("new [Vertical_Div].")
	var vert = Vertical_Div_Struct{}
	var zero float32 = 0 
	vert.SetSubWidth(&zero)
	vert.SetSubHeight(&zero)
	vert.VerticalName = name
	vert.Master_Pos_x = posx
	vert.Master_Pos_y = posy
	vert.Master_Height = height
	vert.Master_Width = width
	vert.ColumnWidths = widths
	vert.Columns = make([]*Box_Struct, len(widths))
	vert.Offset_Pos_x = make([]float32, len(widths))

	var offset = vert.Master_Pos_x
	for i, w := range widths {
		vert.Offset_Pos_x[i] = offset
		var adjusted_w = (w * (*vert.Master_Width))
		vert.Columns[i] = NewBoxComposition(
			fmt.Sprintf("DivStruct_%s_%v", vert.VerticalName, i), 
			offset, vert.Master_Pos_y, 
			&adjusted_w, vert.Master_Height, 
			cons.ALIGN_TOP_LEFT,
		) 
		offset = offset + (w * (*vert.Master_Width))
	}
	return &vert
}

func (v *Vertical_Div_Struct) MoveComponents() {
	//log.Println("moving [VerticalDiv] components.")
	var offset = v.Master_Pos_x
	for i, w := range v.ColumnWidths {
		v.Offset_Pos_x[i] = offset
		v.Columns[i].SetPos(offset, v.Master_Pos_y) 
		offset = offset + (w * (*v.Master_Width))
	}
}

func (v *Vertical_Div_Struct) AddDrawable(drawable intf.Drawable_Interface, index int) {
	v.Columns[index].AddDrawable(drawable)
}

func (v *Vertical_Div_Struct) Draw(){
	for i := 0; i < len(v.Columns); i++ {
		v.Columns[i].Draw()
	}
}
func (v *Vertical_Div_Struct) Redraw(){
	v.MoveComponents()
	for i := 0; i < len(v.Columns); i++ {
		v.Columns[i].Redraw()
	}
}
func (v *Vertical_Div_Struct) SetPos(x, y float32){
	v.Master_Pos_x = x
	v.Master_Pos_y = y
	v.Redraw()
}
func (v *Vertical_Div_Struct) GetPos() (float32, float32){
	return v.Master_Pos_x, v.Master_Pos_y
}
func (v *Vertical_Div_Struct) GetBounds() (float32, float32){
	return *(v.Master_Width), *(v.Master_Height)
}

func (v *Vertical_Div_Struct) SetPosZ(z float32) {
	for i := 0 ; i < len(v.Columns); i++ {
		v.Columns[i].SetPosZ(z)
	}
}
func (v *Vertical_Div_Struct) GetPosZ() float32 {
	return v.Columns[0].GetPosZ()
}