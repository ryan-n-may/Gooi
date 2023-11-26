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
) {
	var input = Input{}
	input.canvas = canvas
	input.name = name
	input.openGLWindowWidth = input.canvas.GetWidth()
	input.openGLWindowHeight = input.canvas.GetHeight()
	input.keylistener = keylistener
	input.writing = NewWriting(
		canvas,
		masterStruct,
		placeholder,
		pos_x, pos_y, pos_z,
		input.openGLWindowWidth, input.openGLWindowHeight,
		font_path, font_name, 
		font_size,
	)
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
	// adding rectangle to cursor nested struct that visualises stroke
	input.cursor.cursorDrawing.CreateRectangle(
		colors.BLUE, 
		2, input.writing.GetHeight(), 
		input.cursor.cursorPosX, pos_y, pos_z,
	)
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
		pos_y, pos_y + input.writing.GetHeight(),
	)
	input.slaveWidth = input.writing.GetWidth()
	input.slaveHeight = input.writing.GetHeight()
	input.masterWidth = input.masterStruct.GetWidth()
	input.masterHeight = input.masterStruct.GetHeight()
}

func (input *Input) selectedInputField(params intf.Paramaters_Interface) {
	input.keylistener.SetFocus(input)
}

func (input *Input) SetDisplayText(text string) { input.writing.SetLabel(text) }
func (input *Input) GetDisplayText() string { return input.writing.GetLabel() }

func (input *Input) GetWidth() float32 { return input.slaveWidth }
func (input *Input) GetHeight() float32 { return input.slaveHeight } 

func (input *Input) GetCursorIndex() int { return input.cursor.cursorIndex }
func (input *Input) SetCursorIndex(index int) { input.cursor.cursorIndex = index }

// glText allows identifying the index of the character selected... from this we can calculate the position on the openGL screen
func (input *Input) ClickIndexToCursorLocation(position float32) int {
	var offset = -input.GetWidth()/2 + input.writing.GetWidth()/2 			// might be wrong now ive changed the code structure.
	var index, _ = input.writing.text.Text.ClickedCharacter(float64(position), float64(offset))
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

func (input *Input) HighlightCharacter(position float32) {
	if input.highlight.leftIndex == -1 {
		input.highlight.leftIndex = input.ClickIndexToCursorLocation(position)
	} else if input.highlight.rightIndex == -1 && position > float32(input.highlight.leftPosX) {
		input.highlight.rightIndex = input.ClickIndexToCursorLocation(position)
	} else {
		if position < float32(input.highlight.leftPosX){
			input.highlight.leftIndex = input.ClickIndexToCursorLocation(position)
		} else {
			input.highlight.rightIndex = input.ClickIndexToCursorLocation(position)
		}
	}
}