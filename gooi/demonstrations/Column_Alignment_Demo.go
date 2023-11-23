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
	var A = windows.NewWindow("Column Alignment Demo", 1000, 600)
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
		200, 50, 300, 50, 
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
		200, 50, 300, 50, 
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
		250, 50, 300, 50, 
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
		300, 50, 300, 50, 
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
		100, 30, 300, 50, 
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
	
	var Col = ompo.NewColumnComposition("Column 1", 
						200, 100, 30, 
						cons.ALIGN_CENTRE_COLUMN)
	
	var Box = ompo.NewBoxComposition("Centre Right Composition", 
						0, 0, 
						A.GetWindowWidth(), 
						A.GetWindowHeight(), 
						cons.ALIGN_CENTRE)

	Col.AddDrawable(Button1)
	Col.AddDrawable(Button2)
	Col.AddDrawable(Button3)
	Col.AddDrawable(Button4)
	Col.AddDrawable(Button5)

	Box.AddDrawable(Col)

	A.GetWindowCanvas().AddComponent(Box)
	
	A.RunWindow()
}
