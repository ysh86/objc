package ns

import (
	"private/objc"
	"unsafe"
)

type metaString struct {
	metaObject
}
type classString struct {
	metaString
	classObject
	// opaque
	// ...
}

type hiddenString struct {
	metaString

	id    objc.ID
	alloc objc.SEL

	init objc.SEL
}

// String is a dummy instance of the meta class of NSString.
// It is used for class methods.
var String *hiddenString

func init() {
	String = &hiddenString{
		id:    objc.LookUpClass("NSString"),
		alloc: objc.SelRegisterName("alloc"),
		init:  objc.SelRegisterName("init"),
	}
}

func (c *metaString) Alloc() *classString {
	id := objc.MsgSend(String.id, String.alloc)
	return (*classString)(unsafe.Pointer(id))
}

func (o *classString) Init() *classString {
	if o == nil {
		return o
	}

	id := objc.MsgSend(objc.ID(unsafe.Pointer(o)), String.init)
	return (*classString)(unsafe.Pointer(id))
}
