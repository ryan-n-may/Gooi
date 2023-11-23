package components 
/**
 * Updated 11/11/2023.
 * TextInput implements -> Drawable -> Component -> Clickable
 **/
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
	gl 		"github.com/go-gl/gl/v4.1-core/gl"
	mgl32   "github.com/go-gl/mathgl/mgl32"
	intf 	"gooi/interfaces"
	event 	"gooi/base/event"
	list 	"gooi/base/listeners"
	log 	"log"
	fmt 	"fmt"
)

type Cursor_Struct struct {
	TextInput 		*Text_Input_Struct
	CursorIndex 	int 
	CursorPosX 		float32
	MasterPosX 		*float32
}
func (c *Cursor_Struct) SetCursorIndexFromClickLocation(pos_x float32){
	var offset = -c.TextInput.Width/2 + c.TextInput.Input_Text.Text.Width()/2
	c.CursorPosX = pos_x
	c.CursorIndex, _ = c.TextInput.Input_Text.Text.ClickedCharacter(float64(pos_x), float64(offset))
	log.Println(fmt.Sprintf("[Custor] index in [TextInput] set to %v from click location at %v.", c.CursorIndex, c.CursorPosX))
}
func (c *Cursor_Struct) GetCursorPosFromIndex() float32 {
	c.CursorPosX = float32(c.TextInput.Input_Text.Text.CharPosition(c.CursorIndex))
	c.CursorPosX = c.CursorPosX + *c.MasterPosX + c.TextInput.Padding + (c.TextInput.Input_Text.Text.Width()/2)
	log.Println(fmt.Sprintf("from [Cursor] index %v, pos is returned as %v", c.CursorIndex, c.CursorPosX))
	return c.CursorPosX
}
type HighLight_Struct struct {
	TextInput 		*Text_Input_Struct
	LeftIndex  		int 
	RightIndex 		int
	LeftPosX 		float32
	RightPosX  		float32
	MasterPosX 		*float32
}
func (h *HighLight_Struct) SetLeftIndexFromClickLocation(pos_x float32){
	var offset = -h.TextInput.Width/2 + h.TextInput.Input_Text.Text.Width()/2
	h.LeftPosX = pos_x
	h.LeftIndex, _ = h.TextInput.Input_Text.Text.ClickedCharacter(float64(pos_x), float64(offset))
	log.Println(fmt.Sprintf("[Highlight] left index in [TextInput] set to %v from click location at %v.", h.LeftIndex, h.LeftPosX))
}
func (h *HighLight_Struct) SetRightIndexFromClickLocation(pos_x float32){
	var offset = -h.TextInput.Width/2 + h.TextInput.Input_Text.Text.Width()/2
	h.RightPosX = pos_x
	h.RightIndex, _ = h.TextInput.Input_Text.Text.ClickedCharacter(float64(pos_x), float64(offset))
	log.Println(fmt.Sprintf("[Highlight] right index in [TextInput] set to %v from click location at %v.", h.RightIndex, h.RightPosX))
}
func (h *HighLight_Struct) GetLeftPosFromIndex() float32 {
	h.LeftPosX = float32(h.TextInput.Input_Text.Text.CharPosition(h.LeftIndex))
	h.LeftPosX = h.LeftPosX + *h.MasterPosX + h.TextInput.Padding + (h.TextInput.Input_Text.Text.Width()/2)
	log.Println(fmt.Sprintf("[Highlight] left pos in [TextInput] is %v from click location at %v.", h.LeftPosX, h.LeftIndex))
	return h.LeftPosX
}
func (h *HighLight_Struct) GetRightPosFromIndex() float32 {
	h.RightPosX = float32(h.TextInput.Input_Text.Text.CharPosition(h.RightIndex))
	h.RightPosX = h.RightPosX + *h.MasterPosX + h.TextInput.Padding + (h.TextInput.Input_Text.Text.Width()/2)
	log.Println(fmt.Sprintf("[Highlight] right pos in [TextInput] is %v from click location at %v.", h.RightPosX, h.RightIndex))
	return h.RightPosX
}
func (h *HighLight_Struct) GetWidth() float32 {
	var width float32 = 0
	if h.RightIndex != -1 {
		width = (h.RightPosX - h.LeftPosX)
	}
	log.Println(fmt.Sprintf("[Highlight] width given as %v.", width))
	return width
}
type Text_Input_Struct struct {
	TextInputName 	string
	err  			error
	XYZ 			[]float32
	RGB 			[]float32 
	DIM 			[]float32
	Canvas 	*Canvas_Struct
	WindowHeight    		*float32
	WindowHeight_Initial 	float32
	WindowWidth   			*float32
	WindowWidth_Initial 	float32
	LabelText 		string
	Pos_x, Pos_y 	float32
	Pos_z 			float32
	Width, Height 	float32
	Radius 			float32
	Padding 		float32
	VAO 			[]intf.Drawing_Struct
	ButtonBorderColour 				[3]float32
	ButtonBackgroundColour_Selected [3]float32
	ButtonBackgroundColour_Idle		[3]float32
	ButtonBackgroundColour 			[3]float32
	ButtonHighlightColour 			[3]float32
	FontName  		string
	FontPath  		string
	FontSize 		int
	Input_Text 		*Text_Struct
	Input_Font 		*Font_Struct
	Bounds_Y_min	int
	Bounds_Y_max	int
	Bounds_X_min 	int
	Bounds_X_max 	int
	Clickable 		bool
	ClickEvent 		intf.Event_Interface
	AnimationTime 	int
	ClickTrigger 	int
	KeyListener 	*list.KeyHandler_Struct
	Highlight 		*HighLight_Struct
	Cursor 			*Cursor_Struct
}
func CreateTextInput(
	name string,
	placeholder string,
	canvas *Canvas_Struct, 
	pos_x, pos_y float32, 
	width, height float32,
	border_radius float32,
	padding float32,
	font_path string,
	font_name string,
	font_size int,
	button_border_colour [3]float32, 
	button_background_colour_selected [3]float32, 
	button_background_colour_idle [3]float32,
	button_highlight_colour [3]float32,
	click_trigger int,
	kl *list.KeyHandler_Struct,
	) *Text_Input_Struct {
	// Creating buttons 
	var i = Text_Input_Struct{}

	i.TextInputName = name
	
	i.Pos_x = pos_x
	i.Pos_y = pos_y
	i.Pos_z = 0.0
	i.Width = width
	i.Height = height
	i.Radius = border_radius
	i.Padding = padding
	i.SetClickableBounds(i.Pos_x, i.Pos_x+i.Width, i.Pos_y, i.Pos_y+i.Height)

	i.Canvas = canvas
	i.LabelText = placeholder

	i.FontPath = font_path
	i.FontName = font_name
	i.FontSize = font_size

	i.ButtonBackgroundColour = button_background_colour_idle
	i.ButtonBackgroundColour_Idle = button_background_colour_idle
	i.ButtonBackgroundColour_Selected = button_background_colour_selected
	i.ButtonBorderColour = button_border_colour
	i.ButtonHighlightColour = button_highlight_colour
	
	i.WindowWidth = canvas.CanvasWindow.GetWindowWidth()
	i.WindowHeight = canvas.CanvasWindow.GetWindowHeight()
	i.WindowWidth_Initial = *canvas.CanvasWindow.GetWindowWidth()
	i.WindowHeight_Initial = *canvas.CanvasWindow.GetWindowHeight()

	i.Input_Font = ReadFontFile(i.FontName, i.FontPath, i.FontSize)
	i.Input_Font.Font.ResizeWindow(*i.WindowWidth, *i.WindowHeight)
	i.Input_Text = CreateText(i.LabelText, i.Input_Font.Font)
	i.Input_Text.Text.SetPosition(mgl32.Vec3{i.Pos_x, i.Pos_y, i.Pos_z})

	i.Clickable = true
	i.AnimationTime = 1
	i.ClickTrigger = click_trigger
	var ClickEventParameters = event.NewEventParameter(&i)
	var Event = event.Event_Struct{
		i.selectedInputField, 
		fmt.Sprintf("selectedInputField_%s", i.TextInputName), 
		ClickEventParameters,
	}
	i.ClickEvent = &Event
	i.GetCanvas().GetEventHandler().RegisterEventToHandler(i.ClickEvent)
	i.KeyListener = kl

	var cursor = Cursor_Struct{}
	i.Cursor = &cursor
	i.Cursor.CursorIndex = 0
	i.Cursor.TextInput = &i
	i.Cursor.MasterPosX = &i.Pos_x

	var highlight = HighLight_Struct{
		&i, -1, -1, 0, 0, &i.Pos_x,
	}
	i.Highlight = &highlight

	i.GeneratePolygons()

   	return &i
}

func (i *Text_Input_Struct) GeneratePolygons() {

	if i.GetKeyListener().GetFocus() == i {
		//fmt.Println("Drawing as if selected")
		i.ButtonBackgroundColour = i.ButtonBackgroundColour_Selected
	} else {
		//fmt.Println("Drawing as if idle")
		i.ButtonBackgroundColour = i.ButtonBackgroundColour_Idle
		i.Highlight.LeftIndex = -1
		i.Highlight.RightIndex = -1
		i.Highlight.LeftPosX = 0
		i.Highlight.RightPosX = 0
	}

	i.VAO = make([]intf.Drawing_Struct, 14)
	// Highlight box
	// Border rectangles 
	i.VAO[0] = intf.Drawing_Struct{ intf.GenerateRectangle(i, i.ButtonBorderColour, i.Width-i.Radius, i.Height, i.Pos_x+i.Radius/2, i.Pos_y, i.Pos_z, i.WindowHeight_Initial, i.WindowWidth_Initial), gl.TRIANGLES }
	i.VAO[1] = intf.Drawing_Struct{ intf.GenerateRectangle(i, i.ButtonBorderColour, i.Width, i.Height-i.Radius, i.Pos_x, i.Pos_y+i.Radius/2, i.Pos_z, i.WindowHeight_Initial, i.WindowWidth_Initial), gl.TRIANGLES }
	// Border corner circles
	i.VAO[2] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBorderColour, i.Pos_x+i.Radius/2, i.Pos_y+i.Radius/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	i.VAO[3] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBorderColour, i.Pos_x+i.Width-i.Radius/2, i.Pos_y+i.Radius/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	i.VAO[4] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBorderColour, i.Pos_x+i.Radius/2, i.Pos_y+i.Height-i.Radius/2, i.Pos_z,  (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	i.VAO[5] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBorderColour, i.Pos_x+i.Width-i.Radius/2, i.Pos_y+i.Height-i.Radius/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	// Background rectangles
	i.VAO[6] = intf.Drawing_Struct{ intf.GenerateRectangle(i, i.ButtonBackgroundColour, i.Width-i.Radius, i.Height-i.Padding, i.Pos_x+i.Radius/2, i.Pos_y+i.Padding/2, i.Pos_z, i.WindowHeight_Initial, i.WindowWidth_Initial), gl.TRIANGLES }
	i.VAO[7] = intf.Drawing_Struct{ intf.GenerateRectangle(i, i.ButtonBackgroundColour, i.Width-i.Padding, i.Height-i.Radius, i.Pos_x+i.Padding/2, i.Pos_y+i.Radius/2, i.Pos_z, i.WindowHeight_Initial, i.WindowWidth_Initial), gl.TRIANGLES }
	// Text cursor and selection highlight	
	i.VAO[8] = intf.Drawing_Struct{ intf.GenerateRectangle(i, i.ButtonHighlightColour, i.Highlight.GetWidth(), i.Input_Text.Text.Height(), i.Highlight.GetLeftPosFromIndex(), i.Pos_y + (i.Height - i.Input_Text.Text.Height())/2, i.Pos_z, i.WindowHeight_Initial, i.WindowWidth_Initial), gl.TRIANGLES }
	i.VAO[9] = intf.Drawing_Struct{ intf.GenerateRectangle(i, i.ButtonBorderColour, 2, i.Input_Text.Text.Height(), i.Cursor.GetCursorPosFromIndex(), i.Pos_y + (i.Height - i.Input_Text.Text.Height())/2, i.WindowHeight_Initial, i.Pos_z, i.WindowWidth_Initial), gl.TRIANGLES }
	// Background corner circles
	i.VAO[10] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBackgroundColour, i.Pos_x+i.Radius/2+i.Padding/2, i.Pos_y+i.Radius/2+i.Padding/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	i.VAO[11] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBackgroundColour, i.Pos_x+i.Width-i.Radius/2-i.Padding/2, i.Pos_y+i.Radius/2+i.Padding/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	i.VAO[12] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBackgroundColour, i.Pos_x+i.Radius/2+i.Padding/2, i.Pos_y+i.Height-i.Radius/2-i.Padding/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	i.VAO[13] = intf.Drawing_Struct{ intf.GenerateCirlce(i, i.ButtonBackgroundColour, i.Pos_x+i.Width-i.Radius/2-i.Padding/2, i.Pos_y+i.Height-i.Radius/2-i.Padding/2, i.Pos_z, (i.Radius)/2, i.WindowWidth_Initial, i.WindowHeight_Initial, 128), gl.TRIANGLE_FAN }
	// Text
	i.Input_Text = CreateText(i.LabelText, i.Input_Font.Font)
	i.Input_Text.Text.SetPosition(mgl32.Vec3{i.Pos_x + i.Input_Text.Text.Width()/2 + i.Padding*2, i.Pos_y + i.Height/2, i.Pos_z})
}

func (i *Text_Input_Struct) GetCursorIndex() int {
	return int(i.Cursor.CursorIndex)
}

func (i *Text_Input_Struct) SetCursorIndex(cursor int) {
	i.Cursor.CursorIndex = cursor
	i.GeneratePolygons()
}

func (i *Text_Input_Struct) SetPos(x, y float32) { 
	i.Pos_x = x 
	i.Pos_y = y 
	i.Input_Text.Text.SetPosition(mgl32.Vec3{ i.Pos_x + i.Input_Text.Text.Width()/2 + i.Padding*2, i.Pos_y + i.Height/2, i.Pos_z})
	i.GeneratePolygons()
	i.SetClickableBounds(i.Pos_x, i.Pos_x+i.Width, i.Pos_y, i.Pos_y+i.Height)
	i.GetCanvas().RefreshCanvas()	
}
func (i *Text_Input_Struct) GetPos() (float32, float32) { return i.Pos_x, i.Pos_y }
func (i *Text_Input_Struct) GetBounds() (float32, float32) { return i.Width, i.Height }

func (i *Text_Input_Struct) SetDisplayText(s string) { i.LabelText = s }
func (i *Text_Input_Struct) GetDisplayText() string { return i.LabelText }

func (i *Text_Input_Struct) Draw() {
	// Obtain program that isnt FontProg
	gl.UseProgram(i.GetCanvas().GetPrograms())
	for _, v := range i.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(i.GetXYZ())/3))
	}
	Draw(i.Input_Text)
}

func (i *Text_Input_Struct) Redraw() {
	i.GeneratePolygons()
	// Obtain program that isnt FontProg
	gl.UseProgram(i.GetCanvas().GetPrograms())
	for _, v := range i.GetVAO(){
		gl.BindVertexArray(v.VAO)
		gl.DrawArrays(v.DrawMode, 0, int32(len(i.GetXYZ())/3))
	}
	Draw(i.Input_Text)
}

func (i *Text_Input_Struct) SetClickableBounds(x_min, x_max, y_min, y_max float32) {
	i.Bounds_X_max = int(x_max)
	i.Bounds_X_min = int(x_min)
	i.Bounds_Y_max = int(y_max)
	i.Bounds_Y_min = int(y_min)
}
func (i *Text_Input_Struct) GetClickableBounds() (int, int, int, int, int) {
	return i.Bounds_X_min, i.Bounds_X_max, i.Bounds_Y_min, i.Bounds_Y_max, int(i.Pos_z)
}

func (i *Text_Input_Struct) Move(delta_x, delta_y float32) {
	i.Pos_y = i.Pos_y + delta_y
	i.Pos_x = i.Pos_x + delta_x
	i.GeneratePolygons()
	i.SetClickableBounds(i.Pos_x, i.Pos_x+i.Width, i.Pos_y, i.Pos_y+i.Height)
} 
func (i *Text_Input_Struct) SetName(s string) { i.TextInputName = s }
func (i *Text_Input_Struct) GetName() string { return i.TextInputName }
func (i *Text_Input_Struct) SetDIM(dim []float32){ i.DIM = dim }
func (i *Text_Input_Struct) GetDIM() []float32 { return i.DIM }
func (i *Text_Input_Struct) SetXYZ(xyz []float32){ i.XYZ = xyz }
func (i *Text_Input_Struct) GetXYZ() []float32 { return i.XYZ }
func (i *Text_Input_Struct) SetRGB(rgb []float32){ i.RGB = rgb }
func (i *Text_Input_Struct) GetRGB() []float32 { return i.RGB }
func (i *Text_Input_Struct) SetVAO(vao []intf.Drawing_Struct){ i.VAO = vao }
func (i *Text_Input_Struct) GetVAO() []intf.Drawing_Struct { return i.VAO }
func (i *Text_Input_Struct) GetCanvas() intf.Canvas_Interface { return i.Canvas }

func (i *Text_Input_Struct) SetClickable(click bool) { i.Clickable = click }
func (i *Text_Input_Struct) GetClickable() bool { return i.Clickable }

func (i *Text_Input_Struct) SetClickEvent(ev intf.Event_Interface){ i.ClickEvent = ev }

func (i *Text_Input_Struct) AnimateTrigger(pressed int, alive *bool){
	if i.Clickable == true && pressed == i.ClickTrigger{
		*alive = true
		defer i.Kill(alive)
		// No animation currently.
		i.ButtonBackgroundColour = i.ButtonBackgroundColour_Selected
		i.GeneratePolygons()
	} 
}
func (i *Text_Input_Struct) TriggerClickEvent(pressed int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if i.ClickEvent.GetMethod() != nil && i.Clickable == true && pressed == i.ClickTrigger{
		i.GetCanvas().GetEventHandler().AddEventToEventQueue(i.ClickEvent.GetName())
		i.Cursor.SetCursorIndexFromClickLocation(pos_x)
		if mod_key == glfw.ModShift {
			i.HighlightCharacter(pos_x)
		} 
  	}
}
func (i *Text_Input_Struct) Kill(alive *bool){
	*alive = false
}

func (i *Text_Input_Struct) SetClickTrigger(trigger int) { i.ClickTrigger = trigger }
func (i *Text_Input_Struct) GetClickTrigger() int { return i.ClickTrigger }

func (i *Text_Input_Struct) GetKeyListener() *list.KeyHandler_Struct { return i.KeyListener }
func (i *Text_Input_Struct) SetKeyListener(kl *list.KeyHandler_Struct) { i.KeyListener = kl }

func (i *Text_Input_Struct) selectedInputField(params intf.Paramaters_Interface) {
	i.GetKeyListener().SetFocus(i)
}

func (i *Text_Input_Struct) HighlightCharacter(pos_x float32) {
	if i.Highlight.LeftIndex == -1 {
		i.Highlight.SetLeftIndexFromClickLocation(float32(pos_x))
	} else if i.Highlight.RightIndex == -1 && pos_x > float32(i.Highlight.LeftPosX) {
		i.Highlight.SetRightIndexFromClickLocation(float32(pos_x))
	} else {
		if pos_x < float32(i.Highlight.LeftPosX){
			i.Highlight.SetLeftIndexFromClickLocation(float32(pos_x))
			fmt.Println("Setting left index")
		} else {
			i.Highlight.SetRightIndexFromClickLocation(float32(pos_x))
			fmt.Println("Setting right index")
		}
	}
	i.GeneratePolygons()
}
func (i *Text_Input_Struct) PosToCursorLocation(pos_x float32) int {
	var offset = -i.Width/2 + i.Input_Text.Text.Width()/2
	var index, _ = i.Input_Text.Text.ClickedCharacter(float64(pos_x), float64(offset))
	return index
}

func (i *Text_Input_Struct) IsMaxLength() bool {
	if i.Input_Text.Text.Width() >= i.Width - i.Padding*2 {
		return true
	} else {
		return false
	}
}
func (i *Text_Input_Struct) SetPosZ(z float32){
	if z < 1.0 {
		i.Hidetext()
	} else {
		i.Showtext()
	}
	i.Pos_z = z
}
func (i *Text_Input_Struct) GetPosZ() float32 {
	return i.Pos_z
}
func (i *Text_Input_Struct) Hidetext(){
	i.Input_Text.Text.Hide()
}
func (i *Text_Input_Struct) Showtext(){
	i.Input_Text.Text.Show()
}