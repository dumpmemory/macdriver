// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of methods that enable the delegate of a path cell object to customize the Open panel or pop-up menu of a path whose style is set to NSPathStylePopUp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate?language=objc
type PPathCellDelegate interface {
	// optional
	PathCellWillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel)
	HasPathCellWillDisplayOpenPanel() bool

	// optional
	PathCellWillPopUpMenu(pathCell PathCell, menu Menu)
	HasPathCellWillPopUpMenu() bool
}

// A delegate implementation builder for the [PPathCellDelegate] protocol.
type PathCellDelegate struct {
	_PathCellWillDisplayOpenPanel func(pathCell PathCell, openPanel OpenPanel)
	_PathCellWillPopUpMenu        func(pathCell PathCell, menu Menu)
}

func (di *PathCellDelegate) HasPathCellWillDisplayOpenPanel() bool {
	return di._PathCellWillDisplayOpenPanel != nil
}

// Implement this method to customize the Open panel shown by a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1526099-pathcell?language=objc
func (di *PathCellDelegate) SetPathCellWillDisplayOpenPanel(f func(pathCell PathCell, openPanel OpenPanel)) {
	di._PathCellWillDisplayOpenPanel = f
}

// Implement this method to customize the Open panel shown by a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1526099-pathcell?language=objc
func (di *PathCellDelegate) PathCellWillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel) {
	di._PathCellWillDisplayOpenPanel(pathCell, openPanel)
}
func (di *PathCellDelegate) HasPathCellWillPopUpMenu() bool {
	return di._PathCellWillPopUpMenu != nil
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1525005-pathcell?language=objc
func (di *PathCellDelegate) SetPathCellWillPopUpMenu(f func(pathCell PathCell, menu Menu)) {
	di._PathCellWillPopUpMenu = f
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1525005-pathcell?language=objc
func (di *PathCellDelegate) PathCellWillPopUpMenu(pathCell PathCell, menu Menu) {
	di._PathCellWillPopUpMenu(pathCell, menu)
}

// ensure impl type implements protocol interface
var _ PPathCellDelegate = (*PathCellDelegateObject)(nil)

// A concrete type for the [PPathCellDelegate] protocol.
type PathCellDelegateObject struct {
	objc.Object
}

func (p_ PathCellDelegateObject) HasPathCellWillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.Sel("pathCell:willDisplayOpenPanel:"))
}

// Implement this method to customize the Open panel shown by a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1526099-pathcell?language=objc
func (p_ PathCellDelegateObject) PathCellWillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel) {
	objc.Call[objc.Void](p_, objc.Sel("pathCell:willDisplayOpenPanel:"), pathCell, openPanel)
}

func (p_ PathCellDelegateObject) HasPathCellWillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.Sel("pathCell:willPopUpMenu:"))
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1525005-pathcell?language=objc
func (p_ PathCellDelegateObject) PathCellWillPopUpMenu(pathCell PathCell, menu Menu) {
	objc.Call[objc.Void](p_, objc.Sel("pathCell:willPopUpMenu:"), pathCell, menu)
}
