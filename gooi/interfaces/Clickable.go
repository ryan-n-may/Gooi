package Interfaces

import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"
)


/**
 * This interface extends the component interface.
 **/
type Clickable_Interface interface {
	SetClickable(bool)
	GetClickable() bool

	SetClickEvent(Event_Interface)
	AnimateTrigger(int, *bool)
	TriggerClickEvent(int, float32, float32, glfw.ModifierKey)	
	Kill(*bool)

	SetClickableBounds(float32, float32, float32, float32)
	GetClickableBounds() (int, int, int, int, int)

	SetClickTrigger(int)
	GetClickTrigger() int

	GetCanvas() Canvas_Interface
}
