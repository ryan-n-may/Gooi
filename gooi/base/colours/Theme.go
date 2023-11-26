package colours

import (
	time "time"
)

// Create GraniteTheme
// ... OceanTheme
// ... BarbieTheme
// ... FrogTheme

type Theme struct {
		// The Window
		WindowBackground 			[3]float32
		// The Button 
		ButtonBackgroundColour		[3]float32
		ButtonBodyColour_Idle 		[3]float32
		ButtonBodyColour_Selected 	[3]float32
		ButtonBodyColour_Disabled 	[3]float32
		ButtonBorderColour 			[3]float32
		ButtonRadius 				float32
		ButtonBorderPadding 		float32
		ButtonBorderThickness 		float32
		ButtonAnimationTime 		time.Duration
		// The Font
		FontPath 					string
		FontName 					string
		FontColour 					[3]float32 // not implemented yet
		// font size is not theme dictated; set manually.
		// The Text field
		TextFieldBackground_Idle    [3]float32
		TextFieldBackground_Selected [3]float32
		TextFieldBorder 			[3]float32
		TextFieldHighlightColour	[3]float32
		TextFieldCursorColour 		[3]float32
		TextFieldRadius 			float32
		TextFieldPadding 			float32
		// Checkbox
		CheckboxColourBackground 	[3]float32
		CheckboxColour_Filled 		[3]float32
		CheckboxColour_Empty 		[3]float32
		CheckBoxColourBorder 		[3]float32
		// Toggle switch 
		ToggleBorderColour 			[3]float32
		ToggleBackgroundColour		[3]float32
		ToggleBackgroundColour_On	[3]float32
		ToggleBackgroundColour_Off 	[3]float32
		ToggleBackgroundColour_Deactivated	[3]float32
		ToggleCircleColour 			[3]float32
		ToggleBorderPadding 		float32
		ToggleAnimationTime 		time.Duration
		// Radiobutton
		RadioButtonPadding 			float32
}

func GrantiteTheme() *Theme {
	var g = Theme{}
	// the window
	g.WindowBackground = [3]float32{50, 50, 50} // dark gray
	// the button
	g.ButtonBackgroundColour = [3]float32{50, 50, 50} // dark gray
	g.ButtonBodyColour_Idle = [3]float32{110, 110, 110} // gray
	g.ButtonBodyColour_Selected = [3]float32{150, 150, 150} // light gray
	g.ButtonBodyColour_Disabled = [3]float32{50, 50, 50} // dark gray
	g.ButtonBorderColour = [3]float32{110, 110, 110}
	g.ButtonRadius = 10
	g.ButtonBorderPadding = 2
	g.ButtonBorderThickness = 2
	g.ButtonAnimationTime = 200
	// Font 
	g.FontPath = "base/components/fonts/luxi.ttf"
	g.FontName = "luxi"
	g.FontColour = [3]float32{210, 210, 210} // very light gray
	// Text field
	g.TextFieldBackground_Idle = [3]float32{50, 50, 50} // dark gray
	g.TextFieldBackground_Selected = [3]float32{110, 110, 110} // gray
	g.TextFieldBorder = [3]float32{110, 110, 110}
	g.TextFieldCursorColour = [3]float32{80, 100, 200} // blue
	g.TextFieldHighlightColour = [3]float32{80, 100, 200} 
	g.TextFieldRadius = 5.0
	g.TextFieldPadding = 2.0
	// Checkbox
	g.CheckboxColourBackground = [3]float32{150, 150, 150} // light gray
	g.CheckboxColour_Filled = [3]float32{80, 100, 200}
	g.CheckboxColour_Empty = [3]float32{150, 150, 150} // light gray
	g.CheckBoxColourBorder = [3]float32{210, 210, 210} // very light gray
	// Toggle Switch
	g.ToggleBorderColour = [3]float32{110, 110, 110}
	g.ToggleBackgroundColour = [3]float32{50, 50, 50}
	g.ToggleBackgroundColour_On = [3]float32{50, 250, 50}
	g.ToggleBackgroundColour_Off = [3]float32{250, 50, 50}
	g.ToggleBackgroundColour_Deactivated = [3]float32{150, 150, 150}
	g.ToggleCircleColour = [3]float32{150, 150, 150} // light gray
	g.ToggleBorderPadding = 4
	g.ToggleAnimationTime = 20
	// Radiobutton
	g.RadioButtonPadding = 2
	return &g
}