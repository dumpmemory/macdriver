// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutBoundarySupplementaryItem] class.
var CollectionLayoutBoundarySupplementaryItemClass = _CollectionLayoutBoundarySupplementaryItemClass{objc.GetClass("NSCollectionLayoutBoundarySupplementaryItem")}

type _CollectionLayoutBoundarySupplementaryItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutBoundarySupplementaryItem] class.
type ICollectionLayoutBoundarySupplementaryItem interface {
	ICollectionLayoutSupplementaryItem
	Alignment() RectAlignment
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	Offset() coregraphics.Point
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
}

// An object used to add headers or footers to a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem?language=objc
type CollectionLayoutBoundarySupplementaryItem struct {
	CollectionLayoutSupplementaryItem
}

func CollectionLayoutBoundarySupplementaryItemFrom(ptr unsafe.Pointer) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItem{
		CollectionLayoutSupplementaryItem: CollectionLayoutSupplementaryItemFrom(ptr),
	}
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[CollectionLayoutBoundarySupplementaryItem](cc, objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:"), objc.Ptr(layoutSize), elementKind, alignment)
	return rv
}

// Creates a boundary supplementary item of the specified size and element kind, with an alignment relative to a section or layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3213819-boundarysupplementaryitemwithlay?language=objc
func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize, elementKind, alignment)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) Alloc() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[CollectionLayoutBoundarySupplementaryItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_Alloc() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.Alloc()
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) New() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[CollectionLayoutBoundarySupplementaryItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutBoundarySupplementaryItem() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.New()
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Init() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[CollectionLayoutBoundarySupplementaryItem](c_, objc.Sel("init"))
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[CollectionLayoutBoundarySupplementaryItem](cc, objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), objc.Ptr(layoutSize), elementKind, objc.Ptr(containerAnchor), objc.Ptr(itemAnchor))
	return rv
}

// Creates a supplementary item of the specified size and element kind, an anchor relative to a container, and an anchor relative to an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3213900-supplementaryitemwithlayoutsize?language=objc
func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize, elementKind, containerAnchor, itemAnchor)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[CollectionLayoutBoundarySupplementaryItem](cc, objc.Sel("itemWithLayoutSize:"), objc.Ptr(layoutSize))
	return rv
}

// Creates an item of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213871-itemwithlayoutsize?language=objc
func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.ItemWithLayoutSize(layoutSize)
}

// The alignment of the boundary supplementary item relative to the section or layout it’s attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3199039-alignment?language=objc
func (c_ CollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	rv := objc.Call[RectAlignment](c_, objc.Sel("alignment"))
	return rv
}

// A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3199044-pintovisiblebounds?language=objc
func (c_ CollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := objc.Call[bool](c_, objc.Sel("pinToVisibleBounds"))
	return rv
}

// A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3199044-pintovisiblebounds?language=objc
func (c_ CollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setPinToVisibleBounds:"), value)
}

// The floating-point value of the boundary supplementary item’s offset from the section or layout it’s attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3199043-offset?language=objc
func (c_ CollectionLayoutBoundarySupplementaryItem) Offset() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("offset"))
	return rv
}

// A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3199040-extendsboundary?language=objc
func (c_ CollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := objc.Call[bool](c_, objc.Sel("extendsBoundary"))
	return rv
}

// A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutboundarysupplementaryitem/3199040-extendsboundary?language=objc
func (c_ CollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setExtendsBoundary:"), value)
}