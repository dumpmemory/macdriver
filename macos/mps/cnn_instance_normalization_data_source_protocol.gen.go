// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// A protocol that defines methods that an instance normalization uses to initialize scale factors and bias terms. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource?language=objc
type PCNNInstanceNormalizationDataSource interface {
	// optional
	InitWithCoder(aDecoder foundation.Coder) objc.Object
	HasInitWithCoder() bool

	// optional
	UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch *foundation.Array) bool
	HasUpdateGammaAndBetaWithInstanceNormalizationStateBatch() bool

	// optional
	Gamma() *float32
	HasGamma() bool

	// optional
	EncodeWithCoder(aCoder foundation.Coder)
	HasEncodeWithCoder() bool

	// optional
	Epsilon() float32
	HasEpsilon() bool

	// optional
	Purge()
	HasPurge() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	Beta() *float32
	HasBeta() bool

	// optional
	CopyWithZoneDevice(zone unsafe.Pointer, device metal.DeviceObject) objc.Object
	HasCopyWithZoneDevice() bool

	// optional
	UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer metal.CommandBufferObject, instanceNormalizationStateBatch *foundation.Array) CNNNormalizationGammaAndBetaState
	HasUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch() bool

	// optional
	Load() bool
	HasLoad() bool

	// optional
	NumberOfFeatureChannels() uint
	HasNumberOfFeatureChannels() bool
}

// ensure impl type implements protocol interface
var _ PCNNInstanceNormalizationDataSource = (*CNNInstanceNormalizationDataSourceObject)(nil)

// A concrete type for the [PCNNInstanceNormalizationDataSource] protocol.
type CNNInstanceNormalizationDataSourceObject struct {
	objc.Object
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasInitWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("initWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2947957-initwithcoder?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) InitWithCoder(aDecoder foundation.Coder) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasUpdateGammaAndBetaWithInstanceNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithInstanceNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953931-updategammaandbetawithinstanceno?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch *foundation.Array) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateGammaAndBetaWithInstanceNormalizationStateBatch:"), instanceNormalizationStateBatch)
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasGamma() bool {
	return c_.RespondsToSelector(objc.Sel("gamma"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953923-gamma?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) Gamma() *float32 {
	rv := objc.Call[*float32](c_, objc.Sel("gamma"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2947953-encodewithcoder?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) EncodeWithCoder(aCoder foundation.Coder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), aCoder)
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasEpsilon() bool {
	return c_.RespondsToSelector(objc.Sel("epsilon"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953925-epsilon?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) Epsilon() float32 {
	rv := objc.Call[float32](c_, objc.Sel("epsilon"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasPurge() bool {
	return c_.RespondsToSelector(objc.Sel("purge"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/3088879-purge?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) Purge() {
	objc.Call[objc.Void](c_, objc.Sel("purge"))
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2952998-label?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasBeta() bool {
	return c_.RespondsToSelector(objc.Sel("beta"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953922-beta?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) Beta() *float32 {
	rv := objc.Call[*float32](c_, objc.Sel("beta"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasCopyWithZoneDevice() bool {
	return c_.RespondsToSelector(objc.Sel("copyWithZone:device:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/3013780-copywithzone?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) CopyWithZoneDevice(zone unsafe.Pointer, device metal.DeviceObject) objc.Object {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithCommandBuffer:instanceNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953926-updategammaandbetawithcommandbuf?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer metal.CommandBufferObject, instanceNormalizationStateBatch *foundation.Array) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("updateGammaAndBetaWithCommandBuffer:instanceNormalizationStateBatch:"), po0, instanceNormalizationStateBatch)
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasLoad() bool {
	return c_.RespondsToSelector(objc.Sel("load"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/3088878-load?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) Load() bool {
	rv := objc.Call[bool](c_, objc.Sel("load"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceObject) HasNumberOfFeatureChannels() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfFeatureChannels"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2947961-numberoffeaturechannels?language=objc
func (c_ CNNInstanceNormalizationDataSourceObject) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}
