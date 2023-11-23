package main

import (
	event 		"GUI/Base/Event"
	comp 		"GUI/Base/Components"
	listeners 	"GUI/Base/Listeners"
	windows 	"GUI/Base/Windows"
	ompo   		"GUI/Base/Compositions"
	cons        "GUI/Base/Constants"
	colours     "GUI/Base/Colours"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Row Alignment Demo", 1000, 600)
	A.OpenWindow()
	A.SetBackgroundColour(colours.LIGHT_GRAY)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating button 1 **/
	// creating the button and assinging it the visual characteristics + the trigger event
	var Button1 = comp.CreateButton(
		"Button1",
		A.WindowCanvas, 
		100, 50, 300, 50, 
		20,
		2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)
	
	var Button2 = comp.CreateButton(
		"Button2",
		A.WindowCanvas, 
		100, 50, 300, 50, 
		20,
		2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Button3 = comp.CreateButton(
		"Button3",
		A.WindowCanvas, 
		150, 50, 300, 50, 
		20,
		2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Button4 = comp.CreateButton(
		"Button4",
		A.WindowCanvas, 
		100, 50, 300, 50, 
		20,
		2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Button5 = comp.CreateButton(
		"Button5",
		A.WindowCanvas, 
		80, 30, 300, 50, 
		20,
		2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)
	
	var Row = ompo.NewRowComposition("Row 1", 
						200, 100, 30, 
						cons.ALIGN_CENTRE_ROW)
	
	var Box = ompo.NewBoxComposition("Centre Right Composition", 
						0, 0, 
						A.GetWindowWidth(), 
						A.GetWindowHeight(), 
						cons.ALIGN_CENTRE)

	Row.AddDrawable(Button1)
	Row.AddDrawable(Button2)
	Row.AddDrawable(Button3)
	Row.AddDrawable(Button4)
	Row.AddDrawable(Button5)

	Box.AddDrawable(Row)

	A.GetWindowCanvas().AddComponent(Box)
	
	A.RunWindow()
}

