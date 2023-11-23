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
	var A = windows.NewWindow("Application Demo", 1000, 600)
	A.OpenWindow()
	A.SetTheme(colours.GrantiteTheme())

/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)

	/** 
	 * Components 
	 **/
	var TitleBar_Label = comp.CreateThemedLabel("Application Title", A.WindowCanvas, 0, 0, 32, colours.GrantiteTheme())

	var LeftBanner_L1 = comp.CreateThemedLabel("Functions", A.WindowCanvas, 0, 0, 16, colours.GrantiteTheme())
	var LeftBanner_B1 = comp.CreateThemedButton("Lorem ipsum", A.WindowCanvas, 150, 25, 0, 0, 8, colours.GrantiteTheme(), &event.NULL_EVENT)
	var LeftBanner_B2 = comp.CreateThemedButton("Consectetur", A.WindowCanvas, 150, 25, 0, 0, 8, colours.GrantiteTheme(), &event.NULL_EVENT)
	var LeftBanner_B3 = comp.CreateThemedButton("Exercitation", A.WindowCanvas, 150, 25, 0, 0, 8, colours.GrantiteTheme(), &event.NULL_EVENT)
	var LeftBanner_B4 = comp.CreateThemedButton("Duis aute", A.WindowCanvas, 100, 40, 0, 0, 8, colours.GrantiteTheme(), &event.NULL_EVENT)
	var LeftBanner_B5 = comp.CreateThemedButton("Excepteur", A.WindowCanvas, 100, 40, 0, 0, 8, colours.GrantiteTheme(), &event.NULL_EVENT)
	var LeftBanner_L2 = comp.CreateThemedLabel("Settings", A.WindowCanvas, 0, 0, 16, colours.GrantiteTheme())
	
	var LeftBanner_T1_L = comp.CreateThemedLabel("Culpa", A.WindowCanvas, 0, 0, 8, colours.GrantiteTheme())
	var LeftBanner_T1 = comp.CreateThemedToggle("Left_Toggle1", A.WindowCanvas, 40, 30, 0, 0, &event.NULL_EVENT, colours.GrantiteTheme())

	var LeftBanner_T2_L = comp.CreateThemedLabel("Anim", A.WindowCanvas, 0, 0, 8, colours.GrantiteTheme())
	var LeftBanner_T2 = comp.CreateThemedToggle("Left_Toggle2", A.WindowCanvas, 40, 30, 0, 0, &event.NULL_EVENT, colours.GrantiteTheme())

	var LeftBanner_T3_L = comp.CreateThemedLabel("Proident", A.WindowCanvas, 0, 0, 8, colours.GrantiteTheme())
	var LeftBanner_T3 = comp.CreateThemedToggle("Left_Toggle3", A.WindowCanvas, 40, 30, 0, 0, &event.NULL_EVENT, colours.GrantiteTheme())

	var RightBanner_L1 = comp.CreateThemedLabel("Culpa anim prodient:", A.WindowCanvas, 0, 0, 16, colours.GrantiteTheme())
	var RightBanner_B1 = comp.CreateThemedButton("Lorem ipsum dolor sit amet", A.WindowCanvas, 250, 30, 0, 0, 16, colours.GrantiteTheme(), &event.NULL_EVENT)
	var RightBanner_B2 = comp.CreateThemedButton("Excepteur sint occaecat", A.WindowCanvas, 250, 30, 0, 0, 16, colours.GrantiteTheme(), &event.NULL_EVENT)
	var RightBanner_R1 = ompo.CreateThemedRadioButton("Right Radio button", A.WindowCanvas, 10, 0, 0, 6, []string{"Flamingo", "Goose", "Chicken", "Camel", "Sheep", "Dog"}, cons.VERTICAL_ORIENT, 16, colours.GrantiteTheme())


	/**
	 * Compositions
	 **/
	var Label_Box_Composition = ompo.NewBoxComposition("Label Box Composition", 0, 0, A.GetWindowWidth(), A.GetWindowHeight(), cons.ALIGN_TOP_CENTRE)
	Label_Box_Composition.AddDrawable(TitleBar_Label)

	var T1_composition = ompo.NewRowComposition("Toggle 1", 0, 0, 10, cons.ALIGN_CENTRE_ROW)
	T1_composition.AddDrawable(LeftBanner_T1)
	T1_composition.AddDrawable(LeftBanner_T1_L)

	var T2_composition = ompo.NewRowComposition("Toggle 2", 0, 0, 10, cons.ALIGN_CENTRE_ROW)
	T2_composition.AddDrawable(LeftBanner_T2)
	T2_composition.AddDrawable(LeftBanner_T2_L)

	var T3_composition = ompo.NewRowComposition("Toggle 3", 0, 0, 10, cons.ALIGN_CENTRE_ROW)
	T3_composition.AddDrawable(LeftBanner_T3)
	T3_composition.AddDrawable(LeftBanner_T3_L)

	var LeftBanner_Column_Composition = ompo.NewColumnComposition("Left Banner Composition", 0, 0, 10, cons.ALIGN_LEFT)
	LeftBanner_Column_Composition.AddDrawable(T3_composition)
	LeftBanner_Column_Composition.AddDrawable(T2_composition)
	LeftBanner_Column_Composition.AddDrawable(T1_composition)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_L2)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_B5)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_B4)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_B3)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_B2)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_B1)
	LeftBanner_Column_Composition.AddDrawable(LeftBanner_L1)

	var RightBanner_Button_Row_Composition = ompo.NewRowComposition("Right Banner row composition", 0, 0, 10, cons.ALIGN_BOTTOM)
	RightBanner_Button_Row_Composition.AddDrawable(RightBanner_B1)
	RightBanner_Button_Row_Composition.AddDrawable(RightBanner_B2)

	var RigthBanner_Column_Composition = ompo.NewColumnComposition("Right Banner Compsoition", 0, 0, 10, cons.ALIGN_CENTRE_COLUMN)
	RigthBanner_Column_Composition.AddDrawable(RightBanner_R1)
	RigthBanner_Column_Composition.AddDrawable(RightBanner_Button_Row_Composition)
	RigthBanner_Column_Composition.AddDrawable(RightBanner_L1)


	var vertical_div_composition = ompo.NewVerticalDivision("Vertical division 1", []float32{0.25, 0.75}, 0, 0, A.GetWindowWidth(), A.GetWindowHeight())
	vertical_div_composition.AddDrawable(LeftBanner_Column_Composition, 0)
	vertical_div_composition.AddDrawable(RigthBanner_Column_Composition, 1)

	var horisontal_div_composition = ompo.NewHorisontalDivision("Horisontal division 1", []float32{0.90, 0.10}, 0, 0, A.GetWindowWidth(), A.GetWindowHeight())
	horisontal_div_composition.AddDrawable(vertical_div_composition, 0)
	horisontal_div_composition.AddDrawable(Label_Box_Composition, 1)
	

	/*var master_box_composition = ompo.NewBoxComposition("Master composition", 
						0, 0, 
						A.GetWindowWidth(), 
						A.GetWindowHeight(), 
						cons.ALIGN_CENTRE)
	master_box_composition.AddDrawable(horisontal_div_composition)*/

	A.GetWindowCanvas().AddComponent(horisontal_div_composition)

/** Telling the mouse handler that the button components are clickable **/
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_B1)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_B2)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_B3)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_B4)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_B5)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_T1)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_T2)
	A.GetMouseHandler().RegisterClickableToHandler(LeftBanner_T3)

	A.RunWindow()
}
