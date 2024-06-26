// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNConvolutionGradientNode] class.
var CNNConvolutionGradientNodeClass = _CNNConvolutionGradientNodeClass{objc.GetClass("MPSCNNConvolutionGradientNode")}

type _CNNConvolutionGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionGradientNode] class.
type ICNNConvolutionGradientNode interface {
	INNGradientFilterNode
}

// A representation of a gradient convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode?language=objc
type CNNConvolutionGradientNode struct {
	NNGradientFilterNode
}

func CNNConvolutionGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientNode {
	return CNNConvolutionGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (cc _CNNConvolutionGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947984-nodewithsourcegradient?language=objc
func CNNConvolutionGradientNode_NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionGradientNode {
	return CNNConvolutionGradientNodeClass.NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
}

func (c_ CNNConvolutionGradientNode) InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947999-initwithsourcegradient?language=objc
func NewCNNConvolutionGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionGradientNode {
	instance := CNNConvolutionGradientNodeClass.Alloc().InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionGradientNodeClass) Alloc() CNNConvolutionGradientNode {
	rv := objc.Call[CNNConvolutionGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNConvolutionGradientNodeClass) New() CNNConvolutionGradientNode {
	rv := objc.Call[CNNConvolutionGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionGradientNode() CNNConvolutionGradientNode {
	return CNNConvolutionGradientNodeClass.New()
}

func (c_ CNNConvolutionGradientNode) Init() CNNConvolutionGradientNode {
	rv := objc.Call[CNNConvolutionGradientNode](c_, objc.Sel("init"))
	return rv
}
