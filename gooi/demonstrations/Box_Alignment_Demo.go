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
	var A = windows.NewWindow("Box Alignment Demo", 800, 400)
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
	var centreComposition = ompo.NewBoxComposition("Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE)
	var topLeftComposition = ompo.NewBoxComposition("Top Left", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_TOP_LEFT)
	var topRightComposition = ompo.NewBoxComposition("Top Right", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_TOP_RIGHT)
	var topCentreComposition = ompo.NewBoxComposition("Top Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_TOP_CENTRE)
	var leftComposition = ompo.NewBoxComposition("Left", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE_LEFT)
	var rightComposition = ompo.NewBoxComposition("Right", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE_RIGHT)
	var bottomLeftComposition = ompo.NewBoxComposition("Bottom Left", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_BOTTOM_LEFT)
	var bottomRightComposition = ompo.NewBoxComposition("Bottom Right", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_BOTTOM_RIGHT)
	var bottomCentreComposition = ompo.NewBoxComposition("Bottom Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_BOTTOM_CENTRE)

	// Now we create the buttons to assign to each composition.. We assign them via addition of the displayable to the composition, 
	// as well as the assignment of the composition as the masterStruct of the button...
	var centreButton = comp.CreateButton( A.WindowCanvas, centreComposition, "Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	centreComposition.AddDisplayable(centreButton)
	var topLeftButton = comp.CreateButton( A.WindowCanvas, topLeftComposition, "Top Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	topLeftComposition.AddDisplayable(topLeftButton)
	var topRightButton = comp.CreateButton( A.WindowCanvas, topRightComposition, "Top Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	topRightComposition.AddDisplayable(topRightButton)
	var topCentreButton = comp.CreateButton( A.WindowCanvas, topCentreComposition, "Top Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	topCentreComposition.AddDisplayable(topCentreButton)
	var leftButton = comp.CreateButton( A.WindowCanvas, leftComposition, "Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	leftComposition.AddDisplayable(leftButton)
	var rightButton = comp.CreateButton( A.WindowCanvas, rightComposition, "Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	rightComposition.AddDisplayable(rightButton)
	var bottomLeftButton = comp.CreateButton( A.WindowCanvas, bottomLeftComposition, "Bottom Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	bottomLeftComposition.AddDisplayable(bottomLeftButton)
	var bottomRightButton = comp.CreateButton( A.WindowCanvas, bottomRightComposition, "Bottom Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	bottomRightComposition.AddDisplayable(bottomRightButton)
	var bottomCentreButton = comp.CreateButton( A.WindowCanvas, bottomCentreComposition, "Bottom Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	bottomCentreComposition.AddDisplayable(bottomCentreButton)

	A.GetWindowCanvas().AddDisplayable(centreComposition)
	A.GetWindowCanvas().AddDisplayable(topLeftComposition)
	A.GetWindowCanvas().AddDisplayable(topRightComposition)
	A.GetWindowCanvas().AddDisplayable(bottomRightComposition)
	A.GetWindowCanvas().AddDisplayable(bottomLeftComposition)
	A.GetWindowCanvas().AddDisplayable(bottomCentreComposition)
	A.GetWindowCanvas().AddDisplayable(topCentreComposition)
	A.GetWindowCanvas().AddDisplayable(leftComposition)
	A.GetWindowCanvas().AddDisplayable(rightComposition)

/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(centreButton)
	
	A.RunWindow()
}


