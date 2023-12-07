package foundation

import (
	glfw 	"github.com/go-gl/glfw/v3.2/glfw"

	beep 	"github.com/faiface/beep"
	mp3 	"github.com/faiface/beep/mp3"
	speaker	"github.com/faiface/beep/speaker"

	intf 	"gooi/interfaces"
	event 	"gooi/base/event"

	time 	"time"
	sync 	"sync"
	os 		"os"
	fmt 	"fmt"
)

type Clickable struct {
	canvas intf.Canvas_Interface

	clickEvent *event.Event_Struct
	clickTrigger int 

	boundsX Coordinate
	boundsY Coordinate

	clickable bool

	buffer *beep.Buffer
	format *beep.Format
}

var (
	audio_wg sync.WaitGroup
)

func NewClickable(canvas intf.Canvas_Interface, clickEvent *event.Event_Struct, clickTrigger int, boundsX_min, boundsX_max, boundsY_min, boundsY_max float32) *Clickable {
	var clickable = Clickable{
		canvas,
		clickEvent,
		clickTrigger,
		Coordinate{boundsX_min, boundsX_max},
		Coordinate{boundsY_min, boundsY_max},
		true,
		nil, 
		nil,
	}
	clickable.canvas.GetEventHandler().RegisterEventToHandler(clickEvent)

	f, _ := os.Open("base/mouseclick.mp3")
	var streamer, format, _ = mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)) // click sound currently isnt working :(((())))
	var buffer = beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()


	clickable.buffer = buffer
	clickable.format = &format

	return &clickable
}

func (clickable *Clickable) TriggerClickEvent(alive *bool, pressAction int, pos_x, pos_y float32, mod_key glfw.ModifierKey) {
	// click events can only be added once the last event has finished being added to the event queue.
	// similar to animations but less consequential.
	
	audio_wg.Add(1)
	go clickable.PlaySound(&audio_wg, *clickable.buffer)

	var kill = func(alive *bool) {
		*alive = false
		return
	}
	if !(*alive) {
		(*alive) = true
		defer kill(alive)
		if clickable.clickEvent.GetMethod() != nil && clickable.clickable == true && pressAction == clickable.clickTrigger{
			clickable.canvas.GetEventHandler().AddEventToEventQueue(clickable.clickEvent.GetName())
  		}
	}

	audio_wg.Wait()
	return
}

func (clickable *Clickable) PlaySound(wg *sync.WaitGroup, buffer beep.Buffer){
	sound := buffer.Streamer(0, buffer.Len())
	speaker.Play(sound)
	fmt.Println("Playing the sound")
	defer wg.Done()
	return
}


/**
 * Public Accessors and Mutators 
 **/
func (clickable *Clickable) SetClickable(boolean bool) { clickable.clickable = boolean }
func (clickable *Clickable) SetClickEvent(event *event.Event_Struct) { clickable.clickEvent = event }
func (clickable *Clickable) SetClickTrigger(trigger int) { clickable.clickTrigger = trigger }
func (clickable *Clickable) SetClickBounds(boundsX_min, boundsX_max, boundsY_min, boundsY_max float32) {
	clickable.boundsX = Coordinate{boundsX_min, boundsX_max}
	clickable.boundsY = Coordinate{boundsY_min, boundsY_max}
}
func (clickable *Clickable) GetClickBounds() (Coordinate, Coordinate) { return clickable.boundsX, clickable.boundsY }
func (clickable *Clickable) GetClickTrigger() int { return clickable.clickTrigger }
func (clickable *Clickable) GetClickable() bool { return clickable.clickable }
