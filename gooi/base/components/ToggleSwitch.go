package components 
/**
 * Updated 11/11/2023.
 * ToggleSwitch implements -> Drawable -> Component -> Clickable
 **/
import (
	glfw 		"github.com/go-gl/glfw/v3.2/glfw"
	intf 		"gooi/interfaces"
	event 		"gooi/base/event"
	cons 		"gooi/base/constants"
	colours 	"gooi/base/colours"
	foundations "gooi/base/components/foundation"

	fmt 		"fmt"
)
type ToggleSwitch_Struct struct {
	canvas 			*Canvas_Struct
	masterStruct 	intf.Displayable
	name 			string

	posX, posY, posZ float32
	width, height float32
	togglePos float32

	openGLWindowWidth, openGLWindowHeight float32

	masterWidth, masterHeight float32
	slaveWidth, slaveHeight float32

	fontName, fontPath string
	fontSize int

	toggleEvent *event.Event_Struct
	toggleState bool

	toggleColour [3]float32

	writing *foundations.Writing
	drawable *foundations.Drawable
	clickable *foundations.Clickable
	animation *foundations.Animation	
}

func CreateToggle(
	canvas *Canvas_Struct, 
	masterStruct intf.Displayable,
	name string,
	width, height, pos_x, pos_y, pos_z float32, 
	font_name, font_path string,
	font_size int, 
	toggle_event *event.Event_Struct,
	) *ToggleSwitch_Struct {
		var t = ToggleSwitch_Struct{
			canvas,
			masterStruct,
			name, 
			pos_x, pos_y, pos_z, 
			width, height,
			pos_x,
			canvas.GetWidth(), canvas.GetHeight(),
			masterStruct.GetWidth(), masterStruct.GetHeight(),
			width, height,
			font_name, font_path, font_size,
			toggle_event,
			false,
			colours.RED,
			nil, nil, nil, nil,
		}
		t.writing = foundations.NewWriting(
			canvas,
			masterStruct,
			name, 
			t.posX + t.width + height*1.5 + 10, t.posY + height/2, t.posZ,
			canvas.GetWidth(), canvas.GetHeight(),
			font_path, font_name, font_size,
		)
		t.drawable = foundations.NewDrawable(
			canvas,
			masterStruct,
			canvas.GetWidth(), canvas.GetHeight(),
		)
		t.clickable = foundations.NewClickable(
			canvas, 
			toggle_event, 
			cons.MOUSE_PRESSED, 
			t.posX, t.posX + width, 
			t.posY, t.posY + height,
		)
		t.animation = foundations.NewAnimation(
			[]func(){
				t.animationf1,
				t.animationf2,
				t.animationf3,
				t.animationf4,
			}, 10,
		)

		t.slaveWidth = t.width + t.height + 10 + t.writing.GetWidth()
		t.slaveHeight = t.height
		if t.height < t.writing.GetHeight() {
			t.slaveHeight = t.writing.GetHeight()
		}

		t.GeneratePolygons()
		t.clickable.SetClickBounds(t.posX, t.posX + t.slaveWidth, t.posY, t.posY + t.slaveHeight)

		canvas.GetWindow().GetMouseHandler().RegisterClickableToHandler(&t)

		return &t	
}

func (t *ToggleSwitch_Struct) animationf1() {
	if t.toggleState {
		t.toggleColour = colours.RED
		t.togglePos -= t.width/4
	} else {
		t.toggleColour = colours.GREEN
		t.togglePos += t.width/4
	}
	t.Redraw()
	t.canvas.RefreshCanvas()
}

func (t *ToggleSwitch_Struct) animationf2() {
	if t.toggleState {
		t.togglePos -= t.width/4
	} else {
		t.togglePos += t.width/4
	}
	t.Redraw()
	t.canvas.RefreshCanvas()
}

func (t *ToggleSwitch_Struct) animationf3() {
	if t.toggleState {
		t.togglePos -= t.width/4
	} else {
		t.togglePos += t.width/4
	}
	t.Redraw()
	t.canvas.RefreshCanvas()
}

func (t *ToggleSwitch_Struct) animationf4() {
	if t.toggleState {
		t.togglePos -= t.width/4
	} else {
		t.togglePos += t.width/4
	}
	t.toggleState = !t.toggleState
	t.Redraw()
	t.canvas.RefreshCanvas()
}

func (t *ToggleSwitch_Struct) GeneratePolygons(){
	fmt.Println("Generating toggle polygons")
	var border float32 = 4
	t.drawable.ClearPolygons()
	// Border shape
	t.drawable.CreateRectangle(colours.DARK_GRAY, 
		t.width, 
		t.height, 
		t.posX + t.height/2, 
		t.posY, 
		t.posZ)
	t.drawable.CreateCircle(colours.DARK_GRAY, 
		t.height/2, 
		t.posX + t.height/2, 
		t.posY + t.height/2, 
		t.posZ)
	t.drawable.CreateCircle(colours.DARK_GRAY, 
		t.height/2, 
		t.posX + t.width + t.height/2, 
		t.posY + t.height/2, 
		t.posZ)
	// Background shape
	t.drawable.CreateRectangle(colours.WHITE, 
		t.width 					+ border, 
		t.height					- border, 
		t.posX + t.height/2			- border/2, 
		t.posY 						+ border/2, 
		t.posZ)
	t.drawable.CreateCircle(colours.WHITE, 
		t.height/2					- border/2, 
		t.posX + t.height/2			+ border/8, 
		t.posY + t.height/2 		+ border/8, 
		t.posZ)
	t.drawable.CreateCircle(colours.WHITE, 
		t.height/2						- border/2, 
		t.posX + t.width + t.height/2	- border/8, 
		t.posY + t.height/2 			+ border/8, 
		t.posZ)
	// Coloured shape
	border = 6
	t.drawable.CreateRectangle(t.toggleColour, 
		t.width 					+ border, 
		t.height					- border, 
		t.posX + t.height/2			- border/2, 
		t.posY 						+ border/2, 
		t.posZ)
	t.drawable.CreateCircle(t.toggleColour, 
		t.height/2					- border/2, 
		t.posX + t.height/2			+ border/8, 
		t.posY + t.height/2 		+ border/16, 
		t.posZ)
	t.drawable.CreateCircle(t.toggleColour, 
		t.height/2						- border/2, 
		t.posX + t.width + t.height/2	- border/8, 
		t.posY + t.height/2 			+ border/16, 
		t.posZ)
	// Toggle shape
	t.drawable.CreateCircle(colours.DARK_GRAY, 
		t.height/2					- border/2, 
		t.togglePos + t.height/2	+ border/8, 
		t.posY + t.height/2 		+ border/16, 
		t.posZ)
	
}

func (t *ToggleSwitch_Struct) Draw() {
	t.drawable.Draw()
	t.writing.Draw()
}

func (t *ToggleSwitch_Struct) Redraw() {
	t.GeneratePolygons()
	t.drawable.Draw()
	t.writing.Draw()
}

func (t *ToggleSwitch_Struct) Click(alive *bool, pressAction int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if pressAction == t.clickable.GetClickTrigger() {
		t.clickable.TriggerClickEvent(alive, pressAction, pos_x, pos_y, mod_key)
		t.animation.RunAnimation(alive)
	}	
}
func (t *ToggleSwitch_Struct) SetPos(x, y, z float32) { 
	t.posX = x
	t.posY = y 
	t.posZ = z
	t.GeneratePolygons()
	t.writing.SetPosition(t.posX + t.width + t.height + 10, t.posY + t.height/2, t.posZ,)
	t.clickable.SetClickBounds(t.posX, t.posX + t.slaveWidth, t.posY, t.posY + t.slaveHeight)
}
func (t *ToggleSwitch_Struct) GetPos() (float32, float32, float32) { return t.posX, t.posY, t.posZ }

func (t *ToggleSwitch_Struct) SetName(s string) { t.name = s }
func (t *ToggleSwitch_Struct) GetName() string { return t.name }

func (t *ToggleSwitch_Struct) GetClickableBounds() (float32, float32, float32, float32, float32) {
	var x, y = t.clickable.GetClickBounds()
	return x.GetP1(), x.GetP2(), y.GetP1(), y.GetP2(), t.posZ
}

func (t *ToggleSwitch_Struct) GetWidth() float32 { return t.slaveWidth }
func (t *ToggleSwitch_Struct) GetHeight() float32 { return t.slaveHeight }

func (t *ToggleSwitch_Struct) GetMasterStruct() intf.Displayable { return t.masterStruct }
func (t *ToggleSwitch_Struct) SetMasterStruct(master intf.Displayable) { t.masterStruct = master }