package crank

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework IOKit
#import <Foundation/Foundation.h>
#import <IOKit/pwr_mgt/IOPMLib.h>

void
wake() {
	CFStringRef* reasonForActivity= CFSTR("Describe Activity Type");
	IOPMAssertionID assertionID;
IOReturn success = IOPMAssertionCreateWithName(kIOPMAssertionTypeNoDisplaySleep,
                                    kIOPMAssertionLevelOn, reasonForActivity, &assertionID);
if (success == kIOReturnSuccess)
{
 
    //Add the work you need to do without
    //  the system sleeping here.
 
    success = IOPMAssertionRelease(assertionID);
    //The system will be able to sleep again.
}
}

void
releaseWake() {

}

*/
import "C"

func Wake() {
	C.wake()
}

func ReleaseWake() {
	C.releaseWake()
}
