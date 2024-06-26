// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNLossDataDescriptor] class.
var CNNLossDataDescriptorClass = _CNNLossDataDescriptorClass{objc.GetClass("MPSCNNLossDataDescriptor")}

type _CNNLossDataDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CNNLossDataDescriptor] class.
type ICNNLossDataDescriptor interface {
	objc.IObject
	Layout() DataLayout
	BytesPerImage() uint
	SetBytesPerImage(value uint)
	Size() metal.Size
	BytesPerRow() uint
	SetBytesPerRow(value uint)
}

// An object that specifies properties used by a loss data descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor?language=objc
type CNNLossDataDescriptor struct {
	objc.Object
}

func CNNLossDataDescriptorFrom(ptr unsafe.Pointer) CNNLossDataDescriptor {
	return CNNLossDataDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CNNLossDataDescriptorClass) Alloc() CNNLossDataDescriptor {
	rv := objc.Call[CNNLossDataDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNLossDataDescriptorClass) New() CNNLossDataDescriptor {
	rv := objc.Call[CNNLossDataDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLossDataDescriptor() CNNLossDataDescriptor {
	return CNNLossDataDescriptorClass.New()
}

func (c_ CNNLossDataDescriptor) Init() CNNLossDataDescriptor {
	rv := objc.Call[CNNLossDataDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951840-cnnlossdatadescriptorwithdata?language=objc
func (cc _CNNLossDataDescriptorClass) CnnLossDataDescriptorWithDataLayoutSize(data []byte, layout DataLayout, size metal.Size) CNNLossDataDescriptor {
	rv := objc.Call[CNNLossDataDescriptor](cc, objc.Sel("cnnLossDataDescriptorWithData:layout:size:"), data, layout, size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951840-cnnlossdatadescriptorwithdata?language=objc
func CNNLossDataDescriptor_CnnLossDataDescriptorWithDataLayoutSize(data []byte, layout DataLayout, size metal.Size) CNNLossDataDescriptor {
	return CNNLossDataDescriptorClass.CnnLossDataDescriptorWithDataLayoutSize(data, layout, size)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951842-layout?language=objc
func (c_ CNNLossDataDescriptor) Layout() DataLayout {
	rv := objc.Call[DataLayout](c_, objc.Sel("layout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951847-bytesperimage?language=objc
func (c_ CNNLossDataDescriptor) BytesPerImage() uint {
	rv := objc.Call[uint](c_, objc.Sel("bytesPerImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951847-bytesperimage?language=objc
func (c_ CNNLossDataDescriptor) SetBytesPerImage(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setBytesPerImage:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951848-size?language=objc
func (c_ CNNLossDataDescriptor) Size() metal.Size {
	rv := objc.Call[metal.Size](c_, objc.Sel("size"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951849-bytesperrow?language=objc
func (c_ CNNLossDataDescriptor) BytesPerRow() uint {
	rv := objc.Call[uint](c_, objc.Sel("bytesPerRow"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951849-bytesperrow?language=objc
func (c_ CNNLossDataDescriptor) SetBytesPerRow(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setBytesPerRow:"), value)
}
