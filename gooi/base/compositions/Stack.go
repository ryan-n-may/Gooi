package compositions
import (
	intf "gooi/interfaces"
	cons "gooi/base/constants"	
	ly 	 "gooi/base/compositions/layout"
	fmt  "fmt"
)
type Stack struct {
	name string
	layout *ly.Layout

	alignment []int
	displayables []intf.Displayable

	slaveWidth, slaveHeight float32
	slaveWidthRatio, slaveHeightRatio float32

	posX, posY, posZ float32
}

func NewStackComposition(
	name string, 
	canvas intf.Canvas_Interface,
	masterStruct intf.Displayable,
	x, y, z, slaveWidthRatio, slaveHeightRatio float32, 
	alignment []int,
) (*Stack) {
		var box = Stack{
			name,
			ly.NewLayout(canvas, masterStruct),
			alignment,

			make([]intf.Displayable, 0),

			slaveWidthRatio * masterStruct.GetWidth(),
			slaveHeightRatio * masterStruct.GetHeight(),
			slaveWidthRatio,
			slaveHeightRatio,

			x, y, z,
		}
		(&box).ArrangeLayout()
		return &box
}

func (box *Stack) AddDisplayable(displayable intf.Displayable){
	fmt.Printf("\t\tAdding displayable\n")
	box.displayables = append(box.displayables, displayable)
	fmt.Printf("\t\tLength of displayables %v\n", len(box.displayables))
	box.ArrangeLayout()
}
func (box *Stack) GetDisplayables() []intf.Displayable {
	return box.displayables
}

func (box *Stack) ArrangeLayout(){

	box.slaveWidth = box.slaveWidthRatio * box.layout.GetMasterStruct().GetWidth()
	box.slaveHeight = box.slaveHeightRatio * box.layout.GetMasterStruct().GetHeight()

	for i, _ := range box.displayables {		
		var componentWidth = box.displayables[i].GetWidth()
		var componentHeight = box.displayables[i].GetHeight()

		if box.alignment[i] == cons.ALIGN_BOTTOM_LEFT {
			box.displayables[i].SetPos(
				box.posX + 0, 
				box.posY,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_BOTTOM_RIGHT {
			box.displayables[i].SetPos(
				box.posX + (box.slaveWidth) - componentWidth, 
				box.posY,
				box.posZ,
		)
		} else if box.alignment[i] == cons.ALIGN_TOP_LEFT {
			box.displayables[i].SetPos(
				box.posX + 0, 
				box.posY + (box.slaveHeight) - componentHeight,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_TOP_RIGHT {
			box.displayables[i].SetPos(
				box.posX + (box.slaveWidth) - componentWidth, 
				box.posY + (box.slaveHeight)-componentHeight,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_CENTRE {
			box.displayables[i].SetPos(
				box.posX + (box.slaveWidth/2) - componentWidth/2, 
				box.posY + (box.slaveHeight)/2 - componentHeight/2,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_CENTRE_LEFT {
			box.displayables[i].SetPos(
				box.posX + 0, 
				box.posY + (box.slaveHeight)/2-componentHeight/2,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_CENTRE_RIGHT {
			box.displayables[i].SetPos(
				box.posX + (box.slaveWidth) - componentWidth, 
				box.posY + (box.slaveHeight)/2-componentHeight/2,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_TOP_CENTRE {
			box.displayables[i].SetPos(
				box.posX + (box.slaveWidth/2) - componentWidth/2, 
				box.posY + (box.slaveHeight)-componentHeight,
				box.posZ,
			)
		} else if box.alignment[i] == cons.ALIGN_BOTTOM_CENTRE {
			box.displayables[i].SetPos(
				box.posX + (box.slaveWidth/2) - componentWidth/2, 
				box.posY,
				box.posZ,
			)
		} else {
			box.displayables[i].SetPos(
				box.posX + box.slaveWidth/2 - componentWidth/2, 
				box.posY + box.slaveHeight/2-componentHeight/2,
				box.posZ,
			)
		}
	}
}
func (box *Stack) Draw(){
	for i, _ := range box.displayables {
		if box.displayables[i] != nil {
			box.displayables[i].Draw()
		}
	}
}
func (box *Stack) Redraw(){
	box.ArrangeLayout()
	for i, _ := range box.displayables {
		if box.displayables[i] != nil {
			box.displayables[i].Draw()
		}
	}
}
func (box *Stack) SetPos(x, y, z float32){
	box.posX = x
	box.posY = y
	box.posZ = z
	box.Redraw()
}

func (box *Stack) GetPos() (float32, float32, float32){
	return box.posX, box.posY, box.posZ
}


func (box *Stack) SetWidth(w float32) { box.slaveWidth = w }
func (box *Stack) SetHeight(h float32) { box.slaveHeight = h }

func (box *Stack) GetWidth() float32 { return box.slaveWidth }
func (box *Stack) GetHeight() float32 { return box.slaveHeight }

func (box *Stack) GetMasterStruct() intf.Displayable { return box.layout.GetMasterStruct() }
func (box *Stack) SetMasterStruct(master intf.Displayable) { box.layout.SetMasterStruct(master) }