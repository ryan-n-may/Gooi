package foundation

type Coordinate struct {
	p1 float32
	p2 float32
}

func (coordinate *Coordinate) GetP1() float32 { return coordinate.p1 }
func (coordinate *Coordinate) GetP2() float32 { return coordinate.p2 }

func (coordinate *Coordinate) InBounds(p3 float32) bool {
	if p3 <= coordinate.p2 && p3 >= coordinate.p1 {
		return true
	} 
	return false
}

type Coordinate3 struct {
	p1 float32
	p2 float32
	p3 float32
}

func (coordinate *Coordinate3) GetP1() float32 { return coordinate.p1 }
func (coordinate *Coordinate3) GetP2() float32 { return coordinate.p2 }
func (coordinate *Coordinate3) GetP3() float32 { return coordinate.p3 }

