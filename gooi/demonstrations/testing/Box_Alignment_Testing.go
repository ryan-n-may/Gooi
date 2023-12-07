package main

import (
	event 		"gooi/base/event"
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
	var centreComposition = ompo.NewBoxComposition("Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE, colours.NONE)
	var topLeftComposition = ompo.NewBoxComposition("Top Left", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_TOP_LEFT, colours.NONE)
	var topRightComposition = ompo.NewBoxComposition("Top Right", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_TOP_RIGHT, colours.NONE)
	var topCentreComposition = ompo.NewBoxComposition("Top Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_TOP_CENTRE, colours.NONE)
	var leftComposition = ompo.NewBoxComposition("Left", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE_LEFT, colours.NONE)
	var rightComposition = ompo.NewBoxComposition("Right", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_CENTRE_RIGHT, colours.NONE)
	var bottomLeftComposition = ompo.NewBoxComposition("Bottom Left", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_BOTTOM_LEFT, colours.NONE)
	var bottomRightComposition = ompo.NewBoxComposition("Bottom Right", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_BOTTOM_RIGHT, colours.NONE)
	var bottomCentreComposition = ompo.NewBoxComposition("Bottom Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, cons.ALIGN_BOTTOM_CENTRE, colours.NONE)

	// Now we create the buttons to assign to each composition.. We assign them via addition of the displayable to the composition, 
	// as well as the assignment of the composition as the masterStruct of the button...
	var centreBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		centreComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	centreComposition.AddDisplayable(centreBox)
	var topLeftBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		topLeftComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	topLeftComposition.AddDisplayable(topLeftBox)
	var topRightBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		topRightComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	topRightComposition.AddDisplayable(topRightBox)
	var topCentreBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		topCentreComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	topCentreComposition.AddDisplayable(topCentreBox)
	var leftBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		leftComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	leftComposition.AddDisplayable(leftBox)
	var rightBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		rightComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	rightComposition.AddDisplayable(rightBox)
	var bottomLeftBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		bottomLeftComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	bottomLeftComposition.AddDisplayable(bottomLeftBox)
	var bottomRightBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		bottomRightComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	bottomRightComposition.AddDisplayable(bottomRightBox)
	var bottomCentreBox = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		bottomCentreComposition,
		0, 0, 0, 0.1, 0.1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	bottomCentreComposition.AddDisplayable(bottomCentreBox)

	A.GetWindowCanvas().AddDisplayable(centreComposition)
	A.GetWindowCanvas().AddDisplayable(topRightComposition)
	A.GetWindowCanvas().AddDisplayable(bottomRightComposition)
	A.GetWindowCanvas().AddDisplayable(bottomLeftComposition)
	A.GetWindowCanvas().AddDisplayable(bottomCentreComposition)
	A.GetWindowCanvas().AddDisplayable(topCentreComposition)
	A.GetWindowCanvas().AddDisplayable(leftComposition)
	A.GetWindowCanvas().AddDisplayable(rightComposition)
	A.GetWindowCanvas().AddDisplayable(topLeftComposition)

	
	A.RunWindow()
}


