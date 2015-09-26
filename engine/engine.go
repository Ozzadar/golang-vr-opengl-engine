/* 	********* 	Engine.go	**********
	**	Contains structures for engine objects and engine functions

*/

package engine

import (
	"unsafe"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

//Vertex
type Vertex struct{
	position [3]float32
	//color [4]byte


}

/*	********** 	Sprite objects 	********** */
type Sprite struct {
	x float32
	y float32
	width float32
	height float32

	vertexData []float32
	vbo uint32	//vertex buffer object
	runtime float32
	previousTime float32

	program uint32
}


func (s *Sprite) Init(x float32, y float32, height float32, width float32, program uint32) {
	s.x = x
	s.y = y
	s.height = height
	s.width = width


	s.vertexData = []float32{

		//x   		   //y   			//z    //r   	//g   	//b  		//a
		s.x + s.width, s.y + s.height ,  0    , 1,  	0.75,  	0.1,		0,
		s.x,		   s.y + s.height,   0    , 0.5,  	0.75,  	0.2, 		0,
		s.x,           s.y,              0,     0.25,  	0.75,  	0.3, 		0,
		s.x,           s.y,              0,     0.25,  	0.75,  	0.3, 		0,
		s.x + s.width, s.y,              0,     0.5,  	0.75,  	0.2, 		0,
		s.x + s.width, s.y + s.height ,  0,     1,  	0.75,  	0.1, 		0 }
 	
 	/*
	//first triangle
	s.vertexData[0] = s.x + s.width
	s.vertexData[1] = s.y + s.height

	s.vertexData[2] = s.x
	s.vertexData[3] = s.y + s.height

	s.vertexData[4] = s.x
	s.vertexData[5] = s.y

	//second triangle
	s.vertexData[6] = s.x
	s.vertexData[7] = s.y

	s.vertexData[8] = s.x	+ s.width
	s.vertexData[9] = s.y

	s.vertexData[10] = s.x + s.width
	s.vertexData[11] = s.y + s.height
*/

	s.runtime = 0

	s.program = program
}

func (s *Sprite) GetVAO() (VAO uint32) {

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	gl.GenBuffers(1, &s.vbo)
  	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)

  	gl.BufferData(gl.ARRAY_BUFFER, len(s.vertexData)*4, gl.Ptr(&s.vertexData[0]), gl.STATIC_DRAW)

  	attrib_loc := uint32(gl.GetAttribLocation(s.program, gl.Str("vert\x00")))
  	color_loc := uint32(gl.GetAttribLocation(s.program, gl.Str("vertColor\x00")))
  	gl.EnableVertexAttribArray(attrib_loc)
 	gl.VertexAttribPointer(attrib_loc, 3, gl.FLOAT, false, int32(unsafe.Sizeof(s.vertexData[0])) * 7, nil)
  	
  	gl.EnableVertexAttribArray(color_loc)
 	gl.VertexAttribPointer(color_loc, 4, gl.FLOAT, false, int32(unsafe.Sizeof(s.vertexData[0])) * 7, gl.PtrOffset(3*4)) 


 	time_loc := gl.GetUniformLocation(s.program, gl.Str("time\x00"))

 	gl.Uniform1f(time_loc, s.runtime)
 

 	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

 	gl.BindVertexArray(0)

	//Update
	time := glfw.GetTime();
	elapsed := float32(time) - s.previousTime
	s.previousTime = float32(time)
	s.runtime = s.runtime + elapsed
	return
}

func (s *Sprite) ReleaseAll() {

	gl.DeleteBuffers(1, &s.vbo)

}
