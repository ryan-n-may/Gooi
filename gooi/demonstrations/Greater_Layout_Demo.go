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
	
/**
 * TABS 
 **/
	var tabsComposition = ompo.NewTabComposition(
		"Tabs",
		A.WindowCanvas,
		A.WindowCanvas,
		E, 
		ML,
		[]string{"Hello", "World"},
		0, 0, 0, 0.5, 0.5,
		"luxi", "base/components/fonts/luxi.ttf", 16,
	)
	var Button1 = comp.CreateButton( A.WindowCanvas, tabsComposition, "Button1",
		200, 60, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	var Button2 = comp.CreateButton( A.WindowCanvas, tabsComposition, "Button2",
		100, 50, 20, 0, 0, 0,
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200,
	)
	tabsComposition.AddDisplayable(Button1, 0)
	tabsComposition.AddDisplayable(Button2, 1)
/** 
 * COLUMN
 **/
 	var columnComposition = ompo.NewColumnComposition(
		"Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, 0.5, 1.0, cons.ALIGN_CENTRE_COLUMN)
	
	// Now we create the buttons to assign to each composition.. We assign them via addition of the displayable to the composition, 
	// as well as the assignment of the composition as the masterStruct of the button...
	var Box1 = ompo.NewBoxComposition( 
		"box1",
		A.WindowCanvas, 
		columnComposition,
		0, 0, 0, 1.0, 0.20, 
		cons.ALIGN_CENTRE,
		colours.GREEN,
	)
	var Box2 = ompo.NewBoxComposition( 
		"box2",
		A.WindowCanvas, 
		columnComposition,
		0, 0, 0, 1.0, 0.20, 
		cons.ALIGN_CENTRE,
		colours.GREEN,
	)
	var Box3 = ompo.NewBoxComposition( 
		"box3",
		A.WindowCanvas, 
		columnComposition,
		0, 0, 0, 1.0, 0.20, 
		cons.ALIGN_CENTRE,
		colours.GREEN,
	)
	var Box4 = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		columnComposition,
		0, 0, 0, 1.0, 0.20, 
		cons.ALIGN_CENTRE,
		colours.GREEN,
	)
	columnComposition.AddDisplayable(Box1)
	columnComposition.AddDisplayable(Box2)
	columnComposition.AddDisplayable(Box3)
	columnComposition.AddDisplayable(Box4)
/**
 * ROW 
 **/
 	// Here we create multiple stacked compositions. Each box composition has a differing align constant.
	var rowComposition = ompo.NewRowComposition(
		"row", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, 0.5, 0.5, cons.ALIGN_CENTRE_COLUMN)
	
	// Now we create the buttons to assign to each composition.. We assign them via addition of the displayable to the composition, 
	// as well as the assignment of the composition as the masterStruct of the button...
	var Box5 = ompo.NewBoxComposition( 
		"box1",
		A.WindowCanvas, 
		rowComposition,
		0, 0, 0, 0.2, 1, 
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	var Box6 = ompo.NewBoxComposition( 
		"box2",
		A.WindowCanvas, 
		rowComposition,
		0, 0, 0, 0.2, 1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	var Box7 = ompo.NewBoxComposition( 
		"box3",
		A.WindowCanvas, 
		rowComposition,
		0, 0, 0, 0.2, 1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	var Box8 = ompo.NewBoxComposition( 
		"box4",
		A.WindowCanvas, 
		rowComposition,
		0, 0, 0, 0.2, 1,
		cons.ALIGN_CENTRE,
		colours.RED,
	)
	rowComposition.AddDisplayable(Box5)
	rowComposition.AddDisplayable(Box6)
	rowComposition.AddDisplayable(Box7)
	rowComposition.AddDisplayable(Box8)


/** 
 * MASTER STACK 
 **/

 	var alignments = []int{
		cons.ALIGN_BOTTOM_RIGHT,
		cons.ALIGN_BOTTOM_LEFT, 
		cons.ALIGN_TOP_LEFT,
	}
	var stack = ompo.NewStackComposition(
		"Stack", 
		A.WindowCanvas, A.WindowCanvas, 
		0, 0, 0, 1.0, 1.0, alignments)
	
	stack.AddDisplayable(columnComposition)
	stack.AddDisplayable(tabsComposition)
	stack.AddDisplayable(rowComposition)

	A.GetWindowCanvas().AddDisplayable(stack)

	A.RunWindow()
}


