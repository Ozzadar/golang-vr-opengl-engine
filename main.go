
package main
//Let's try importing the ovr c library

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/lib/ -L/usr/local/lib/ -lovr -lglfw3 -lX11 -lXxf86vm -lXrandr -lpthread -lrt -lGL -lm -lstdc++
#include <OVR_CAPI.h>
#include <OVR_CAPI_GL.h>
*/
import "C"


import "fmt"

func main() {

	fmt.Printf("Importing worked.\n")

	C.ovr_Initialize(nil)

	hmdCount := (int)(C.ovrHmd_Detect())
	fmt.Println(hmdCount)
	fmt.Printf("Found " + string(hmdCount) + " connected Rift device(s)\n")

	for i:= 0; i < hmdCount; i++ {
		hmd := C.ovrHmd_Create((C.int)(i))
		fmt.Printf(C.GoString(hmd.ProductName))
		C.ovrHmd_Destroy(hmd);

	}


	hmd := C.ovrHmd_CreateDebug(6)
	fmt.Printf(C.GoString(hmd.ProductName))
	C.ovrHmd_Destroy(hmd)

	C.ovr_Shutdown();
	
}