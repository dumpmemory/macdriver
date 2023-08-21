// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DirectoryEnumerator] class.
var DirectoryEnumeratorClass = _DirectoryEnumeratorClass{objc.GetClass("NSDirectoryEnumerator")}

type _DirectoryEnumeratorClass struct {
	objc.Class
}

// An interface definition for the [DirectoryEnumerator] class.
type IDirectoryEnumerator interface {
	IEnumerator
	SkipDescendents()
	SkipDescendants()
	Level() uint
	IsEnumeratingDirectoryPostOrder() bool
	DirectoryAttributes() map[FileAttributeKey]objc.Object
	FileAttributes() map[FileAttributeKey]objc.Object
}

// An object that enumerates the contents of a directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator?language=objc
type DirectoryEnumerator struct {
	Enumerator
}

func DirectoryEnumeratorFrom(ptr unsafe.Pointer) DirectoryEnumerator {
	return DirectoryEnumerator{
		Enumerator: EnumeratorFrom(ptr),
	}
}

func (dc _DirectoryEnumeratorClass) Alloc() DirectoryEnumerator {
	rv := objc.Call[DirectoryEnumerator](dc, objc.Sel("alloc"))
	return rv
}

func DirectoryEnumerator_Alloc() DirectoryEnumerator {
	return DirectoryEnumeratorClass.Alloc()
}

func (dc _DirectoryEnumeratorClass) New() DirectoryEnumerator {
	rv := objc.Call[DirectoryEnumerator](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDirectoryEnumerator() DirectoryEnumerator {
	return DirectoryEnumeratorClass.New()
}

func (d_ DirectoryEnumerator) Init() DirectoryEnumerator {
	rv := objc.Call[DirectoryEnumerator](d_, objc.Sel("init"))
	return rv
}

// Causes the receiver to skip recursion into the most recently obtained subdirectory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator/1412990-skipdescendents?language=objc
func (d_ DirectoryEnumerator) SkipDescendents() {
	objc.Call[objc.Void](d_, objc.Sel("skipDescendents"))
}

// Causes the receiver to skip recursion into the most recently obtained subdirectory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator/1409644-skipdescendants?language=objc
func (d_ DirectoryEnumerator) SkipDescendants() {
	objc.Call[objc.Void](d_, objc.Sel("skipDescendants"))
}

// The number of levels deep the current object is in the directory hierarchy being enumerated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator/1408465-level?language=objc
func (d_ DirectoryEnumerator) Level() uint {
	rv := objc.Call[uint](d_, objc.Sel("level"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator/3088819-isenumeratingdirectorypostorder?language=objc
func (d_ DirectoryEnumerator) IsEnumeratingDirectoryPostOrder() bool {
	rv := objc.Call[bool](d_, objc.Sel("isEnumeratingDirectoryPostOrder"))
	return rv
}

// A dictionary with the attributes of the directory at which enumeration started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator/1411357-directoryattributes?language=objc
func (d_ DirectoryEnumerator) DirectoryAttributes() map[FileAttributeKey]objc.Object {
	rv := objc.Call[map[FileAttributeKey]objc.Object](d_, objc.Sel("directoryAttributes"))
	return rv
}

// A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdirectoryenumerator/1413284-fileattributes?language=objc
func (d_ DirectoryEnumerator) FileAttributes() map[FileAttributeKey]objc.Object {
	rv := objc.Call[map[FileAttributeKey]objc.Object](d_, objc.Sel("fileAttributes"))
	return rv
}