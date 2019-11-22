package main

import (
	"fmt"

	"private/objc"
	"private/objc/ns"
)

func main() {
	cn := objc.LookUpClass("NSObject")
	cs := objc.LookUpClass("NSString")
	fmt.Printf("NSObject class: %p, %#v\n", cn, cn)
	fmt.Printf("NSString class: %p, %#v\n", cs, cs)

	o := ns.Object.Alloc()
	fmt.Printf("NSObject alloc: %#v\n", o)

	s := ns.String.Alloc()
	fmt.Printf("NSString alloc: %#v\n", s)

	//ss := s.Init()
	ss := s.Alloc().Init()
	fmt.Printf("NSString init: %#v\n", ss)
}
