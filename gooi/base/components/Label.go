package components 
/**
 * Updated 11/11/2023.
 * Label -> Implements Drawable
 **/
import (
	mgl32   "github.com/go-gl/mathgl/mgl32"
	intf 	"gooi/interfaces"
	colours "gooi/base/colours"
	log 	"log"
)
type Label_Struct struct {
	err  			error
	LabelText 		string
	Pos_x, Pos_y 	float32
	Pos_z 			float32
	LabelCanvas 	*Canvas_Struct
	Label_Text 		*Text_Struct
	Label_Font 		*Font_Struct
}
func CreateLabel(
	name string,
	canvas *Canvas_Struct, 
	pos_x, pos_y float32, 
	font_path string,
	font_name string,
	font_size int,
	) *Label_Struct {
	log.Println("creating [Label] struct.")
	var l = Label_Struct{}
		l.Pos_x = pos_x
		l.Pos_y = pos_y
		l.Pos_z = 1.0
		l.LabelCanvas = canvas
		l.LabelText = name
		// Window characteristics
		var window_width = canvas.CanvasWindow.GetWindowWidth()
		var window_height = canvas.CanvasWindow.GetWindowHeight()
		// Font and text
		l.Label_Font = ReadFontFile(font_name, font_path, font_size)
		l.Label_Font.Font.ResizeWindow(*window_width, *window_height)
		l.Label_Text = CreateText(l.LabelText, l.Label_Font.Font)
		l.Label_Text.Text.SetPosition(mgl32.Vec3{pos_x, pos_y, l.Pos_z})
	   	return &l
}
func CreateThemedLabel(
	name string, 
	canvas *Canvas_Struct, 
	pos_x, pos_y float32,
	font_size int,
	theme *colours.Theme,
	) *Label_Struct {
		log.Println("creating themed [Label] struct.")
		var l = Label_Struct{}
		l.Pos_x = pos_x
		l.Pos_y = pos_y
		l.Pos_z = 1.0
		l.LabelCanvas = canvas
		l.LabelText = name
		// Window characteristics
		var window_width = canvas.CanvasWindow.GetWindowWidth()
		var window_height = canvas.CanvasWindow.GetWindowHeight()
		// Font and text
		l.Label_Font = ReadFontFile(theme.FontName, theme.FontPath, font_size)
		l.Label_Font.Font.ResizeWindow(*window_width, *window_height)
		l.Label_Text = CreateText(l.LabelText, l.Label_Font.Font)
		l.Label_Text.Text.SetPosition(mgl32.Vec3{pos_x, pos_y, l.Pos_z})
	   	return &l
	}

func (l *Label_Struct) Draw() {	Draw(l.Label_Text) }
func (l *Label_Struct) Redraw() { 
	Draw(l.Label_Text) 
}
/**
 * Setters and Getters 
 **/
func (l *Label_Struct) SetPos(x, y float32) { 
	l.Pos_x = x + l.Label_Text.Text.Width()/2
	l.Pos_y = y + l.Label_Text.Text.Height()/2 
	l.Label_Text.Text.SetPosition(mgl32.Vec3{l.Pos_x, l.Pos_y, l.Pos_z})
}
func (l *Label_Struct) GetPos() (float32, float32) { return l.Pos_x -  l.Label_Text.Text.Width()/2, l.Pos_y -  l.Label_Text.Text.Height()/2 }
func (l *Label_Struct) GetBounds() (float32, float32) { return l.Label_Text.Text.Width(), l.Label_Text.Text.Height() }
func (l *Label_Struct) GetTextCanvas() intf.Canvas_Interface { return l.LabelCanvas }
func (l *Label_Struct) SetDisplayText(s string) { l.LabelText = s }
func (l *Label_Struct) GetDisplayText() string { return l.LabelText }

func (l *Label_Struct) SetPosZ(z float32) { 
	if z < 1.0 {
		l.Hidetext()
	} else {
		l.Showtext()
	}
	l.Pos_z = z
}
func (l *Label_Struct) GetPosZ() float32 { 
	return l.Pos_z
}
func (l *Label_Struct) Hidetext(){
	l.Label_Text.Text.Hide()
}
func (l *Label_Struct) Showtext(){
	l.Label_Text.Text.Show()
}
