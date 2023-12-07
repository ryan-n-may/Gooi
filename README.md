# Gooi: A Golang GUI Framework
An openGL based GUI framework built for desktop Golang applications. 

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![OpenGL](https://img.shields.io/badge/OpenGL-%23FFFFFF.svg?style=for-the-badge&logo=opengl)

## Contents 
- Dependencies 
# Usage 
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

![Button](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/button.png)

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

![Button](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/button.png)

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

![Button](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/button.png)

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

![Button](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/button.png)

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

![Button](https://github.com/ryan-n-may/Gooi/blob/main/readme/screenshots/button.png)

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







