// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ExtensionItem] class.
var ExtensionItemClass = _ExtensionItemClass{objc.GetClass("NSExtensionItem")}

type _ExtensionItemClass struct {
	objc.Class
}

// An interface definition for the [ExtensionItem] class.
type IExtensionItem interface {
	objc.IObject
	AttributedTitle() AttributedString
	SetAttributedTitle(value IAttributedString)
	Attachments() []ItemProvider
	SetAttachments(value []IItemProvider)
	AttributedContentText() AttributedString
	SetAttributedContentText(value IAttributedString)
	UserInfo() Dictionary
	SetUserInfo(value Dictionary)
}

// An immutable collection of values representing different aspects of an item for an extension to act upon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem?language=objc
type ExtensionItem struct {
	objc.Object
}

func ExtensionItemFrom(ptr unsafe.Pointer) ExtensionItem {
	return ExtensionItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionItemClass) Alloc() ExtensionItem {
	rv := objc.Call[ExtensionItem](ec, objc.Sel("alloc"))
	return rv
}

func (ec _ExtensionItemClass) New() ExtensionItem {
	rv := objc.Call[ExtensionItem](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionItem() ExtensionItem {
	return ExtensionItemClass.New()
}

func (e_ ExtensionItem) Init() ExtensionItem {
	rv := objc.Call[ExtensionItem](e_, objc.Sel("init"))
	return rv
}

// An optional title for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1416592-attributedtitle?language=objc
func (e_ ExtensionItem) AttributedTitle() AttributedString {
	rv := objc.Call[AttributedString](e_, objc.Sel("attributedTitle"))
	return rv
}

// An optional title for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1416592-attributedtitle?language=objc
func (e_ ExtensionItem) SetAttributedTitle(value IAttributedString) {
	objc.Call[objc.Void](e_, objc.Sel("setAttributedTitle:"), value)
}

// An optional array of media data associated with the extension item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1416690-attachments?language=objc
func (e_ ExtensionItem) Attachments() []ItemProvider {
	rv := objc.Call[[]ItemProvider](e_, objc.Sel("attachments"))
	return rv
}

// An optional array of media data associated with the extension item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1416690-attachments?language=objc
func (e_ ExtensionItem) SetAttachments(value []IItemProvider) {
	objc.Call[objc.Void](e_, objc.Sel("setAttachments:"), value)
}

// An optional string describing the extension item content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1408297-attributedcontenttext?language=objc
func (e_ ExtensionItem) AttributedContentText() AttributedString {
	rv := objc.Call[AttributedString](e_, objc.Sel("attributedContentText"))
	return rv
}

// An optional string describing the extension item content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1408297-attributedcontenttext?language=objc
func (e_ ExtensionItem) SetAttributedContentText(value IAttributedString) {
	objc.Call[objc.Void](e_, objc.Sel("setAttributedContentText:"), value)
}

// An optional dictionary of keys and values corresponding to the extension item’s properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1414953-userinfo?language=objc
func (e_ ExtensionItem) UserInfo() Dictionary {
	rv := objc.Call[Dictionary](e_, objc.Sel("userInfo"))
	return rv
}

// An optional dictionary of keys and values corresponding to the extension item’s properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/1414953-userinfo?language=objc
func (e_ ExtensionItem) SetUserInfo(value Dictionary) {
	objc.Call[objc.Void](e_, objc.Sel("setUserInfo:"), value)
}
