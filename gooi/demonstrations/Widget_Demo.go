package main

import (
	fmt  		"fmt"
	event 		"GUI/Base/Event"
	comp 		"GUI/Base/Components"
	listeners 	"GUI/Base/Listeners"
	windows 	"GUI/Base/Windows"
	intf 		"GUI/Interfaces"
	ompo   		"GUI/Base/Compositions"
	cons        "GUI/Base/Constants"
	colours     "GUI/Base/Colours"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Widget Demo", 1000, 600)
	A.OpenWindow()
	A.SetBackgroundColour(colours.LIGHT_GRAY)
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
	var Button_1 = comp.CreateButton(
		"Button1",
		A.WindowCanvas, 
		100, 50, 300, 50, 
		20,
		2, 2,
		250,
		test_event,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)
/** Creating button 2 **/
	event_arguments = event.NewEventParameter(Button_1)
	var disable_event = &event.Event_Struct{
		disablebutton1, 
		"disablebutton1",
		event_arguments,
	}
	E.RegisterEventToHandler(disable_event)
	var Button_2 = comp.CreateButton(
		"Button2",
		A.WindowCanvas, 
		100, 30, 300, 50, 
		10,
		2, 2,
		250,
		disable_event,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Button_3 = comp.CreateThemedButton(
		"Button2B",
		A.WindowCanvas, 
		150, 50, 300, 50, 
		16,
		colours.GrantiteTheme(),	
		&event.NULL_EVENT,
	)


	var Label = comp.CreateLabel(
		"Label Component",
		A.WindowCanvas, 
		140.0, 30.0, 
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		32,
	)
	
	var Toggle = comp.CreateToggleSwitch(
		"Toggle Switch 1",
		A.WindowCanvas, 
		30, 20, 150, 150, 
		2,
		10,
		&event.NULL_EVENT,
		colours.GREEN,
		colours.RED,
		colours.LIGHT_GRAY,
		colours.GRAY,
		colours.WHITE,
	) 

	var Checkbox = comp.CreateCheckbox(
		"Checkbox 1",
		A.WindowCanvas,
		8, 50, 50, 
		2,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.LIGHT_BLUE,
		colours.WHITE,
		colours.GRAY,
		colours.LIGHT_GRAY,
		colours.WHITE,
	) 

	var RadioButton = ompo.CreateRadioButton(
		"Radio Button",
		A.WindowCanvas,
		8, 50, 100, 
		2,
		4,
		cons.VERTICAL_ORIENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		[]string{"Green", "Red", "Blue", "Orange"},
		colours.LIGHT_BLUE,
		colours.WHITE,
		colours.GRAY,
		colours.LIGHT_GRAY,
		colours.WHITE,
	)

	var KeyListener = listeners.CreateKeyListener("KeyListener", A.GetWindowCanvas())

	var text_input = comp.CreateTextInput(
		"Text Input",
		"Placeholder",
		A.GetWindowCanvas(),
		0, 0, 300, 30,
		10,
		2,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.GRAY,
		colours.WHITE,
		colours.LIGHT_GRAY,
		colours.LIGHT_BLUE,
		cons.MOUSE_PRESSED,
		KeyListener,
	)

	var text_input_2 = comp.CreateTextInput(
		"Text Input 2",
		"Placeholder",
		A.GetWindowCanvas(),
		0, 0, 100, 30,
		10,
		2,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.GRAY,
		colours.WHITE,
		colours.LIGHT_GRAY,
		colours.LIGHT_BLUE,
		cons.MOUSE_PRESSED,
		KeyListener,
	)

	var row_composition = ompo.NewRowComposition("Row 1", 150, 50, 30, cons.ALIGN_CENTRE_ROW)
	row_composition.AddDrawable(Button_1)
	row_composition.AddDrawable(Button_2)
	row_composition.AddDrawable(Button_3)
	
	var column_composition = ompo.NewColumnComposition("Column 1", 200, 100, 30, cons.ALIGN_CENTRE_COLUMN)
	column_composition.AddDrawable(RadioButton)
	column_composition.AddDrawable(Toggle)
	column_composition.AddDrawable(Checkbox)
	column_composition.AddDrawable(text_input)
	column_composition.AddDrawable(text_input_2)
	column_composition.AddDrawable(row_composition)
	column_composition.AddDrawable(Label)
	
	var master_box_composition = ompo.NewBoxComposition("Master composition", 
						0, 0, 
						A.GetWindowWidth(), 
						A.GetWindowHeight(), 
						cons.ALIGN_CENTRE)
	master_box_composition.AddDrawable(column_composition)

	A.GetWindowCanvas().AddComponent(master_box_composition)

/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(Button_1)
	A.GetMouseHandler().RegisterClickableToHandler(Button_2)	
	A.GetMouseHandler().RegisterClickableToHandler(Button_3)	
	A.GetMouseHandler().RegisterClickableToHandler(Toggle)	
	A.GetMouseHandler().RegisterClickableToHandler(Checkbox)	
	A.GetMouseHandler().RegisterClickableToHandler(text_input)
	A.GetMouseHandler().RegisterClickableToHandler(text_input_2)

	A.RunWindow()
}

func test(param intf.Paramaters_Interface){
	fmt.Println(param.GetParameters().([]string)[0], param.GetParameters().([]string)[1])
}

func disablebutton1(param intf.Paramaters_Interface){
	param.GetParameters().(intf.Clickable_Interface).SetClickable(false)
	fmt.Println(param.GetParameters().(intf.Component_Interface).GetName(), " set to no-clickable")
}
