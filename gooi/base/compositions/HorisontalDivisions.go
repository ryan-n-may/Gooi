package compositions
import (
	intf 	"gooi/interfaces"
	cons 	"gooi/base/constants"
	fmt 	"fmt"
	log 	"log"
)
type Horisontal_Div_Struct struct {
	HorisontalName 		string

	Master_Pos_x 		float32
	Master_Pos_y 		float32
	Master_Height 		*float32
	Master_Width 		*float32
	
	RowHeights	 		[]float32
	Rows    			[]*Box_Struct

	Offset_Pos_y 		[]float32
}

func NewHorisontalDivision(name string, heights []float32, posx, posy float32, width, height *float32) *Horisontal_Div_Struct {
	log.Println("new [Horisontal_Div].")
	var hori = Horisontal_Div_Struct{}
	hori.HorisontalName = name
	hori.Master_Pos_x = posx
	hori.Master_Pos_y = posy
	hori.Master_Height = height
	hori.Master_Width = width
	hori.RowHeights = heights
	hori.Rows = make([]*Box_Struct, len(heights))
	hori.Offset_Pos_y = make([]float32, len(heights))

	var offset = hori.Master_Pos_y
	for i, height := range heights {
		hori.Offset_Pos_y[i] = offset
		var adjusted_h = height * (*hori.Master_Height)
		hori.Rows[i] = NewBoxComposition(
			fmt.Sprintf("DivStruct_%s_%v", hori.HorisontalName, i), 
			hori.Master_Pos_x, offset, 
			hori.Master_Width, &adjusted_h,
			cons.ALIGN_TOP_LEFT,
		) 
		offset = offset + (height * (*hori.Master_Height))
	}
	return &hori
}

func (h *Horisontal_Div_Struct) MoveComponents() {
	var offset = h.Master_Pos_y
	for i, height := range h.RowHeights {
		h.Offset_Pos_y[i] = offset
		h.Rows[i].SetPos(h.Master_Pos_x, offset) 
		h.Rows[i].SetHeight(height * (*h.Master_Height))
		offset = offset + (height * (*h.Master_Height))
	}
}


func (h *Horisontal_Div_Struct) AddDrawable(drawable intf.Drawable_Interface, index int) {
	h.Rows[index].AddDrawable(drawable)
}
func (h *Horisontal_Div_Struct) Draw(){
	for i := 0; i < len(h.Rows); i++ {
		h.Rows[i].Draw()
	}
}
func (h *Horisontal_Div_Struct) Redraw(){
	h.MoveComponents()
	for i := 0; i < len(h.Rows); i++ {
		h.Rows[i].Redraw()
	}
}

func (h *Horisontal_Div_Struct) SetPos(x, y float32){
	h.Master_Pos_x = x
	h.Master_Pos_y = y
	h.MoveComponents()
}
func (h *Horisontal_Div_Struct) GetPos() (float32, float32){
	return h.Master_Pos_x, h.Master_Pos_y
}
func (h *Horisontal_Div_Struct) GetBounds() (float32, float32){
	return (*h.Master_Width), (*h.Master_Height)
}

func (h *Horisontal_Div_Struct) SetPosZ(z float32) {
	for i := 0 ; i < len(h.Rows); i++ {
		h.Rows[i].SetPosZ(z)
	}
}
func (h *Horisontal_Div_Struct) GetPosZ() float32 {
	return h.Rows[0].GetPosZ()
}