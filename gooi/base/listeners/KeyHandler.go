package listeners
import (
	glfw "github.com/go-gl/glfw/v3.2/glfw"
	intf "gooi/interfaces"
	strings "strings"
	reg "regexp"
	"fmt"
)
var (
	Alphabetical = "^[A-Z]*$"
	Numerical = "^[0-9]*$"
	Numeric_Upper = map[string]string {
		"1": "!", "2": "@",
		"3": "#", "4": "$",
		"5": "%", "6": "^",
		"7": "&", "8": "*",
		"9": "(", "0": ")",
	}
)
type KeyHandler_Struct struct {
	Name 		string
	KeyQueue 	[]Key_Struct
	Canvas 		intf.Canvas_Interface
	Focus		intf.Editable_Interface
}
type Key_Struct struct {
	Key glfw.Key
	Mod glfw.ModifierKey
}
func CreateKeyListener(s string, canvas intf.Canvas_Interface) *KeyHandler_Struct {
	var kh = KeyHandler_Struct{}
	kh.KeyQueue = make([]Key_Struct, 0)
	kh.Name = s
	kh.Canvas = canvas
	kh.GetCanvas().GetWindow().GetWindow().SetKeyCallback(kh.keyCallback)
	return &kh
}
func (kh *KeyHandler_Struct) AddKeyEvent(key glfw.Key, mod glfw.ModifierKey){
	var key_struct = Key_Struct{key, mod}
	kh.KeyQueue = append(kh.KeyQueue, key_struct)

	if kh.GetFocus() != nil {
		var prev_text = kh.GetFocus().GetDisplayText()
		var cursor_pos = kh.GetFocus().GetCursorIndex()
		// Arrow key right
		if kh.GetFocus() != nil && key_struct.Key == glfw.KeyRight{
			if (cursor_pos != len(prev_text)) {
				kh.GetFocus().SetCursorIndex(cursor_pos+1)
			} else {
				kh.GetFocus().SetCursorIndex(cursor_pos)
			}
		// Arrow key left
		} else if kh.GetFocus() != nil && key_struct.Key == glfw.KeyLeft {
			if (cursor_pos != 0) {
				kh.GetFocus().SetCursorIndex(cursor_pos-1)
			} else {
				kh.GetFocus().SetCursorIndex(cursor_pos)
			}
			kh.GetFocus().GeneratePolygons()
		// Shift key
		} else if kh.GetFocus() != nil && key_struct.Key == glfw.KeyLeftShift {
			// Do nothing
		// Backspace
		} else if kh.GetFocus() != nil && key_struct.Key == glfw.KeyBackspace && cursor_pos != 0 {
			kh.GetFocus().SetDisplayText(prev_text[:cursor_pos-1] + prev_text[cursor_pos:])
			kh.GetFocus().SetCursorIndex(cursor_pos-1) 
			kh.GetFocus().GeneratePolygons()
			fmt.Println("Backspace")
		// Delete
		} else if kh.GetFocus() != nil && key_struct.Key == glfw.KeyDelete && cursor_pos != len(prev_text) {
			kh.GetFocus().SetDisplayText(prev_text[:cursor_pos] + prev_text[cursor_pos+1:]) 
			kh.GetFocus().GeneratePolygons()
		// Character!
		} else if kh.GetFocus() != nil && cursor_pos != len(prev_text) && !kh.GetFocus().IsMaxLength() {
			kh.GetFocus().SetDisplayText(prev_text[:cursor_pos] + kh.KeyQueueToString() + prev_text[cursor_pos:])
			kh.GetFocus().SetCursorIndex(cursor_pos+1)
			kh.GetFocus().GeneratePolygons()
			fmt.Println("Adding character")
		} else if kh.GetFocus() != nil && cursor_pos == len(prev_text) && !kh.GetFocus().IsMaxLength() {
			kh.GetFocus().SetDisplayText(prev_text + kh.KeyQueueToString())
			kh.GetFocus().SetCursorIndex(cursor_pos+1)
			kh.GetFocus().GeneratePolygons()
			fmt.Println("Adding character at end")
		} else {
			fmt.Println(kh.GetFocus().IsMaxLength())
			fmt.Println("Not sure")
		}
	}
	kh.FlushKeyQueue()
}
func (kh *KeyHandler_Struct) ReadKeyQueue() []Key_Struct { return kh.KeyQueue }
func (kh *KeyHandler_Struct) KeyQueueToString() string {
	var converted_string = ""
	for _, k := range kh.KeyQueue {
		converted_string += kh.KeyLookup(k)
	}
	return converted_string
}
func (kh *KeyHandler_Struct) KeyLookup(key Key_Struct) string {
	var key_string = string(key.Key)
	
	var isAlphabetic, _ = reg.MatchString(Alphabetical, key_string)
	var isNumeric, _ = reg.MatchString(Numerical, key_string)

	if isAlphabetic {
		fmt.Println("Is Alphabetic")
		if key.Mod == glfw.ModShift {
			key_string = strings.ToUpper(key_string)
		} else {
			key_string = strings.ToLower(key_string)
		}
	} else if isNumeric {
		fmt.Println("Is Numeric")
		if key.Mod == glfw.ModShift {
			key_string = Numeric_Upper[key_string]
		} else {
			key_string = (key_string)
		}
	}
	
	return key_string
}
func (kh *KeyHandler_Struct) FlushKeyQueue() { kh.KeyQueue = make([]Key_Struct, 0) }
func (kh *KeyHandler_Struct) SetName(s string) { kh.Name = s }
func (kh *KeyHandler_Struct) GetName() string { return kh.Name }
func (kh *KeyHandler_Struct) SetCanvas(canvas intf.Canvas_Interface) { kh.Canvas = canvas }
func (kh *KeyHandler_Struct) GetCanvas() intf.Canvas_Interface { return kh.Canvas }
func (kh *KeyHandler_Struct) SetFocus(d intf.Editable_Interface) { 
	kh.Focus = d 
	if kh.Focus != nil {		
		kh.FlushKeyQueue()
	}
}
func (kh *KeyHandler_Struct) GetFocus() intf.Editable_Interface { return kh.Focus }

func (kh* KeyHandler_Struct) keyCallback(
	window *glfw.Window, 
	key glfw.Key, 
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey) {
	if action == glfw.Press {
		fmt.Println("Key pressed")
		kh.AddKeyEvent(key, mods)
		if (key == glfw.KeyEscape) || (key == glfw.KeyEnter) {
			var focus = kh.GetFocus()
			kh.SetFocus(nil)
			focus.GeneratePolygons()
			fmt.Println("Exit key pressed")
		}
	}
}