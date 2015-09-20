
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
	"github.com/go-gl/mathgl/mgl32"
)


const WindowWidth = 1000
const WindowHeight = 800

func init() {

	runtime.LockOSThread()
}

	/* *****************  KEYBOARD INPUT ******************** */
func onKey(w *glfw.Window, key glfw.Key, scancode int,
			action glfw.Action, mods glfw.ModifierKey) {

		if key == glfw.KeyEscape && action == glfw.Press {
			fmt.Println("Close Window")
			w.SetShouldClose(true)
		}
}


/* *****************  MAIN FUNCTION  ************************ */
func main() {

	
	// *************  OPENGL / GLFW INIT CODE  ************** */

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

	//Print OpenGL Version to console
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version)

	//keyboard input callback
	window.SetKeyCallback(onKey);

	/* **************************************************** */



	/* ****************    OVR INIT CODE  ***************** */

	C.ovr_Initialize(nil)

	//create an HMD for reference.
	var hmd C.ovrHmd = nil

	// find number of headsets
	hmdCount := (int)(C.ovrHmd_Detect())
	// print headset count
	fmt.Println(hmdCount)
	fmt.Printf("Found " + string(hmdCount) + " connected Rift device(s)\n")

	// grab the first headset
	if hmdCount > 0 {
		for i:= 0; i < 1; i++ {
			hmd = C.ovrHmd_Create((C.int)(i))
			//Print headset name
			fmt.Println(C.GoString(hmd.ProductName))
		}
	}

	//if there is no headset connected, create a new debug.
	if hmd == nil {

		fmt.Println("Unable to open rift device\n Creating debug device.");
		hmd = C.ovrHmd_CreateDebug(C.ovrHmd_DK2)
	}

	//Starts the sensor device
	if C.ovrHmd_ConfigureTracking(hmd, C.ovrTrackingCap_Orientation, 0) == 0 {
		fmt.Println("Unable to start Rift head tracker")

	}



	

	/* ***************************************************** */

	//previousTime := glfw.GetTime()


	/* *******************   MAIN LOOP  ******************** */

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		
		//Update
		//time := glfw.GetTime();
		//elapsed := time - previousTime
		//previousTime = time

		state := C.ovrHmd_GetTrackingState(hmd, 0)
		orientation := state.HeadPose.ThePose.Orientation

		var q mgl32.Quat
		q.W = (float32) (orientation.w)
		q.V[0] = (float32) (orientation.x)
		q.V[1] = (float32) (orientation.y)
		q.V[2] = (float32) (orientation.z)


		fmt.Printf("w: %f X: %f Y: %f Z: %f\n", q.W, q.X(), q.Y(), q.Z())
		//Swap buffers
		window.SwapBuffers()
		//Poll for events (keyboard, resize, etc)
		glfw.PollEvents()

	}


	/* *****************  OVR SHUTDOWN  ******************** */
	C.ovrHmd_Destroy(hmd)
	C.ovr_Shutdown()
	
}