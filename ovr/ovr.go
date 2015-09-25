/* OVR SPECIFIC FUNCTIONS AND WRAPPER THINGS */
package ovr


//Let's try importing the ovr c library

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/lib/ -L/usr/local/lib/ -lovr -lglfw3 -lX11 -lXxf86vm -lXrandr -lpthread -lrt -lGL -lm -lstdc++
#include <OVR_CAPI.h>
#include <OVR_CAPI_GL.h>
*/
import "C"