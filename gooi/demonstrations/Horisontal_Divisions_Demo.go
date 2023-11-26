package main

import (
	fmt  		"fmt"
	event 		"gooi/base/event"
	comp 		"gooi/base/components"
	listeners 	"gooi/base/listeners"
	windows 	"gooi/base/windows"
	intf 		"gooi/interfaces"
	ompo   		"gooi/base/compositions"
	cons        "gooi/base/constants"
	colours     "gooi/base/colours"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Widget Demo", 1000, 600)
	A.OpenWindow()
	A.SetBackgroundColour(colours.LIGHT_GRAY)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)

	var Button_1 = comp.CreateButton(
		"Button1",
		A.WindowCanvas, 
		100, 50, 300, 50, 
		20,
		2, 2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Button_2 = comp.CreateButton(
		"Button2",
		A.WindowCanvas, 
		100, 30, 300, 50, 
		10,
		2, 2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var Button_3 = comp.CreateButton(
		"Button3",
		A.WindowCanvas, 
		100, 30, 300, 50, 
		10,
		2, 2,
		250,
		&event.NULL_EVENT,
		"Base/Components/Fonts/luxi.ttf",
		"luxi", 
		16,
		colours.BLUE,
		colours.DARK_BLUE,
		colours.GRAY,
		colours.GRAY,
		colours.WHITE,
	)

	var hori_div_composition = ompo.NewHorisontalDivision("Horisontal division 1", []float32{0.25, 0.25, 0.50}, 0, 0, A.GetWindowWidth(), A.GetWindowHeight())
	hori_div_composition.AddDrawable(Button_1, 0)
	hori_div_composition.AddDrawable(Button_2, 1)
	hori_div_composition.AddDrawable(Button_3, 2)
	
	
	var master_box_composition = ompo.NewBoxComposition("Master composition", 
						0, 0, 
						A.GetWindowWidth(), 
						A.GetWindowHeight(), 
						cons.ALIGN_CENTRE)
	master_box_composition.AddDrawable(hori_div_composition)

	A.GetWindowCanvas().AddComponent(master_box_composition)

/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(Button_1)
	A.GetMouseHandler().RegisterClickableToHandler(Button_2)

	A.RunWindow()
}

func test(param intf.Paramaters_Interface){
	fmt.Println(param.GetParameters().([]string)[0], param.GetParameters().([]string)[1])
}

func disablebutton1(param intf.Paramaters_Interface){
	param.GetParameters().(intf.Clickable_Interface).SetClickable(false)
	fmt.Println(param.GetParameters().(intf.Component_Interface).GetName(), " set to no-clickable")
}
