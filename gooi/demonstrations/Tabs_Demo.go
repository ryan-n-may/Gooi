package main

import (
	event 		"gooi/base/event"
	comp 		"gooi/base/components"
	listeners 	"gooi/base/listeners"
	windows 	"gooi/base/windows"
	ompo   		"gooi/base/compositions"
	cons        "gooi/base/constants"
	colours     "gooi/base/colours"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Tabs Alignment Demo", 800, 500)
	A.OpenWindow()
	A.SetBackgroundColour(colours.BLUE)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
	
	// Here we create multiple stacked compositions. Each box composition has a differing align constant.
	var tabsComposition = ompo.NewTabComposition(
		"Tabs",
		A.WindowCanvas,
		A.WindowCanvas,
		E, 
		ML,
		[]string{"Hello", "World"},
		0, 0, 0, 1.0, 1.0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
	)

	
	var column = ompo.NewColumnComposition(
		"Column", A.WindowCanvas, tabsComposition, 0, 0, 0, 0.5, 0.5, cons.ALIGN_LEFT)

	var Button1 = comp.CreateButton( 
		A.WindowCanvas, column, "Button1",
		200, 60, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var Button2 = comp.CreateButton( A.WindowCanvas, tabsComposition, "Button2",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)

	column.AddDisplayable(Button1)
	
	tabsComposition.AddDisplayable(column, 0)
	tabsComposition.AddDisplayable(Button2, 1)
	
	A.GetWindowCanvas().AddDisplayable(tabsComposition)

	A.RunWindow()
}


