package main

import (
	event 		"gooi/base/event"
	comp 		"gooi/base/components"
	listeners 	"gooi/base/listeners"
	windows 	"gooi/base/windows"
	ompo   		"gooi/base/compositions"
	colours     "gooi/base/colours"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Tabs Alignment Demo", 1000, 800)
	A.OpenWindow()
	A.SetBackgroundColour(colours.LIGHT_GRAY)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	ML.SetCurrentZLayer(0.0)
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating button 1 **/
	var window_w = A.GetWindowWidth()
	var window_h = A.GetWindowHeight()


	var Button1 = comp.CreateButton(
		"B1",
		A.WindowCanvas, 
		200, 20, 300, 50, 
		20,
		2,
		2,
		250,
		&event.NULL_EVENT,
		"base/components/fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Label = comp.CreateLabel(
		"Label Component",
		A.WindowCanvas, 
		140.0, 30.0, 
		"base/components/fonts/luxi.ttf",
		"luxi", 
		32,
	)

	var tabs = ompo.NewTabs("Tabs", A.WindowCanvas, E, ML, []string{"Settings", "Display"}, 0.0, 0.0, window_w, window_h)

	tabs.AddDrawable(Button1, 1)
	tabs.AddDrawable(Label, 0)
	
	A.GetWindowCanvas().AddComponent(tabs)

/** Telling the mouse handler that the button components are clickable **/
	//A.GetMouseHandler().RegisterClickableToHandler(component1)
	
	A.RunWindow()
}


