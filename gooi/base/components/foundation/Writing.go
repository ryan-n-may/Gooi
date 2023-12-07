package foundation 

import (
	mgl32   "github.com/go-gl/mathgl/mgl32"

	intf 	"gooi/interfaces"
	font 	"gooi/base/font"

	fmt 	"fmt"
)

type Writing struct {
	canvas intf.Canvas_Interface

	label string

	textPosition Coordinate3

	masterStruct intf.Displayable
	masterWidth float32
	masterHeight float32

	slaveWidth float32
	slaveHeight	float32

	openGLWindowHeight float32
	openGLWindowWidth float32

	font *font.Font_Struct
	text *font.Text_Struct
}

func NewWriting(canvas intf.Canvas_Interface, 
	masterStruct intf.Displayable, 
	label string, 
	pos_x, pos_y, pos_z, openGLWindowWidth, openGLWindowHeight float32, 
	font_path, font_name string, 
	font_size int) *Writing {
	var writing = Writing{
		canvas, 
		label, 
		Coordinate3{pos_x, pos_y, pos_z},
		masterStruct, 
		masterStruct.GetWidth(),
		masterStruct.GetHeight(),
		0, 0, 
		openGLWindowHeight, 
		openGLWindowWidth,
		font.ReadFontFile(font_name, font_path, font_size),
		nil,
	}
	writing.font.Font.ResizeWindow(openGLWindowWidth, openGLWindowHeight)
	writing.text = font.CreateText(writing.label, writing.font.Font)
	writing.text.Text.SetPosition(mgl32.Vec3{
		writing.textPosition.p1,
		writing.textPosition.p2,
		writing.textPosition.p3,
	})
	writing.slaveWidth = writing.text.Text.Width()
	writing.slaveHeight = writing.text.Text.Height()
	return &writing
}

func (writing *Writing) Draw(){
	font.Draw(writing.text)
}

func (writing *Writing) SetLabel(newLabel string) {
	writing.label = newLabel
	writing.text = font.CreateText(writing.label, writing.font.Font)
	writing.text.Text.SetPosition(mgl32.Vec3{
		writing.textPosition.p1,
		writing.textPosition.p2,
		writing.textPosition.p3,
	})
	writing.slaveWidth = writing.text.Text.Width()
	writing.slaveHeight = writing.text.Text.Height()
}

func (writing *Writing) SetPosition(pos_x, pos_y, pos_z float32) {
	writing.textPosition = Coordinate3{pos_x, pos_y, pos_z}
	writing.text.Text.SetPosition(mgl32.Vec3{
		writing.textPosition.p1,
		writing.textPosition.p2,
		writing.textPosition.p3,
	})
}

func (writing *Writing) GetPosition() Coordinate3 {
	return writing.textPosition 
}

func (writing *Writing) GetLabel() string {
	return writing.label
}

func (writing *Writing) GetWidth() float32 {
	fmt.Printf("Width is %v\n", writing.slaveWidth)
	return writing.slaveWidth
}

func (writing *Writing) GetHeight() float32 {
	return writing.slaveHeight
}

func (writing *Writing) GetText() *font.Text_Struct {
	return writing.text
}