package Interfaces


/**
 * This interface extends the component interface.
 **/
type Editable_Interface interface {
	SetDisplayText(string)
	GetDisplayText() string

	GetCursorIndex() int
	SetCursorIndex(int) 

	IsMaxLength() bool

	GeneratePolygons()

}

