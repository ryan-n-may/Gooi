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
	var A = windows.NewWindow("Row Alignment Demo", 800, 300)
	A.OpenWindow()
	A.SetBackgroundColour(colours.BLUE)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating button 1 **/
	var window_w float32 = 1.0
	var window_h float32 = 1.0

	// Here we create multiple stacked compositions. Each box composition has a differing align constant.
	var rowComposition = ompo.NewRowComposition(
		"Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE_COLUMN)
	
	// Now we create the buttons to assign to each composition.. We assign them via addition of the displayable to the composition, 
	// as well as the assignment of the composition as the masterStruct of the button...
	var Button1 = comp.CreateButton( A.WindowCanvas, rowComposition, "Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var Button2 = comp.CreateButton( A.WindowCanvas, rowComposition, "Top Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var Button3 = comp.CreateButton( A.WindowCanvas, rowComposition, "Top Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var Button4 = comp.CreateButton( A.WindowCanvas, rowComposition, "Top Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	rowComposition.AddDisplayable(Button1)
	rowComposition.AddDisplayable(Button2)
	rowComposition.AddDisplayable(Button3)
	rowComposition.AddDisplayable(Button4)
	
	A.GetWindowCanvas().AddDisplayable(rowComposition)

/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(Button1)
	
	A.RunWindow()
}


