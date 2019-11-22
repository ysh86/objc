package ns

import (
	"private/objc"
	"unsafe"
)

type metaObject struct {
}
type classObject struct {
	metaObject
	isa objc.Class // TODO: これは何を指しているのか？alloc, init で値が変わる
	// opaque
	// ...
}

type hiddenObject struct {
	metaObject

	id    objc.ID
	alloc objc.SEL
}

// Object is a dummy instance of the meta class of NSObject.
// It is used for class methods.
var Object *hiddenObject

func init() {
	Object = &hiddenObject{
		id:    objc.LookUpClass("NSObject"),
		alloc: objc.SelRegisterName("alloc"),
	}
}

func (c *metaObject) Alloc() *classObject {
	id := objc.MsgSend(Object.id, Object.alloc)
	return (*classObject)(unsafe.Pointer(id))
}
