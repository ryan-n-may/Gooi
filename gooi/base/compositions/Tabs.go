package compositions
import (
	intf 	"gooi/interfaces"
	cons 	"gooi/base/constants"
	comp 	"gooi/base/components"
	event 	"gooi/base/event"
	ly 		"gooi/base/compositions/layout"
	colours "gooi/base/colours"
	fmt 	"fmt"
)
type Tabs struct {
	name string
	layout *ly.Layout

	labels 				[]string

	mouseHandler 		intf.MouseHandler_Interface
	eventHandler 		intf.EventHandler_Interface

	slaveWidth, slaveHeight float32
	slaveWidthRatio, slaveHeightRatio float32

	posX, posY, posZ float32

	// master box that stores all components of tabs 
	// and column structure to store button row and stacked tabs
	masterTabBox 		*Box
	masterColumn		*Column	
	// button row to store buttons for each tab
	tabButtonBox 		*Box
	tabButtonRow 		*Row
	tabButton 			[]*comp.Button
	// actual boxes (super imposed)
	// needs a rectangular background element to do correct z-position overlaying. 
	tabStackBox 		*Box
	tabStack  	 		*Stack
	tabs  	 			[]*Box
}

func NewTabComposition(
	name string, 
	canvas intf.Canvas_Interface, 
	masterStruct intf.Displayable,
	eventhandler intf.EventHandler_Interface, 
	mousehandler intf.MouseHandler_Interface, 
	labels []string, 
	x, y, z, slaveWidthRatio, slaveHeightRatio float32,
	font_name, font_path string,
	font_size int,
) (*Tabs) {
	var tabs = Tabs{
		name,
		ly.NewLayout(canvas, masterStruct),
		labels,
		mousehandler, eventhandler,

		slaveWidthRatio * masterStruct.GetWidth(), 
		slaveHeightRatio * masterStruct.GetHeight(),
		slaveWidthRatio, slaveHeightRatio,

		x, y, z,
		// master struct of these nested structs are the struct being created now.
		// so we must initialise these later to avoid throwing an error. 
		nil, nil, nil, nil, nil, nil, nil, nil,
	}

	/** Creating container that holds everything **/
	(&tabs).masterTabBox = NewBoxComposition(
		fmt.Sprintf("%s_masterBox", name),
		canvas, 
		masterStruct,
		0, 0, 0, slaveWidthRatio, slaveHeightRatio,
		cons.ALIGN_TOP_LEFT,
		colours.NONE,
	)

	(&tabs).masterColumn = NewColumnComposition(
		fmt.Sprintf("%s_masterColumn", name),
		canvas,
		(&tabs).masterTabBox,
		0, 0, 0,
		cons.ALIGN_TOP_LEFT,
	)
	
	(&tabs).tabButtonRow = NewRowComposition(
		fmt.Sprintf("%s_tabs_button_row", name),
		canvas,
		(&tabs).masterColumn,
		0, 0, 0, 
		cons.ALIGN_TOP,
	)

	(&tabs).tabButton = make([]*comp.Button, len(labels))
	for i, label := range labels {
		fmt.Printf("Creating tab button with label %s\n", label)
		var event_arguments = event.NewEventParameter(i)
		// creating event 
		var selectTab_event = &event.Event_Struct{
			(&tabs).MoveToFront, 
			fmt.Sprintf("selectTab_%v", i),
			event_arguments,
		}
		(&tabs).tabButton[i] = comp.NewButton(
			canvas,
			(&tabs).tabButtonRow,
			label,
			100, 20,
			5,
			0, 0, 0, 
			font_name, font_path,
			font_size, 
			selectTab_event,
			100,
		)
		canvas.GetWindow().GetMouseHandler().RegisterClickableToHandler((&tabs).tabButton[i])
	}
	

	(&tabs).tabStackBox = NewBoxComposition(
		fmt.Sprintf("%s_tabs_stack_box", name),
		canvas,
		(&tabs).masterColumn,
		0, 0, 0, 1.0, 0.9,
		cons.ALIGN_TOP_LEFT,
		colours.LIGHT_BLUE,
	)

	/** Creating tabs (boxes) based upon given labels **/
	var alignment = []int{}
	for i := 0; i < len(labels) ; i++ {
		alignment =  append(alignment, cons.ALIGN_TOP_LEFT)
	}

	(&tabs).tabStack = NewStackComposition(
		fmt.Sprintf("%s_tabBox", name),
		canvas, 
		tabs.tabStackBox,
		0, 0, 0, 1, 1, 
		alignment,
	)

	(&tabs).tabs = make([]*Box, len(labels))
	for i, label := range labels {
		fmt.Printf("Creating tab box for label %s %v\n", label, i)
		(&tabs).tabs[i] = NewBoxComposition(
			fmt.Sprintf("%s_tab_%s", name, label), 
			canvas, 
			(&tabs).tabStack,	
			0, 0, 0, 1, 1, 
			cons.ALIGN_TOP_LEFT,
			colours.LIGHT_GRAY,
		)
	}

	for _, tab := range (&tabs).tabs { (&tabs).tabStack.AddDisplayable(tab) }
	(&tabs).tabStackBox.AddDisplayable((&tabs).tabStack)

	for _, button := range (&tabs).tabButton { (&tabs).tabButtonRow.AddDisplayable(button) }

	(&tabs).masterColumn.AddDisplayable((&tabs).tabStackBox)
	(&tabs).masterColumn.AddDisplayable((&tabs).tabButtonRow)

	(&tabs).masterTabBox.AddDisplayable((&tabs).masterColumn)

	return &tabs
}

func (tabs *Tabs) MoveToFront(param intf.Paramaters_Interface) {
	var Index = param.GetParameters().(int)
	fmt.Printf("Moving %v to the front\n", Index)
	for i, _ := range tabs.tabs {
		var x, y, _ = tabs.tabs[i].GetPos()
		tabs.tabs[i].SetPos(x, y, 0.0) 
	}
	var x, y, _ = tabs.tabs[Index].GetPos()
	tabs.tabs[Index].SetPos(x, y, 1.0) 
	tabs.Draw()
}

func (tabs *Tabs) ArrangeLayout() {
	tabs.masterTabBox.ArrangeLayout()
	tabs.slaveWidth = tabs.masterTabBox.GetWidth()
	tabs.slaveHeight = tabs.masterTabBox.GetHeight()	
}

func (tabs *Tabs) Draw(){
	tabs.masterTabBox.Draw()
}

func (tabs *Tabs) Redraw(){
	tabs.masterTabBox.Redraw()
}

func (tabs *Tabs) SetPos(x, y, z float32){
	tabs.posX = x
	tabs.posY = y
	tabs.posZ = z
	tabs.Redraw()
}
func (tabs *Tabs) GetPos() (float32, float32, float32){
	return tabs.posX, tabs.posY, tabs.posZ
}

func (tabs *Tabs) AddDisplayable(displayable intf.Displayable, tab int) {
	tabs.tabs[tab].AddDisplayable(displayable)
}

func (tabs *Tabs) GetWidth() float32 { return tabs.masterTabBox.GetWidth() }
func (tabs *Tabs) GetHeight() float32 { return tabs.masterTabBox.GetHeight() }

func (tabs *Tabs) SetWidth(w float32) { tabs.slaveWidth = w }
func (tabs *Tabs) SetHeight(h float32) { tabs.slaveHeight = h }

func (tabs *Tabs) GetMasterStruct() intf.Displayable { return tabs.layout.GetMasterStruct() }
func (tabs *Tabs) SetMasterStruct(master intf.Displayable) { tabs.layout.SetMasterStruct(master) }