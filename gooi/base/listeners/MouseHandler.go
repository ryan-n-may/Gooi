package listeners
import (
	intf "gooi/interfaces"
	cons "gooi/base/constants"
	glfw "github.com/go-gl/glfw/v3.2/glfw"
)
type MosueHandler_Struct struct {
	Clickables 	[]intf.Clickable
	Name 		string
	MousePressed bool
	MousePressed_Previous bool
	Click_x float64
	Click_y float64	
	Current_Z_Layer float32
}
func CreateMouseHandler(s string) *MosueHandler_Struct {
	var mh = MosueHandler_Struct{}
	mh.Clickables = make([]intf.Clickable, 0)
	mh.Name = s
	mh.MousePressed = false
	mh.Current_Z_Layer = 0
	return &mh
}
func (mh *MosueHandler_Struct) RegisterClickableToHandler(ci intf.Clickable) {
	mh.Clickables = append(mh.Clickables, ci)
}
func (mh *MosueHandler_Struct) SetCurrentZLayer(z float32){
	mh.Current_Z_Layer = z
}
func (mh *MosueHandler_Struct) GetCurrentZLayer() float32 {
	return mh.Current_Z_Layer
}
func (mh *MosueHandler_Struct) GetClickData(window *glfw.Window) (float32, float32, int) {

	mh.Click_x, mh.Click_y = window.GetCursorPos()

	var _, height = window.GetSize()
	mh.Click_y = float64(height) - mh.Click_y

	mh.MousePressed_Previous = mh.MousePressed

	if window.GetMouseButton(glfw.MouseButton1) == glfw.Press {
		mh.MousePressed = true
	} else if window.GetMouseButton(glfw.MouseButton1) == glfw.Release {
		mh.MousePressed = false
	}

	var mousePressChange = cons.NO_CHANGE
	if mh.MousePressed_Previous == true && mh.MousePressed == false {
		mousePressChange = cons.MOUSE_RELEASED
	} else if mh.MousePressed_Previous == false && mh.MousePressed == true {
		mousePressChange = cons.MOUSE_PRESSED
	}

	return float32(mh.Click_x), float32(mh.Click_y), mousePressChange
}
func (mh *MosueHandler_Struct) CheckClick(posx, posy float32, pressed int, mod_key glfw.ModifierKey){
	var alive = false
	for _, ci := range mh.Clickables {
		var x_min, x_max, y_min, y_max, z = ci.GetClickableBounds()
		if posx >= x_min && posx <= x_max && posy <= y_max && posy >= y_min && z == mh.GetCurrentZLayer(){
			if !alive{
				ci.Click(&alive, pressed, float32(posx), float32(posy), mod_key)
			} 
		}
	}
}
func (mh *MosueHandler_Struct) GetClickables() []intf.Clickable {
	return mh.Clickables
}
func (mh *MosueHandler_Struct) SetName(s string) {
	mh.Name = s
}
func (mh *MosueHandler_Struct) GetName() string {
	return mh.Name
}