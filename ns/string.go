package ns

import (
	"github.com/ysh86/objc"
)

type classString struct {
	classObject // super & isa member/class object
}

// InstanceString corresponds to an instance of NSString class.
type InstanceString struct {
	classString
	// isa: classString.id
	// opaque members
	// ...
}

type cachedClassString struct {
	classString // class methods
	// class object: classString.id

	// class methods
	// ...
	// instance methods
	// ...
}

// String is a dummy(cached) object of the NSString class.
// It is used for class methods & selector cache.
var String *cachedClassString

func init() {
	String = &cachedClassString{classString: classString{classObject: classObject{
		id: objc.LookUpClass("NSString"),
	}},
	// methods...
	}
}

// Alloc just casts the result from the super.
func (c *classString) Alloc() *InstanceString {
	obj := c.classObject.Alloc()
	return (*InstanceString)(objc.ID(obj))
}

// Init just casts the result from the super.
func (o *classString) Init() *InstanceString {
	if o == nil {
		return nil
	}

	obj := o.classObject.Init()
	return (*InstanceString)(objc.ID(obj))
}
