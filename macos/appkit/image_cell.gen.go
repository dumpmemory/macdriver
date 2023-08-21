// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageCell] class.
var ImageCellClass = _ImageCellClass{objc.GetClass("NSImageCell")}

type _ImageCellClass struct {
	objc.Class
}

// An interface definition for the [ImageCell] class.
type IImageCell interface {
	ICell
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
}

// An NSImageCell object displays a single image (encapsulated in an NSImage object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell?language=objc
type ImageCell struct {
	Cell
}

func ImageCellFrom(ptr unsafe.Pointer) ImageCell {
	return ImageCell{
		Cell: CellFrom(ptr),
	}
}

func (ic _ImageCellClass) Alloc() ImageCell {
	rv := objc.Call[ImageCell](ic, objc.Sel("alloc"))
	return rv
}

func ImageCell_Alloc() ImageCell {
	return ImageCellClass.Alloc()
}

func (ic _ImageCellClass) New() ImageCell {
	rv := objc.Call[ImageCell](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageCell() ImageCell {
	return ImageCellClass.New()
}

func (i_ ImageCell) Init() ImageCell {
	rv := objc.Call[ImageCell](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageCell) InitImageCell(image IImage) ImageCell {
	rv := objc.Call[ImageCell](i_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func NewImageCellImageCell(image IImage) ImageCell {
	instance := ImageCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

func (i_ ImageCell) InitTextCell(string_ string) ImageCell {
	rv := objc.Call[ImageCell](i_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func NewImageCellTextCell(string_ string) ImageCell {
	instance := ImageCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

// The alignment of the receiver’s image relative to its frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/1524421-imagealignment?language=objc
func (i_ ImageCell) ImageAlignment() ImageAlignment {
	rv := objc.Call[ImageAlignment](i_, objc.Sel("imageAlignment"))
	return rv
}

// The alignment of the receiver’s image relative to its frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/1524421-imagealignment?language=objc
func (i_ ImageCell) SetImageAlignment(value ImageAlignment) {
	objc.Call[objc.Void](i_, objc.Sel("setImageAlignment:"), value)
}

// The scaling mode used to fit the receiver's image into the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/1532559-imagescaling?language=objc
func (i_ ImageCell) ImageScaling() ImageScaling {
	rv := objc.Call[ImageScaling](i_, objc.Sel("imageScaling"))
	return rv
}

// The scaling mode used to fit the receiver's image into the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/1532559-imagescaling?language=objc
func (i_ ImageCell) SetImageScaling(value ImageScaling) {
	objc.Call[objc.Void](i_, objc.Sel("setImageScaling:"), value)
}

// The style of the frame that borders the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/1526164-imageframestyle?language=objc
func (i_ ImageCell) ImageFrameStyle() ImageFrameStyle {
	rv := objc.Call[ImageFrameStyle](i_, objc.Sel("imageFrameStyle"))
	return rv
}

// The style of the frame that borders the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/1526164-imageframestyle?language=objc
func (i_ ImageCell) SetImageFrameStyle(value ImageFrameStyle) {
	objc.Call[objc.Void](i_, objc.Sel("setImageFrameStyle:"), value)
}