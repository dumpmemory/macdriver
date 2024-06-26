// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CommandBuffer] class.
var CommandBufferClass = _CommandBufferClass{objc.GetClass("MPSCommandBuffer")}

type _CommandBufferClass struct {
	objc.Class
}

// An interface definition for the [CommandBuffer] class.
type ICommandBuffer interface {
	objc.IObject
	CommitAndContinue()
	PrefetchHeapForWorkloadSize(size uint)
	CommandBuffer() metal.CommandBufferObject
	Predicate() Predicate
	SetPredicate(value IPredicate)
	HeapProvider() HeapProviderObject
	SetHeapProvider(value PHeapProvider)
	SetHeapProviderObject(valueObject objc.IObject)
	RootCommandBuffer() metal.CommandBufferObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer?language=objc
type CommandBuffer struct {
	objc.Object
}

func CommandBufferFrom(ptr unsafe.Pointer) CommandBuffer {
	return CommandBuffer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CommandBuffer) InitWithCommandBuffer(commandBuffer metal.PCommandBuffer) CommandBuffer {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CommandBuffer](c_, objc.Sel("initWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114031-initwithcommandbuffer?language=objc
func NewCommandBufferWithCommandBuffer(commandBuffer metal.PCommandBuffer) CommandBuffer {
	instance := CommandBufferClass.Alloc().InitWithCommandBuffer(commandBuffer)
	instance.Autorelease()
	return instance
}

func (cc _CommandBufferClass) CommandBufferWithCommandBuffer(commandBuffer metal.PCommandBuffer) CommandBuffer {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CommandBuffer](cc, objc.Sel("commandBufferWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114030-commandbufferwithcommandbuffer?language=objc
func CommandBuffer_CommandBufferWithCommandBuffer(commandBuffer metal.PCommandBuffer) CommandBuffer {
	return CommandBufferClass.CommandBufferWithCommandBuffer(commandBuffer)
}

func (cc _CommandBufferClass) CommandBufferFromCommandQueue(commandQueue metal.PCommandQueue) CommandBuffer {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	rv := objc.Call[CommandBuffer](cc, objc.Sel("commandBufferFromCommandQueue:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114029-commandbufferfromcommandqueue?language=objc
func CommandBuffer_CommandBufferFromCommandQueue(commandQueue metal.PCommandQueue) CommandBuffer {
	return CommandBufferClass.CommandBufferFromCommandQueue(commandQueue)
}

func (cc _CommandBufferClass) Alloc() CommandBuffer {
	rv := objc.Call[CommandBuffer](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CommandBufferClass) New() CommandBuffer {
	rv := objc.Call[CommandBuffer](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCommandBuffer() CommandBuffer {
	return CommandBufferClass.New()
}

func (c_ CommandBuffer) Init() CommandBuffer {
	rv := objc.Call[CommandBuffer](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3152524-commitandcontinue?language=objc
func (c_ CommandBuffer) CommitAndContinue() {
	objc.Call[objc.Void](c_, objc.Sel("commitAndContinue"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229858-prefetchheapforworkloadsize?language=objc
func (c_ CommandBuffer) PrefetchHeapForWorkloadSize(size uint) {
	objc.Call[objc.Void](c_, objc.Sel("prefetchHeapForWorkloadSize:"), size)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114028-commandbuffer?language=objc
func (c_ CommandBuffer) CommandBuffer() metal.CommandBufferObject {
	rv := objc.Call[metal.CommandBufferObject](c_, objc.Sel("commandBuffer"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114032-predicate?language=objc
func (c_ CommandBuffer) Predicate() Predicate {
	rv := objc.Call[Predicate](c_, objc.Sel("predicate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114032-predicate?language=objc
func (c_ CommandBuffer) SetPredicate(value IPredicate) {
	objc.Call[objc.Void](c_, objc.Sel("setPredicate:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229857-heapprovider?language=objc
func (c_ CommandBuffer) HeapProvider() HeapProviderObject {
	rv := objc.Call[HeapProviderObject](c_, objc.Sel("heapProvider"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229857-heapprovider?language=objc
func (c_ CommandBuffer) SetHeapProvider(value PHeapProvider) {
	po0 := objc.WrapAsProtocol("MPSHeapProvider", value)
	objc.Call[objc.Void](c_, objc.Sel("setHeapProvider:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229857-heapprovider?language=objc
func (c_ CommandBuffer) SetHeapProviderObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setHeapProvider:"), valueObject)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3166772-rootcommandbuffer?language=objc
func (c_ CommandBuffer) RootCommandBuffer() metal.CommandBufferObject {
	rv := objc.Call[metal.CommandBufferObject](c_, objc.Sel("rootCommandBuffer"))
	return rv
}
