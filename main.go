
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

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)


const WindowWidth = 1000
const WindowHeight = 800

func init() {

	runtime.LockOSThread()
}

func onKey(w *glfw.Window, key glfw.Key, scancode int,
			action glfw.Action, mods glfw.ModifierKey) {

		if key == glfw.KeyEscape && action == glfw.Press {
			fmt.Println("Close Window")
			w.SetShouldClose(true)
		}
}
func main() {

	
	// *************  OPENGL / GLFW INIT CODE  ****************

	if err:=glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	
	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "LinuxVR", nil, nil)

	if gl.Init(); err != nil {
		panic(err)
	}

	window.MakeContextCurrent();

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version)

	//keyboard input callback
	window.SetKeyCallback(onKey);

	/* **************************************************** */



	/* ****************    OVR INIT CODE  ***************** */

	C.ovr_Initialize(nil)

	hmdCount := (int)(C.ovrHmd_Detect())
	fmt.Println(hmdCount)
	fmt.Printf("Found " + string(hmdCount) + " connected Rift device(s)\n")

	for i:= 0; i < hmdCount; i++ {
		hmd := C.ovrHmd_Create((C.int)(i))
		fmt.Println(C.GoString(hmd.ProductName))
		C.ovrHmd_Destroy(hmd);

	}

	hmd := C.ovrHmd_CreateDebug(6)
	fmt.Println(C.GoString(hmd.ProductName))
	C.ovrHmd_Destroy(hmd)

	

	/* ***************************************************** */

	//previousTime := glfw.GetTime()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		//Update
		//time := glfw.GetTime();
		//elapsed := time - previousTime
		//previousTime = time


		//Swap buffers
		window.SwapBuffers()
		glfw.PollEvents()

	}


	/* *****************  OVR SHUTDOWN  ******************** */
	C.ovr_Shutdown()
	
}