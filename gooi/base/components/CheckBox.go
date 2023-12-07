package components 
/**
 * Updated 11/11/2023.
 * Checkbox Component.
 * Implements Drawable -> Component -> Clickable.
 **/
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"

	intf 	"gooi/interfaces"
	event 	"gooi/base/event"
	cons 	"gooi/base/constants"
	colours "gooi/base/colours"
	foundations "gooi/base/components/foundation"

	fmt "fmt"
)
type CheckBox_Struct struct {
	canvas intf.Canvas_Interface
	masterStruct intf.Displayable

	name string

	radius 			float32
	posX, posY, posZ float32

	openGLWindowWidth, openGLWindowHeight float32

	filledColour [3]float32
	filledState bool

	masterWidth, masterHeight float32
	slaveWidth, slaveHeight float32

	clickable *foundations.Clickable
	animation *foundations.Animation
	drawable *foundations.Drawable
	writing *foundations.Writing

	animationFunctions []func()
}

func NewCheckbox(
	canvas 					intf.Canvas_Interface, 
	masterStruct 			intf.Displayable,
	name 					string,

	radius, pos_x, pos_y, pos_z float32, 
	
	check_event 			*event.Event_Struct,
	
	font_name 				string,
	font_path 				string,
	font_size 				int,
	) *CheckBox_Struct {
	// Creating buttons 
	var b = CheckBox_Struct{
		canvas, 
		masterStruct,
		name, 
		radius, pos_x, pos_y, pos_z,
		canvas.GetWidth(), canvas.GetHeight(),
		colours.WHITE, false,
		masterStruct.GetWidth(), masterStruct.GetHeight(),
		0, 0, // slave dimensions initially set to 0
		nil, nil, nil, nil, // foundations not yet initialised.
		[]func(){},
	}

	b.animationFunctions = []func(){b.animationf1}

	b.writing = foundations.NewWriting(
		canvas, 
		masterStruct,
		name,
		pos_x, pos_y, pos_z, 
		canvas.GetWidth(), canvas.GetHeight(),
		font_path, font_name, font_size,
	)

	b.clickable = foundations.NewClickable(
		canvas,
		check_event,
		cons.MOUSE_PRESSED,
		pos_x, pos_x + b.radius*2,
		pos_y, pos_y + b.radius*2,
	)

	b.drawable = foundations.NewDrawable(
		canvas, 
		masterStruct, 
		canvas.GetWidth(), canvas.GetHeight(),
	)

	b.animation = foundations.NewAnimation(b.animationFunctions, 0)

	b.SetPos(pos_x, pos_y, pos_z)

	if b.writing.GetHeight() > radius*2 {
		b.slaveHeight = b.writing.GetHeight() 
	} else {
		b.slaveHeight = radius*2
	}
	b.slaveWidth = b.writing.GetWidth() + radius*3

	b.writing.SetPosition(b.posX + b.writing.GetWidth()/2 + b.radius*2, b.posY + b.slaveHeight/2 , b.posZ)

	b.GeneratePolygons()

	canvas.GetWindow().GetMouseHandler().RegisterClickableToHandler(&b)
	
   	return &b
}

func (b *CheckBox_Struct) animationf1() {
	fmt.Println("Running checkbox animation")
	if b.filledState {
		b.filledColour = colours.WHITE
	} else {
		b.filledColour = colours.BLUE
	}
	b.filledState = !b.filledState
	b.Redraw()
}

func (b *CheckBox_Struct) GeneratePolygons(){
	var x = b.posX + b.radius
	var y = b.posY + b.radius
	b.drawable.ClearPolygons()
	var border float32 = 1.0
	b.drawable.CreateCircle(colours.DARK_GRAY, b.radius, x, y, b.posZ)
	b.drawable.CreateCircle(colours.WHITE, b.radius-border*2, x+border/4, y+border/4, b.posZ)
	b.drawable.CreateCircle(b.filledColour, b.radius-border*4, x+border/2, y+border/2, b.posZ)
}
func (b *CheckBox_Struct) Draw() {
	b.drawable.Draw()
	b.writing.Draw()
}
func (b *CheckBox_Struct) Redraw() {
	b.GeneratePolygons()
	b.drawable.Draw()
	b.writing.Draw()
}
func (b *CheckBox_Struct) ToggleFilledState(state bool){
	if state {
		b.filledColour = colours.BLUE
	} else {
		b.filledColour = colours.WHITE
	}
	b.filledState = state
	b.GeneratePolygons()
}
func (b *CheckBox_Struct) GetFilledState() bool { return b.filledState }

// Triggers a click event if the checkbox is clickable, pressed is correct (MOUSE_PRESS, MOUSE_RELEASE)
// and the click event doesnt have a nil method.
func (b *CheckBox_Struct) Click(alive *bool, pressAction int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if pressAction == b.clickable.GetClickTrigger() {
		b.clickable.TriggerClickEvent(alive, pressAction, pos_x, pos_y, mod_key)
		b.animation.RunAnimation(alive)
	}
}

func (b *CheckBox_Struct) SetPos(x, y, z float32) { 
	b.posX = x 
	b.posY = y 
	b.posZ = z
	b.writing.SetPosition(b.posX + b.writing.GetWidth()/2 + b.radius*3, b.posY + b.writing.GetHeight()/2 , b.posZ)
	b.GeneratePolygons()
	b.clickable.SetClickBounds(b.posX, b.posX+b.radius*2, b.posY, b.posY+b.radius*2)
	b.Redraw()
}
func (b *CheckBox_Struct) GetPos() (float32, float32, float32) { return b.posX, b.posY, b.posZ }

func (b *CheckBox_Struct) SetName(s string) { b.name = s }
func (b *CheckBox_Struct) GetName() string { return b.name }

func (b *CheckBox_Struct) GetCanvas() intf.Canvas_Interface { return b.canvas }

func (b *CheckBox_Struct) GetClickableBounds() (float32, float32, float32, float32, float32){
	var x, y = b.clickable.GetClickBounds()
	return x.GetP1(), x.GetP2(), y.GetP1(), y.GetP2(), b.posZ
}

func (b *CheckBox_Struct) GetHeight() float32 { return b.slaveHeight }
func (b *CheckBox_Struct) GetWidth() float32 { return b.slaveWidth }

func (b *CheckBox_Struct) GetMasterStruct() intf.Displayable { return b.masterStruct }
func (b *CheckBox_Struct) SetMasterStruct(master intf.Displayable) { b.masterStruct = master }