package foundation 

import (
	intf 	"gooi/interfaces"
	list 	"gooi/base/listeners"
	event 	"gooi/base/event"
	colors 	"gooi/base/colours"

	fmt 	"fmt"
)

type Cursor struct {
	input 		*Input
	cursorIndex int 
	cursorPosX 	float32
	masterPosX 	*float32

	cursorDrawing *Drawable
}

func (cursor *Cursor) GetCursorPosFromIndex(index int) float32 {
	var position = float32(cursor.input.writing.GetText().Text.CharPosition(index))
	var writing_coordinates = cursor.input.writing.GetPosition()
	position = position + writing_coordinates.GetP1()
	return position
}

func (cursor *Cursor) GeneratePolygons(){
	var coordinate3 = cursor.input.writing.GetPosition()
	cursor.cursorDrawing.ClearPolygons()
	cursor.cursorDrawing.CreateRectangle(
		colors.BLUE, 
		2, cursor.input.writing.GetHeight(), 
		cursor.cursorPosX, coordinate3.GetP2() - cursor.input.writing.GetHeight()/2, coordinate3.GetP3(),
	)
	return
}

type Highlight struct {
	input 		*Input
	leftIndex  	int 
	rightIndex 	int
	leftPosX 	float32
	rightPosX  	float32
	masterPosX 	*float32

	highlightDrawing *Drawable
}

type Input struct {
	canvas 		intf.Canvas_Interface
	keylistener *list.KeyHandler_Struct

	name string

	posX, posY, posZ float32

	cursor *Cursor
	highlight *Highlight
	writing *Writing 
	clickable *Clickable
	clickEvent *event.Event_Struct

	masterStruct intf.Displayable
	masterWidth float32
	masterHeight float32

	slaveWidth float32
	slaveHeight float32

	openGLWindowWidth float32
	openGLWindowHeight float32
}

func NewInput(
	canvas intf.Canvas_Interface, 
	masterStruct intf.Displayable,
	keylistener *list.KeyHandler_Struct,
	name string,
	placeholder string, 
	pos_x, pos_y, pos_z float32, 
	font_path, font_name string,
	font_size int,
	click_trigger int,
) *Input {
	var input = Input{}
	input.canvas = canvas
	input.name = name
	input.openGLWindowWidth = input.canvas.GetWidth()
	input.openGLWindowHeight = input.canvas.GetHeight()
	input.keylistener = keylistener
	input.masterStruct = masterStruct
	input.masterWidth = input.masterStruct.GetWidth()	
	input.masterHeight = input.masterStruct.GetHeight()
	input.posX = pos_x
	input.posY = pos_y
	input.posZ = pos_z
	input.writing = NewWriting(
		canvas,
		masterStruct,
		placeholder,
		pos_x, 
		pos_y, 
		pos_z,
		input.openGLWindowWidth, input.openGLWindowHeight,
		font_path, font_name, 
		font_size,
	)
	input.writing.SetPosition(pos_x + input.writing.GetWidth()/2, pos_y + input.writing.GetHeight()/2 + (input.masterHeight - input.writing.GetHeight())/2, pos_z)
	fmt.Printf("Position set as %v, %v\n", pos_x + input.writing.GetWidth()/2, pos_y + input.writing.GetHeight()/2)
	// creating cursor nested struct
	input.cursor = &Cursor{
		&input, 
		0, 0, &pos_x,
		NewDrawable(
			canvas, 
			masterStruct, 
			input.openGLWindowWidth, input.openGLWindowHeight,
		),
	}
	input.cursor.GeneratePolygons()
	// creating highlight nested struct
	input.highlight = &Highlight {
		&input, 
		0, 0, 0, 0, &pos_x, 
		NewDrawable(
			canvas, 
			masterStruct, 
			input.openGLWindowWidth, input.openGLWindowHeight,
		),
	}
	// adding rectangle to highlight that visualises selection
	input.highlight.highlightDrawing.CreateRectangle(
		colors.LIGHT_BLUE, 
		(input.highlight.rightPosX - input.highlight.leftPosX), 
		input.writing.GetHeight(), 
		input.highlight.leftPosX, pos_y, pos_z,
	)
	// adding the click event to the text field that sets focus in the key handler...
	input.clickEvent = &event.Event_Struct{
		input.selectedInputField, 
		fmt.Sprintf("selectedInputField_%s", input.name), 
		event.NewEventParameter(&input),
	}
	// creating clickable nested struct
	input.clickable = NewClickable(
		input.canvas,
		input.clickEvent,
		click_trigger,
		pos_x, pos_x + input.writing.GetWidth(),
		pos_y + (input.masterHeight - input.writing.GetHeight())/2, 
		pos_y + input.writing.GetHeight() + (input.masterHeight - input.writing.GetHeight())/2,
	)

	input.slaveWidth = input.writing.GetWidth()
	input.slaveHeight = input.writing.GetHeight()

	return &input
}

func (input *Input) GeneratePolygons(){
	input.cursor.GeneratePolygons()
}

func (input *Input) selectedInputField(params intf.Paramaters_Interface) {
	input.keylistener.SetFocus(input)
	var glfwWindow = input.canvas.GetWindow().GetWindow()
	var x, _, _ = input.canvas.GetWindow().GetMouseHandler().GetClickData(glfwWindow)
	input.MoveCursorToClickLocation(x)
	fmt.Printf("Cursor index is %s\n.", input.cursor.cursorIndex)
	fmt.Printf("Cursor location is %s\n.", input.cursor.cursorPosX)
	input.GeneratePolygons()
}

func (input *Input) Draw() { 
	input.writing.Draw() 
	input.cursor.cursorDrawing.Draw()
}

func (input *Input) SetPos(x, y, z float32) {
	input.posX = x
	input.posY = y
	input.posZ = z
	input.writing.SetPosition(x + input.writing.GetWidth()/2, y + input.writing.GetHeight()/2 + (input.masterHeight - input.writing.GetHeight())/2, z)
	input.clickable.SetClickBounds(x, x + input.writing.GetWidth(), y, y + input.writing.GetHeight() + (input.masterHeight - input.writing.GetHeight())/2)
}

func (input *Input) GetPos() (float32, float32, float32) {
	return input.posX, input.posY, input.posZ
}

func (input *Input) SetDisplayText(text string) { 
	input.writing.SetLabel(text) 
	input.SetPos(input.posX, input.posY, input.posZ)
}
func (input *Input) GetDisplayText() string { return input.writing.GetLabel() }

func (input *Input) GetWidth() float32 { return input.slaveWidth }
func (input *Input) GetHeight() float32 { return input.slaveHeight } 

func (input *Input) GetCursorIndex() int { return input.cursor.cursorIndex }
func (input *Input) SetCursorIndex(index int) { 
	input.cursor.cursorIndex = index 
	var position = input.cursor.GetCursorPosFromIndex(index)
	input.cursor.cursorPosX = position
	input.GeneratePolygons()
}

// glText allows identifying the index of the character selected... from this we can calculate the position on the openGL screen
func (input *Input) ClickLocationToCursorIndex(position float32) int {
	//var coordinate = input.writing.GetPosition()	
	var offset = input.writing.GetWidth()/2
	fmt.Printf("OFFSET = %v\n", offset)
	var index, _ = input.writing.text.Text.ClickedCharacter(float64(position), float64(offset))
	fmt.Printf("Cursor CLICKED is %v\n", index)
	return index
}

// the key listener wont allow continued input if the max length is reached. max length is designed by master struct (canvas, or other drawable)
func (input *Input) IsMaxLength() bool {
	// tests if the text is longer than the master struct width - aka the width of an input field with a background rectangle...
	if input.writing.GetWidth() >= input.masterStruct.GetWidth() { // removed padding - might look weird.
		return true
	} 
	return false
}

func (input *Input) MoveCursorToClickLocation(position float32) {
	var index = input.ClickLocationToCursorIndex(position)
	input.cursor.cursorIndex = index
	input.cursor.cursorPosX = input.cursor.GetCursorPosFromIndex(index)
	return
}

func (input *Input) HighlightCharacter(position float32) {
	if input.highlight.leftIndex == -1 {
		input.highlight.leftIndex = input.ClickLocationToCursorIndex(position)
	} else if input.highlight.rightIndex == -1 && position > float32(input.highlight.leftPosX) {
		input.highlight.rightIndex = input.ClickLocationToCursorIndex(position)
	} else {
		if position < float32(input.highlight.leftPosX){
			input.highlight.leftIndex = input.ClickLocationToCursorIndex(position)
		} else {
			input.highlight.rightIndex = input.ClickLocationToCursorIndex(position)
		}
	}
	return
}

func (input *Input) GetClickable() *Clickable { return input.clickable }
func (input *Input) GetWriting() *Writing { return input.writing }