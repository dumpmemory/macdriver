// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A role-based protocol that declares the minimum interface necessary to act as a container for other user interface elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitygroup?language=objc
type PAccessibilityGroup interface {
}

// ensure impl type implements protocol interface
var _ PAccessibilityGroup = (*AccessibilityGroupObject)(nil)

// A concrete type for the [PAccessibilityGroup] protocol.
type AccessibilityGroupObject struct {
	objc.Object
}
