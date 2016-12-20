package crank

// #include <windows.h>
import "C"

func Wake(reason string) {
	C.SetThreadExecutionState(C.ES_CONTINUOUS | C.ES_SYSTEM_REQUIRED);
}
