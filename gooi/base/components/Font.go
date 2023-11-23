package components
/**
 * Updated 11/11/2023.
 * Font 
 * Doesn't implement interface.
 * Extension of gltext.
 **/
import (
	fixed   "golang.org/x/image/math/fixed"
	mgl32   "github.com/go-gl/mathgl/mgl32"
	gltext 	"gooi/gltext"
	v41 	"gooi/gltext/v41"
	os 		"os"
	fmt     "fmt"
)
/** 
 * Read font file and save contents to struct. 
 **/
type Font_Struct struct {
	Font *v41.Font
}
func ReadFontFile(fontName, fontPath string, scale int) (*Font_Struct){
	//gltext.IsDebug = true
	var font_struct = Font_Struct{}
	// Attempt to read font configuration file
	var font *v41.Font
	config_name := fmt.Sprintf("%s_%v", fontName, scale)
	config, err := gltext.LoadTruetypeFontConfig("fontconfigs", config_name)
	if err == nil {
		// If no error, font is given by the configuration
		font, _ = v41.NewFont(config)
	} else {
		// Otherwise, we open the TTF font format and generate a configuration
		var fd, _ = os.Open(fontPath)
		defer fd.Close()
		runeRanges 	:= make(gltext.RuneRanges, 0)
		runeRange 	:= gltext.RuneRange{Low: 1, High: 127}
		runeRanges 	= append(runeRanges, runeRange)
		scale 		:= fixed.Int26_6(scale)
		runesPerRow := fixed.Int26_6(128)
		config, _ = gltext.NewTruetypeFontConfig(fd, scale, runeRanges, runesPerRow, 5)
		// Save configuration
		config.Save("fontconfigs", config_name)
		font, _ = v41.NewFont(config)
	}
	font_struct.Font = font
	return &font_struct
}
/**
 * Generate Text using string and font, save text in struct
 **/
type Text_Struct struct {
	Text *v41.Text
}
func CreateText(str string, font *v41.Font) (*Text_Struct) {
	var text_struct = Text_Struct{}
	var scaleMin, scaleMax = float32(0.01), float32(10.1)
	text := v41.NewText(font, scaleMin, scaleMax)
	text.SetString(str)
	text.MaxRuneCount = 10
	text.SetColor(mgl32.Vec3{0, 0, 0})
	text.SetPosition(mgl32.Vec3{0, 0, 0})
	text_struct.Text = text
	return &text_struct
}	
// Draw text struct. 
func Draw(text_struct *Text_Struct) {
	text_struct.Text.Draw()
	text_struct.Text.Show()
}	


