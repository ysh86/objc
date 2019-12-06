package main

import (
	"fmt"

	"github.com/ysh86/objc"
	"github.com/ysh86/objc/ns"
)

func main() {
	co := objc.LookUpClass("NSObject")
	cs := objc.LookUpClass("NSString")
	fmt.Printf("NSObject [%s class]: %p, %#v\n", objc.ClassGetName(objc.Class(co)), co, (*ns.InstanceObject)(co))
	fmt.Printf("NSString [%s class]: %p, %#v\n", objc.ClassGetName(objc.Class(cs)), cs, (*ns.InstanceString)(cs))

	oo := ns.Object.Alloc()
	ioo := (*ns.InstanceObject)(oo)
	coo := ioo.Class()
	fmt.Printf("NSObject [%s alloc]: %p, %#v\n", objc.ClassGetName(coo), ioo, ioo)

	ioo.Release()
	ioo = nil

	ooo := ns.String.Alloc()
	iss := (*ns.InstanceString)(ooo)
	css := iss.Class()
	fmt.Printf("NSString  [%s alloc]: %p, %#v\n", objc.ClassGetName(css), iss, iss)
	oooo := iss.Init()
	iss = (*ns.InstanceString)(objc.ID(oooo))
	css = iss.Class()
	fmt.Printf("NSString [%s alloc] init]: %p, %#v\n", objc.ClassGetName(css), iss, iss)

	sss := iss.Alloc().Init()
	csss := sss.Class()
	fmt.Printf("NSString [%s alloc] init]: %p, %#v\n", objc.ClassGetName(csss), sss, sss)

	iss.Release()
	iss = nil
	sss.Release()
	sss = nil

	dele := ns.Object.Alloc().Init()

	dele.Release()
}
