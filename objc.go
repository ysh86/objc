// +build darwin

package objc

/*
#cgo LDFLAGS: -framework Cocoa
#include <objc/objc-class.h>
#include <stdlib.h> // for free()

static inline id objc_msgSend0(id self, SEL op) {
	return objc_msgSend(self, op);
}
*/
import "C"

import (
	"fmt"
	"io"
	"unsafe"
)

type Class C.Class
type ID C.id
type SEL C.SEL

func LookUpClass(name string) ID {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))

	return ID(unsafe.Pointer(C.objc_lookUpClass(n)))
}

func SelRegisterName(str string) SEL {
	s := C.CString(str)
	defer C.free(unsafe.Pointer(s))

	return SEL(C.sel_registerName(s))
}

func MsgSend(self ID, op SEL, args ...interface{}) ID {
	id := C.objc_msgSend0(C.id(self), C.SEL(op))
	return ID(id)
}

// -------------------- trial ------------------------

func dumpAllMethods(w io.Writer, cls C.Class) {
	var count C.uint
	methodsC := C.class_copyMethodList(cls, &count)
	defer C.free(unsafe.Pointer(methodsC))

	fmt.Fprintf(w, "  count: %d\n", count)
	methods := (*[1<<31 - 1]C.Method)(unsafe.Pointer(methodsC))[:count:count] // no copy
	for i, m := range methods {
		selC := C.method_getName(m)
		nameC := C.sel_getName(selC)
		name := C.GoString(nameC)
		typesC := C.method_getTypeEncoding(m)
		types := C.GoString(typesC)
		fmt.Fprintf(w, "  %d: %s, %s\n", i, name, types)
	}
}

func dumpAllClasses(w io.Writer) {
	numClasses := C.objc_getClassList(nil, 0)
	fmt.Fprintf(w, "numClasses: %d\n", numClasses)

	classes := make([]C.Class, numClasses)
	numClasses = C.objc_getClassList(&classes[0], numClasses)

	for i, cls := range classes {
		nameC := C.class_getName(cls)
		name := C.GoString(nameC)
		fmt.Fprintf(w, "%d: %s\n", i, name)

		if name == "NSString" {
			dumpAllMethods(w, cls)
		}
	}
}
