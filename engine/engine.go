/* 	********* 	Engine.go	**********
	**	Contains structures for engine objects and engine functions

*/

package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

/*	********** 	Sprite objects 	********** */
type Sprite struct {
	x float32
	y float32
	width float32
	height float32
	vertexData [12]float32
	vbo uint32	//vertex buffer object
	program uint32
}


func (s *Sprite) Init(x float32, y float32, height float32, width float32, program uint32) {
	s.x = x
	s.y = y
	s.height = height
	s.width = width

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

	s.program = program
}

func (s *Sprite) GetVAO() (VAO uint32) {

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	gl.GenBuffers(1, &s.vbo)
  	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)

  	gl.BufferData(gl.ARRAY_BUFFER, len(s.vertexData)*4, gl.Ptr(&s.vertexData[0]), gl.STATIC_DRAW)

  	attrib_loc := uint32(gl.GetAttribLocation(s.program, gl.Str("vert\x00")))

  	gl.EnableVertexAttribArray(attrib_loc)
 	gl.VertexAttribPointer(attrib_loc, 2, gl.FLOAT, false, 0, nil)
 	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

 	gl.BindVertexArray(0)
	return
}

func (s *Sprite) ReleaseAll() {

	gl.DeleteBuffers(1, &s.vbo)

}
