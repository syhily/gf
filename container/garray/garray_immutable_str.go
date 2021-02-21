// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package garray

// NewImmutableStrArray would create a immutable int array from go slice.
// We recommend using the StrArray.Immutable() instead of this for better customization.
func NewImmutableStrArray(array ...string) ImmutableStrArray {
	return NewStrArrayFromCopy(array, false)
}

// ImmutableStrArray is the StrArray which couldn't be modified after creating.
// All the functions from StrArray would be extends by ImmutableStrArray except the modify method.
type ImmutableStrArray interface {

	// Get returns the value by the specified index.
	// If the given <index> is out of range of the array, the <found> is false.
	Get(index int) (value string, found bool)

	// Sum returns the sum of values in an array.
	Sum() (sum int)

	// Range picks and returns items by range, like array[start:end].
	// Notice, if in concurrent-safe usage, it returns a copy of slice;
	// else a pointer to the underlying data.
	//
	// If <end> is negative, then the offset will start from the end of array.
	// If <end> is omitted, then the sequence will have everything from start up
	// until the end of the array.
	Range(start int, end ...int) []string

	// SubSlice returns a slice of elements from the array as specified
	// by the <offset> and <size> parameters.
	// If in concurrent safe usage, it returns a copy of the slice; else a pointer.
	//
	// If offset is non-negative, the sequence will start at that offset in the array.
	// If offset is negative, the sequence will start that far from the end of the array.
	//
	// If length is given and is positive, then the sequence will have up to that many elements in it.
	// If the array is shorter than the length, then only the available array elements will be present.
	// If length is given and is negative then the sequence will stop that many elements from the end of the array.
	// If it is omitted, then the sequence will have everything from offset up until the end of the array.
	//
	// Any possibility crossing the left border of array, it will fail.
	SubSlice(offset int, length ...int) []string

	// Len returns the length of array.
	Len() int

	// Slice returns the underlying data of array.
	// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
	// or else a pointer to the underlying data.
	Slice() []string

	// Interfaces returns current array as []interface{}.
	Interfaces() []interface{}

	// Contains checks whether a value exists in the array.
	Contains(value string) bool

	// ContainsI checks whether a value exists in the array with case-insensitively.
	// Note that it internally iterates the whole array to do the comparison with case-insensitively.
	ContainsI(value string) bool

	// Search searches array by <value>, returns the index of <value>,
	// or returns -1 if not exists.
	Search(value string) int

	// Chunk splits an array into multiple arrays,
	// the size of each array is determined by <size>.
	// The last chunk may contain less than size elements.
	Chunk(size int) [][]string

	// Rand randomly returns one item from array(no deleting).
	Rand() (string, bool)

	// Rands randomly returns <size> items from array(no deleting).
	Rands(size int) []string

	// Join joins array elements with a string <glue>.
	Join(glue string) string

	// CountValues counts the number of occurrences of all values in the array.
	CountValues() map[string]int

	// Iterator is alias of IteratorAsc.
	Iterator(f func(k int, v string) bool)

	// IteratorAsc iterates the array readonly in ascending order with given callback function <f>.
	// If <f> returns true, then it continues iterating; or false to stop.
	IteratorAsc(f func(k int, v string) bool)

	// IteratorDesc iterates the array readonly in descending order with given callback function <f>.
	// If <f> returns true, then it continues iterating; or false to stop.
	IteratorDesc(f func(k int, v string) bool)

	// String returns current array as a string, which implements like json.Marshal does.
	String() string

	// MarshalJSON implements the interface MarshalJSON for json.Marshal.
	// Note that do not use pointer as its receiver here.
	MarshalJSON() ([]byte, error)

	// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
	UnmarshalJSON(b []byte) error

	// UnmarshalValue is an interface implement which sets any type of value for array.
	UnmarshalValue(value interface{}) error

	// IsEmpty checks whether the array is empty.
	IsEmpty() bool
}