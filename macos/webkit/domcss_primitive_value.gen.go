// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [DOMCSSPrimitiveValue] class.
var DOMCSSPrimitiveValueClass = _DOMCSSPrimitiveValueClass{objc.GetClass("DOMCSSPrimitiveValue")}

type _DOMCSSPrimitiveValueClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSPrimitiveValue] class.
type IDOMCSSPrimitiveValue interface {
	IDOMCSSValue
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssprimitivevalue?language=objc
type DOMCSSPrimitiveValue struct {
	DOMCSSValue
}

func DOMCSSPrimitiveValueFrom(ptr unsafe.Pointer) DOMCSSPrimitiveValue {
	return DOMCSSPrimitiveValue{
		DOMCSSValue: DOMCSSValueFrom(ptr),
	}
}

func (dc _DOMCSSPrimitiveValueClass) Alloc() DOMCSSPrimitiveValue {
	rv := objc.Call[DOMCSSPrimitiveValue](dc, objc.Sel("alloc"))
	return rv
}

func (dc _DOMCSSPrimitiveValueClass) New() DOMCSSPrimitiveValue {
	rv := objc.Call[DOMCSSPrimitiveValue](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSPrimitiveValue() DOMCSSPrimitiveValue {
	return DOMCSSPrimitiveValueClass.New()
}

func (d_ DOMCSSPrimitiveValue) Init() DOMCSSPrimitiveValue {
	rv := objc.Call[DOMCSSPrimitiveValue](d_, objc.Sel("init"))
	return rv
}
