package Interfaces

type Component_Interface interface {
	SetName(string)
	GetName() string

    SetXYZ([]float32) 
	GetXYZ() []float32

	SetRGB([]float32) 
	GetRGB() []float32

	SetDIM([]float32) 
	GetDIM() []float32

	SetVAO([]Drawing) 
	GetVAO() []Drawing	

	Move(float32, float32)

	Draw()
}
