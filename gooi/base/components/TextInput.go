package components 

import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
	
	intf 	"gooi/interfaces"
	list 	"gooi/base/listeners"
	cons 	"gooi/base/constants"
	foundations "gooi/base/components/foundation"
	
	fmt 	"fmt"
)

type TextInput struct {
	canvas 			intf.Canvas_Interface
	masterStruct 	intf.Displayable
	name 			string
	placeholder 	string

	inputbox 		*Rectangle_Struct

	keylistener 	*list.KeyHandler_Struct
	input 			*foundations.Input
	prompt 			*foundations.Writing

	posX, posY, posZ float32

	masterWidth, masterHeight float32
	slaveWidth, slaveHeight float32
}

func NewTextInput(
	canvas 					intf.Canvas_Interface,
	masterStruct			intf.Displayable,
	name 					string,
	placeholder 			string,
	keylistener 			*list.KeyHandler_Struct,
	width, height  			float32,
	pos_x, pos_y, pos_z 	float32,
	radius 					float32,
	colour 					[3]float32,
	font_name, font_path 	string,
	font_size 				int,
) *TextInput {
	
	var textInput = TextInput{}
	
	textInput.inputbox = NewRectangle(
		canvas, 
		masterStruct,
		"TextInput_Background",
		width, height,
		pos_x, pos_y, pos_z, 
		radius,
		colour,
		cons.NO_FILL,
		cons.MATCH_MASTER_POSITION,
	)

	textInput.prompt = foundations.NewWriting(
		canvas, 
		masterStruct,
		name,
		pos_x, pos_y, pos_z, 
		canvas.GetWidth(), canvas.GetHeight(),
		font_path, font_name, font_size,
	)

	textInput.canvas = canvas
	textInput.masterStruct = textInput.inputbox
	textInput.name = name
	textInput.placeholder = placeholder
	textInput.keylistener = keylistener
	textInput.posX = pos_x
	textInput.posY = pos_y
	textInput.posZ = pos_z

	textInput.input = foundations.NewInput(
		textInput.canvas, 
		textInput.masterStruct,
		textInput.keylistener,
		textInput.name,
		textInput.placeholder, 
		textInput.posX, textInput.posY, textInput.posZ, 
		font_path, font_name, font_size, 
		cons.MOUSE_PRESSED,
	)

	textInput.masterWidth 	= masterStruct.GetWidth()
	textInput.masterHeight 	= masterStruct.GetHeight()
	textInput.slaveWidth 	= textInput.input.GetWidth()
	textInput.slaveHeight 	= textInput.input.GetHeight()

	canvas.GetWindow().GetMouseHandler().RegisterClickableToHandler(&textInput)

	textInput.GeneratePolygons()

	return &textInput
}

func (t *TextInput) GeneratePolygons(){
	fmt.Println("Generating polygons")
	t.inputbox.GeneratePolygons()
	t.canvas.RefreshCanvas()
	return
}

func (t *TextInput) GetWidth() float32 { return t.input.GetWidth() }
func (t *TextInput) GetHeight() float32 { return t.input.GetHeight() }

func (t *TextInput) Draw() { 
	t.prompt.Draw()
	t.inputbox.Draw()
	t.input.Draw() 
}
func (t *TextInput) Redraw() { 
	t.GeneratePolygons()
	t.prompt.Draw()
	t.inputbox.Draw()
	t.input.Draw() 
}

func (t *TextInput) SetPos(x, y, z float32) {
	var padding float32 = 10
	t.prompt.SetPosition(x + t.prompt.GetWidth()/2, y + t.prompt.GetHeight()/2, z)
	t.inputbox.SetPos(x + t.prompt.GetWidth() + padding, y, z)
	t.input.SetPos(x + t.prompt.GetWidth() + padding, y, z)
}

func (t *TextInput) GetPos() (float32, float32, float32) {
	var x, y, z = t.input.GetPos()
	return x, y, z
}

func (t *TextInput) Click(alive *bool, pressAction int, pos_x, pos_y float32, mod_key glfw.ModifierKey){
	if pressAction == t.input.GetClickable().GetClickTrigger() {
		t.input.GetClickable().TriggerClickEvent(alive, pressAction, pos_x, pos_y, mod_key)
		//b.animation.RunAnimation(alive)
	}
}

func (t *TextInput) GetClickableBounds() (float32, float32, float32, float32, float32) {
	var x, y = t.input.GetClickable().GetClickBounds()
	return x.GetP1(), x.GetP2(), y.GetP1(), y.GetP2(), t.posZ
}

func (t *TextInput) GetMasterStruct() intf.Displayable { return t.masterStruct }
func (t *TextInput) SetMasterStruct(master intf.Displayable) { t.masterStruct = master }