package compositions
import (
	intf 	"gooi/interfaces"
	cons 	"gooi/base/constants"
	ly 		"gooi/base/compositions/layout"
)
type Row struct {
	name string
	layout *ly.Layout

	alignment int
	displayables []intf.Displayable

	slaveWidth, slaveHeight float32

	posX, posY, posZ float32
}

func NewRowComposition(
	name string, 
	canvas intf.Canvas_Interface,
	masterStruct intf.Displayable,
	x, y, z float32, 
	alignment int,
) (*Row) {
		var row = Row{
			name,
			ly.NewLayout(canvas, masterStruct),
			alignment, 
			make([]intf.Displayable, 0),
			0, 0,
			x, y, z,
		}
		return &row
}
func (row *Row) AddDisplayable(displayable intf.Displayable){
	row.displayables = append(row.displayables, displayable)
	row.ArrangeLayout()
}
func (row *Row) ArrangeLayout(){

	var padding float32 = 10

	//row.slaveWidth = row.slaveWidthRatio * row.layout.GetMasterStruct().GetWidth()
	//row.slaveHeight = row.slaveHeightRatio * row.layout.GetMasterStruct().GetHeight()

	var alignment float32 = 0
	if row.alignment == cons.ALIGN_TOP {
		alignment = 1
	} else if row.alignment == cons.ALIGN_BOTTOM {
		alignment = 0
	} else if row.alignment == cons.ALIGN_CENTRE_ROW {
		alignment = 0.5
	} else {
		alignment = 0.5
	}
	if len(row.displayables) != 0 {
		var current_x = row.posX
		var current_y = row.posY
		
		var progressWidth float32 = 0
		var greatestHeight float32 = 0
		
		var width = row.displayables[0].GetWidth()
		var height = row.displayables[0].GetHeight()
		if height > greatestHeight {
			greatestHeight = height
		}
		progressWidth += (width + padding * 2)
		row.displayables[0].SetPos(current_x, current_y - height*alignment, row.posZ)
		for i := 1 ; i < len(row.displayables); i++ {
			var prev_width = row.displayables[i-1].GetWidth()
			
			var current_width = row.displayables[i].GetWidth()
			var current_height = row.displayables[i].GetHeight()

			progressWidth += (current_width + padding)
			current_x = current_x + prev_width + padding
			row.displayables[i].SetPos(current_x, current_y - current_height*alignment, row.posZ)
			if current_height > greatestHeight {
				greatestHeight = current_height
			}
		}

		row.slaveHeight = greatestHeight
		row.slaveWidth = progressWidth
		
		
		for i := 0 ; i < len(row.displayables); i++ {
			var x, y, _ = row.displayables[i].GetPos()
			if row.alignment == cons.ALIGN_TOP {
				row.displayables[i].SetPos(x, y + row.slaveHeight, row.posZ)
			} else if row.alignment == cons.ALIGN_BOTTOM {
				row.displayables[i].SetPos(x, y + 0, row.posZ)
			} else if row.alignment == cons.ALIGN_CENTRE_ROW {
				row.displayables[i].SetPos(x, y + row.slaveHeight/2, row.posZ)
			} else {
				row.displayables[i].SetPos(x, y + row.slaveHeight/2, row.posZ)
			}
		}
		

	}
}
func (row *Row) Draw(){
	for i := 0 ; i < len(row.displayables); i++ {
		row.displayables[i].Draw()
	}
}
func (row *Row) Redraw(){
	row.ArrangeLayout()
	for i := 0 ; i < len(row.displayables); i++ {
		row.displayables[i].Redraw()
	}
}
func (row *Row) SetPos(x, y, z float32){
	row.posX = x
	row.posY = y
	row.posZ = z
	row.Redraw()
}
func (row *Row) GetPos() (float32, float32, float32){
	return row.posX, row.posY, row.posZ
}

func (row *Row) GetWidth() float32 { return row.slaveWidth }
func (row *Row) GetHeight() float32 { return row.slaveHeight }

func (row *Row) GetMasterStruct() intf.Displayable { return row.layout.GetMasterStruct() }
func (row *Row) SetMasterStruct(displayable intf.Displayable) { row.layout.SetMasterStruct(displayable) }