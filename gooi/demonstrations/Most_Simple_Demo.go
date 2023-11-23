package main

import (
	event 		"GUI/Base/Event"
	listeners 	"GUI/Base/Listeners"
	windows 	"GUI/Base/Windows"
	colours     "GUI/Base/Colours"
	comp 		"GUI/Base/Components"
	cons 		"GUI/Base/Constants"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Column Alignment Demo", 400, 200)
	A.OpenWindow()
	A.SetBackgroundColour(colours.LIGHT_GRAY)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
	var KeyListener = listeners.CreateKeyListener("KeyListener", A.GetWindowCanvas())
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating button 1 **/

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

	A.GetMouseHandler().RegisterClickableToHandler(text_input)
	A.GetWindowCanvas().AddComponent(text_input)
	// creating the button and assinging it the visual characteristics + the trigger event
	A.RunWindow()
}
