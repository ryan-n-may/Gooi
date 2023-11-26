package base

import (
	gl 	 "github.com/go-gl/gl/v4.1-core/gl"
	intf "gooi/interfaces"
	math "math"
)

func GenerateCircle(int intf.Drawable, RGBColour [3]float32, pos_x, pos_y, pos_z, radius float32, window_width, window_height float32, fidelity float64) uint32 {
	//Generating XYZ
	var angle = 2*math.Pi / fidelity // angle delta in radians
	var triagleCount = fidelity - 2
	var xyz = make([]float32, 0)
	var rgb = make([]float32, 0)
	var dim = make([]float32, 0)
	for i := 0.0; i < triagleCount; i++ {
		var currentAngle = angle * i
		var x = radius * float32(math.Cos(currentAngle)) + pos_x
		var y = radius * float32(math.Sin(currentAngle)) + pos_y
		var z = pos_z

		xyz = append(xyz, x)
		xyz = append(xyz, y)
		xyz = append(xyz, z)
		rgb = append(rgb, RGBColour[0]/255)
		rgb = append(rgb, RGBColour[1]/255)
		rgb = append(rgb, RGBColour[2]/255)
		dim = append(dim, window_width)
		dim = append(dim, window_height)
		dim = append(dim, 1.0)
	}

	int.SetXYZ(xyz)
   	int.SetDIM(dim)
   	int.SetRGB(rgb)

   	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(xyz), gl.Ptr(xyz), gl.DYNAMIC_DRAW)
    gl.EnableVertexAttribArray(0)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

    var vbo2 uint32
    gl.GenBuffers(1, &vbo2)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo2)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(rgb), gl.Ptr(rgb), gl.DYNAMIC_DRAW)
    gl.EnableVertexAttribArray(1)
    gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

    var vbo3 uint32
    gl.GenBuffers(1, &vbo3)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo3)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(dim), gl.Ptr(dim), gl.DYNAMIC_DRAW)
    gl.EnableVertexAttribArray(2)
    gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func GenerateRectangle(int intf.Drawable, RGBColour [3]float32, width, height, pos_x, pos_y, pos_z float32, window_height, window_width float32) uint32 {
	var rgb = []float32{
    		RGBColour[0]/255, RGBColour[1]/255, RGBColour[2]/255,
	    	RGBColour[0]/255, RGBColour[1]/255, RGBColour[2]/255,
	    	RGBColour[0]/255, RGBColour[1]/255, RGBColour[2]/255,

	    	RGBColour[0]/255, RGBColour[1]/255, RGBColour[2]/255,
	    	RGBColour[0]/255, RGBColour[1]/255, RGBColour[2]/255,
	    	RGBColour[0]/255, RGBColour[1]/255, RGBColour[2]/255,
   	}
   	var xyz = []float32{
       		pos_x, 			pos_y, 			pos_z,	 
       		pos_x+width, 	pos_y,  		pos_z, 
       		pos_x,			pos_y+height,  	pos_z, 

       		pos_x+width, 	pos_y+height, 	pos_z,	 
       		pos_x, 			pos_y+height,  	pos_z, 
      		pos_x+width,	pos_y,   		pos_z, 
  	}

  	var dim = []float32{
  			window_width, window_height, 1.0,
  			window_width, window_height, 1.0,
  			window_width, window_height, 1.0,

  			window_width, window_height, 1.0,
  			window_width, window_height, 1.0,
  			window_width, window_height, 1.0,
	}
   	int.SetXYZ(xyz)
   	int.SetDIM(dim)
   	int.SetRGB(rgb)

   	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(xyz), gl.Ptr(xyz), gl.STATIC_DRAW)
    gl.EnableVertexAttribArray(0)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

    var vbo2 uint32
    gl.GenBuffers(1, &vbo2)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo2)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(rgb), gl.Ptr(rgb), gl.STATIC_DRAW)
    gl.EnableVertexAttribArray(1)
    gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

    var vbo3 uint32
    gl.GenBuffers(1, &vbo3)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo3)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(dim), gl.Ptr(dim), gl.STATIC_DRAW)
    gl.EnableVertexAttribArray(2)
    gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)

	return vao
}
