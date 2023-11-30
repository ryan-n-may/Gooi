package layout 

import (
	intf 		"gooi/interfaces"
)

type Layout struct {
	canvas intf.Canvas_Interface

	masterStruct intf.Displayable
	masterWidth float32
	masterHeight float32

	slaveWidth float32
	slaveHeight	float32

	openGLWindowWidth float32
	openGLWindowHeight float32
}

func NewLayout(canvas intf.Canvas_Interface, masterStruct intf.Displayable) *Layout {
	var layout = Layout{
		canvas, 
		masterStruct, 

		masterStruct.GetWidth(),
		masterStruct.GetHeight(),

		0, 0, 

		canvas.GetWidth(), 
		canvas.GetHeight(),
	}
	return &layout
}

func (layout *Layout) GetWidth() float32 { return layout.slaveWidth }
func (layout *Layout) GetHeight() float32 { return layout.slaveHeight }

func (layout *Layout) SetWidth(w float32) { layout.slaveWidth = w }
func (layout *Layout) SetHeight(h float32) { layout.slaveHeight = h }

func (layout *Layout) GetMasterStruct() intf.Displayable { return layout.masterStruct }
func (layout *Layout) SetMasterStruct(master intf.Displayable) { 
	layout.masterStruct = master 
	layout.masterWidth = master.GetWidth()
	layout.masterHeight = master.GetHeight()
}

func (layout *Layout) GetCanvas() intf.Canvas_Interface { return layout.canvas }