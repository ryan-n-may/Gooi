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
	var A = windows.NewWindow("Application Demo", 900, 600)
	A.OpenWindow()
	A.SetTheme(colours.GrantiteTheme())

/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)

	/** Application Title Bar **/
	var TitleBar_Label = comp.CreateThemedLabel("System Settings", A.WindowCanvas, 0, 0, 32, colours.GrantiteTheme())
	/** Sidebar Buttons **/
	var Operations_Label = comp.CreateThemedLabel("Operations", A.WindowCanvas, 0, 0, 16, colours.GrantiteTheme())
	var Reset_Button = comp.CreateThemedButton("Reset to Defaults", A.WindowCanvas, 150, 25, 0, 0, 16, colours.GrantiteTheme(), &event.NULL_EVENT)
	var Apply_Button = comp.CreateThemedButton("Apply Settings", A.WindowCanvas, 150, 25, 0, 0, 16, colours.GrantiteTheme(), &event.NULL_EVENT)
	var Save_Button = comp.CreateThemedButton("Save Settings", A.WindowCanvas, 150, 25, 0, 0, 16, colours.GrantiteTheme(), &event.NULL_EVENT)
	/** Main Tabs Page **/
	var App_Tabs = ompo.NewTabs("Tabs", A.WindowCanvas, E, ML, []string{"Personalisation", "Display", "About"}, 0.0, 0.0, A.GetWindowWidth(), A.GetWindowHeight())

	var Toggle1 = comp.CreateThemedToggle("Toggle1", A.WindowCanvas, 40, 30, 0, 0, 16, &event.NULL_EVENT, colours.GrantiteTheme())
	var Toggle2 = comp.CreateThemedToggle("Toggle2", A.WindowCanvas, 40, 30, 0, 0, 16, &event.NULL_EVENT, colours.GrantiteTheme())
	var Toggle3 = comp.CreateThemedToggle("Toggle3", A.WindowCanvas, 40, 30, 0, 0, 16, &event.NULL_EVENT, colours.GrantiteTheme())

	/**
	 * Compositions
	 **/
	var TitleBoxComposition = ompo.NewBoxComposition("Label Box Composition", 0, 0, A.GetWindowWidth(), A.GetWindowHeight(), cons.ALIGN_TOP_CENTRE)
	TitleBoxComposition.AddDrawable(TitleBar_Label)

	var LeftSideComposition = ompo.NewColumnComposition("Left Side Composition", 0, 0, 10, cons.ALIGN_LEFT)
	LeftSideComposition.AddDrawable(Reset_Button)
	LeftSideComposition.AddDrawable(Apply_Button)
	LeftSideComposition.AddDrawable(Save_Button)
	LeftSideComposition.AddDrawable(Operations_Label)

	var Personalisation_Column = ompo.NewColumnComposition("Personalisation", 0, 0, 10, cons.ALIGN_LEFT)
	Personalisation_Column.AddDrawable(Toggle1)

	var Display_Column = ompo.NewColumnComposition("Personalisation", 0, 0, 10, cons.ALIGN_LEFT)
	Display_Column.AddDrawable(Toggle2)

	var About_Column = ompo.NewColumnComposition("Personalisation", 0, 0, 10, cons.ALIGN_LEFT)
	About_Column.AddDrawable(Toggle3)

	App_Tabs.AddDrawable(Personalisation_Column, 0)
	App_Tabs.AddDrawable(Display_Column, 1)
	App_Tabs.AddDrawable(About_Column, 2)

	var VerticalDivision = ompo.NewVerticalDivision("Vertical division 1", []float32{0.25, 0.75}, 0, 0, A.GetWindowWidth(), A.GetWindowHeight())
	VerticalDivision.AddDrawable(LeftSideComposition, 0)
	VerticalDivision.AddDrawable(App_Tabs, 1)

	var HorisontalDivision = ompo.NewHorisontalDivision("Horisontal division 1", []float32{0.90, 0.10}, 0, 0, A.GetWindowWidth(), A.GetWindowHeight())
	HorisontalDivision.AddDrawable(VerticalDivision, 0)
	HorisontalDivision.AddDrawable(TitleBoxComposition, 1)


	A.GetWindowCanvas().AddComponent(HorisontalDivision)

/** Telling the mouse handler that the button components are clickable **/
	
	A.GetMouseHandler().RegisterClickableToHandler(Toggle1)
	A.GetMouseHandler().RegisterClickableToHandler(Toggle2)
	A.GetMouseHandler().RegisterClickableToHandler(Toggle3)
	A.GetMouseHandler().RegisterClickableToHandler(Reset_Button)
	A.GetMouseHandler().RegisterClickableToHandler(Apply_Button)
	A.GetMouseHandler().RegisterClickableToHandler(Save_Button)

	A.RunWindow()
}
