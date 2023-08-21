// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextInputContext] class.
var TextInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}

type _TextInputContextClass struct {
	objc.Class
}

// An interface definition for the [TextInputContext] class.
type ITextInputContext interface {
	objc.IObject
	Activate()
	DiscardMarkedText()
	HandleEvent(event IEvent) bool
	Deactivate()
	InvalidateCharacterCoordinates()
	KeyboardInputSources() []TextInputSourceIdentifier
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	SelectedKeyboardInputSource() TextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier)
	Client() TextInputClientWrapper
}

// An object that represents the Cocoa text input system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext?language=objc
type TextInputContext struct {
	objc.Object
}

func TextInputContextFrom(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextInputContext) InitWithClient(client PTextInputClient) TextInputContext {
	po0 := objc.WrapAsProtocol("NSTextInputClient", client)
	rv := objc.Call[TextInputContext](t_, objc.Sel("initWithClient:"), po0)
	return rv
}

// The designated initializer [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1532777-initwithclient?language=objc
func NewTextInputContextWithClient(client PTextInputClient) TextInputContext {
	instance := TextInputContextClass.Alloc().InitWithClient(client)
	instance.Autorelease()
	return instance
}

func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := objc.Call[TextInputContext](tc, objc.Sel("alloc"))
	return rv
}

func TextInputContext_Alloc() TextInputContext {
	return TextInputContextClass.Alloc()
}

func (tc _TextInputContextClass) New() TextInputContext {
	rv := objc.Call[TextInputContext](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextInputContext() TextInputContext {
	return TextInputContextClass.New()
}

func (t_ TextInputContext) Init() TextInputContext {
	rv := objc.Call[TextInputContext](t_, objc.Sel("init"))
	return rv
}

// Activates the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1534378-activate?language=objc
func (t_ TextInputContext) Activate() {
	objc.Call[objc.Void](t_, objc.Sel("activate"))
}

// Tells the Cocoa text input system to discard the current conversion session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1528752-discardmarkedtext?language=objc
func (t_ TextInputContext) DiscardMarkedText() {
	objc.Call[objc.Void](t_, objc.Sel("discardMarkedText"))
}

// Tells the Cocoa text input system to handle mouse or key events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1528602-handleevent?language=objc
func (t_ TextInputContext) HandleEvent(event IEvent) bool {
	rv := objc.Call[bool](t_, objc.Sel("handleEvent:"), objc.Ptr(event))
	return rv
}

// Returns the display name for the given text input source identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1526644-localizednameforinputsource?language=objc
func (tc _TextInputContextClass) LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	rv := objc.Call[string](tc, objc.Sel("localizedNameForInputSource:"), inputSourceIdentifier)
	return rv
}

// Returns the display name for the given text input source identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1526644-localizednameforinputsource?language=objc
func TextInputContext_LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	return TextInputContextClass.LocalizedNameForInputSource(inputSourceIdentifier)
}

// Deactivates the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1535505-deactivate?language=objc
func (t_ TextInputContext) Deactivate() {
	objc.Call[objc.Void](t_, objc.Sel("deactivate"))
}

// Notifies the Cocoa text input system that the position information previously queried via methods like firstRectForCharacterRange:actualRange: needs to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1535165-invalidatecharactercoordinates?language=objc
func (t_ TextInputContext) InvalidateCharacterCoordinates() {
	objc.Call[objc.Void](t_, objc.Sel("invalidateCharacterCoordinates"))
}

// The array of keyboard text input source identifier strings available to the receiver. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1529156-keyboardinputsources?language=objc
func (t_ TextInputContext) KeyboardInputSources() []TextInputSourceIdentifier {
	rv := objc.Call[[]TextInputSourceIdentifier](t_, objc.Sel("keyboardInputSources"))
	return rv
}

// The set of keyboard input source locales allowed when this input context is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1532284-allowedinputsourcelocales?language=objc
func (t_ TextInputContext) AllowedInputSourceLocales() []string {
	rv := objc.Call[[]string](t_, objc.Sel("allowedInputSourceLocales"))
	return rv
}

// The set of keyboard input source locales allowed when this input context is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1532284-allowedinputsourcelocales?language=objc
func (t_ TextInputContext) SetAllowedInputSourceLocales(value []string) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowedInputSourceLocales:"), value)
}

// A Boolean value that indicates whether the client handles NSGlyphInfoAttributeName or not. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1534420-acceptsglyphinfo?language=objc
func (t_ TextInputContext) AcceptsGlyphInfo() bool {
	rv := objc.Call[bool](t_, objc.Sel("acceptsGlyphInfo"))
	return rv
}

// A Boolean value that indicates whether the client handles NSGlyphInfoAttributeName or not. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1534420-acceptsglyphinfo?language=objc
func (t_ TextInputContext) SetAcceptsGlyphInfo(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAcceptsGlyphInfo:"), value)
}

// The identifier string for the selected keyboard text input source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1533970-selectedkeyboardinputsource?language=objc
func (t_ TextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier {
	rv := objc.Call[TextInputSourceIdentifier](t_, objc.Sel("selectedKeyboardInputSource"))
	return rv
}

// The identifier string for the selected keyboard text input source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1533970-selectedkeyboardinputsource?language=objc
func (t_ TextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedKeyboardInputSource:"), value)
}

// Returns the current, activated, text input context object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1527409-currentinputcontext?language=objc
func (tc _TextInputContextClass) CurrentInputContext() TextInputContext {
	rv := objc.Call[TextInputContext](tc, objc.Sel("currentInputContext"))
	return rv
}

// Returns the current, activated, text input context object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1527409-currentinputcontext?language=objc
func TextInputContext_CurrentInputContext() TextInputContext {
	return TextInputContextClass.CurrentInputContext()
}

// The owner of this input context. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/1533583-client?language=objc
func (t_ TextInputContext) Client() TextInputClientWrapper {
	rv := objc.Call[TextInputClientWrapper](t_, objc.Sel("client"))
	return rv
}