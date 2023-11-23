package event
import (
	list "container/list"
	intf "gooi/interfaces"
)
var (
	NULL_EVENT = Event_Struct{
		nil,
		"",
		&Event_Paramaters{
			nil,
		},
	}
)
type Event_Paramaters struct {
	paramaters any
}
func (p *Event_Paramaters) SetParameters(param any) { p.paramaters = param }
func (p *Event_Paramaters) GetParameters() any { return p.paramaters }
func NewEventParameter(p any) intf.Paramaters_Interface{
	var s =  Event_Paramaters{p}
	return &s
}
type Event_Struct struct {
	Method 	func(intf.Paramaters_Interface)
	Name 	string
	Params	intf.Paramaters_Interface
}
func (e *Event_Struct) SetName(s string) { e.Name = s }
func (e *Event_Struct) GetName() string { return e.Name }
func (e *Event_Struct) SetMethod(f func(intf.Paramaters_Interface)) { e.Method = f }
func (e *Event_Struct) GetMethod() func(intf.Paramaters_Interface) { return e.Method }
func (e *Event_Struct) SetParameterStruct(p intf.Paramaters_Interface) { e.Params = p }
func (e *Event_Struct) GetParameterStruct() intf.Paramaters_Interface { return e.Params }
type EventHandler_Struct struct {
	EventRegistry map[string]func(intf.Paramaters_Interface)
	ParamaterRegistry map[string]intf.Paramaters_Interface
	EventQueue *list.List
}
func NewEvent(name string, f func(intf.Paramaters_Interface), params intf.Paramaters_Interface) *Event_Struct{
	var event = Event_Struct{ f, name, params }
	return &event
}
func NewEventHandler() *EventHandler_Struct {
	var e = EventHandler_Struct{}
	e.EventRegistry = make(map[string]func(intf.Paramaters_Interface))
	e.ParamaterRegistry = make(map[string]intf.Paramaters_Interface)
	e.EventQueue = list.New()
	return &e
}
func (e *EventHandler_Struct) RegisterEventToHandler(ev intf.Event_Interface){
	e.EventRegistry[ev.GetName()] = ev.GetMethod()
	e.ParamaterRegistry[ev.GetName()] = ev.GetParameterStruct()
}
func (e *EventHandler_Struct) AddEventToEventQueue(ev string){
	e.EventQueue.PushBack(ev)
}
func (e *EventHandler_Struct) ExecuteNextEvent() {
	if e.EventQueue.Len() > 0 {
		var elem = e.EventQueue.Front()
		e.EventQueue.Remove(elem)
		var method = e.EventRegistry[(elem.Value).(string)]
		var paramaters = e.ParamaterRegistry[(elem.Value).(string)]
		method(paramaters)
	}
}
func (e *EventHandler_Struct) SkipNextEvent() {
	if e.EventQueue.Len() > 0 {
		var elem = e.EventQueue.Front()
		e.EventQueue.Remove(elem)
	}
}