package main

import (
	event 		"gooi/base/event"
	comp 		"gooi/base/components"
	listeners 	"gooi/base/listeners"
	windows 	"gooi/base/windows"
	ompo   		"gooi/base/compositions"
	colours     "gooi/base/colours"
	cons 		"gooi/base/constants"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Tabs Alignment Demo", 800, 400)
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


	var Button0 = comp.CreateButton(
		"B0",
		A.WindowCanvas, 
		200, 50, 300, 50, 
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


	var Button1 = comp.CreateButton(
		"B1",
		A.WindowCanvas, 
		200, 50, 300, 50, 
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

	var Button2 = comp.CreateButton(
		"B2",
		A.WindowCanvas, 
		200, 50, 300, 50, 
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

	
	var tabs = ompo.NewTabs("Tabs", A.WindowCanvas, E, ML, []string{"Settings", "Display"}, 0.0, 0.0, A.GetWindowWidth(), A.GetWindowHeight())

	tabs.AddDrawable(Button1, 1)
	tabs.AddDrawable(Button2, 0)
	
	var hori_div_composition = ompo.NewHorisontalDivision("Horisontal division 1", []float32{0.80, 0.20}, 0, 0, A.GetWindowWidth(), A.GetWindowHeight())
	hori_div_composition.AddDrawable(Button0, 1)
	hori_div_composition.AddDrawable(tabs, 0)

	var centred_box = ompo.NewBoxComposition("Master Box", 0, 0, A.GetWindowWidth(), A.GetWindowHeight(), cons.ALIGN_CENTRE)
	centred_box.AddDrawable(hori_div_composition)

	A.GetWindowCanvas().AddComponent(centred_box)

/** Telling the mouse handler that the button components are clickable **/
	//A.GetMouseHandler().RegisterClickableToHandler(component1)
	
	A.RunWindow()
}


