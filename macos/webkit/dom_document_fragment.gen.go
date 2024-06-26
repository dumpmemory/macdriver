// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [DOMDocumentFragment] class.
var DOMDocumentFragmentClass = _DOMDocumentFragmentClass{objc.GetClass("DOMDocumentFragment")}

type _DOMDocumentFragmentClass struct {
	objc.Class
}

// An interface definition for the [DOMDocumentFragment] class.
type IDOMDocumentFragment interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domdocumentfragment?language=objc
type DOMDocumentFragment struct {
	DOMNode
}

func DOMDocumentFragmentFrom(ptr unsafe.Pointer) DOMDocumentFragment {
	return DOMDocumentFragment{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMDocumentFragmentClass) Alloc() DOMDocumentFragment {
	rv := objc.Call[DOMDocumentFragment](dc, objc.Sel("alloc"))
	return rv
}

func (dc _DOMDocumentFragmentClass) New() DOMDocumentFragment {
	rv := objc.Call[DOMDocumentFragment](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMDocumentFragment() DOMDocumentFragment {
	return DOMDocumentFragmentClass.New()
}

func (d_ DOMDocumentFragment) Init() DOMDocumentFragment {
	rv := objc.Call[DOMDocumentFragment](d_, objc.Sel("init"))
	return rv
}
