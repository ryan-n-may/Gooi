package components 
/**
 * Updated 11/11/2023.
 * Button Component.
 * Implements Drawable -> Component -> Clickable.
 **/
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	mgl32   "github.com/go-gl/mathgl/mgl32"
	
	intf 	"gooi/interfaces"
	event 	"gooi/base/event"
	cons 	"gooi/base/constants"
	colours "gooi/base/colours"
	
	time 	"time"
	log 	"log"
	fmt 	"fmt"
)
type Button_Struct struct {
	err  			error
	ButtonName 		string
	XYZ 			[]float32
	RGB 			[]float32 
	DIM 			[]float32
	VAO 			[]intf.Drawing_Struct
	ButtonCanvas 	*Canvas_Struct
	Clickable 		bool
	ClickEvent 		intf.Event_Interface
	ClickTrigger 	int
	Bounds_X_max 	int 
	Bounds_X_min 	int
	Bounds_Y_max 	int
	Bounds_Y_min 	int
	ButtonText 		string
	Width, Height 	float32
	Pos_x, Pos_y 	float32
	Pos_z 			float32
	Radius 			float32
	Padding			float32
	BorderThickness float32
	WindowHeight    		*float32
	WindowHeight_Initial 	float32
	WindowWidth   			*float32
	WindowWidth_Initial 	float32
	AnimationTime  	time.Duration
	Button_Text 	*Text_Struct
	Button_Font 	*Font_Struct
	ButtonBodyColour 			[3]float32
	ButtonBodyColour_Active		[3]float32
	ButtonBodyColour_Clicked 	[3]float32
	ButtonBodyColour_Disabled 	[3]float32
	ButtonBorderColour   		[3]float32
	ButtonBackgroundColour   	[3]float32
}
/**
 * CreateButton
 * 	Creates a new button composable.
 **/
func CreateThemedButton(
	name 				string,
	canvas 				*Canvas_Struct, 
	width, height 		float32,
	pos_x, pos_y 		float32, 
	font_size 			int, 
	theme  				*colours.Theme,
	button_event 		*event.Event_Struct,
) *Button_Struct {
	log.Println("creating new thembed [Button] struct.")
	var b = Button_Struct{}
	// Specified paramaters
	b.ButtonCanvas = canvas
	b.Width = width
	b.Height = height
	b.Pos_x = pos_x
	b.Pos_y = pos_y
	b.Pos_z = 0.0
	b.ClickEvent = button_event
	b.Clickable = true
	b.ClickTrigger = cons.MOUSE_PRESSED
	b.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
	b.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
	b.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
	b.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()
	b.ButtonName = name
	// Themed paramaters
	b.Radius = theme.ButtonRadius
	b.Padding = theme.ButtonBorderPadding
	b.BorderThickness = theme.ButtonBorderThickness
	b.AnimationTime = theme.ButtonAnimationTime
	b.Button_Font = ReadFontFile(theme.FontName, theme.FontPath, font_size)
	b.Button_Font.Font.ResizeWindow(*b.WindowWidth, *b.WindowHeight)
	b.Button_Text = CreateText(b.ButtonName, b.Button_Font.Font)
	b.Button_Text.Text.SetPosition(mgl32.Vec3{pos_x + width/2, pos_y + height/2, 1.0})
	b.ButtonBodyColour = theme.ButtonBodyColour_Idle
	b.ButtonBodyColour_Active = theme.ButtonBodyColour_Idle
	b.ButtonBodyColour_Clicked = theme.ButtonBodyColour_Selected
	b.ButtonBodyColour_Disabled = theme.ButtonBodyColour_Disabled
	b.ButtonBorderColour = theme.ButtonBorderColour
	b.ButtonBackgroundColour = theme.ButtonBackgroundColour
	b.GeneratePolygons()
	b.SetClickableBounds(b.Pos_x, b.Pos_x+b.Width, b.Pos_y, b.Pos_y+b.Height)
   	return &b
}
func CreateButton(
	name 				string,
	canvas 				*Canvas_Struct, 
	width, height 		float32,
	pos_x, pos_y 		float32, 
	radius 				float32,
	padding 			float32,
	border_thickness    float32,
	animation_time 		time.Duration,
	button_event 		*event.Event_Struct,
	font_path 			string,
	font_name 			string,
	font_size 			int,
	button_body_colour 				[3]float32,
	button_body_colour_clicked		[3]float32,
	button_body_colour_deactivated	[3]float32,
	button_border_colour	 		[3]float32,
	button_background_colour 		[3]float32,
	) *Button_Struct {
	log.Println("creating new custom [Button] struct.")
	var b = Button_Struct{}
	b.Width = width
	b.Height = height
	b.Pos_x = pos_x
	b.Pos_y = pos_y
	b.Pos_z = 0.0
	b.Radius = radius
	b.BorderThickness = border_thickness
	b.Padding = padding
	b.ButtonCanvas = canvas
	b.AnimationTime = animation_time
	// Clickable interface charactistics
	b.ClickEvent = button_event
	b.Clickable = true
	b.ClickTrigger = cons.MOUSE_PRESSED
	// Window characteristics
 	b.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
	b.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
	b.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
	b.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()
	// Component interface characteristics
	b.ButtonName = name
	// Button Label
	b.Button_Font = ReadFontFile(font_name, font_path, font_size)
	b.Button_Font.Font.ResizeWindow(*b.WindowWidth, *b.WindowHeight)
	b.Button_Text = CreateText(b.ButtonName, b.Button_Font.Font)
	b.Button_Text.Text.SetPosition(mgl32.Vec3{pos_x + width/2, pos_y + height/2, 1.0})
	b.ButtonBodyColour = button_body_colour
	b.ButtonBodyColour_Active = button_body_colour
	b.ButtonBodyColour_Clicked = button_body_colour_clicked
	b.ButtonBodyColour_Disabled = button_body_colour_deactivated
	b.ButtonBorderColour = button_border_colour
	b.ButtonBackgroundColour = button_background_colour
	b.GeneratePolygons()
	b.SetClickableBounds(b.Pos_x, b.Pos_x+b.Width, b.Pos_y, b.Pos_y+b.Height)
   	return &b
}
// GeneratePolygons()
// Generates the VAO array of the polygons used to draw the button. 
// Stores the VAO in intf.Drawing_Struct alongisde the drawing mode (gl.TRIANGLE or gl.TRIANGLE_FAN)
func (b *Button_Struct) GeneratePolygons(){
	b.VAO = make([]intf.Drawing_Struct, 18)
	// Border rectangles 
	b.VAO[0] = intf.Drawing_Struct{ intf.GenerateRectangle(b, b.ButtonBorderColour, b.Width-b.Radius, b.Height, b.Pos_x+b.Radius/2, b.Pos_y, b.Pos_z, b.WindowHeight_Initial, b.WindowWidth_Initial), gl.TRIANGLES }
	b.VAO[1] = intf.Drawing_Struct{ intf.GenerateRectangle(b, b.ButtonBorderColour, b.Width, b.Height-b.Radius, b.Pos_x, b.Pos_y+b.Radius/2, b.Pos_z, b.WindowHeight_Initial, b.WindowWidth_Initial), gl.TRIANGLES }
	// Border corner circles
	b.VAO[2] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBorderColour, b.Pos_x+b.Radius/2, b.Pos_y+b.Radius/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[3] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBorderColour, b.Pos_x+b.Width-b.Radius/2, b.Pos_y+b.Radius/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[4] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBorderColour, b.Pos_x+b.Radius/2, b.Pos_y+b.Height-b.Radius/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[5] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBorderColour, b.Pos_x+b.Width-b.Radius/2, b.Pos_y+b.Height-b.Radius/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	// Background rectangles
	b.VAO[6] = intf.Drawing_Struct{ intf.GenerateRectangle(b, b.ButtonBackgroundColour, b.Width-b.Radius, b.Height-b.BorderThickness, b.Pos_x+b.Radius/2, b.Pos_y+b.BorderThickness/2, b.Pos_z, b.WindowHeight_Initial, b.WindowWidth_Initial), gl.TRIANGLES }
	b.VAO[7] = intf.Drawing_Struct{ intf.GenerateRectangle(b, b.ButtonBackgroundColour, b.Width-b.BorderThickness, b.Height-b.Radius, b.Pos_x+b.BorderThickness/2, b.Pos_y+b.Radius/2, b.Pos_z, b.WindowHeight_Initial, b.WindowWidth_Initial), gl.TRIANGLES }
	// Background corner circles
	b.VAO[8] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBackgroundColour, b.Pos_x+b.Radius/2+b.BorderThickness/2, b.Pos_y+b.Radius/2+b.BorderThickness/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[9] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBackgroundColour, b.Pos_x+b.Width-b.Radius/2-b.BorderThickness/2, b.Pos_y+b.Radius/2+b.BorderThickness/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[10] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBackgroundColour, b.Pos_x+b.Radius/2+b.BorderThickness/2, b.Pos_y+b.Height-b.Radius/2-b.BorderThickness/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[11] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBackgroundColour, b.Pos_x+b.Width-b.Radius/2-b.BorderThickness/2, b.Pos_y+b.Height-b.Radius/2-b.BorderThickness/2, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	// Button Body rectangles
	b.VAO[12] = intf.Drawing_Struct{ intf.GenerateRectangle(b, b.ButtonBodyColour, b.Width-b.Radius, b.Height-b.Padding*2, b.Pos_x+b.Radius/2, b.Pos_y+b.Padding, b.Pos_z, b.WindowHeight_Initial, b.WindowWidth_Initial), gl.TRIANGLES }
	b.VAO[13] = intf.Drawing_Struct{ intf.GenerateRectangle(b, b.ButtonBodyColour, b.Width-b.Padding*2, b.Height-b.Radius, b.Pos_x+b.Padding, b.Pos_y+b.Radius/2, b.Pos_z, b.WindowHeight_Initial, b.WindowWidth_Initial), gl.TRIANGLES }
	// Button Body corner circles
	b.VAO[14] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBodyColour, b.Pos_x+b.Radius/2+b.Padding, b.Pos_y+b.Radius/2+b.Padding, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[15] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBodyColour, b.Pos_x+b.Width-b.Radius/2-b.Padding, b.Pos_y+b.Radius/2+b.Padding, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[16] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBodyColour, b.Pos_x+b.Radius/2+b.Padding, b.Pos_y+b.Height-b.Radius/2-b.Padding, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	b.VAO[17] = intf.Drawing_Struct{ intf.GenerateCirlce(b, b.ButtonBodyColour, b.Pos_x+b.Width-b.Radius/2-b.Padding, b.Pos_y+b.Height-b.Radius/2-b.Padding, b.Pos_z, (b.Radius)/2, b.WindowWidth_Initial, b.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
}
// Draw()
// This method draws the VAO array to gl using the canvas program.
func (b *Button_Struct) Draw() {
	// Obtain program that isnt FontProg
	gl.UseProgram(b.GetCanvas().GetPrograms())
	for _, v := range b.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(b.GetXYZ())/3))
	}
	// Draw button text (Uses FontProg program)
	// Modified glText implementation:
	Draw(b.Button_Text)
}
func (b *Button_Struct) Redraw() {
	b.GeneratePolygons()
	gl.UseProgram(b.GetCanvas().GetPrograms())
	for _, v := range b.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(b.GetXYZ())/3))
	}
	// Draw button text (Uses FontProg program)
	// Modified glText implementation:
	//b.Button_Text.Text.Font.ResizeWindow(*b.WindowWidth, *b.WindowHeight)
	Draw(b.Button_Text)
}
// Move()
// This method moves the button and the clickable area of the button.
func (b *Button_Struct) Move(delta_x, delta_y float32) {
	log.Println(fmt.Sprintf("moving [Button] by %v, %v.", delta_x, delta_y))
	b.Pos_y = b.Pos_y + delta_y
	b.Pos_x = b.Pos_x + delta_x
	b.GeneratePolygons()
	b.SetClickableBounds(b.Pos_x, b.Pos_x+b.Width, b.Pos_y, b.Pos_y+b.Height)
} 
// Kill(*bool)
// This method sets alive to false: indicates to the animation process that the 
// animation has completed.
func (b *Button_Struct) Kill(alive *bool){
	*alive = false
}
// StatClickAnimation()
// Set button colour to clicked state; refresh canvas and update polyons.
func (b *Button_Struct) StartClickAnimation(){
	b.ButtonBodyColour = b.ButtonBodyColour_Clicked
	b.GeneratePolygons()
   	b.GetCanvas().RefreshCanvas()	
}
// EndClickAnimation()
// Set button colour to active state (clickable but idle); refresh canvas and update polygons
func (b *Button_Struct) EndClickAnimation(){
	b.ButtonBodyColour = b.ButtonBodyColour_Active
   	b.GeneratePolygons()
   	b.GetCanvas().RefreshCanvas()	
}
// AnimateTrigger(int, *bool)
// This method animates the click operation
func (b *Button_Struct) AnimateTrigger(pressed int, alive *bool){
	if b.Clickable == true && pressed == b.ClickTrigger{
		*alive = true
		defer b.Kill(alive)
		b.StartClickAnimation()
		time.Sleep(b.AnimationTime * time.Millisecond)
		b.EndClickAnimation()
	} 
}
// TriggerCLickEvent(pressed)
// This method add's the button Event to the EventHandler queue;
// IF -> the button has a method, it is clickable, and the trigger is correct;
// Trigger can be MOUSE_PRESSED, MOUSE_RELEASED, or NO_CHANGE
func (b *Button_Struct) TriggerClickEvent(pressed int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if b.ClickEvent.GetMethod() != nil && b.Clickable == true && pressed == b.ClickTrigger{
		log.Println("triggering click event in [Button]")
		b.GetCanvas().GetEventHandler().AddEventToEventQueue(b.ClickEvent.GetName())
  	}
}
// SetClickable(bool) & GetClickable() bool
// Sets the button to clickable or un-clickable.
// Regenerates polygons to update colour, and refresh canvas.
func (b *Button_Struct) SetClickable(clickable bool){
	log.Println(fmt.Sprintf("[Button] set clickable to %v.", clickable))
	b.Clickable = clickable
	if !clickable{ b.ButtonBodyColour = b.ButtonBodyColour_Disabled
   	} else { b.ButtonBodyColour = b.ButtonBodyColour_Active }
   	b.GeneratePolygons()
   	b.GetCanvas().RefreshCanvas()	  
}
func (b *Button_Struct) GetClickable() bool { return b.Clickable }
// SetPos(float32, float32) & GetPos() float32, float32
// Sets the position on the window (absolute)
// Origin starting in the bottom left corner of the window.
// Updates clickable bounds and re-draws.
func (b *Button_Struct) SetPos(x, y float32) { 
	b.Pos_x = x
	b.Pos_y = y
	b.Button_Text.Text.SetPosition(mgl32.Vec3{b.Pos_x + b.Width/2, b.Pos_y + b.Height/2, 1.0}) 
	b.GeneratePolygons()
	b.SetClickableBounds(b.Pos_x, b.Pos_x+b.Width, b.Pos_y, b.Pos_y+b.Height)
	b.GetCanvas().RefreshCanvas()	
}
func (b *Button_Struct) GetPos() (float32, float32) { return b.Pos_x, b.Pos_y }
/**
 * Other Setter and Getter Methods
 **/
func (b *Button_Struct) GetBounds() (float32, float32) { return b.Width, b.Height }
func (b *Button_Struct) SetButtonText(s string) { b.ButtonText = s }
func (b *Button_Struct) GetButtonText() string { return b.ButtonText }
func (b *Button_Struct) SetName(s string) { b.ButtonName = s }
func (b *Button_Struct) GetName() string { return b.ButtonName }
func (b *Button_Struct) SetDIM(dim []float32){ b.DIM = dim }
func (b *Button_Struct) GetDIM() []float32 { return b.DIM }
func (b *Button_Struct) SetXYZ(xyz []float32){ b.XYZ = xyz }
func (b *Button_Struct) GetXYZ() []float32 { return b.XYZ }
func (b *Button_Struct) SetRGB(rgb []float32){ b.RGB = rgb }
func (b *Button_Struct) GetRGB() []float32 { return b.RGB }
func (b *Button_Struct) SetVAO(vao []intf.Drawing_Struct){ b.VAO = vao }
func (b *Button_Struct) GetVAO() []intf.Drawing_Struct { return b.VAO }
func (b *Button_Struct) SetClickEvent(ev intf.Event_Interface){ b.ClickEvent = ev }
func (b *Button_Struct) SetClickTrigger(t int) { b.ClickTrigger = t }
func (b *Button_Struct) GetClickTrigger() int {	return b.ClickTrigger }
func (b *Button_Struct) GetCanvas() intf.Canvas_Interface { return b.ButtonCanvas }
func (b *Button_Struct) SetClickableBounds(x_min, x_max, y_min, y_max float32) {
	b.Bounds_X_max = int(x_max)
	b.Bounds_X_min = int(x_min)
	b.Bounds_Y_max = int(y_max)
	b.Bounds_Y_min = int(y_min)
}
func (b *Button_Struct) GetClickableBounds() (int, int, int, int, int) {
	return b.Bounds_X_min, b.Bounds_X_max, b.Bounds_Y_min, b.Bounds_Y_max, int(b.Pos_z)
}
func (b *Button_Struct) SetPosZ(z float32){
	if z >= 1.0 {
		b.Hidetext()
	} else {
		b.Showtext()
	}
	b.Pos_z = z
}
func (b *Button_Struct) GetPosZ() float32 {
	return b.Pos_z
}
func (b *Button_Struct) Hidetext(){
	log.Println("hiding [Button] text.")
	b.Button_Text.Text.Hide()
}
func (b *Button_Struct) Showtext(){
	log.Println("showing [Button] text.")
	b.Button_Text.Text.Show()
}
