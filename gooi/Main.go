package main 

import (
	comp "GUI/Base/Components"
	glfw "github.com/go-gl/glfw/v3.2/glfw"
	gl 			"github.com/go-gl/gl/v4.1-core/gl"
	"fmt"
	"runtime"
)

func main(){
	runtime.LockOSThread()

	glfw.Init()
	defer glfw.Terminate()
	
	glfw.WindowHint(glfw.Resizable, 0)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)

	var Window, _ = glfw.CreateWindow(600, 300, "Text Test", nil, nil) 

	Window.MakeContextCurrent()

	gl.Init()
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("OpenGL version: \" %s \"\n", version)

	var font_struct = comp.ReadFontFile()
	width, height := Window.GetSize()
	font_struct.Font.ResizeWindow(float32(width), float32(height))
	var text_struct = comp.CreateText("大好きどなに", font_struct.Font)

	for !Window.ShouldClose(){
		fmt.Print(".")	
		gl.Clear(gl.COLOR_BUFFER_BIT)

		comp.DrawText(text_struct)
		
		Window.SwapBuffers()
		glfw.PollEvents()
	}
	
	text_struct.Text.Release()
	font_struct.Font.Release()
}
