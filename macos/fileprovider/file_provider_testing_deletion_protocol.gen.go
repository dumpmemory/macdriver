// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/darwinkit/objc"
)

// An operation that syncs the deletion of the source item to the target location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion?language=objc
type PFileProviderTestingDeletion interface {
	// optional
	SourceItemIdentifier() FileProviderItemIdentifier
	HasSourceItemIdentifier() bool

	// optional
	TargetItemBaseVersion() FileProviderItemVersion
	HasTargetItemBaseVersion() bool

	// optional
	TargetItemIdentifier() FileProviderItemIdentifier
	HasTargetItemIdentifier() bool

	// optional
	TargetSide() FileProviderTestingOperationSide
	HasTargetSide() bool

	// optional
	DomainVersion() FileProviderDomainVersion
	HasDomainVersion() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderTestingDeletion = (*FileProviderTestingDeletionObject)(nil)

// A concrete type for the [PFileProviderTestingDeletion] protocol.
type FileProviderTestingDeletionObject struct {
	objc.Object
}

func (f_ FileProviderTestingDeletionObject) HasSourceItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("sourceItemIdentifier"))
}

// The unique identifier for the source item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736236-sourceitemidentifier?language=objc
func (f_ FileProviderTestingDeletionObject) SourceItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("sourceItemIdentifier"))
	return rv
}

func (f_ FileProviderTestingDeletionObject) HasTargetItemBaseVersion() bool {
	return f_.RespondsToSelector(objc.Sel("targetItemBaseVersion"))
}

// The version of the deleted item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736237-targetitembaseversion?language=objc
func (f_ FileProviderTestingDeletionObject) TargetItemBaseVersion() FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](f_, objc.Sel("targetItemBaseVersion"))
	return rv
}

func (f_ FileProviderTestingDeletionObject) HasTargetItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("targetItemIdentifier"))
}

// The unique identifier for the target item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736238-targetitemidentifier?language=objc
func (f_ FileProviderTestingDeletionObject) TargetItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("targetItemIdentifier"))
	return rv
}

func (f_ FileProviderTestingDeletionObject) HasTargetSide() bool {
	return f_.RespondsToSelector(objc.Sel("targetSide"))
}

// The target location for the delete operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736239-targetside?language=objc
func (f_ FileProviderTestingDeletionObject) TargetSide() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("targetSide"))
	return rv
}

func (f_ FileProviderTestingDeletionObject) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// The domain’s version when the source location deleted the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736235-domainversion?language=objc
func (f_ FileProviderTestingDeletionObject) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}
