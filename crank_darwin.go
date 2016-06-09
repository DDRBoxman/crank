package crank

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework IOKit
#import <Foundation/Foundation.h>
#import <IOKit/pwr_mgt/IOPMLib.h>

static IOPMAssertionID assertionID = 0;

void
wake(char* reason) {
	if (assertionID) return;

	CFStringRef reasonForActivity= CFStringCreateWithCString(NULL, reason, kCFStringEncodingUTF8);

	IOReturn success = IOPMAssertionCreateWithName(kIOPMAssertionTypeNoDisplaySleep,
                                    kIOPMAssertionLevelOn, reasonForActivity, &assertionID);
    free(reason);
}

void
releaseWake() {
	if (!assertionID) return;

	IOPMAssertionRelease(assertionID);

	assertionID = 0;
}

*/
import "C"

func Wake(reason string) {
	C.wake(C.CString(reason))
}

func ReleaseWake() {
	C.releaseWake()
}
