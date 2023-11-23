package compositions
import (
	comp 	"gooi/base/components"
	cons 	"gooi/base/constants"
	event 	"gooi/base/event"
	intf 	"gooi/interfaces"
	colours "gooi/base/colours"
)
type RadioButton_Struct struct {
	RadioButtonName string
	Buttons 		[]*comp.CheckBox_Struct
	ButtonCount     int
	SelectedIndex 	int
	Pos_x 			float32
	Pos_y    		float32
	Radius   		float32
	Padding   		float32
	RadioCanvas   	*comp.Canvas_Struct
	CheckBodyColour_Empty 		[3]float32
	CheckBodyColour_Filled 		[3]float32
	CheckBodyColour_Deactivated [3]float32
	CheckBorderColour 			[3]float32
	CheckBackgroundColour 		[3]float32
	FontPath 		string
	FontName  		string
	FontSize       	int
	ChangeEvent 	*event.Event_Struct
	Orientation 	int
	Labels 			[]string
	Master_Height 	float32
	Master_Width 	float32
}
func CreateThemedRadioButton(
	radioButtonName 				string,
	canvas 							*comp.Canvas_Struct,
	radius, pos_x, pos_y 			float32, 
	buttonCount 					int,
	labels  						[]string,
	orientation 					int,
	font_size  						int,
	theme 							*colours.Theme,
) *RadioButton_Struct {
	var radiobutton = RadioButton_Struct{}
	// Specified
	radiobutton.RadioButtonName = radioButtonName
	radiobutton.Master_Height = 0
	radiobutton.Master_Width = 0
	radiobutton.Pos_x = pos_x
	radiobutton.Pos_y = pos_y
	radiobutton.Radius = radius
	radiobutton.RadioCanvas = canvas 
	radiobutton.ButtonCount = buttonCount
	radiobutton.Buttons = make([]*comp.CheckBox_Struct, buttonCount)
	radiobutton.Labels = labels
	radiobutton.Orientation = orientation
	radiobutton.FontSize = font_size
	var event_arguments = event.NewEventParameter(&radiobutton)
	var radio_button_event = &event.Event_Struct{
		radioButtonClicked, 
		"radiobuttonclicked",
		event_arguments,
	}
	radiobutton.ChangeEvent = radio_button_event
	radiobutton.RadioCanvas.GetEventHandler().RegisterEventToHandler(radiobutton.ChangeEvent)
	// Themed
	radiobutton.Padding = theme.RadioButtonPadding
	radiobutton.CheckBodyColour_Filled = theme.CheckboxColour_Filled
	radiobutton.CheckBodyColour_Empty = theme.CheckboxColour_Empty
	radiobutton.CheckBodyColour_Deactivated = theme.CheckboxColour_Empty
	radiobutton.CheckBorderColour = theme.CheckBoxColourBorder
	radiobutton.CheckBackgroundColour = theme.CheckboxColourBackground
	radiobutton.FontName = theme.FontName
	radiobutton.FontPath = theme.FontPath
	radiobutton.CreateRadioButtons()
	return &radiobutton
}
func CreateRadioButton(
	radioButtonName 				string,
	canvas 							*comp.Canvas_Struct,
	radius, pos_x, pos_y 			float32, 
	padding 						float32,
	buttonCount 					int,
	orientation 					int,
	font_path string,
	font_name string,
	font_size int, 
	labels  						[]string,
	check_body_colour_filled		[3]float32,
	check_body_colour_empty 		[3]float32,
	check_body_colour_deactivated	[3]float32,
	check_border_colour	 			[3]float32,
	check_background_colour 		[3]float32,
) *RadioButton_Struct {
	var radiobutton = RadioButton_Struct{}
	radiobutton.RadioButtonName = radioButtonName
	radiobutton.Master_Height = 0
	radiobutton.Master_Width = 0
	radiobutton.Pos_x = pos_x
	radiobutton.Pos_y = pos_y
	radiobutton.Radius = radius
	radiobutton.Padding = padding
	radiobutton.RadioCanvas = canvas 
	radiobutton.CheckBodyColour_Filled = check_body_colour_filled
	radiobutton.CheckBodyColour_Empty = check_body_colour_empty
	radiobutton.CheckBodyColour_Deactivated = check_body_colour_deactivated
	radiobutton.CheckBorderColour = check_border_colour
	radiobutton.CheckBackgroundColour = check_background_colour
	radiobutton.ButtonCount = buttonCount
	radiobutton.Buttons = make([]*comp.CheckBox_Struct, buttonCount)
	radiobutton.FontName = font_name
	radiobutton.FontPath = font_path
	radiobutton.FontSize = font_size
	radiobutton.Labels = labels
	radiobutton.Orientation = orientation
	var event_arguments = event.NewEventParameter(&radiobutton)
	var radio_button_event = &event.Event_Struct{
		radioButtonClicked, 
		"radiobuttonclicked",
		event_arguments,
	}
	radiobutton.ChangeEvent = radio_button_event
	radiobutton.RadioCanvas.GetEventHandler().RegisterEventToHandler(radiobutton.ChangeEvent)
	radiobutton.CreateRadioButtons()
	return &radiobutton
}
func (radiobutton *RadioButton_Struct) CreateRadioButtons(){
	var button_pos_x float32 = 0
	var button_pos_y float32 = 0
	for i := 0; i < radiobutton.ButtonCount; i++ {
		radiobutton.Buttons[i] = comp.CreateCheckbox(
				radiobutton.Labels[i],
				radiobutton.RadioCanvas, 
				radiobutton.Radius, 0, 0, radiobutton.Padding,
				radiobutton.ChangeEvent,
				radiobutton.FontPath,
				radiobutton.FontName,
				radiobutton.FontSize,
				radiobutton.CheckBodyColour_Filled,
				radiobutton.CheckBodyColour_Empty,
				radiobutton.CheckBodyColour_Deactivated,
				radiobutton.CheckBorderColour,
				radiobutton.CheckBackgroundColour,
				)
		if radiobutton.Orientation == cons.HORISONT_ORIENT {
			button_pos_x = radiobutton.Pos_x + radiobutton.Buttons[i].CheckBox_Text.Text.Width()*float32(i) + radiobutton.Radius*4*float32(i)
			button_pos_y = radiobutton.Pos_y 
		} else if radiobutton.Orientation ==  cons.VERTICAL_ORIENT{
			button_pos_x = radiobutton.Pos_x  
			button_pos_y = radiobutton.Pos_y + radiobutton.Radius*3*float32(i)
		}
		radiobutton.Buttons[i].SetPos(button_pos_x, button_pos_y)
		radiobutton.Buttons[i].GetCanvas().GetWindow().GetMouseHandler().RegisterClickableToHandler(radiobutton.Buttons[i])
	}
	radiobutton.SelectedIndex = 0
	radiobutton.Buttons[radiobutton.SelectedIndex].ToggleFilledState(true)
}
func (radiobutton *RadioButton_Struct) MoveRadioButtons(){
	var button_pos_x float32 = radiobutton.Pos_x
	var button_pos_y float32 = radiobutton.Pos_y
	for i := 0; i < radiobutton.ButtonCount; i++ {
		if radiobutton.Orientation == cons.HORISONT_ORIENT {
			button_pos_x = radiobutton.Pos_x + radiobutton.Buttons[i].CheckBox_Text.Text.Width()*float32(i) + radiobutton.Radius*4*float32(i)
			button_pos_y = radiobutton.Pos_y 
		} else if radiobutton.Orientation == cons.VERTICAL_ORIENT{
			button_pos_x = radiobutton.Pos_x  
			button_pos_y = radiobutton.Pos_y + radiobutton.Radius*3*float32(i)
		}
		radiobutton.Buttons[i].SetPos(button_pos_x, button_pos_y)
	}
}
func (b *RadioButton_Struct) Draw() {
	for _, button := range b.Buttons {
		button.Draw()
	}
}
func (b *RadioButton_Struct) Redraw() {
	for _, button := range b.Buttons {
		button.Redraw()
	}
}
func (b *RadioButton_Struct) GetSelectedIndex() int {
	return b.SelectedIndex
}
func radioButtonClicked(param intf.Paramaters_Interface){
	var radiobutton = param.GetParameters().(*RadioButton_Struct)
	var newIndex = 0
	for index, button := range radiobutton.Buttons {
		if button.FilledState && index != radiobutton.SelectedIndex{
			newIndex = index
			radiobutton.Buttons[index].ToggleFilledState(true)
		} else if button.FilledState && index == radiobutton.SelectedIndex{
			radiobutton.Buttons[index].ToggleFilledState(false)
		}
	}
	radiobutton.SelectedIndex = newIndex
	radiobutton.Draw()
}
func (r *RadioButton_Struct) SetPos(x, y float32){
	r.Pos_x = x
	r.Pos_y = y
	r.MoveRadioButtons()
	r.Draw()
}
func (r *RadioButton_Struct) GetPos() (float32, float32){
	return r.Pos_x, r.Pos_y
}
func (r *RadioButton_Struct) GetBounds() (float32, float32){
	if r.Orientation == cons.VERTICAL_ORIENT {
		var max_width float32 = 0
		var sum_height float32 = 0
		for _, c := range r.Buttons {
			var w, h = c.GetBounds()
			if w > max_width {
				max_width = w
			} 
			sum_height += h 
			sum_height += c.Radius
		}
		r.Master_Width = max_width
		r.Master_Height = sum_height
	} else {
		var sum_width float32 = 0
		var max_height float32 = 0
		for _, c := range r.Buttons {
			var w, h = c.GetBounds()
			if h > max_height {
				max_height = h
			} 
			sum_width += w
		}
		r.Master_Width = sum_width
		r.Master_Height = max_height
	}
	return r.Master_Width, r.Master_Height
}

func (b *RadioButton_Struct) SetPosZ(z float32) {
	for i := 0 ; i < len(b.Buttons); i++ {
		b.Buttons[i].SetPosZ(z)
	}
}
func (b *RadioButton_Struct) GetPosZ() float32 {
	return b.Buttons[0].GetPosZ()
}