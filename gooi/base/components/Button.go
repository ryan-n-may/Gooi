package components 
/**
 * Updated 11/11/2023.
 * Button Component.
 **/
import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"

	intf 		"gooi/interfaces"
	event 		"gooi/base/event"
	cons 		"gooi/base/constants"
	colours 	"gooi/base/colours"
	foundation 	"gooi/base/components/foundation"

	time 	"time"
	fmt 	"fmt"
)

type Button struct {
	canvas 			intf.Canvas_Interface
	masterStruct 	intf.Displayable
	componentName	string

	radius float32
	posX float32
	posY float32
	posZ float32

	masterWidth float32
	masterHeight float32

	slaveWidth float32 
	slaveHeight float32

	openGLWindowWidth	float32
	openGLWindowHeight 	float32

	animationFunctions []func()
	
	clickable 		*foundation.Clickable
	drawable 		*foundation.Drawable
	writing 		*foundation.Writing
	animation 		*foundation.Animation	

	buttonBodyColour [3]float32
}

func NewButton(
	canvas 					intf.Canvas_Interface, 
	masterStruct 			intf.Displayable,
	name 					string,
	width, height 			float32,
	radius					float32,
	pos_x, pos_y, pos_z		float32, 
	font_name, font_path 	string,
	font_size 				int, 
	button_event 			*event.Event_Struct,
	animation_time 			time.Duration,
) *Button {
	var b = Button{
		canvas,
		masterStruct,
		name,
		radius,
		pos_x, pos_y, pos_z,
		masterStruct.GetWidth(), masterStruct.GetHeight(), // master struct dimensions
		width, height, // current slave dimensions
		canvas.GetWidth(), canvas.GetHeight(), // openGL window dimensions
		[]func(){},
		foundation.NewClickable(canvas, button_event, cons.MOUSE_PRESSED, pos_x, pos_x+width, pos_y, pos_y+width),
		foundation.NewDrawable(canvas, masterStruct, canvas.GetWidth(), canvas.GetHeight()),
		foundation.NewWriting(canvas, masterStruct, name, pos_x, pos_y, pos_z, canvas.GetWidth(), canvas.GetHeight(), font_path, font_name, font_size), 
		nil,
		colours.LIGHT_GRAY,
	}
	b.animationFunctions = []func(){b.animationf1, b.animationf2}
	b.animation = foundation.NewAnimation(b.animationFunctions, animation_time)
	b.writing.SetPosition(pos_x + width/2, pos_y + height/2, b.posZ)
	b.GeneratePolygons()
	canvas.GetWindow().GetMouseHandler().RegisterClickableToHandler(&b)
   	return &b
}

func (b *Button) animationf1() {
	fmt.Println("Animation1")
	b.buttonBodyColour = colours.DARK_GRAY
	b.GeneratePolygons()
	b.GetCanvas().RefreshCanvas()	
}
func (b *Button) animationf2() {
	fmt.Println("Animation2")
	b.buttonBodyColour = colours.LIGHT_GRAY
	b.GeneratePolygons()
	b.GetCanvas().RefreshCanvas()	
}

// GeneratePolygons()
// Generates the VAO array of the polygons used to draw the button. 
// Stores the VAO in intf.Drawing_Struct alongisde the drawing mode (gl.TRIANGLE or gl.TRIANGLE_FAN)
func (b *Button) GeneratePolygons(){
	b.drawable.ClearPolygons()
	var border float32 = 2
	/** Border **/
	b.drawable.CreateRoundedRectangle(colours.DARK_BLUE, b.slaveWidth, b.slaveHeight, b.posX, b.posY, b.posZ, b.radius)
	/** Background Rectangles **/
	b.drawable.CreateRoundedRectangle(colours.WHITE, b.slaveWidth-border/2, b.slaveHeight-border/2, b.posX+border/2, b.posY+border/2, b.posZ, b.radius-border/2)
	/** Button Body **/
	b.drawable.CreateRoundedRectangle(b.buttonBodyColour, b.slaveWidth-border, b.slaveHeight-border, b.posX+border, b.posY+border, b.posZ, b.radius-border)
}
// Draw()
// This method draws the VAO array to gl using the canvas program.
func (b *Button) Draw() {
	// Obtain program that isnt gltext font drawing program
	b.drawable.Draw()
	// Draw button text (Uses FontProg program) and modified glText
	b.writing.Draw()
}
func (b *Button) Redraw() {
	b.GeneratePolygons()
	b.drawable.Draw()
	// Draw button text (Uses FontProg program) and modified glText
	b.writing.Draw()
}

func (b *Button) SetPos(x, y, z float32) { 
	b.posX = x
	b.posY = y
	b.posZ = z
	// relocate writing
	b.writing.SetPosition(x + b.GetWidth()/2, y + b.GetHeight()/2, b.posZ)
	// relocate clickable bounds
	b.clickable.SetClickBounds(x, x + b.slaveWidth, y, y + b.slaveHeight)
	// refresh drawables in foundation
	b.Redraw()
}
func (b *Button) GetPos() (float32, float32, float32) { return b.posX, b.posY, b.posZ }

func (b *Button) GetBounds() (foundation.Coordinate, foundation.Coordinate) { return b.clickable.GetClickBounds() }

func (b *Button) SetButtonText(s string) { b.writing.SetLabel(s) }
func (b *Button) GetButtonText() string { return b.writing.GetLabel() }

func (b *Button) SetName(s string) { b.componentName = s }
func (b *Button) GetName() string { return b.componentName }

func (b *Button) SetClickEvent(ev *event.Event_Struct){ b.clickable.SetClickEvent(ev) }
func (b *Button) SetClickTrigger(t int) { b.clickable.SetClickTrigger(t) }
func (b *Button) GetClickTrigger() int {	return b.clickable.GetClickTrigger() }

func (b *Button) GetCanvas() intf.Canvas_Interface { return b.canvas }

func (b *Button) GetClickableBounds() (float32, float32, float32, float32, float32) {
	var x, y = b.clickable.GetClickBounds()
	return x.GetP1(), x.GetP2(), y.GetP1(), y.GetP2(), b.posZ
}

func (b *Button) GetWidth() float32 {
	return b.slaveWidth
}
func (b *Button) GetHeight() float32 {
	return b.slaveHeight
}

func (b *Button) GetClickable() *foundation.Clickable {
	return b.clickable
}

func (b *Button) Click(alive *bool, pressAction int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if pressAction == b.clickable.GetClickTrigger() {
		b.clickable.TriggerClickEvent(alive, pressAction, pos_x, pos_y, mod_key)
		b.animation.RunAnimation(alive)
	}
}

func (b *Button) GetMasterStruct() intf.Displayable { return b.masterStruct }
func (b *Button) SetMasterStruct(master intf.Displayable) { b.masterStruct = master }