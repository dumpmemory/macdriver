// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNArithmeticGradientNode] class.
var NNArithmeticGradientNodeClass = _NNArithmeticGradientNodeClass{objc.GetClass("MPSNNArithmeticGradientNode")}

type _NNArithmeticGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNArithmeticGradientNode] class.
type INNArithmeticGradientNode interface {
	INNGradientFilterNode
	MaximumValue() float32
	SetMaximumValue(value float32)
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	Bias() float32
	SetBias(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	IsSecondarySourceFilter() bool
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
}

// A representation of the base class for gradient arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode?language=objc
type NNArithmeticGradientNode struct {
	NNGradientFilterNode
}

func NNArithmeticGradientNodeFrom(ptr unsafe.Pointer) NNArithmeticGradientNode {
	return NNArithmeticGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNArithmeticGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](n_, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), gradientImages, filter, isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952980-initwithgradientimages?language=objc
func NewNNArithmeticGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	instance := NNArithmeticGradientNodeClass.Alloc().InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages, filter, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (n_ NNArithmeticGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956166-initwithsourcegradient?language=objc
func NewNNArithmeticGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	instance := NNArithmeticGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNArithmeticGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956167-nodewithsourcegradient?language=objc
func NNArithmeticGradientNode_NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	return NNArithmeticGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
}

func (nc _NNArithmeticGradientNodeClass) Alloc() NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNArithmeticGradientNodeClass) New() NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNArithmeticGradientNode() NNArithmeticGradientNode {
	return NNArithmeticGradientNodeClass.New()
}

func (n_ NNArithmeticGradientNode) Init() NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952986-maximumvalue?language=objc
func (n_ NNArithmeticGradientNode) MaximumValue() float32 {
	rv := objc.Call[float32](n_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952986-maximumvalue?language=objc
func (n_ NNArithmeticGradientNode) SetMaximumValue(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952994-secondarystrideinpixelsy?language=objc
func (n_ NNArithmeticGradientNode) SecondaryStrideInPixelsY() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952994-secondarystrideinpixelsy?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryStrideInPixelsY(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952968-secondarystrideinpixelsx?language=objc
func (n_ NNArithmeticGradientNode) SecondaryStrideInPixelsX() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952968-secondarystrideinpixelsx?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryStrideInPixelsX(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952981-secondaryscale?language=objc
func (n_ NNArithmeticGradientNode) SecondaryScale() float32 {
	rv := objc.Call[float32](n_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952981-secondaryscale?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryScale(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952988-bias?language=objc
func (n_ NNArithmeticGradientNode) Bias() float32 {
	rv := objc.Call[float32](n_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952988-bias?language=objc
func (n_ NNArithmeticGradientNode) SetBias(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952989-minimumvalue?language=objc
func (n_ NNArithmeticGradientNode) MinimumValue() float32 {
	rv := objc.Call[float32](n_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952989-minimumvalue?language=objc
func (n_ NNArithmeticGradientNode) SetMinimumValue(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952987-issecondarysourcefilter?language=objc
func (n_ NNArithmeticGradientNode) IsSecondarySourceFilter() bool {
	rv := objc.Call[bool](n_, objc.Sel("isSecondarySourceFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952984-secondarystrideinfeaturechannels?language=objc
func (n_ NNArithmeticGradientNode) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952984-secondarystrideinfeaturechannels?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952993-primaryscale?language=objc
func (n_ NNArithmeticGradientNode) PrimaryScale() float32 {
	rv := objc.Call[float32](n_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952993-primaryscale?language=objc
func (n_ NNArithmeticGradientNode) SetPrimaryScale(value float32) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimaryScale:"), value)
}
