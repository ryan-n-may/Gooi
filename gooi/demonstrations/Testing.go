package main

import (
	fmt  		"fmt"
	event 		"gooi/base/event"
	comp 		"gooi/base/components"
	listeners 	"gooi/base/listeners"
	cons 		"gooi/base/constants"
	windows 	"gooi/base/windows"
	intf 		"gooi/interfaces"
	colours     "gooi/base/colours"

)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Testing", 800, 600)
	A.OpenWindow()
	A.SetBackgroundColour(colours.WHITE)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating button 1 **/
	// creating event arguments for button method
	var event_arguments = event.NewEventParameter([]string{"Hello World", " and Jupiter!"})
	// creating event 
	var test_event = &event.Event_Struct{
		test, 
		"test",
		event_arguments,
	}
	// registering the event to the event handler
	E.RegisterEventToHandler(test_event)
	// creating the button and assinging it the visual characteristics + the trigger event
	var Button = comp.CreateButton(
		A.WindowCanvas, 
		A.WindowCanvas, 
		"Button1",
		100, 50, 20, 
		0, 0, 0,
		"luxi", 
		"base/components/fonts/luxi.ttf",
		16,
		test_event,
		200,
	)

	var checkbox = comp.CreateCheckbox(
		A.WindowCanvas, 
		A.WindowCanvas,
		"Checkbox Testing",
		10, 200, 200, 0.0,
		&event.NULL_EVENT, 
		"luxi",
		"base/components/fonts/luxi.ttf",
		16, 
	)

	var label = comp.CreateLabel(
		A.WindowCanvas,
		A.WindowCanvas,
		"label",
		100, 300, 0, 
		"luxi", 
		"base/components/fonts/luxi.ttf",
		16,
	)

	var toggle = comp.CreateToggle(
		A.WindowCanvas, 
		A.WindowCanvas, 
		"toggle", 
		30, 30, 
		300, 300, 0.0, 
		"luxi", 
		"base/components/fonts/luxi.ttf",
		16,
		&event.NULL_EVENT,
	)
	
	var rectangle = comp.CreateRectangle(
		A.WindowCanvas, 
		toggle, 
		"label rectangle",
		0, 0, 
		0, 0, 0, 
		colours.GRAY,
		cons.FILL_MASTER_DIMENSIONS,
	)

	A.GetWindowCanvas().AddComponent(Button)
	A.GetWindowCanvas().AddComponent(checkbox)
	A.GetWindowCanvas().AddComponent(label)
	A.GetWindowCanvas().AddComponent(rectangle)
	A.GetWindowCanvas().AddComponent(toggle)

/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(Button)
	A.GetMouseHandler().RegisterClickableToHandler(checkbox)
	A.GetMouseHandler().RegisterClickableToHandler(toggle)

	A.RunWindow()
}

func test(param intf.Paramaters_Interface){
	fmt.Println(param.GetParameters().([]string)[0], param.GetParameters().([]string)[1])
}