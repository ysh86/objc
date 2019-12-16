package ns

import (
	"github.com/ysh86/objc"
)

type classString struct {
	classObject // super
}

// InstanceString corresponds to an instance of NSString class.
type InstanceString struct {
	classString    // class methods
	InstanceObject // super & isa member
	// opaque members
	// ...
}

type cachedClassString struct {
	classString         // class methods
	cls         objc.ID // class object

	// class methods
	// ...
	// instance methods
	// ...
}

// String is a dummy(cached) object of the NSString class.
// It is used for class methods & selector cache.
var String *cachedClassString

func init() {
	String = &cachedClassString{
		cls: objc.LookUpClass("NSString"),
		// methods...
	}
}

// Alloc just casts the result from the super.
func (c *classString) Alloc() *InstanceString {
	obj := c.classObject.Alloc()
	return (*InstanceString)(objc.ID(obj))
}

// Init just casts the result from the super.
func (o *InstanceString) Init() *InstanceString {
	if o == nil {
		return nil
	}

	obj := o.InstanceObject.Init()
	return (*InstanceString)(objc.ID(obj))
}
