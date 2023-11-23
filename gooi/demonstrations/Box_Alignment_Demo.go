package main

import (
	event 		"GUI/Base/Event"
	comp 		"GUI/Base/Components"
	listeners 	"GUI/Base/Listeners"
	windows 	"GUI/Base/Windows"
	ompo   		"GUI/Base/Compositions"
	cons        "GUI/Base/Constants"
	colours     "GUI/Base/Colours"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Box Alignment Demo", 1000, 800)
	A.OpenWindow()
	A.SetBackgroundColour(colours.LIGHT_GRAY)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating button 1 **/
	var window_w = A.GetWindowWidth()
	var window_h = A.GetWindowHeight()


	
	// creating the button and assinging it the visual characteristics + the trigger event
	var Button_Centre = comp.CreateButton(
		"ALIGN_CENTRE",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Centre_Composition = ompo.NewBoxComposition("Top Left Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_CENTRE)
	Centre_Composition.AddDrawable(Button_Centre)

	var Button_Top_Left = comp.CreateButton(
		"ALIGN_TOP_LEFT",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Top_Left_Composition = ompo.NewBoxComposition("Top Left Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_TOP_LEFT)
	Top_Left_Composition.AddDrawable(Button_Top_Left)

	var Button_Top_Right = comp.CreateButton(
		"ALIGN_TOP_RIGHT",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Top_Right_Composition = ompo.NewBoxComposition("Top Left Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_TOP_RIGHT)
	Top_Right_Composition.AddDrawable(Button_Top_Right)

	var Button_Bottom_Right = comp.CreateButton(
		"ALIGN_BOTTOM_RIGHT",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Bottom_Right_Composition = ompo.NewBoxComposition("Bottom Right Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_BOTTOM_RIGHT)
	Bottom_Right_Composition.AddDrawable(Button_Bottom_Right)

	var Button_Bottom_Left = comp.CreateButton(
		"ALIGN_BOTTOM_LEFT",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Bottom_Left_Composition = ompo.NewBoxComposition("Bottom Left Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_BOTTOM_LEFT)
	Bottom_Left_Composition.AddDrawable(Button_Bottom_Left)

	var Button_Bottom_Centre = comp.CreateButton(
		"ALIGN_BOTTOM_CENTRE",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Bottom_Centre_Composition = ompo.NewBoxComposition("Bottom Centre Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_BOTTOM_CENTRE)
	Bottom_Centre_Composition.AddDrawable(Button_Bottom_Centre)

	var Button_Top_Centre = comp.CreateButton(
		"ALIGN_TOP_CENTRE",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Top_Centre_Composition = ompo.NewBoxComposition("Top Centre Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_TOP_CENTRE)
	Top_Centre_Composition.AddDrawable(Button_Top_Centre)

	var Button_Centre_Left = comp.CreateButton(
		"ALIGN_CENTRE_LEFT",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Centre_Left_Composition = ompo.NewBoxComposition("Centre Left Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_CENTRE_LEFT)
	Centre_Left_Composition.AddDrawable(Button_Centre_Left)

	var Button_Centre_Right = comp.CreateButton(
		"ALIGN_CENTRE_RIGHT",
		A.WindowCanvas, 
		200, 50, 300, 50, 
		20,
		2,
		2,
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
	var Centre_Right_Composition = ompo.NewBoxComposition("Centre Right Composition", 
						0, 0, 
						window_w, 
						window_h, 
						cons.ALIGN_CENTRE_RIGHT)
	Centre_Right_Composition.AddDrawable(Button_Centre_Right)

	A.GetWindowCanvas().AddComponent(Centre_Composition)
	A.GetWindowCanvas().AddComponent(Top_Left_Composition)
	A.GetWindowCanvas().AddComponent(Top_Right_Composition)
	A.GetWindowCanvas().AddComponent(Bottom_Right_Composition)
	A.GetWindowCanvas().AddComponent(Bottom_Left_Composition)
	A.GetWindowCanvas().AddComponent(Bottom_Centre_Composition)
	A.GetWindowCanvas().AddComponent(Top_Centre_Composition)
	A.GetWindowCanvas().AddComponent(Centre_Left_Composition)
	A.GetWindowCanvas().AddComponent(Centre_Right_Composition)

/** Telling the mouse handler that the button components are clickable **/
	//A.GetMouseHandler().RegisterClickableToHandler(component1)
	
	A.RunWindow()
}


