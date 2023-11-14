// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionRenderer] class.
var CaptionRendererClass = _CaptionRendererClass{objc.GetClass("AVCaptionRenderer")}

type _CaptionRendererClass struct {
	objc.Class
}

// An interface definition for the [CaptionRenderer] class.
type ICaptionRenderer interface {
	objc.IObject
	CaptionSceneChangesInRange(consideredTimeRange coremedia.TimeRange) []CaptionRendererScene
	RenderInContextForTime(ctx coregraphics.ContextRef, time coremedia.Time)
	Captions() []Caption
	SetCaptions(value []ICaption)
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
}

// An object that renders captions for display at a particular time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer?language=objc
type CaptionRenderer struct {
	objc.Object
}

func CaptionRendererFrom(ptr unsafe.Pointer) CaptionRenderer {
	return CaptionRenderer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionRendererClass) Alloc() CaptionRenderer {
	rv := objc.Call[CaptionRenderer](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CaptionRendererClass) New() CaptionRenderer {
	rv := objc.Call[CaptionRenderer](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionRenderer() CaptionRenderer {
	return CaptionRendererClass.New()
}

func (c_ CaptionRenderer) Init() CaptionRenderer {
	rv := objc.Call[CaptionRenderer](c_, objc.Sel("init"))
	return rv
}

// Determine render time ranges within an enclosing time range to account for visual changes among captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer/3752969-captionscenechangesinrange?language=objc
func (c_ CaptionRenderer) CaptionSceneChangesInRange(consideredTimeRange coremedia.TimeRange) []CaptionRendererScene {
	rv := objc.Call[[]CaptionRendererScene](c_, objc.Sel("captionSceneChangesInRange:"), consideredTimeRange)
	return rv
}

// Draw the captions for the time you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer/3752971-renderincontext?language=objc
func (c_ CaptionRenderer) RenderInContextForTime(ctx coregraphics.ContextRef, time coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("renderInContext:forTime:"), ctx, time)
}

// The captions to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer/3752970-captions?language=objc
func (c_ CaptionRenderer) Captions() []Caption {
	rv := objc.Call[[]Caption](c_, objc.Sel("captions"))
	return rv
}

// The captions to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer/3752970-captions?language=objc
func (c_ CaptionRenderer) SetCaptions(value []ICaption) {
	objc.Call[objc.Void](c_, objc.Sel("setCaptions:"), value)
}

// The drawing bounds of caption scenes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer/3752968-bounds?language=objc
func (c_ CaptionRenderer) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("bounds"))
	return rv
}

// The drawing bounds of caption scenes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrenderer/3752968-bounds?language=objc
func (c_ CaptionRenderer) SetBounds(value coregraphics.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("setBounds:"), value)
}