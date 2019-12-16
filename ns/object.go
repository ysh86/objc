package ns

import (
	"unsafe"

	"github.com/ysh86/objc"
)

type classObject struct {
}

// InstanceObject corresponds to an instance of NSObject class.
type InstanceObject struct {
	classObject // class methods
	isa         objc.ID
	// opaque members
	// ...
}

type cachedClassObject struct {
	classObject         // class methods
	cls         objc.ID // class object

	// class methods
	alloc objc.SEL
	// instance methods
	class   objc.SEL
	init    objc.SEL
	release objc.SEL
}

// Object is a dummy(cached) object of the NSObject class.
// It is used for class methods & selector chache.
var Object *cachedClassObject

func init() {
	Object = &cachedClassObject{
		cls:     objc.LookUpClass("NSObject"),
		alloc:   objc.SelRegisterName("alloc"),
		class:   objc.SelRegisterName("class"),
		init:    objc.SelRegisterName("init"),
		release: objc.SelRegisterName("release"),
	}
}

func (c *classObject) Alloc() *InstanceObject {
	cls := ((*cachedClassObject)(unsafe.Pointer(c))).cls
	id := objc.MsgSend(cls, Object.alloc)
	return (*InstanceObject)(id)
}

func (o *InstanceObject) Class() objc.Class {
	if o == nil {
		return nil
	}

	id := objc.MsgSend(objc.ID(o), Object.class)
	return objc.Class(id)
}

func (o *InstanceObject) Init() *InstanceObject {
	if o == nil {
		return nil
	}

	id := objc.MsgSend(objc.ID(o), Object.init)
	return (*InstanceObject)(id)
}

func (o *InstanceObject) Release() {
	if o != nil {
		objc.MsgSend(objc.ID(o), Object.release)
	}
}
