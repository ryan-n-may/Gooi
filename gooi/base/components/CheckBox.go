package components 
/**
 * Updated 11/11/2023.
 * Checkbox Component.
 * Implements Drawable -> Component -> Clickable.
 **/
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	mgl32   "github.com/go-gl/mathgl/mgl32"
	intf 	"gooi/interfaces"
	event 	"gooi/base/event"
	cons 	"gooi/base/constants"
	log 	"log"
)
type CheckBox_Struct struct {
	err  			error
	CheckBoxName	string
	XYZ 			[]float32
	RGB 			[]float32 
	DIM 			[]float32
	VAO 			[]intf.Drawing_Struct
	Clickable 		bool
	ClickEvent 		intf.Event_Interface
	ClickTrigger 	int
	Bounds_X_max 	int 
	Bounds_X_min 	int
	Bounds_Y_max 	int
	Bounds_Y_min 	int
	Radius 			float32
	Pos_x, Pos_y 	float32
	Pos_z 			float32
	Padding			float32
	WindowHeight    		*float32
	WindowHeight_Initial 	float32
	WindowWidth   			*float32
	WindowWidth_Initial 	float32
	CheckBoxCanvas 	*Canvas_Struct
	CheckBodyColour 			[3]float32
	CheckBodyColour_Filled		[3]float32
	CheckBodyColour_Empty		[3]float32
	CheckBodyColour_Disabled 	[3]float32
	CheckBorderColour   		[3]float32
	CheckBackgroundColour   	[3]float32
	CheckBox_Text 	*Text_Struct
	CheckBox_Font 	*Font_Struct
	FilledState bool
}
func CreateCheckbox(
	name 					string,
	canvas 					*Canvas_Struct, 
	radius, pos_x, pos_y 	float32, 
	padding 				float32,
	check_event 			*event.Event_Struct,
	font_path 				string,
	font_name 				string,
	font_size 				int,
	check_body_colour_filled		[3]float32,
	check_body_colour_empty 		[3]float32,
	check_body_colour_deactivated	[3]float32,
	check_border_colour	 			[3]float32,
	check_background_colour 		[3]float32,
	) *CheckBox_Struct {
	log.Println("creating [Checkbox] struct.")
	// Creating buttons 
	var b = CheckBox_Struct{}
	b.CheckBoxName = name
	// Button specific characteristics 
	b.Radius = radius
	b.Pos_x = pos_x
	b.Pos_y = pos_y
	b.Pos_z = 0.0
	b.Padding = padding
	b.CheckBoxCanvas = canvas
	// Clickable interface charactistics
	b.ClickEvent = check_event
	b.Clickable = true
	b.ClickTrigger = cons.MOUSE_PRESSED
	// Window characteristics
 	b.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
	b.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
	b.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
	b.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()
	// Component colours
	b.CheckBodyColour = check_body_colour_empty
	b.CheckBodyColour_Filled = check_body_colour_filled
	b.CheckBodyColour_Empty = check_body_colour_empty
	b.CheckBodyColour_Disabled = check_body_colour_deactivated
	b.CheckBorderColour = check_border_colour
	b.CheckBackgroundColour = check_background_colour
	// Checkbox text
	b.CheckBox_Font = ReadFontFile(font_name, font_path, font_size)
	b.CheckBox_Font.Font.ResizeWindow(*b.WindowWidth, *b.WindowHeight)
	b.CheckBox_Text = CreateText(b.CheckBoxName, b.CheckBox_Font.Font)
	b.CheckBox_Text.Text.SetPosition(mgl32.Vec3{pos_x + b.CheckBox_Text.Text.Width()/2 + b.Radius*2, pos_y, 0.0})
	b.FilledState = false
	b.GeneratePolygons()
	b.SetClickableBounds(b.Pos_x-b.Radius, b.Pos_x+b.Radius, b.Pos_y-b.Radius, b.Pos_y+b.Radius)
   	return &b
}
func (b *CheckBox_Struct) GeneratePolygons(){
	b.VAO = make([]intf.Drawing_Struct, 3)
	// Border circle
	b.VAO[0] = intf.Drawing_Struct{intf.GenerateCirlce(b, b.CheckBorderColour, b.Pos_x, b.Pos_y, b.Pos_z, b.Radius, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	// Background circle
	b.VAO[1] = intf.Drawing_Struct{intf.GenerateCirlce(b, b.CheckBackgroundColour, b.Pos_x+b.Padding/16, b.Pos_y+b.Padding/16, b.Pos_z, b.Radius-b.Padding/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	// Background colour
	b.VAO[2] = intf.Drawing_Struct{intf.GenerateCirlce(b, b.CheckBodyColour, b.Pos_x+b.Padding/8, b.Pos_y+b.Padding/8, b.Pos_z, b.Radius-b.Padding, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }	
}
func (b *CheckBox_Struct) Draw() {
	gl.UseProgram(b.GetCanvas().GetPrograms())
	for _, v := range b.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(b.GetXYZ())/3))
	}
	// Draw text part of checkbox (using slightly modified gltext)
	Draw(b.CheckBox_Text)
}
func (b *CheckBox_Struct) Redraw() {
	b.GeneratePolygons()
	gl.UseProgram(b.GetCanvas().GetPrograms())
	for _, v := range b.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(b.GetXYZ())/3))
	}
	// Draw button text (Uses FontProg program)
	// Modified glText implementation:
	Draw(b.CheckBox_Text)
}
func (b *CheckBox_Struct) Move(delta_x, delta_y float32) {
	b.Pos_y = b.Pos_y + delta_y
	b.Pos_x = b.Pos_x + delta_x
	b.GeneratePolygons()
} 
func (b *CheckBox_Struct) ToggleFilledState(state bool){
	if state {
		b.CheckBodyColour = b.CheckBodyColour_Filled
	} else {
		b.CheckBodyColour = b.CheckBodyColour_Empty
	}
	b.FilledState = state
	b.GeneratePolygons()
}
func (b *CheckBox_Struct) GetFilledState() bool { return b.FilledState }
// Animation methods 
// Currently theres not animation- just a simple toggle.
func (b *CheckBox_Struct) Kill(alive *bool){ *alive = false }
func (b *CheckBox_Struct) StartClickAnimation(){
	if b.FilledState{ b.ToggleFilledState(false)
	} else { b.ToggleFilledState(true) }
	b.GeneratePolygons()
   	b.GetCanvas().RefreshCanvas()	
}
func (b *CheckBox_Struct) AnimateTrigger(pressed int, alive *bool){
	if b.Clickable == true && pressed == b.ClickTrigger{
		// alive isn't vital here as theres no animation right now.
		// keep alive present encase we animate in future. 
		*alive = true
		defer b.Kill(alive)
		b.StartClickAnimation()
	} 
}
// Triggers a click event if the checkbox is clickable, pressed is correct (MOUSE_PRESS, MOUSE_RELEASE)
// and the click event doesnt have a nil method.
func (b *CheckBox_Struct) TriggerClickEvent(pressed int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if b.ClickEvent.GetMethod() != nil && b.Clickable == true && pressed == b.ClickTrigger{
		b.GetCanvas().GetEventHandler().AddEventToEventQueue(b.ClickEvent.GetName())
  	}
}
/**
 * Other Setters and Getters
 **/
func (b *CheckBox_Struct) SetClickable(clickable bool){
	b.Clickable = clickable
	if !clickable{ b.CheckBodyColour = b.CheckBodyColour_Disabled
   	} else if b.FilledState { b.CheckBodyColour = b.CheckBodyColour_Filled 
   	} else if !b.FilledState { b.CheckBodyColour = b.CheckBodyColour_Empty } 
   	b.GeneratePolygons()
   	b.GetCanvas().RefreshCanvas()	  
}
func (b *CheckBox_Struct) GetClickable() bool { return b.Clickable }
func (b *CheckBox_Struct) SetPos(x, y float32) { 
	b.Pos_x = x + b.Radius 
	b.Pos_y = y
	b.CheckBox_Text.Text.SetPosition(mgl32.Vec3{b.Pos_x + b.CheckBox_Text.Text.Width()/2 + b.Radius*3, b.Pos_y, 0.0})
	b.GeneratePolygons()
	b.SetClickableBounds(b.Pos_x-b.Radius, b.Pos_x+b.Radius, b.Pos_y-b.Radius, b.Pos_y+b.Radius)
}
func (b *CheckBox_Struct) GetPos() (float32, float32) { return b.Pos_x, b.Pos_y }
func (b *CheckBox_Struct) GetBounds() (float32, float32) { return b.Radius*2 + b.CheckBox_Text.Text.Width(), b.Radius*2 }
func (b *CheckBox_Struct) SetName(s string) { b.CheckBoxName = s }
func (b *CheckBox_Struct) GetName() string { return b.CheckBoxName }
func (b *CheckBox_Struct) SetDIM(dim []float32){ b.DIM = dim }
func (b *CheckBox_Struct) GetDIM() []float32 { return b.DIM }
func (b *CheckBox_Struct) SetXYZ(xyz []float32){ b.XYZ = xyz }
func (b *CheckBox_Struct) GetXYZ() []float32 { return b.XYZ }
func (b *CheckBox_Struct) SetRGB(rgb []float32){ b.RGB = rgb }
func (b *CheckBox_Struct) GetRGB() []float32 { return b.RGB }
func (b *CheckBox_Struct) SetVAO(vao []intf.Drawing_Struct){ b.VAO = vao }
func (b *CheckBox_Struct) GetVAO() []intf.Drawing_Struct { return b.VAO }
func (b *CheckBox_Struct) SetClickEvent(ev intf.Event_Interface){ b.ClickEvent = ev }
func (b *CheckBox_Struct) SetClickTrigger(t int) { b.ClickTrigger = t }
func (b *CheckBox_Struct) GetClickTrigger() int {	return b.ClickTrigger }
func (b *CheckBox_Struct) GetCanvas() intf.Canvas_Interface { return b.CheckBoxCanvas }
func (b *CheckBox_Struct) SetClickableBounds(x_min, x_max, y_min, y_max float32) {
	b.Bounds_X_max = int(x_max)
	b.Bounds_X_min = int(x_min)
	b.Bounds_Y_max = int(y_max)
	b.Bounds_Y_min = int(y_min)
}
func (b *CheckBox_Struct) GetClickableBounds() (int, int, int, int, int) {
	return b.Bounds_X_min, b.Bounds_X_max, b.Bounds_Y_min, b.Bounds_Y_max, int(b.Pos_z)
}
func (b *CheckBox_Struct) SetPosZ(z float32){
	b.Pos_z = z
}
func (b *CheckBox_Struct) GetPosZ() float32 {
	return b.Pos_z
}
func (b *CheckBox_Struct) Hidetext(){
	b.CheckBox_Text.Text.Hide()
}
func (b *CheckBox_Struct) Showtext(){
	b.CheckBox_Text.Text.Show()
}

