// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A container for pipeline state descriptors and their associated compiled shader code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive?language=objc
type PBinaryArchive interface {
	// optional
	SerializeToURLError(url foundation.URL, error unsafe.Pointer) bool
	HasSerializeToURLError() bool

	// optional
	AddTileRenderPipelineFunctionsWithDescriptorError(descriptor TileRenderPipelineDescriptor, error unsafe.Pointer) bool
	HasAddTileRenderPipelineFunctionsWithDescriptorError() bool

	// optional
	AddComputePipelineFunctionsWithDescriptorError(descriptor ComputePipelineDescriptor, error unsafe.Pointer) bool
	HasAddComputePipelineFunctionsWithDescriptorError() bool

	// optional
	AddRenderPipelineFunctionsWithDescriptorError(descriptor RenderPipelineDescriptor, error unsafe.Pointer) bool
	HasAddRenderPipelineFunctionsWithDescriptorError() bool

	// optional
	AddFunctionWithDescriptorLibraryError(descriptor FunctionDescriptor, library LibraryObject, error unsafe.Pointer) bool
	HasAddFunctionWithDescriptorLibraryError() bool

	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// ensure impl type implements protocol interface
var _ PBinaryArchive = (*BinaryArchiveObject)(nil)

// A concrete type for the [PBinaryArchive] protocol.
type BinaryArchiveObject struct {
	objc.Object
}

func (b_ BinaryArchiveObject) HasSerializeToURLError() bool {
	return b_.RespondsToSelector(objc.Sel("serializeToURL:error:"))
}

// Writes the contents of the archive to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553928-serializetourl?language=objc
func (b_ BinaryArchiveObject) SerializeToURLError(url foundation.URL, error unsafe.Pointer) bool {
	rv := objc.Call[bool](b_, objc.Sel("serializeToURL:error:"), url, error)
	return rv
}

func (b_ BinaryArchiveObject) HasAddTileRenderPipelineFunctionsWithDescriptorError() bool {
	return b_.RespondsToSelector(objc.Sel("addTileRenderPipelineFunctionsWithDescriptor:error:"))
}

// Adds a description of a tile renderer pipeline to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3564419-addtilerenderpipelinefunctionswi?language=objc
func (b_ BinaryArchiveObject) AddTileRenderPipelineFunctionsWithDescriptorError(descriptor TileRenderPipelineDescriptor, error unsafe.Pointer) bool {
	rv := objc.Call[bool](b_, objc.Sel("addTileRenderPipelineFunctionsWithDescriptor:error:"), descriptor, error)
	return rv
}

func (b_ BinaryArchiveObject) HasAddComputePipelineFunctionsWithDescriptorError() bool {
	return b_.RespondsToSelector(objc.Sel("addComputePipelineFunctionsWithDescriptor:error:"))
}

// Adds a description of a compute pipeline to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553924-addcomputepipelinefunctionswithd?language=objc
func (b_ BinaryArchiveObject) AddComputePipelineFunctionsWithDescriptorError(descriptor ComputePipelineDescriptor, error unsafe.Pointer) bool {
	rv := objc.Call[bool](b_, objc.Sel("addComputePipelineFunctionsWithDescriptor:error:"), descriptor, error)
	return rv
}

func (b_ BinaryArchiveObject) HasAddRenderPipelineFunctionsWithDescriptorError() bool {
	return b_.RespondsToSelector(objc.Sel("addRenderPipelineFunctionsWithDescriptor:error:"))
}

// Adds a description of a render pipeline to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553925-addrenderpipelinefunctionswithde?language=objc
func (b_ BinaryArchiveObject) AddRenderPipelineFunctionsWithDescriptorError(descriptor RenderPipelineDescriptor, error unsafe.Pointer) bool {
	rv := objc.Call[bool](b_, objc.Sel("addRenderPipelineFunctionsWithDescriptor:error:"), descriptor, error)
	return rv
}

func (b_ BinaryArchiveObject) HasAddFunctionWithDescriptorLibraryError() bool {
	return b_.RespondsToSelector(objc.Sel("addFunctionWithDescriptor:library:error:"))
}

// Adds a description of a function to the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3750523-addfunctionwithdescriptor?language=objc
func (b_ BinaryArchiveObject) AddFunctionWithDescriptorLibraryError(descriptor FunctionDescriptor, library LibraryObject, error unsafe.Pointer) bool {
	po1 := objc.WrapAsProtocol("MTLLibrary", library)
	rv := objc.Call[bool](b_, objc.Sel("addFunctionWithDescriptor:library:error:"), descriptor, po1, error)
	return rv
}

func (b_ BinaryArchiveObject) HasDevice() bool {
	return b_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the binary archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553926-device?language=objc
func (b_ BinaryArchiveObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](b_, objc.Sel("device"))
	return rv
}

func (b_ BinaryArchiveObject) HasSetLabel() bool {
	return b_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553927-label?language=objc
func (b_ BinaryArchiveObject) SetLabel(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setLabel:"), value)
}

func (b_ BinaryArchiveObject) HasLabel() bool {
	return b_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchive/3553927-label?language=objc
func (b_ BinaryArchiveObject) Label() string {
	rv := objc.Call[string](b_, objc.Sel("label"))
	return rv
}
