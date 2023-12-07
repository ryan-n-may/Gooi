# Gooi: A Golang GUI Framework

An openGL based GUI framework built for desktop Golang applications. 


![mascot](Gooi.png)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![OpenGL](https://img.shields.io/badge/OpenGL-%23FFFFFF.svg?style=for-the-badge&logo=opengl)
![Go-gl](https://avatars.githubusercontent.com/u/2505184?s=48&v=4)

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
The Button composition implements in Animation, Clickable, Drawable, and Writing compositional struct. 
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
```golang
components.NewButton(intf.Canvas_Interface, intf.Displayable, Name (String), Width, Height, Radius, PosX, PosY, PosZ (Float32), FontName, FontPath (String), FontSize (Int), ButtonEvent (*event.Event_Struct), AniamtionTime (time.Duration))
```


