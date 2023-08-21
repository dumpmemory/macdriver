// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureBoundingBoxGeometryDescriptor] class.
var AccelerationStructureBoundingBoxGeometryDescriptorClass = _AccelerationStructureBoundingBoxGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureBoundingBoxGeometryDescriptor")}

type _AccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureBoundingBoxGeometryDescriptor] class.
type IAccelerationStructureBoundingBoxGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	BoundingBoxBufferOffset() uint
	SetBoundingBoxBufferOffset(value uint)
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	BoundingBoxBuffer() BufferWrapper
	SetBoundingBoxBuffer(value PBuffer)
	SetBoundingBoxBufferObject(valueObject objc.IObject)
}

// A description of a list of bounding boxes to turn into an acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor?language=objc
type AccelerationStructureBoundingBoxGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

func AccelerationStructureBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureBoundingBoxGeometryDescriptor {
	return AccelerationStructureBoundingBoxGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) Descriptor() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureBoundingBoxGeometryDescriptor](ac, objc.Sel("descriptor"))
	return rv
}

// Creates a new bounding box descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553864-descriptor?language=objc
func AccelerationStructureBoundingBoxGeometryDescriptor_Descriptor() AccelerationStructureBoundingBoxGeometryDescriptor {
	return AccelerationStructureBoundingBoxGeometryDescriptorClass.Descriptor()
}

func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureBoundingBoxGeometryDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructureBoundingBoxGeometryDescriptor_Alloc() AccelerationStructureBoundingBoxGeometryDescriptor {
	return AccelerationStructureBoundingBoxGeometryDescriptorClass.Alloc()
}

func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) New() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureBoundingBoxGeometryDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureBoundingBoxGeometryDescriptor() AccelerationStructureBoundingBoxGeometryDescriptor {
	return AccelerationStructureBoundingBoxGeometryDescriptorClass.New()
}

func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) Init() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureBoundingBoxGeometryDescriptor](a_, objc.Sel("init"))
	return rv
}

// The offset, in bytes, to the first bounding box in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553861-boundingboxbufferoffset?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBufferOffset() uint {
	rv := objc.Call[uint](a_, objc.Sel("boundingBoxBufferOffset"))
	return rv
}

// The offset, in bytes, to the first bounding box in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553861-boundingboxbufferoffset?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBufferOffset(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxBufferOffset:"), value)
}

// The stride, in bytes, between bounding boxes in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553863-boundingboxstride?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Call[uint](a_, objc.Sel("boundingBoxStride"))
	return rv
}

// The stride, in bytes, between bounding boxes in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553863-boundingboxstride?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxStride:"), value)
}

// The number of bounding boxes in the bounding box buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553862-boundingboxcount?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Call[uint](a_, objc.Sel("boundingBoxCount"))
	return rv
}

// The number of bounding boxes in the bounding box buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553862-boundingboxcount?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxCount:"), value)
}

// A buffer that contains bounding box data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553860-boundingboxbuffer?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](a_, objc.Sel("boundingBoxBuffer"))
	return rv
}

// A buffer that contains bounding box data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553860-boundingboxbuffer?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxBuffer:"), po0)
}

// A buffer that contains bounding box data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureboundingboxgeometrydescriptor/3553860-boundingboxbuffer?language=objc
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxBuffer:"), objc.Ptr(valueObject))
}