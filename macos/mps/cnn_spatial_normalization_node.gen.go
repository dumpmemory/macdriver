// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSpatialNormalizationNode] class.
var CNNSpatialNormalizationNodeClass = _CNNSpatialNormalizationNodeClass{objc.GetClass("MPSCNNSpatialNormalizationNode")}

type _CNNSpatialNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNSpatialNormalizationNode] class.
type ICNNSpatialNormalizationNode interface {
	ICNNNormalizationNode
	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
}

// A representation of a spatial normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode?language=objc
type CNNSpatialNormalizationNode struct {
	CNNNormalizationNode
}

func CNNSpatialNormalizationNodeFrom(ptr unsafe.Pointer) CNNSpatialNormalizationNode {
	return CNNSpatialNormalizationNode{
		CNNNormalizationNode: CNNNormalizationNodeFrom(ptr),
	}
}

func (c_ CNNSpatialNormalizationNode) InitWithSource(sourceNode INNImageNode) CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866502-initwithsource?language=objc
func NewCNNSpatialNormalizationNodeWithSource(sourceNode INNImageNode) CNNSpatialNormalizationNode {
	instance := CNNSpatialNormalizationNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (c_ CNNSpatialNormalizationNode) InitWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](c_, objc.Sel("initWithSource:kernelSize:"), objc.Ptr(sourceNode), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866438-initwithsource?language=objc
func NewCNNSpatialNormalizationNodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNSpatialNormalizationNode {
	instance := CNNSpatialNormalizationNodeClass.Alloc().InitWithSourceKernelSize(sourceNode, kernelSize)
	instance.Autorelease()
	return instance
}

func (cc _CNNSpatialNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](cc, objc.Sel("nodeWithSource:kernelSize:"), objc.Ptr(sourceNode), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866401-nodewithsource?language=objc
func CNNSpatialNormalizationNode_NodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNSpatialNormalizationNode {
	return CNNSpatialNormalizationNodeClass.NodeWithSourceKernelSize(sourceNode, kernelSize)
}

func (cc _CNNSpatialNormalizationNodeClass) Alloc() CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNSpatialNormalizationNodeClass) New() CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSpatialNormalizationNode() CNNSpatialNormalizationNode {
	return CNNSpatialNormalizationNodeClass.New()
}

func (c_ CNNSpatialNormalizationNode) Init() CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNSpatialNormalizationNodeClass) NodeWithSource(sourceNode INNImageNode) CNNSpatialNormalizationNode {
	rv := objc.Call[CNNSpatialNormalizationNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866460-nodewithsource?language=objc
func CNNSpatialNormalizationNode_NodeWithSource(sourceNode INNImageNode) CNNSpatialNormalizationNode {
	return CNNSpatialNormalizationNodeClass.NodeWithSource(sourceNode)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866424-kernelheight?language=objc
func (c_ CNNSpatialNormalizationNode) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866424-kernelheight?language=objc
func (c_ CNNSpatialNormalizationNode) SetKernelHeight(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelHeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866402-kernelwidth?language=objc
func (c_ CNNSpatialNormalizationNode) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationnode/2866402-kernelwidth?language=objc
func (c_ CNNSpatialNormalizationNode) SetKernelWidth(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelWidth:"), value)
}