package components 
/**
 * Updated 11/11/2023.
 * ToggleSwitch implements -> Drawable -> Component -> Clickable
 **/
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	intf 	"gooi/interfaces"
	event 	"gooi/base/event"
	cons 	"gooi/base/constants"
	colours "gooi/base/colours"
	time 	"time"
)
type ToggleSwitch_Struct struct {
	err  					error
	// Component interface 
	ToggleSwitchName 		string
	XYZ 					[]float32
	RGB 					[]float32 
	DIM 					[]float32
	VAO 					[]intf.Drawing_Struct
	// Clickable interface
	Clickable 				bool
	ClickEvent 				intf.Event_Interface
	ClickTrigger 			int
	Bounds_X_max 			int 
	Bounds_X_min 			int
	Bounds_Y_max 			int
	Bounds_Y_min 			int
	// Button specific 
	Width, Height 			float32
	Pos_x, Pos_y 			float32
	Pos_z 					float32
	Padding  				float32
	WindowHeight    		*float32
	WindowHeight_Initial 	float32
	WindowWidth   			*float32
	WindowWidth_Initial 	float32
	ToggleCanvas 			*Canvas_Struct
	AnimationTime  			time.Duration
	// Colours
	ToggleColour  				[3]float32
	ToggleColour_True 			[3]float32
	ToggleColour_False			[3]float32
	ToggleColour_Deactivated 	[3]float32
	ToggleBorderColour   		[3]float32
	ToggleBackgroundColour   	[3]float32

	TogglePosition 				float32
	ToggleState 				bool
}

func CreateThemedToggle(
	name string,
	canvas *Canvas_Struct, 
	width, height, pos_x, pos_y float32, 
	toggle_event *event.Event_Struct,
	theme *colours.Theme,
	) *ToggleSwitch_Struct {
		// Specified
		var t = ToggleSwitch_Struct{}
		t.ToggleSwitchName = name
		t.Width = width
		t.Height = height
		t.Pos_x = pos_x
		t.Pos_y = pos_y
		t.Pos_z = 0.0
		t.ToggleCanvas = canvas
		t.ClickEvent = toggle_event
		// Window characteristics
	 	t.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
		t.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
		t.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
		t.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()
		// Hard coded
		t.ToggleState = false
		t.TogglePosition = 0
		t.Clickable = true
		t.ClickTrigger = cons.MOUSE_PRESSED
		// Themed
		t.Padding = theme.ToggleBorderPadding
		t.AnimationTime = theme.ToggleAnimationTime
		
		t.ToggleColour 				= theme.ToggleCircleColour
		t.ToggleColour_True 		= theme.ToggleBackgroundColour_On
		t.ToggleColour_False		= theme.ToggleBackgroundColour_Off
		t.ToggleColour_Deactivated 	= theme.ToggleBackgroundColour_Deactivated
		t.ToggleBorderColour   		= theme.ToggleBorderColour
		t.ToggleBackgroundColour 	= theme.ToggleBackgroundColour

		return &t	
}

func CreateToggleSwitch(
	name string,
	canvas *Canvas_Struct, 
	width, height, pos_x, pos_y float32, 
	padding float32,
	animation_time time.Duration,
	toggle_event *event.Event_Struct,
	toggle_body_colour_true			[3]float32,
	toggle_body_colour_false		[3]float32,
	toggle_body_colour_deactivated  [3]float32,
	toggle_border_colour	 		[3]float32,
	toggle_background_colour 		[3]float32,
	) *ToggleSwitch_Struct {
	// Creating buttons 
	var t = ToggleSwitch_Struct{}
	// Button specific characteristics 
	t.Width = width
	t.Height = height
	t.Pos_x = pos_x
	t.Pos_y = pos_y
	t.Pos_z = 0.0
	t.Padding = padding
	t.ToggleCanvas = canvas
	t.AnimationTime = animation_time
	// Clickable interface charactistics
	t.ClickEvent = toggle_event
	t.Clickable = true
	t.ClickTrigger = cons.MOUSE_PRESSED
	// Window characteristics
 	t.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
	t.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
	t.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
	t.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()
	// Component interface characteristics
	t.ToggleSwitchName = name
	t.ToggleState = false
	
	t.ToggleColour 				= toggle_body_colour_false
	t.ToggleColour_True 		= toggle_body_colour_true
	t.ToggleColour_False		= toggle_body_colour_false
	t.ToggleColour_Deactivated 	= toggle_body_colour_deactivated
	t.ToggleBorderColour   		= toggle_border_colour
	t.ToggleBackgroundColour 	= toggle_background_colour

	t.TogglePosition 			= 0

	t.GeneratePolygons()
	t.SetClickableBounds(t.Pos_x - t.Height/2, t.Pos_x + t.Width + t.Height/2, t.Pos_y, t.Pos_y + t.Height)

   	return &t
}

func (t *ToggleSwitch_Struct) GeneratePolygons(){
	t.VAO = make([]intf.Drawing_Struct, 10)
	// Border rounded rectangle
	t.VAO[0] = intf.Drawing_Struct{
		intf.GenerateRectangle(t, t.ToggleBorderColour, t.Width, t.Height, t.Pos_x, t.Pos_y, t.Pos_z, t.WindowHeight_Initial, t.WindowWidth_Initial),
		gl.TRIANGLES }
	t.VAO[1] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleBorderColour, t.Pos_x, t.Pos_y+t.Height/2, t.Pos_z, t.Height/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }	
	t.VAO[2] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleBorderColour, t.Pos_x+t.Width, t.Pos_y+t.Height/2, t.Pos_z, t.Height/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }	
	// Background rounded rectangle	
	t.VAO[3] = intf.Drawing_Struct{
		intf.GenerateRectangle(t, t.ToggleBackgroundColour, t.Width-t.Padding, t.Height-t.Padding, t.Pos_x+(t.Padding/2), t.Pos_y+(t.Padding/2), t.Pos_z, t.WindowHeight_Initial, t.WindowWidth_Initial),
		gl.TRIANGLES }
	t.VAO[4] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleBackgroundColour, t.Pos_x, t.Pos_y+t.Height/2, t.Pos_z, (t.Height-t.Padding)/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }	
	t.VAO[5] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleBackgroundColour, t.Pos_x+t.Width, t.Pos_y+t.Height/2, t.Pos_z, (t.Height-t.Padding)/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }
	//Coloured rounded rectangle
	t.VAO[6] = intf.Drawing_Struct{
		intf.GenerateRectangle(t, t.ToggleColour, t.Width-t.Padding*2, t.Height-t.Padding*2, t.Pos_x+t.Padding, t.Pos_y+t.Padding, t.Pos_z, t.WindowHeight_Initial, t.WindowWidth_Initial),
		gl.TRIANGLES }
	t.VAO[7] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleColour, t.Pos_x, t.Pos_y+t.Height/2, t.Pos_z, (t.Height-t.Padding*2)/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }	
	t.VAO[8] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleColour, t.Pos_x+t.Width, t.Pos_y+t.Height/2, t.Pos_z, (t.Height-t.Padding*2)/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }
	//Toggle	
	t.VAO[9] = intf.Drawing_Struct{
		intf.GenerateCirlce(t, t.ToggleBorderColour, t.Pos_x+t.TogglePosition, t.Pos_y+t.Height/2, t.Pos_z, (t.Height-t.Padding*2)/2, t.WindowWidth_Initial, t.WindowHeight_Initial, 128),
		gl.TRIANGLE_FAN }	
}

func (t *ToggleSwitch_Struct) Draw() {
	gl.UseProgram(t.GetCanvas().GetPrograms())
	for _, v := range t.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(t.GetXYZ())/3))
	}
}

func (t *ToggleSwitch_Struct) Redraw() {
	t.GeneratePolygons()
	gl.UseProgram(t.GetCanvas().GetPrograms())
	for _, v := range t.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(t.GetXYZ())/3))
	}
}


func (t *ToggleSwitch_Struct) Move(delta_x, delta_y float32) {
	t.Pos_y = t.Pos_y + delta_y
	t.Pos_x = t.Pos_x + delta_x
	t.GeneratePolygons()
} 

func (t *ToggleSwitch_Struct) Kill(alive *bool){
	*alive = false
}

func (t *ToggleSwitch_Struct) StartClickAnimation(){
	if !t.ToggleState{
		for i := 0.0; i < 1.0; i += 0.1 {
			t.TogglePosition = t.Width * float32(i)
			t.ToggleColour = t.ToggleColour_True
			time.Sleep(t.AnimationTime * time.Millisecond)
			t.GeneratePolygons()
	   		t.GetCanvas().RefreshCanvas()	
		}
	} else {
		for i := 1.0; i >= 0.0; i -= 0.1 {
			t.TogglePosition = t.Width * float32(i)
			t.ToggleColour = t.ToggleColour_False
			time.Sleep(t.AnimationTime * time.Millisecond)
			t.GeneratePolygons()
	   		t.GetCanvas().RefreshCanvas()	
		}
	}
	t.ToggleState = !t.ToggleState
}

func (t *ToggleSwitch_Struct) AnimateTrigger(pressed int, alive *bool){
	if t.Clickable == true && pressed == t.ClickTrigger{
		*alive = true
		defer t.Kill(alive)
		t.StartClickAnimation()
	} 
}

func (t *ToggleSwitch_Struct) TriggerClickEvent(pressed int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if t.ClickEvent.GetMethod() != nil && t.Clickable == true && pressed == t.ClickTrigger{
		t.GetCanvas().GetEventHandler().AddEventToEventQueue(t.ClickEvent.GetName())
  	}
}

func (t *ToggleSwitch_Struct) SetClickable(clickable bool){
	t.Clickable = clickable
	if !clickable{ t.ToggleColour = t.ToggleColour_Deactivated
   	} else { t.ToggleColour = t.ToggleColour_True }
   	t.GeneratePolygons()
   	t.GetCanvas().RefreshCanvas()	  
}
func (t *ToggleSwitch_Struct) GetClickable() bool { return t.Clickable }
func (t *ToggleSwitch_Struct) SetPos(x, y float32) { 
	t.Pos_x = x + t.Height/2
	t.Pos_y = y 
	t.GeneratePolygons()
	t.SetClickableBounds(t.Pos_x - t.Height/2, t.Pos_x + t.Width + t.Height/2, t.Pos_y, t.Pos_y + t.Height)
}
func (t *ToggleSwitch_Struct) GetPos() (float32, float32) { return t.Pos_x, t.Pos_y }
func (t *ToggleSwitch_Struct) GetBounds() (float32, float32) { return t.Width + t.Height, t.Height }
func (t *ToggleSwitch_Struct) SetName(s string) { t.ToggleSwitchName = s }
func (t *ToggleSwitch_Struct) GetName() string { return t.ToggleSwitchName }
func (t *ToggleSwitch_Struct) SetDIM(dim []float32){ t.DIM = dim }
func (t *ToggleSwitch_Struct) GetDIM() []float32 { return t.DIM }
func (t *ToggleSwitch_Struct) SetXYZ(xyz []float32){ t.XYZ = xyz }
func (t *ToggleSwitch_Struct) GetXYZ() []float32 { return t.XYZ }
func (t *ToggleSwitch_Struct) SetRGB(rgb []float32){ t.RGB = rgb }
func (t *ToggleSwitch_Struct) GetRGB() []float32 { return t.RGB }
func (t *ToggleSwitch_Struct) SetVAO(vao []intf.Drawing_Struct){ t.VAO = vao }
func (t *ToggleSwitch_Struct) GetVAO() []intf.Drawing_Struct{ return t.VAO }
func (t *ToggleSwitch_Struct) SetClickEvent(ev intf.Event_Interface){ t.ClickEvent = ev }
func (t *ToggleSwitch_Struct) SetClickTrigger(trigger int) { t.ClickTrigger = trigger }
func (t *ToggleSwitch_Struct) GetClickTrigger() int { return t.ClickTrigger }
func (t *ToggleSwitch_Struct) GetCanvas() intf.Canvas_Interface { return t.ToggleCanvas }
func (t *ToggleSwitch_Struct) SetClickableBounds(x_min, x_max, y_min, y_max float32) {
	t.Bounds_X_max = int(x_max)
	t.Bounds_X_min = int(x_min)
	t.Bounds_Y_max = int(y_max)
	t.Bounds_Y_min = int(y_min)
}
func (t *ToggleSwitch_Struct) GetClickableBounds() (int, int, int, int, int) {
	return t.Bounds_X_min, t.Bounds_X_max, t.Bounds_Y_min, t.Bounds_Y_max, int(t.Pos_z)
}
func (t *ToggleSwitch_Struct) SetPosZ(z float32){
	t.Pos_z = z
}
func (t *ToggleSwitch_Struct) GetPosZ() float32 {
	return t.Pos_z
}