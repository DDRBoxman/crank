package crank

// #include <windows.h>
import "C"

func Wake() {
	C.SetThreadExecutionState(ES_CONTINUOUS | ES_SYSTEM_REQUIRED);
}
