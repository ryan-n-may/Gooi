package Interfaces

type Paramaters_Interface interface {
	SetParameters(any)
	GetParameters() any
}

type Event_Interface interface {
	SetMethod(func(Paramaters_Interface))
	GetMethod() func(Paramaters_Interface)

	SetName(string)
	GetName() string

	SetParameterStruct(Paramaters_Interface)
	GetParameterStruct() Paramaters_Interface
}

type EventHandler_Interface interface {
	// Adds an event so the listener is informed
	// on what to do when a given event occurs.
	RegisterEventToHandler(Event_Interface)

	// Add the name of the event to the event queue
	// events are executed, 1 event at a time, in a 
	// queue. 
	AddEventToEventQueue(string)

	// Execute the next event, or skip event.
	ExecuteNextEvent()
	SkipNextEvent()
}
