/* REFACTORING OF GOLANG-VR-LINUX TO SUPPORT A MORE ROBUST DEVELOPMENT PIPELINE 
	STILL BASICALLY IDENTICAL TO LEGACY VERSION, REFACTORING TO COME AFTER OPENGL
	CALLS ARE MADE

*/


package main

//Let's try importing the ovr c library

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/lib/ -L/usr/local/lib/ -lovr -lglfw3 -lX11 -lXxf86vm -lXrandr -lpthread -lrt -lGL -lm -lstdc++
#include <OVR_CAPI.h>
#include <OVR_CAPI_GL.h>
*/
import "C"

import (

	"fmt"
	"runtime"
	"log"
	"./engine"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	//"github.com/go-gl/mathgl/mgl32"
	 "./common"  //This includes some shader loaders and program binding functions
	// "./ovr" 	   //Go wrapper for common OVR functions
)

var firstSprite engine.Sprite

func init() {
	runtime.LockOSThread()  //allows for OpenGL

}

/* *****************  MAIN FUNCTION  ************************ */
func main() {

	if err:=glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw")
	}
	
	defer glfw.Terminate()


	// *************  OPENGL / GLFW INIT CODE  ************** */

	
	glfw.WindowHint(glfw.Decorated, 0)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	
	window, err := glfw.CreateWindow(1000, 1000, "GoGL", nil, nil)

	if gl.Init(); err != nil {
		panic(err)
	}

	window.MakeContextCurrent();

	//Print OpenGL Version to console
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version, "\n\n")

	//keyboard input callback
	window.SetKeyCallback(onKey);
	spriteprogram := common.LoadShaderProgram("./shaders/vertexShader", "./shaders/fragmentShader")
	firstSprite.Init(-0.25,-0.25,0.5,0.5,spriteprogram)



	/* *******************   MAIN LOOP  ******************** */

	for !window.ShouldClose() {

		

		update()

		render(window, spriteprogram);



		//Poll for events (keyboard, resize, etc)
		


	}
	
}

func update() {


}

func render(w *glfw.Window, program uint32) {
	gl.ClearColor(0,0,0,1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	
	gl.UseProgram(program)
	gl.BindVertexArray(firstSprite.GetVAO())
	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	//Swap buffers
	w.SwapBuffers()
	glfw.PollEvents()
	gl.UseProgram(0)
}

/* *****************  KEYBOARD INPUT ******************** */
func onKey(w *glfw.Window, key glfw.Key, scancode int,
			action glfw.Action, mods glfw.ModifierKey) {

		if key == glfw.KeyEscape && action == glfw.Press {
			fmt.Println("Close Window")
			w.SetShouldClose(true)
		}
}



