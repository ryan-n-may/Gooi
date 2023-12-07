package compositions
import (
	intf			"gooi/interfaces"
	cons 			"gooi/base/constants"	
	ly 	 			"gooi/base/compositions/layout"
	comp 			"gooi/base/components"

	"fmt"
)
type Box struct {
	name 			string
	layout  		*ly.Layout

	alignment 		int
	displayables 	[]intf.Displayable

	slaveWidth, slaveHeight float32
	slaveWidthRatio, slaveHeightRatio float32

	posX, posY, posZ float32

	backgroundRectangle *comp.Rectangle_Struct
}

func NewBoxComposition(
	name string, 
	canvas intf.Canvas_Interface, 
	masterStruct intf.Displayable,
	x, y, z, slaveWidthRatio, slaveHeightRatio float32, 
	alignment int, 
	colour [3]float32,
) (*Box) {
		var box = Box{
			name,
			ly.NewLayout(canvas, masterStruct),
			alignment,

			make([]intf.Displayable, 1),

			slaveWidthRatio * masterStruct.GetWidth(), slaveHeightRatio * masterStruct.GetHeight(),
			slaveWidthRatio, slaveHeightRatio,
			x, y, z, 
			nil,
		}
		(&box).backgroundRectangle = comp.NewRectangle(
			canvas, 
			&box,
			fmt.Sprintf("%s_rectangle", name),
			0, 0, 
			0, 0, 0, 
			0, 
			colour,
			cons.FILL_MASTER_DIMENSIONS,
			cons.MATCH_MASTER_POSITION,
		)
		(&box).ArrangeLayout()
		return &box
}
func (box *Box) AddDisplayable(drawable intf.Displayable){
	box.displayables[0] = drawable
	box.ArrangeLayout()
}
func (box *Box) GetDisplayables() intf.Displayable {
	return box.displayables[0]
}

func (box *Box) ArrangeLayout(){
	
	box.slaveWidth = box.slaveWidthRatio * box.layout.GetMasterStruct().GetWidth()
	box.slaveHeight = box.slaveHeightRatio * box.layout.GetMasterStruct().GetHeight()

	fmt.Printf("Box slave width %v\n", box.slaveWidth)
	fmt.Printf("Box slave height %v\n", box.slaveHeight)

	if box.displayables[0] != nil {
		var componentWidth = box.displayables[0].GetWidth()
		var componentHeight = box.displayables[0].GetHeight()

		if box.alignment == cons.ALIGN_BOTTOM_LEFT {
			box.displayables[0].SetPos(
				box.posX + 0, 
				box.posY + 0,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_BOTTOM_RIGHT {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth - componentWidth, 
				box.posY + 0,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_TOP_LEFT {
			box.displayables[0].SetPos(
				box.posX + 0, 
				box.posY + box.slaveHeight - componentHeight,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_TOP_RIGHT {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth - componentWidth, 
				box.posY + box.slaveHeight - componentHeight,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_CENTRE {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth/2 - componentWidth/2, 
				box.posY + box.slaveHeight/2 - componentHeight/2,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_CENTRE_LEFT {
			box.displayables[0].SetPos(
				box.posX + 0, 
				box.posY + box.slaveHeight/2 - componentHeight/2,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_CENTRE_RIGHT {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth - componentWidth, 
				box.posY + box.slaveHeight/2 - componentHeight/2,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_TOP_CENTRE {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth/2 - componentWidth/2, 
				box.posY + box.slaveHeight - componentHeight,
				box.posZ,
			)
		} else if box.alignment == cons.ALIGN_BOTTOM_CENTRE {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth/2 - componentWidth/2, 
				box.posY + 0,
				box.posZ,
			)
		} else {
			box.displayables[0].SetPos(
				box.posX + box.slaveWidth/2 - componentWidth/2, 
				box.posY + box.slaveHeight/2 - componentHeight/2,
				box.posZ,
			)
		}
	}
}
func (box *Box) Draw(){
	box.backgroundRectangle.Draw()
	if box.displayables[0] != nil {
		box.displayables[0].Draw()
	}
}
func (box *Box) Redraw(){
	box.ArrangeLayout()
	box.backgroundRectangle.Redraw()
	if box.displayables[0] != nil {
		box.displayables[0].Redraw()
	}
}
func (box *Box) SetPos(x, y, z float32){
	box.posX = x
	box.posY = y
	box.posZ = z 
	box.backgroundRectangle.SetPos(x, y, z)
	box.Redraw()
}
func (box *Box) GetPos() (float32, float32, float32){ return box.posX, box.posY, box.posZ }

func (box *Box) SetWidth(w float32) { box.slaveWidth = w }
func (box *Box) SetHeight(h float32) { box.slaveHeight = h }

func (box *Box) GetWidth() float32 { return box.slaveWidth }
func (box *Box) GetHeight() float32 { return box.slaveHeight }

func (box *Box) GetMasterStruct() intf.Displayable { return box.layout.GetMasterStruct() }
func (box *Box) SetMasterStruct(master intf.Displayable) { box.layout.SetMasterStruct(master) }