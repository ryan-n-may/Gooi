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
	var A = windows.NewWindow("Column Alignment Demo", 800, 500)
	A.OpenWindow()
	A.SetBackgroundColour(colours.BLUE)
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
	var KL = listeners.CreateKeyListener("KeyListener", A.GetWindowCanvas())

	var columnComposition = ompo.NewColumnComposition(
		"Centre", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, cons.ALIGN_LEFT)
	
	var Checkbox1 = comp.NewCheckbox(A.WindowCanvas, columnComposition, "box1", 10, 0, 0, 0, &event.NULL_EVENT, "luxi", "base/components/fonts/luxi.ttf", 16)
	var Box1 = ompo.NewBoxComposition("box2", A.WindowCanvas, columnComposition, 0, 0, 0, 1.0, 0.25, cons.ALIGN_CENTRE, colours.DARK_GRAY)
	var Input1 = comp.NewTextInput(A.WindowCanvas, columnComposition, "   Input:", "Placeholder", KL, 200, 20, 0, 0, 0, 10, colours.WHITE, "luxi", "base/components/fonts/luxi.ttf", 16)

	var Row1 = ompo.NewRowComposition("row1", A.WindowCanvas, columnComposition, 0, 0, 0, cons.ALIGN_CENTRE_ROW)
	Row1.AddDisplayable(comp.NewCheckbox(A.WindowCanvas, Row1, "cb_1", 10, 0, 0, 0, &event.NULL_EVENT, "luxi", "base/components/fonts/luxi.ttf", 16))
	Row1.AddDisplayable(comp.NewCheckbox(A.WindowCanvas, Row1, "cb_2", 10, 0, 0, 0, &event.NULL_EVENT, "luxi", "base/components/fonts/luxi.ttf", 16))
	Row1.AddDisplayable(comp.NewCheckbox(A.WindowCanvas, Row1, "cb_3", 10, 0, 0, 0, &event.NULL_EVENT, "luxi", "base/components/fonts/luxi.ttf", 16))
	
	var Row2 = ompo.NewRowComposition("row2", A.WindowCanvas, columnComposition, 0, 0, 0, cons.ALIGN_CENTRE_ROW)
	Row2.AddDisplayable(comp.NewButton(A.WindowCanvas, Row2, "button_1", 100, 30, 5, 0, 0, 0, "luxi", "base/components/fonts/luxi.ttf", 16, &event.NULL_EVENT, 200))
	Row2.AddDisplayable(comp.NewButton(A.WindowCanvas, Row2, "button_2", 100, 30, 5, 0, 0, 0, "luxi", "base/components/fonts/luxi.ttf", 16, &event.NULL_EVENT, 200))
	Row2.AddDisplayable(comp.NewButton(A.WindowCanvas, Row2, "button_3", 100, 30, 5, 0, 0, 0, "luxi", "base/components/fonts/luxi.ttf", 16, &event.NULL_EVENT, 200))
	
	var Label1 = comp.NewLabel(A.WindowCanvas, columnComposition, "Label1   ", 0, 0, 0, "luxi", "base/components/fonts/luxi.ttf", 16)

	var ToggleSwitch1 = comp.NewToggle(A.WindowCanvas, columnComposition, "Toggle:", 50, 30, 0, 0, 0, "luxi", "base/components/fonts/luxi.ttf", 16, &event.NULL_EVENT)

	columnComposition.AddDisplayable(Checkbox1)
	columnComposition.AddDisplayable(Box1)
	columnComposition.AddDisplayable(Input1)
	columnComposition.AddDisplayable(Label1)
	columnComposition.AddDisplayable(ToggleSwitch1)
	columnComposition.AddDisplayable(Row2)
	columnComposition.AddDisplayable(Row1)

	var alignments = []int{
 		cons.ALIGN_TOP_LEFT,
	}
	var stack = ompo.NewStackComposition(
		"Stack", 
		A.WindowCanvas, A.WindowCanvas, 
		0, 0, 0, 1.0, 1.0, alignments)
	
	stack.AddDisplayable(columnComposition)
	A.GetWindowCanvas().AddDisplayable(stack)
	A.RunWindow()
}


