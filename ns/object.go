package ns

import (
	"github.com/ysh86/objc"
)

type classObject struct {
	id objc.ID // isa member/class object
}

// InstanceObject corresponds to an instance of NSObject class.
type InstanceObject struct {
	classObject
	// isa: classObject.id
	// opaque members
	// ...
}

type cachedClassObject struct {
	classObject // class methods
	// class object: classObject.id

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
	Object = &cachedClassObject{classObject: classObject{
		id: objc.LookUpClass("NSObject"),
	},
		alloc:   objc.SelRegisterName("alloc"),
		class:   objc.SelRegisterName("class"),
		init:    objc.SelRegisterName("init"),
		release: objc.SelRegisterName("release"),
	}
}

func (c *classObject) Alloc() *InstanceObject {
	id := objc.MsgSend(c.id, Object.alloc)
	return (*InstanceObject)(id)
}

func (o *classObject) Class() objc.Class {
	if o == nil {
		return nil
	}

	id := objc.MsgSend(objc.ID(o), Object.class)
	return objc.Class(id)
}

func (o *classObject) Init() *InstanceObject {
	if o == nil {
		return nil
	}

	id := objc.MsgSend(objc.ID(o), Object.init)
	return (*InstanceObject)(id)
}

func (o *classObject) Release() {
	if o != nil {
		objc.MsgSend(objc.ID(o), Object.release)
	}
}
