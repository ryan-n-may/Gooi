package compositions
import (
	intf "gooi/interfaces"
	cons "gooi/base/constants"	
	ly 	 "gooi/base/compositions/layout"
	"fmt"
)
type Column struct {
	name 			string
	layout  		*ly.Layout

	alignment 		int
	displayables 	[]intf.Displayable

	slaveWidth, slaveHeight float32

	posX, posY, posZ float32
}

func NewColumnComposition(
	name string, 
	canvas intf.Canvas_Interface,
	masterStruct intf.Displayable,
	x, y, z float32,
	alignment int,
) (*Column) {
		var col = Column{
			name, 
			ly.NewLayout(canvas, masterStruct),
			alignment,
			make([]intf.Displayable, 0),
			0, 0,
			x, y, z,
		}
		return &col
}
func (col *Column) AddDisplayable(displayable intf.Displayable){
	col.displayables = append(col.displayables, displayable)
	col.ArrangeLayout()
}
func (col *Column) GetDisplayables() []intf.Displayable {
	return col.displayables
}
func (col *Column) SetDisplayables(displayables []intf.Displayable) {
	col.displayables = displayables
}

func (col *Column) ArrangeLayout(){
	var padding float32 = 10

	//col.slaveWidth = col.slaveWidthRatio * col.layout.GetMasterStruct().GetWidth()
	//col.slaveHeight = col.slaveHeightRatio * col.layout.GetMasterStruct().GetHeight()
	
	var alignment float32 = 0
	if col.alignment == cons.ALIGN_CENTRE_COLUMN {
		alignment = 0.5
	} else if col.alignment == cons.ALIGN_LEFT {
		alignment = 0
	} else if col.alignment == cons.ALIGN_RIGHT {
		alignment = 1
	} else {
		alignment = 0.5
	}

	if len(col.displayables) != 0 {
		var current_x = col.posX
		var current_y = col.posY

		var progressHeight float32 = 0
		var greatestWidth float32 = 0
		
		for i := 0 ; i < len(col.displayables); i++ {
			var current_width = col.displayables[i].GetWidth()
			if current_width > greatestWidth {
				greatestWidth = current_width
			}
		}
		var width = col.displayables[0].GetWidth()
		var height = col.displayables[0].GetHeight()

		fmt.Printf("\t\tHeight is = %v\n", height)

		if width > greatestWidth {
			greatestWidth = width
		}
		progressHeight += (height + padding * 2)
		col.displayables[0].SetPos(current_x - width*alignment + col.slaveWidth*alignment, current_y, col.posZ)
		for i := 1 ; i < len(col.displayables); i++ {

			var prev_height = col.displayables[i-1].GetHeight()

			fmt.Printf("\t\tHeight is = %v\n", prev_height)
			
			var current_width = col.displayables[i].GetWidth()
			var current_height = col.displayables[i].GetHeight()
			
			progressHeight += (current_height + padding)
			current_y = current_y + prev_height + padding
			col.displayables[i].SetPos(current_x - current_width*alignment + col.slaveWidth*alignment, current_y, col.posZ)
			if current_width > greatestWidth {
				greatestWidth = current_width
			}
		}

		col.slaveHeight = progressHeight 
		col.slaveWidth = greatestWidth
	}
}
func (col *Column) Draw(){
	for i := 0 ; i < len(col.displayables); i++ {
		col.displayables[i].Draw()
	}
}
func (col *Column) Redraw(){
	col.ArrangeLayout()
	for i := 0 ; i < len(col.displayables); i++ {
		col.displayables[i].Redraw()
	}
}
func (col *Column) SetPos(x, y, z float32){
	col.posX = x
	col.posY = y
	col.posZ = z
	col.Redraw()
}
func (col *Column) GetPos() (float32, float32, float32){
	return col.posX, col.posY, col.posZ
}

func (col *Column) GetWidth() float32 { return col.slaveWidth }
func (col *Column) GetHeight() float32 { return col.slaveHeight }

func (col *Column) GetMasterStruct() intf.Displayable { return col.layout.GetMasterStruct() }
func (col *Column) SetMasterStruct(displayable intf.Displayable) { col.layout.SetMasterStruct(displayable) }