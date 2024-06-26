// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [IndexSet] class.
var IndexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}

type _IndexSetClass struct {
	objc.Class
}

// An interface definition for the [IndexSet] class.
type IIndexSet interface {
	objc.IObject
	ContainsIndex(value uint) bool
	IndexesInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	EnumerateIndexesWithOptionsUsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool))
	IndexGreaterThanOrEqualToIndex(value uint) uint
	ContainsIndexesInRange(range_ Range) bool
	IndexLessThanIndex(value uint) uint
	EnumerateRangesWithOptionsUsingBlock(opts EnumerationOptions, block func(range_ Range, stop *bool))
	IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint
	EnumerateIndexesUsingBlock(block func(idx uint, stop *bool))
	CountOfIndexesInRange(range_ Range) uint
	IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet
	GetIndexesMaxCountInIndexRange(indexBuffer *uint, bufferSize uint, range_ RangePointer) uint
	EnumerateIndexesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(idx uint, stop *bool))
	IndexWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexGreaterThanIndex(value uint) uint
	EnumerateRangesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(range_ Range, stop *bool))
	IndexLessThanOrEqualToIndex(value uint) uint
	ContainsIndexes(indexSet IIndexSet) bool
	IndexesWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	IndexInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IsEqualToIndexSet(indexSet IIndexSet) bool
	EnumerateRangesUsingBlock(block func(range_ Range, stop *bool))
	IntersectsIndexesInRange(range_ Range) bool
	FirstIndex() uint
	Count() uint
	LastIndex() uint
}

// An immutable collection of unique integer values that represent indexes in another collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset?language=objc
type IndexSet struct {
	objc.Object
}

func IndexSetFrom(ptr unsafe.Pointer) IndexSet {
	return IndexSet{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ IndexSet) InitWithIndex(value uint) IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("initWithIndex:"), value)
	return rv
}

// Initializes an allocated NSIndexSet object with an index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1416501-initwithindex?language=objc
func NewIndexSetWithIndex(value uint) IndexSet {
	instance := IndexSetClass.Alloc().InitWithIndex(value)
	instance.Autorelease()
	return instance
}

func (ic _IndexSetClass) IndexSetWithIndex(value uint) IndexSet {
	rv := objc.Call[IndexSet](ic, objc.Sel("indexSetWithIndex:"), value)
	return rv
}

// Creates an index set with an index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1427254-indexsetwithindex?language=objc
func IndexSet_IndexSetWithIndex(value uint) IndexSet {
	return IndexSetClass.IndexSetWithIndex(value)
}

func (ic _IndexSetClass) IndexSetWithIndexesInRange(range_ Range) IndexSet {
	rv := objc.Call[IndexSet](ic, objc.Sel("indexSetWithIndexesInRange:"), range_)
	return rv
}

// Creates an index set with an index range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1427274-indexsetwithindexesinrange?language=objc
func IndexSet_IndexSetWithIndexesInRange(range_ Range) IndexSet {
	return IndexSetClass.IndexSetWithIndexesInRange(range_)
}

func (i_ IndexSet) InitWithIndexesInRange(range_ Range) IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("initWithIndexesInRange:"), range_)
	return rv
}

// Initializes an allocated NSIndexSet object with an index range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1414013-initwithindexesinrange?language=objc
func NewIndexSetWithIndexesInRange(range_ Range) IndexSet {
	instance := IndexSetClass.Alloc().InitWithIndexesInRange(range_)
	instance.Autorelease()
	return instance
}

func (i_ IndexSet) InitWithIndexSet(indexSet IIndexSet) IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("initWithIndexSet:"), indexSet)
	return rv
}

// Initializes an allocated NSIndexSet object with an index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1415602-initwithindexset?language=objc
func NewIndexSetWithIndexSet(indexSet IIndexSet) IndexSet {
	instance := IndexSetClass.Alloc().InitWithIndexSet(indexSet)
	instance.Autorelease()
	return instance
}

func (ic _IndexSetClass) IndexSet() IndexSet {
	rv := objc.Call[IndexSet](ic, objc.Sel("indexSet"))
	return rv
}

// Creates an empty index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1427281-indexset?language=objc
func IndexSet_IndexSet() IndexSet {
	return IndexSetClass.IndexSet()
}

func (ic _IndexSetClass) Alloc() IndexSet {
	rv := objc.Call[IndexSet](ic, objc.Sel("alloc"))
	return rv
}

func (ic _IndexSetClass) New() IndexSet {
	rv := objc.Call[IndexSet](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIndexSet() IndexSet {
	return IndexSetClass.New()
}

func (i_ IndexSet) Init() IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("init"))
	return rv
}

// Indicates whether the index set contains a specific index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1410176-containsindex?language=objc
func (i_ IndexSet) ContainsIndex(value uint) bool {
	rv := objc.Call[bool](i_, objc.Sel("containsIndex:"), value)
	return rv
}

// Returns an NSIndexSet containing the receiving index set’s objects in the specified range that pass the Block test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1412153-indexesinrange?language=objc
func (i_ IndexSet) IndexesInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("indexesInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

// Executes a given Block over the index set’s indexes, using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1414545-enumerateindexeswithoptions?language=objc
func (i_ IndexSet) EnumerateIndexesWithOptionsUsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool)) {
	objc.Call[objc.Void](i_, objc.Sel("enumerateIndexesWithOptions:usingBlock:"), opts, block)
}

// Returns either the closest index in the index set that is greater than or equal to a specific index or the not-found indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1407870-indexgreaterthanorequaltoindex?language=objc
func (i_ IndexSet) IndexGreaterThanOrEqualToIndex(value uint) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexGreaterThanOrEqualToIndex:"), value)
	return rv
}

// Indicates whether the index set contains the indexes represented by an index range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1408511-containsindexesinrange?language=objc
func (i_ IndexSet) ContainsIndexesInRange(range_ Range) bool {
	rv := objc.Call[bool](i_, objc.Sel("containsIndexesInRange:"), range_)
	return rv
}

// Returns either the closest index in the index set that is less than a specific index or the not-found indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1417609-indexlessthanindex?language=objc
func (i_ IndexSet) IndexLessThanIndex(value uint) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexLessThanIndex:"), value)
	return rv
}

// Executes a given block using each object in the index set, in the specified ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1412673-enumeraterangeswithoptions?language=objc
func (i_ IndexSet) EnumerateRangesWithOptionsUsingBlock(opts EnumerationOptions, block func(range_ Range, stop *bool)) {
	objc.Call[objc.Void](i_, objc.Sel("enumerateRangesWithOptions:usingBlock:"), opts, block)
}

// Returns the index of the first object that passes the predicate Block test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1408471-indexpassingtest?language=objc
func (i_ IndexSet) IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexPassingTest:"), predicate)
	return rv
}

// Executes a given Block using each object in the index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1411395-enumerateindexesusingblock?language=objc
func (i_ IndexSet) EnumerateIndexesUsingBlock(block func(idx uint, stop *bool)) {
	objc.Call[objc.Void](i_, objc.Sel("enumerateIndexesUsingBlock:"), block)
}

// Returns the number of indexes in the index set that are members of a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1410114-countofindexesinrange?language=objc
func (i_ IndexSet) CountOfIndexesInRange(range_ Range) uint {
	rv := objc.Call[uint](i_, objc.Sel("countOfIndexesInRange:"), range_)
	return rv
}

// Returns an NSIndexSet containing the receiving index set’s objects that pass the Block test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1407357-indexespassingtest?language=objc
func (i_ IndexSet) IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("indexesPassingTest:"), predicate)
	return rv
}

// The index set fills an index buffer with the indexes contained both in the index set and in an index range, returning the number of indexes copied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1409332-getindexes?language=objc
func (i_ IndexSet) GetIndexesMaxCountInIndexRange(indexBuffer *uint, bufferSize uint, range_ RangePointer) uint {
	rv := objc.Call[uint](i_, objc.Sel("getIndexes:maxCount:inIndexRange:"), indexBuffer, bufferSize, range_)
	return rv
}

// Executes a given Block using the indexes in the specified range, using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1408162-enumerateindexesinrange?language=objc
func (i_ IndexSet) EnumerateIndexesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(idx uint, stop *bool)) {
	objc.Call[objc.Void](i_, objc.Sel("enumerateIndexesInRange:options:usingBlock:"), range_, opts, block)
}

// Returns the index of the first object that passes the predicate Block test using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1415860-indexwithoptions?language=objc
func (i_ IndexSet) IndexWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns either the closest index in the index set that is greater than a specific index or the not-found indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1407192-indexgreaterthanindex?language=objc
func (i_ IndexSet) IndexGreaterThanIndex(value uint) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexGreaterThanIndex:"), value)
	return rv
}

// Enumerates over the ranges in the range of objects using the block [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1416352-enumeraterangesinrange?language=objc
func (i_ IndexSet) EnumerateRangesInRangeOptionsUsingBlock(range_ Range, opts EnumerationOptions, block func(range_ Range, stop *bool)) {
	objc.Call[objc.Void](i_, objc.Sel("enumerateRangesInRange:options:usingBlock:"), range_, opts, block)
}

// Returns either the closest index in the index set that is less than or equal to a specific index or the not-found indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1412299-indexlessthanorequaltoindex?language=objc
func (i_ IndexSet) IndexLessThanOrEqualToIndex(value uint) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexLessThanOrEqualToIndex:"), value)
	return rv
}

// Indicates whether the receiving index set contains a superset of the indexes in another index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1414823-containsindexes?language=objc
func (i_ IndexSet) ContainsIndexes(indexSet IIndexSet) bool {
	rv := objc.Call[bool](i_, objc.Sel("containsIndexes:"), indexSet)
	return rv
}

// Returns an NSIndexSet containing the receiving index set’s objects that pass the Block test using the specified enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1412401-indexeswithoptions?language=objc
func (i_ IndexSet) IndexesWithOptionsPassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](i_, objc.Sel("indexesWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the index of the first object in the specified range that passes the predicate Block test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1415301-indexinrange?language=objc
func (i_ IndexSet) IndexInRangeOptionsPassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

// Indicates whether the indexes in the receiving index set are the same indexes contained in another index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1412212-isequaltoindexset?language=objc
func (i_ IndexSet) IsEqualToIndexSet(indexSet IIndexSet) bool {
	rv := objc.Call[bool](i_, objc.Sel("isEqualToIndexSet:"), indexSet)
	return rv
}

// Executes a given block using each object in the index set, in the specified ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1409668-enumeraterangesusingblock?language=objc
func (i_ IndexSet) EnumerateRangesUsingBlock(block func(range_ Range, stop *bool)) {
	objc.Call[objc.Void](i_, objc.Sel("enumerateRangesUsingBlock:"), block)
}

// Indicates whether the index set contains any of the indexes in a range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1413352-intersectsindexesinrange?language=objc
func (i_ IndexSet) IntersectsIndexesInRange(range_ Range) bool {
	rv := objc.Call[bool](i_, objc.Sel("intersectsIndexesInRange:"), range_)
	return rv
}

// The first index in the index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1410814-firstindex?language=objc
func (i_ IndexSet) FirstIndex() uint {
	rv := objc.Call[uint](i_, objc.Sel("firstIndex"))
	return rv
}

// The number of indexes in the index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1409648-count?language=objc
func (i_ IndexSet) Count() uint {
	rv := objc.Call[uint](i_, objc.Sel("count"))
	return rv
}

// The last index in the index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexset/1415020-lastindex?language=objc
func (i_ IndexSet) LastIndex() uint {
	rv := objc.Call[uint](i_, objc.Sel("lastIndex"))
	return rv
}
