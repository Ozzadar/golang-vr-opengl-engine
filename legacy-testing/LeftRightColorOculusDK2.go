
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




func init() {

	init_resources();
	runtime.LockOSThread()
}

func init_resources() {

}

func draw(w *glfw.Window) {

	
	//Get the horizontal split size of the window
	sizex, sizey := w.GetSize()

	var eyesize mgl32.Vec2
	eyesize[0] = float32(sizex / 2.0)
	eyesize[1] = float32(sizey)

	gl.Enable(gl.SCISSOR_TEST)

	gl.Scissor(0,0,int32(eyesize[0]), int32(eyesize[1]))
	gl.ClearColor(1,0,0,1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Scissor(int32(eyesize[0]),0,int32(eyesize[0]), int32(eyesize[1]))
	gl.ClearColor(0,0,1,1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Disable(gl.SCISSOR_TEST)
	

	
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

	/* ****************    OVR INIT CODE  ***************** */

	if err:=glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw")
	}

	defer glfw.Terminate()
	C.ovr_Initialize(nil)

	//create an HMD for reference.
	var hmd C.ovrHmd = nil

	// find number of headsets
	hmdCount := (int)(C.ovrHmd_Detect())
	// print headset count
	fmt.Println(hmdCount)
	fmt.Printf("Found %d connected Rift device(s)\n\n", hmdCount)

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

		fmt.Println("Unable to open rift device\n Creating debug device.\n");
		hmd = C.ovrHmd_CreateDebug(C.ovrHmd_DK2)
	}

	//Starts the sensor device
	if C.ovrHmd_ConfigureTracking(hmd, C.ovrTrackingCap_Orientation | C.ovrTrackingCap_Position, 0) == 0 {
		fmt.Println("Unable to start Rift head tracker\n")

	}

	//extendedMode := C.ovrHmdCap_ExtendDesktop & hmd.HmdCaps

	//positioning of window and size of window
	var outposition mgl32.Vec2
	outposition[0] = (float32) (hmd.WindowsPos.x)
	outposition[1] = (float32) (hmd.WindowsPos.y)


	//TODO: Change this to output at chosen resolution, not necessarily native pg. 76 Oculus Rift in action
	var outsize mgl32.Vec2
	outsize[0] = (float32) (hmd.Resolution.w)
	outsize[1] = (float32) (hmd.Resolution.h)
	
	//print position and sizes to console
	fmt.Printf("Rift position:\t\t %f \t %f \nRift Size:\t\t %f \t %f \n\n", outposition.X(), outposition.Y(), outsize.X(), outsize.Y())



	monitors := glfw.GetMonitors()
	var riftIndex int;
	//loop over the monitors
	for index, element := range monitors {
		//print the monitor positions
		posX,posY := element.GetPos();
		fmt.Printf("Monitor Position:\t\t %d \t %d\n", posX, posY)

		if float32(posX) == outposition.X() && float32(posY) == outposition.Y() {

			riftIndex = index;
		}
	}

	//Get video mode of monitor
	mode := monitors[riftIndex].GetVideoMode()
	outsize[0] = float32(mode.Width)
	outsize[1] = float32(mode.Height)

	/* ***************************************************** */

	// *************  OPENGL / GLFW INIT CODE  ************** */

	
	glfw.WindowHint(glfw.Decorated, 0)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	
	window, err := glfw.CreateWindow(int(outsize.X()), int(outsize.Y()), "LinuxVR", nil, nil)
	window.SetPos(int(outposition.X()), int(outposition.Y()))

	if gl.Init(); err != nil {
		panic(err)
	}

	window.MakeContextCurrent();

	//Print OpenGL Version to console
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version, "\n\n")

	//keyboard input callback
	window.SetKeyCallback(onKey);

	/* **************************************************** */



	previousTime := glfw.GetTime()
	totalTime := 0.0000

	/* *******************   MAIN LOOP  ******************** */

	for !window.ShouldClose() {

		glfw.PollEvents()

		//Update
		time := glfw.GetTime();
		elapsed := time - previousTime
		previousTime = time
		totalTime = totalTime + elapsed

		//get current head state
		state := C.ovrHmd_GetTrackingState(hmd, 0)
		orientation := state.HeadPose.ThePose.Orientation

		//convert to go type float32
		var q mgl32.Quat
		q.W = (float32) (orientation.w)
		q.V[0] = (float32) (orientation.x)
		q.V[1] = (float32) (orientation.y)
		q.V[2] = (float32) (orientation.z)

		//publish tracking information once a second
		if totalTime >= 1 {
			fmt.Printf("w: %f X: %f Y: %f Z: %f\n", q.W, q.X(), q.Y(), q.Z())
			totalTime = 0
		}

		//basic opengl things



		draw(window);

		//Swap buffers
		window.SwapBuffers()

		//Poll for events (keyboard, resize, etc)
		


	}


	/* *****************  OVR SHUTDOWN  ******************** */
	C.ovrHmd_Destroy(hmd)
	C.ovr_Shutdown()
	
}
