package compositions
import (
	intf 	"gooi/interfaces"
	cons 	"gooi/base/constants"
	comp 	"gooi/base/components"
	colours "gooi/base/colours"
	event 	"gooi/base/event"
	fmt 	"fmt"
	log 	"log"
)
type Tabs_Struct struct {
	TabCanvas 			*comp.Canvas_Struct

	TabsName 			string

	Master_Pos_x 		float32
	Master_Pos_y 		float32
	Master_Pos_z 		float32
	Master_Height 		*float32
	Master_Width 		*float32

	Labels 				[]string
	
	TabButtonRow 		*Row_Struct
	Buttons 			[]*comp.Button_Struct
	
	Tabs  	 			[]*Box_Struct
	TabStack 			*Stack_Struct
	TabContainer 		*Horisontal_Div_Struct

	MouseHandler 		intf.MouseHandler_Interface
}

func NewTabs(name string, canvas *comp.Canvas_Struct, eventhandler intf.EventHandler_Interface, mousehandler intf.MouseHandler_Interface, labels []string, posx, posy float32, width, height *float32) *Tabs_Struct {
	log.Println("new [Tabs].")
	var tabs = Tabs_Struct{}
	tabs.TabsName = name
	tabs.TabCanvas = canvas
	tabs.Master_Pos_x = posx
	tabs.Master_Pos_y = posy
	tabs.Master_Height = height
	tabs.Master_Width = width
	tabs.MouseHandler = mousehandler
	tabs.Labels = labels

	tabs.Tabs = make([]*Box_Struct, len(tabs.Labels))
	tabs.TabStack = NewStackComposition(fmt.Sprintf("TabStack_%s", tabs.TabsName), tabs.Master_Pos_x, tabs.Master_Pos_y, tabs.Master_Width, tabs.Master_Height, cons.ALIGN_TOP_LEFT)
	
	tabs.Buttons = make([]*comp.Button_Struct, len(tabs.Labels))
	tabs.TabButtonRow = NewRowComposition(fmt.Sprintf("TabButtonRow_%s", tabs.TabsName), tabs.Master_Pos_x, tabs.Master_Pos_y, 2, cons.ALIGN_CENTRE_ROW)

	tabs.TabContainer = NewHorisontalDivision(fmt.Sprintf("TabConatiner_%s", tabs.TabsName), []float32{0.9, 0.1}, tabs.Master_Pos_x, tabs.Master_Pos_y, tabs.Master_Width, tabs.Master_Height)
	


	for i, _ := range tabs.Tabs {
		tabs.Tabs[i] = NewBoxComposition(
			fmt.Sprintf("Tab_%s_%v", tabs.TabsName, i), 
			tabs.Master_Pos_x, tabs.Master_Pos_y, 
			tabs.Master_Width, tabs.Master_Height, 
			cons.ALIGN_TOP_LEFT,
		) 

		var move_to_front_arguments = event.NewEventParameter(i)
		// creating event 
		var move_to_front_event = &event.Event_Struct{
			tabs.MoveToFront, 
			fmt.Sprintf("MoveTabToFront_%s_%v", tabs.TabsName, i),
			move_to_front_arguments,
		}
		// registering the event to the event handler
		eventhandler.RegisterEventToHandler(move_to_front_event)

		// THIS BUTTON IS A PLACEHOLDER
		tabs.Buttons[i] = comp.CreateButton(
			tabs.Labels[i],
			tabs.TabCanvas, 
			100, 30, 300, 50, 
			10,
			2, 2,
			250,
			move_to_front_event,
			"Base/Components/Fonts/luxi.ttf",
			"luxi", 
			16,
			colours.BLUE,
			colours.DARK_BLUE,
			colours.GRAY,
			colours.GRAY,
			colours.WHITE,
		)
	}
	return &tabs
}

func (tabs *Tabs_Struct) MoveToFront(param intf.Paramaters_Interface) {
	var Index = param.GetParameters().(int)
	for i, _ := range tabs.Tabs {
		tabs.Tabs[i].SetPosZ(0.0) 
		log.Println(fmt.Sprintf("Setting tab in position %v to 0.0", i))
	}
	log.Println(fmt.Sprintf("Setting tab in position %v to 1.0", Index))
	tabs.Tabs[Index].SetPosZ(1.0)
	tabs.Redraw()
}

func (tabs *Tabs_Struct) MoveComponents() {
	
}

func (tabs *Tabs_Struct) AddDrawable(drawable intf.Drawable_Interface, index int) {
	tabs.Tabs[index].AddDrawable(drawable)
	tabs.TabButtonRow.AddDrawable(tabs.Buttons[index])
	tabs.MouseHandler.RegisterClickableToHandler(tabs.Buttons[index])
	tabs.TabStack.AddDrawable(tabs.Tabs[index])
	tabs.TabContainer.AddDrawable(tabs.TabStack, 0)
	tabs.TabContainer.AddDrawable(tabs.TabButtonRow, 1)
}

func (tabs *Tabs_Struct) Draw(){
	tabs.TabContainer.Draw()
}

func (tabs *Tabs_Struct) Redraw(){
	tabs.TabContainer.Redraw()
}

func (tabs *Tabs_Struct) SetPos(x, y float32){
	tabs.Master_Pos_x = x
	tabs.Master_Pos_y = y
	tabs.Redraw()
}
func (tabs *Tabs_Struct) GetPos() (float32, float32){
	return tabs.Master_Pos_x, tabs.Master_Pos_y
}
func (tabs *Tabs_Struct) GetBounds() (float32, float32){
	return *(tabs.Master_Width), *(tabs.Master_Height)
}

func (tabs *Tabs_Struct) SetPosZ(z float32) {
	tabs.Master_Pos_z = z
}
func (tabs *Tabs_Struct) GetPosZ() float32 {
	return tabs.Master_Pos_z
}