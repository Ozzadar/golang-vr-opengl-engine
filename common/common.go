/* 	**********	Common.h 	**********
	** Common functions for an OpenGL engine
	** LoadShader(string, uint32) shader -- Creates a shader from file, returns shader uint32
	** LoadShaderProgram (string, string) -- Creates programs from file, returns program unit32

*/
package common


import (
  	"io/ioutil"
	"log"
	"strings"
	"github.com/go-gl/gl/v4.5-core/gl"
)

func LoadShader(filename string, shader_type uint32) (shader uint32) {
  var shader_bytes []byte
  var shader_string string
  var shader_err error
  var status int32
  shader = gl.CreateShader(shader_type)
  shader_bytes, shader_err = ioutil.ReadFile(filename)
  shader_bytes = append(shader_bytes, []byte("\x00")[0])
  if shader_err != nil {
    log.Fatal("Could not load shader from file: ", filename)
  }

  shader_string = string(shader_bytes)
  csource := gl.Str(shader_string)
  gl.ShaderSource(shader, 1, &csource, nil)
  gl.CompileShader(shader)

  gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    log.Printf("Compile error in shader %s:\n", filename)
    var logLength int32
    gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

    l := strings.Repeat("\x00", int(logLength+1))
    gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(l))

    log.Fatal(l)
  }

  return
}

func LoadShaderProgram(vertex_shader string, fragment_shader string) (program uint32) {
  program = gl.CreateProgram()

  vs := LoadShader(vertex_shader, gl.VERTEX_SHADER)
  fs := LoadShader(fragment_shader, gl.FRAGMENT_SHADER)
  gl.AttachShader(program, vs)
  gl.AttachShader(program, fs)

  gl.LinkProgram(program)

  gl.DetachShader(program, vs)
  gl.DetachShader(program, fs)

  return
}
