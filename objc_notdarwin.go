// +build !darwin

package objc

import (
	"unsafe"
)

type Class unsafe.Pointer // = C.Class
type ID unsafe.Pointer // = C.id
type SEL unsafe.Pointer // = C.SEL

func LookUpClass(name string) ID {
	return nil
}

func ClassGetName(cls Class) string {
	return ""
}

func SelRegisterName(str string) SEL {
	return nil
}

func MsgSend(self ID, op SEL, args ...interface{}) ID {
	return nil
}
