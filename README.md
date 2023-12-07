# Gooi: A Golang GUI Framework
An openGL based GUI framework built for desktop Golang applications. 

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![OpenGL](https://img.shields.io/badge/OpenGL-%23FFFFFF.svg?style=for-the-badge&logo=opengl)

## Contents 
- Window and Canvas
- Constants
- Components
- Compositions
- Events
- Key and Mouse Listeners
  
# Usage 
## The Window and Canvas
### Application window

> import gooi/base/windows

```golang
windows.NewWindow(title string, width, height float32) *intf.Window_Interface
```

<details>
  <summary>Window Methods</summary>

> #### Window operational methods 
> ```golang
> (A *ApplicationWindow).CloseWindow()
> (A *ApplicationWindow).OpenWindow()
> (A *ApplicationWindow).RunWindow()
> ```
> #### Set/Get glfw window
> ```golang
> (A *ApplicationWindow).SetWindow(w *glfw.Window)
> (A *ApplicationWindow).GetWindow() *glfw.Window
> ```
> #### Set/Get window height
> ```golang
> (A *ApplicationWindow).SetWindowHeight(h float32)
> (A *ApplicationWindow).GetWindowHeight() float32
> ```
> #### Set/Get window width
> ```golang
> (A *ApplicationWindow).SetWindowWidth(w float32)
> (A *ApplicationWindow).GetWindowWidth() float32
> ```
> #### Set background colour 
> ```golang
> (A *ApplicationWindow).SetBackgroundColour(colour [3]float32)
> ```
> #### Set/Get resizable 
> ```golang
> (A *ApplicationWindow).SetResizable(r int)
> (A *ApplicationWindow).GetResizable() int
> ```
> #### Set/Get title 
> ```golang
> (A *ApplicationWindow).SetWindowTitle(t string)
> (A *ApplicationWindow).GetWindowTitle() string
> ```
> #### Set/Get open method 
> ```golang
> (A *ApplicationWindow).SetOpenMethod(f func())
> (A *ApplicationWindow).GetOpenMethod() func()
> ```
> #### Set/Get close method 
> ```golang
> (A *ApplicationWindow).SetCloseMethod(f func())
> (A *ApplicationWindow).GetCloseMethod() func()
> ```
> #### Set/Get window canvas 
> ```golang
> (A *ApplicationWindow).SetWindowCanvas(c intf.Canvas_Interface)
> (A *ApplicationWindow).GetWindowCanvas() intf.Canvas_Interface
> ```
> #### Set/Get mouse handler 
> ```golang
> (A *ApplicationWindow).SetMouseHandler(c intf.MouseHandler_Interface)
> (A *ApplicationWindow).GetWindowCanvas() intf.MouseHandler_Interface
> ```
> #### Set/Get mouse handler 
> ```golang
> (A *ApplicationWindow).SetMouseHandler(c intf.MouseHandler_Interface)
> (A *ApplicationWindow).GetWindowCanvas() intf.MouseHandler_Interface
> ```

</details>

### Canvas

> import gooi/base/components

```golang
components.NewCanvas(window intf.Window_Interface) *Canvas_Struct
```

<details>
  <summary>Canvas Methods</summary>

> #### Add/Get/Set displayable
> ```golang
> (c *Canvas_Struct).AddDisplayable(a intf.Displayable)
> (c *Canvas_Struct).GetDisplayable(a intf.Displayable)
> (c *Canvas_Struct).SetDisplayable(a intf.Displayable)
> ```
> #### Count components
> ```golang
> (c *Canvas_Struct).CountComponents() int
> ```
> #### Refresh/Redraw/Draw
> ```golang
> (c *Canvas_Struct).RefreshCanvas()
> (c *Canvas_Struct).Redraw()
> (c *Canvas_Struct).Draw()
> ```
> - Refresh: selects corrent openGL program, clears the openGL screen, calles Draw() and swaps buffers.
> - Redraw: called Redraw() on all displayables.
> - Draw: called Draw() on all displayables
> #### Compile canvas shader
> ```golang
> (c *Canvas_Struct).CompileCanvasShader(source string, shaderType uint32) uint32
> ```
> #### Set background colour
> ```golang
> (c *Canvas_Struct).SetBackgroundColour(colour [3]float32)
> ```
> #### Get programs
> ```golang
> (c *Canvas_Struct).GetPrograms() uint32
> ```
> #### Get/Set event handler
> ```golang
> (c *Canvas_Struct).GetEventHandler() intf.EventHandler_Interface
> (c *Canvas_Struct).SetEventHandler(intf.EventHandler_Interface)
> ```
> #### Get/Set window interface
> ```golang
> (c *Canvas_Struct).GetWindow() intf.Window_Interface
> (c *Canvas_Struct).SetWindow(intf.Window_Interface)
> ```
> #### Get width/height
> ```golang
> (c *Canvas_Struct).GetWidth() float32
> (c *Canvas_Struct).GetHeight() float32
> ```
> #### Set/Get position
> ```golang
> (c *Canvas_Struct).SetPos(x, y, z float32)
> (c *Canvas_Struct).GetPos() (float32, float32, float32)
> ```
> - GetPos always returns 0, 0, 0 as the canvas is at the root of the ApplicationWindow.
> #### Get/Set master struct
> ```golang
> (c *Canvas_Struct).GetMasterStruct() intf.Displayable
> (c *Canvas_Struct).SetMasterStruct(intf.Displayable)
> ```

</details>

## Important Interfaces 

> import gooi/interfaces

### Displayable
All displayable structs implement the `intf.Displayable` interface. This includes `compositions` and `components`. The interface methods are as follows: 
```golang
(i intf.Displayable).SetPos(float32, float32, float32)
(i intf.Displayable).GetPos() (float32, float32, float32)
(i intf.Displayable).GetWidth() float32
(i intf.Displayable).GetHeight() float32
(i intf.Displayable).GetMasterStruct() intf.Displayable
(i intf.Displayable).SetMasterStruct() intf.Displayable
(i intf.Displayable).Draw()
(i intf.Displayable).Redraw()
```
### Clickable
Anything that is composed of the clickable struct in the foundation package implements the clickable interface. 
```golang
(c intf.Clickable).GetClickableBounds() (x_min float32, x_max float32, y_min float32, y_max float32, z float32)
(c intf.Clickable).Click(click_alive *bool, press_action int, x float32, y float32, glfw.ModifierKey)
```
### Drawable 
Anything that is composed of the drawable struct in the foundation package implements the drawable interface. 
```golang
(d intf.Drawable).SetXYZ([]float32)
(d intf.Drawable).GetXYZ() []float32
(d intf.Drawable).SetRGB([]float32)
(d intf.Drawable).GetRGB() []float32
(d intf.Drawable).SetDIM([]float32)
(d intf.Drawable).GetDIM() []float32
(d intf.Drawable).SetVAO([]intf.Drawing)
(d intf.Drawable).GetVAO() []intf.Drawing
(d intf.Drawable).GetWidth() float32
(d intf.Drawable).GetHeight() float32
```

## Constants 
> import "gooi/base/constants"
```golang
const (
  // Mouse click pressed conditions
  MOUSE_RELEASED = 0
  MOUSE_PRESSED = 1
  NO_CHANGE = 2
  // Orientation of ratio button component
  VERTICAL_ORIENT = 3
  HORISONT_ORIENT = 4
  // Box alignment constants
  ALIGN_CENTRE = 5
  ALIGN_TOP_CENTRE = 6
  ALIGN_BOTTOM_CENTRE = 7
  ALIGN_TOP_LEFT = 8
  ALIGN_TOP_RIGHT = 9
  ALIGN_BOTTOM_LEFT = 10
  ALIGN_BOTTOM_RIGHT = 11
  ALIGN_CENTRE_LEFT = 12
  ALIGN_CENTRE_RIGHT = 13
  // Column alignment constants 
  ALIGN_LEFT = 14
  ALIGN_RIGHT = 15
  ALIGN_CENTRE_COLUMN = 16
  // Row alignment constants
  ALIGN_TOP = 17
  ALIGN_BOTTOM = 18
  ALIGN_CENTRE_ROW = 19
  // Rectangle constants
  FILL_MASTER_DIMENSIONS = 20
  NO_FILL = 21
  MATCH_MASTER_POSITION  = 22
  NO_MATCH_POSITION = 23
)
```
> import "gooi/base/colours"
```golang
var (
  WHITE
  BLACK
  BLUE
  DARK_BLUE
  LIGHT_BLUE
  GRAY
  DARK_GRAY
  LIGHT_GRAY
  GREEN
  RED
  NONE
)
```

## Components 
Gooi components are displayable structs that can be placed on the canvas to achieve some functionality or communicate some information.
### The foundation package
These components are based upon the foundation structs (see foundation package in source).  These foundations implement compositional structs for "Clickable", "Animation", "Drawable", "Input", and "Writing". These labels are largely self explanatory. 
- "Clickable" handles the triggering and assigning of events to clickable areas.
- "Animation" handles the execution of animation-specific functions with time-delay.
- "Drawable" handles the generation of openGL graphics.
- "Input" handles the implementation of editable text input-fields.
- "Writing" handles the implementation of text (using the GLtext package).

The implementation of these compositional structs are listed for each of the displayable components in this section.
### Button 
The Button composition is composed of the Animation, Clickable, Drawable, and Writing compositional structs. 

![Button](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/button.png)

<details>
  <summary>Button struct</summary>
  
```golang
type Button struct {
canvas intf.Canvas_Interface
masterStruct intf.Displayable
componentName string
radius float32
posX, posY, posZ float32
masterWidth, masterHeight float32
slaveWidth, slaveHeight float32
openGLWindowWidth, openGLWindowHeight float32
animationFunctions []func()
clickable *foundation.Clickable
drawable *foundation.Drawable
writing *foundation.Writing
animation *foundation.Animation	
buttonBodyColour [3]float32
}
```

</details>

```golang
components.NewButton(intf.Canvas_Interface, MasterStruct intf.Displayable, Name string, Width, Height, Radius, PosX, PosY, PosZ float32, FontName, FontPath string, FontSize int, ButtonEvent *event.Event_Struct, AniamtionTime time.Duration)
```
- "Radius" refers to the radius of the rounded corners of the button.
- PosX, PosY, and PosZ may be set to 0 if component is arranged inside a compositional structure.

### Label 
The Label composition implements the Writing compositional struct. 

![Label](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/label.png)

<details>
  <summary>Label struct</summary>
  
```golang
type Label struct {
canvas intf.Canvas_Interface
masterStruct intf.Displayable
name string
posX, posY, posZ float32
masterWidth, masterHeight float32 
slaveWidth, slaveHeight float32 
openGLWindowWidth float32 
openGLWindowHeight float32
writing *foundations.Writing
}
```

</details>

```golang
components.NewLabel(intf.Canvas_Interface, MasterStruct intf.Displayable, Name string, PosX, PosY, PosZ float32, FontName, FontPath string, FontSize int)
```

### CheckBox 
The CheckBox composition is composed of the Drawable, Clickable, Writing, and Animation compositional structs. 

![Checkbox](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/checkbox.png)

<details>
  <summary>CheckBox struct</summary>
  
```golang
type CheckBox struct {
canvas intf.Canvas_Interface
masterStruct intf.Displayable
name string
radius, posX, posY, posZ float32
openGLWindowWidth float32 
openGLWindowHeight float32
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
```

</details>

```golang
components.NewCheckBox(intf.Canvas_Interface, MasterStruct intf.Displayable, Name string, Radius, PosX, PosY, PosZ float32, CheckEvent *event.Event_Struct, FontName, FontPath string, FontSize int)
```
- Here "Radius" refers to the literal radius of the circular checkbox.
- PosX, PosY, and PosZ may be set to 0 if component is arranged inside a compositional structure.

### Rectangle 
The Rectangle composition is composed of the Drawable compositional struct.

<details>
  <summary>Rectangle struct</summary>
  
```golang
type Rectangle struct {
canvas intf.Canvas_Interface
masterStruct intf.Displayable
name string
colour [3]float32
posX, posY, posZ, Radius float32
openGLWindowWidth float32 
openGLWindowHeight float32
masterWidth, masterHeight float32 
slaveWidth, slaveHeight float32 
drawable *foundations.Drawable
fillStyle, positionStyle int
}
```

</details>

```golang
components.NewRectangle(intf.Canvas_Interface, MasterStruct intf.Displayable, Name string, Width, Height, PosX, PosY, PosZ, Radius float32, Colour [3]float32, FillStyle, PositionStyle integer)
```
- FillStyle and PositionStyle are integer constants (see constant section of this readme / constants package in source).
  - constants dictate whether the rectangle immitates the position and dimensions of the MasterStruct.  
- PosX, PosY, and PosZ may be set to 0 if component is arranged inside a compositional structure.

### TextInput 
The TextInput composition is composed of the Input struct. By extension, it is indirectly composed of the Drawable, Writing, and Clickable structs. Additionally, it explicitly implements the Writing struct via the input prompt (see TextInput struct). 

![TextInput](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/input.png)

<details>
  <summary>TextInput struct</summary>
  
```golang
type TextInput struct {
canvas intf.Canvas_Interface
masterStruct intf.Displayable
name, placeholder string
inputbox *Rectangle
keylistener *listeners.KeyHandler_Struct
input *foundaitons.Input
prompt *foundaitons.Writing
posX, posY, posZ, Radius float32
masterWidth, masterHeight float32 
slaveWidth, slaveHeight float32 
}
```

</details>

```golang
components.NewTextInput(intf.Canvas_Interface, MasterStruct intf.Displayable, Name, Placeholder string, KeyListener *listeners.KeyHandler_Struct, Width, Height, PosX, PosY, PosZ, Radius float32, Colour [3]float32, FontName, FontPath string, FontSize int)
```
- Placeholder is the text present in the input field by default.
- Name is the text present in the text field prompt.
- Colour refers to the colour of the text field background rectangle.

### ToggleSwitch 
The ToggleSwitch composition is composed of the Drawable, Animation, Writing, and Clickable compositional structs. 

![ToggleSwitch](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/toggle.png)

<details>
  <summary>ToggleSwitch struct</summary>
  
```golang
type ToggleSwitch struct {
canvas intf.Canvas_Interface
masterStruct intf.Displayable
name, placeholder string
posX, posY, posZ float32
width, height float32
togglePos float32
openGLWindowWidth, openGLWindowHeight float32
masterWidth, masterHeight float32 
slaveWidth, slaveHeight float32
fontName, fontPath string
fontSize int
toggleEvent *event.Event_Struct
toggleState bool
toggleColour [3]float32
writing *foundations.Writing
drawable *foundations.Drawable
clickable *foundations.Clickable
animation *foundations.Animation
}
```

</details>

```golang
components.NewToggle(intf.Canvas_Interface, MasterStruct intf.Displayable, Name string, Width, Height, PosX, PosY, PosZ, Radius float32, FontName, FontPath string, FontSize int, toggleEvent *event.Event_Struct)
```
## Compositions 
### Box composition
- Simple composition that holds one `intf.Displayable`.
- Alignment may be
  - `constants.ALIGN_CENTRE`
  - `constants.ALIGN_TOP_CENTRE` / `constants.ALIGN_BOTTOM_CENTRE`
  - `constants.ALIGN_TOP_LEFT`/  `constants.ALIGN_TOP_RIGHT`
  - `constants.ALIGN_BOTTOM_LEFT` / `constants.ALIGN_BOTTOM_RIGHT`
  - `constants.ALIGN_CENTRE_LEFT` / `constants.ALIGN_CENTRE_RIGHT`
```golang
compositions.NewBoxComposition(name string, canvas intf.Canvas_Interface, masterStruct intf.Displayable, x, y, z, slaveWidthRatio, slaveHeightRatio float32, alignment int, colour [3]float32) *Box
```
- `slaveWidthRatio` and `slaveHeightRatio` refers to a fractional proportion of the masterStruct dimensions.
- ie: `(0.5, 0.5)` would indicate that the box composition takes up 1/4 of the masterStruct dimensisons.
![ColumnAlignment](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/box_alignment.png)
### Stack composition
- Stack composition is an extension of box that allows for multiple `intf.Displayable` structs to the displayed in the same composition, with unique alignments.
```golang
compositions.NewStackComposition(name string, canvas intf.Canvas_Interface, x, y, z slaveWidthRatio, slaveHeightRatio float32, alignment []int)
```
- alignments are constants defined in an array for each of the planned `intf.Displayables` to be added to the composition. 
### Column composition
- The column composition organises `intf.Displayable` structs in a vertical stack.
- the slaveWidth and slaveHeight of this composition are not defined explicity, but instead obtained via the addition of the slave dimensions of all composition displayables plus padding.
- Alignment can be stated as:
  - `ALIGN_LEFT`
  - `ALIGN_RIGHT`
  - `ALIGN_CENTRE_COLUMN`
```golang
compositions.NewColumnComposition(name string, canvas intf.Canvas_Interface, masterStruct intf.Displayable, x, y, z float32, alignment int)
```
![ColumnAlignment](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/column_alignment.png)
### Row composition
- The row composition organises `intf.Displayable` structs in a horisontal line.
- the slaveWidth and slaveHeight of this composition are not defined explicity, but instead obtained via the addition of the slave dimensions of all composition displayables plus padding.
- Alignment can be stated as:
  - `ALIGN_TOP`
  - `ALIGN_BOTTOM`
  - `ALIGN_CENTRE_ROW`
```golang
compositions.NewRowComposition(name string, canvas intf.Canvas_Interface, masterStruct intf.Displayable, x, y, z float32, alignment int)
```
![ColumnAlignment](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/row_alignment.png)
### Tabs composition (Alpha)
- The tabbed composition allows for multiple box compositions to be switched between via the use of a row of selection buttons.
``` golang
compositions.NewTabComposition(name string, canvas intf.Canvas_Interface, masterStruct intf.Displayable, eventHandler intf.EventHandler_Interface, mouseHandler intf.MouseHandler_Interface, labels []string, x, y, z, slaveWithRatio, slaveHeightRatio float32, font_name, font_path, font_size string)
```






