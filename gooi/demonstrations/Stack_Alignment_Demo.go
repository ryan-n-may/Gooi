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
	var A = windows.NewWindow("Stack Alignment Demo", 800, 400)
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
	var alignments = []int{
		cons.ALIGN_CENTRE,
		cons.ALIGN_TOP_LEFT, 
		cons.ALIGN_TOP_RIGHT,
		cons.ALIGN_TOP_CENTRE,
		cons.ALIGN_CENTRE_LEFT,
		cons.ALIGN_CENTRE_RIGHT,
		cons.ALIGN_BOTTOM_LEFT, 
		cons.ALIGN_BOTTOM_RIGHT, 
		cons.ALIGN_BOTTOM_CENTRE,
	}
	var stack = ompo.NewStackComposition("Stack", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, window_w, window_h, alignments)
	
	// Now we create the buttons to assign to each composition.. We assign them via addition of the displayable to the composition, 
	// as well as the assignment of the composition as the masterStruct of the button...
	var centreButton = comp.CreateButton( A.WindowCanvas, stack, "Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var topLeftButton = comp.CreateButton( A.WindowCanvas, stack, "Top Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var topRightButton = comp.CreateButton( A.WindowCanvas, stack, "Top Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var topCentreButton = comp.CreateButton( A.WindowCanvas, stack, "Top Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var leftButton = comp.CreateButton( A.WindowCanvas, stack, "Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var rightButton = comp.CreateButton( A.WindowCanvas, stack, "Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var bottomLeftButton = comp.CreateButton( A.WindowCanvas, stack, "Bottom Left",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var bottomRightButton = comp.CreateButton( A.WindowCanvas, stack, "Bottom Right",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var bottomCentreButton = comp.CreateButton( A.WindowCanvas, stack, "Bottom Centre",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	
	stack.AddDisplayable(centreButton)
	stack.AddDisplayable(topLeftButton)
	stack.AddDisplayable(topRightButton)
	stack.AddDisplayable(topCentreButton)
	stack.AddDisplayable(leftButton)
	stack.AddDisplayable(rightButton)
	stack.AddDisplayable(bottomLeftButton)
	stack.AddDisplayable(bottomRightButton)
	stack.AddDisplayable(bottomCentreButton)

	A.GetWindowCanvas().AddDisplayable(stack)
	
/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(centreButton)
	
	A.RunWindow()
}


