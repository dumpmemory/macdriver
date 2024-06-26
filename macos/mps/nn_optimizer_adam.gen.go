// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNOptimizerAdam] class.
var NNOptimizerAdamClass = _NNOptimizerAdamClass{objc.GetClass("MPSNNOptimizerAdam")}

type _NNOptimizerAdamClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerAdam] class.
type INNOptimizerAdam interface {
	INNOptimizer
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, maximumVelocityVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, maximumVelocityVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, maximumVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, maximumVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferObjectBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferObjectBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	Beta2() float64
	Epsilon() float32
	Beta1() float64
	TimeStep() uint
	SetTimeStep(value uint)
}

// An optimization layer that performs an Adam pdate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam?language=objc
type NNOptimizerAdam struct {
	NNOptimizer
}

func NNOptimizerAdamFrom(ptr unsafe.Pointer) NNOptimizerAdam {
	return NNOptimizerAdam{
		NNOptimizer: NNOptimizerFrom(ptr),
	}
}

func (n_ NNOptimizerAdam) InitWithDeviceLearningRate(device metal.PDevice, learningRate float32) NNOptimizerAdam {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("initWithDevice:learningRate:"), po0, learningRate)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966719-initwithdevice?language=objc
func NewNNOptimizerAdamWithDeviceLearningRate(device metal.PDevice, learningRate float32) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().InitWithDeviceLearningRate(device, learningRate)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerAdam) InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device metal.PDevice, beta1 float64, beta2 float64, epsilon float32, timeStep uint, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerAdam {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:"), po0, beta1, beta2, epsilon, timeStep, optimizerDescriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966718-initwithdevice?language=objc
func NewNNOptimizerAdamWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device metal.PDevice, beta1 float64, beta2 float64, epsilon float32, timeStep uint, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device, beta1, beta2, epsilon, timeStep, optimizerDescriptor)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerAdamClass) Alloc() NNOptimizerAdam {
	rv := objc.Call[NNOptimizerAdam](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNOptimizerAdamClass) New() NNOptimizerAdam {
	rv := objc.Call[NNOptimizerAdam](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerAdam() NNOptimizerAdam {
	return NNOptimizerAdamClass.New()
}

func (n_ NNOptimizerAdam) Init() NNOptimizerAdam {
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizerAdam) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerAdam {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNOptimizerAdam_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerAdam) InitWithDevice(device metal.PDevice) NNOptimizerAdam {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerAdamWithDevice(device metal.PDevice) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175016-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, maximumVelocityVector IVector, resultValuesVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:"), po0, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, maximumVelocityVector, resultValuesVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175016-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, maximumVelocityVector IVector, resultValuesVector IVector) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:"), commandBufferObject, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, maximumVelocityVector, resultValuesVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197825-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, maximumVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:"), po0, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, maximumVelocityMatrix, resultValuesMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197825-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, maximumVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:"), commandBufferObject, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, maximumVelocityMatrix, resultValuesMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3019334-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:"), po0, batchNormalizationState, inputMomentumVectors, inputVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3019334-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBufferObject, batchNormalizationState, inputMomentumVectors, inputVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3013782-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), po0, convolutionGradientState, convolutionSourceState, inputMomentumVectors, inputVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3013782-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBufferObject, convolutionGradientState, convolutionSourceState, inputMomentumVectors, inputVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966716-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, resultValuesVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:"), po0, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, resultValuesVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966716-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, inputVelocityVector IVector, resultValuesVector IVector) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:"), commandBufferObject, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, resultValuesVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197826-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), po0, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, resultValuesMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197826-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), commandBufferObject, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, resultValuesMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175015-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), po0, convolutionGradientState, convolutionSourceState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175015-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBufferObject, convolutionGradientState, convolutionSourceState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175014-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), po0, batchNormalizationState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175014-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBufferObject, batchNormalizationState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3013781-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), po0, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, inputVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3013781-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBufferObject, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, inputVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175013-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), po0, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3175013-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, inputVelocityVectors []IVector, maximumVelocityVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBufferObject, batchNormalizationGradientState, batchNormalizationSourceState, inputMomentumVectors, inputVelocityVectors, maximumVelocityVectors, resultState)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966715-beta2?language=objc
func (n_ NNOptimizerAdam) Beta2() float64 {
	rv := objc.Call[float64](n_, objc.Sel("beta2"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966717-epsilon?language=objc
func (n_ NNOptimizerAdam) Epsilon() float32 {
	rv := objc.Call[float32](n_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966714-beta1?language=objc
func (n_ NNOptimizerAdam) Beta1() float64 {
	rv := objc.Call[float64](n_, objc.Sel("beta1"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966720-timestep?language=objc
func (n_ NNOptimizerAdam) TimeStep() uint {
	rv := objc.Call[uint](n_, objc.Sel("timeStep"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966720-timestep?language=objc
func (n_ NNOptimizerAdam) SetTimeStep(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setTimeStep:"), value)
}
