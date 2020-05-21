// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package datatype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testTypeInfo = newElementTypeInfo("test")
var test2TypeInfo = newElementTypeInfo("test2")

func TestCollectionImplementsAccessor(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Implements(t, (*CollectionAccessor)(nil), c)
}

func TestCollectionDataType(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Equal(t, CollectionDataType, c.DataType())
}

func TestCollectionTypeInfo(t *testing.T) {
	c := NewCollection(testTypeInfo)
	i := c.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "Collection", i.String())
		assert.Nil(t, i.FQBaseName(), "no base name expected")
	}
}

func TestNewCollectionNoItemType(t *testing.T) {
	assert.Panics(t, func() { NewCollection(nil) })
}

func TestNewCollection(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Same(t, testTypeInfo, c.ItemTypeInfo())
	assert.True(t, c.Empty(), "new collection must be empty")
	assert.Equal(t, 0, c.Count())
}

func TestNewCollectionUndefined(t *testing.T) {
	c := NewCollectionUndefined()
	assert.Same(t, undefinedTypeInfo, c.ItemTypeInfo())
	assert.True(t, c.Empty(), "new collection must be empty")
	assert.Equal(t, 0, c.Count())
}

func TestNewCollectionGetEmpty(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Panics(t, func() { c.Get(0) })
}

func TestCollectionAddGet(t *testing.T) {
	item1 := newAccessorMock()
	item2 := newAccessorMock()
	c := NewCollection(testTypeInfo)
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
}

func TestCollectionAddUniqueEmpty(t *testing.T) {
	item1 := newAccessorMockWithValue(0)
	item2 := newAccessorMockWithValue(1)
	c := NewCollection(testTypeInfo)
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, true, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
}

func TestCollectionAddUniqueDupNil(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Equal(t, true, c.AddUnique(nil))
	assert.Equal(t, false, c.AddUnique(nil))
	assert.Equal(t, true, c.AddUnique(NewString("test")))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
}

func TestCollectionAddUniqueDiscard(t *testing.T) {
	item1 := newAccessorMockWithValue(1)
	item2 := newAccessorMockWithValue(1)
	c := NewCollection(testTypeInfo)
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, false, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
}

func TestCollectionAddUniqueExistingNil(t *testing.T) {
	item1 := newAccessorMockWithValue(1)
	c := NewCollection(testTypeInfo)
	assert.Equal(t, true, c.AddUnique(nil))
	assert.Equal(t, true, c.AddUnique(item1))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Nil(t, c.Get(0))
	assert.Same(t, item1, c.Get(1))
}

func TestCollectionStringAddUniqueEmpty(t *testing.T) {
	item1 := NewString("test1")
	item2 := NewString("test2")
	c := NewCollection(testTypeInfo)
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, true, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
}

func TestCollectionStringAddUniqueDiscard(t *testing.T) {
	item1 := NewString("test1")
	item2 := NewString("test1")
	c := NewCollection(testTypeInfo)
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, false, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
}

func TestCollectionAddAllUnique(t *testing.T) {
	item1 := NewString("test1")
	item2 := NewString("test2")
	item3 := NewString("test3")
	item4 := NewString("test2")
	item5 := NewString("test4")

	c1 := NewCollection(testTypeInfo)
	assert.Equal(t, true, c1.AddUnique(item1))
	assert.Equal(t, true, c1.AddUnique(item2))
	assert.Equal(t, true, c1.AddUnique(item3))

	c2 := NewCollection(testTypeInfo)
	assert.Equal(t, true, c2.AddUnique(item4))
	assert.Equal(t, true, c2.AddUnique(item5))

	c2.AddAllUnique(c1)
	if assert.Equal(t, 4, c2.Count()) {
		assert.Same(t, item4, c2.Get(0))
		assert.Same(t, item5, c2.Get(1))
		assert.Same(t, item1, c2.Get(2))
		assert.Same(t, item3, c2.Get(3))
	}
}

func TestCollectionAddAll(t *testing.T) {
	item1 := NewString("test1")
	item2 := NewString("test2")
	item3 := NewString("test3")
	item4 := NewString("test2")
	item5 := NewString("test4")

	c1 := NewCollection(testTypeInfo)
	assert.Equal(t, true, c1.AddUnique(item1))
	assert.Equal(t, true, c1.AddUnique(item2))
	assert.Equal(t, true, c1.AddUnique(item3))

	c2 := NewCollection(testTypeInfo)
	assert.Equal(t, true, c2.AddUnique(item4))
	assert.Equal(t, true, c2.AddUnique(item5))

	c2.AddAll(c1)
	if assert.Equal(t, 5, c2.Count()) {
		assert.Same(t, item4, c2.Get(0))
		assert.Same(t, item5, c2.Get(1))
		assert.Same(t, item1, c2.Get(2))
		assert.Same(t, item2, c2.Get(3))
		assert.Same(t, item3, c2.Get(4))
	}
}

func TestCollectionEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewCollection(testTypeInfo).Equal(newAccessorMock()))
	assert.Equal(t, false, NewCollection(testTypeInfo).Equivalent(newAccessorMock()))
}

func TestCollectionEqualNil(t *testing.T) {
	assert.Equal(t, false, NewCollection(testTypeInfo).Equal(nil))
	assert.Equal(t, false, NewCollection(testTypeInfo).Equivalent(nil))
}

func TestCollectionEqualEmpty(t *testing.T) {
	assert.Equal(t, true, NewCollection(testTypeInfo).Equal(NewCollection(testTypeInfo)))
	assert.Equal(t, true, NewCollection(testTypeInfo).Equivalent(NewCollection(testTypeInfo)))
}

func TestCollectionEqualItemTypeDiffers(t *testing.T) {
	assert.Equal(t, true, NewCollection(testTypeInfo).Equal(NewCollection(test2TypeInfo)))
	assert.Equal(t, true, NewCollection(testTypeInfo).Equivalent(NewCollection(test2TypeInfo)))
}

func TestCollectionEqual(t *testing.T) {
	c1 := NewCollection(testTypeInfo)
	c1.Add(newAccessorMockWithValue(0))
	c1.Add(newAccessorMockWithValue(1))
	c2 := NewCollection(testTypeInfo)
	c2.Add(newAccessorMockWithValue(0))
	c2.Add(newAccessorMockWithValue(1))
	assert.Equal(t, true, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestCollectionEqualOrderDiffers(t *testing.T) {
	c1 := NewCollection(testTypeInfo)
	c1.Add(newAccessorMockWithValue(0))
	c1.Add(newAccessorMockWithValue(1))
	c2 := NewCollection(testTypeInfo)
	c2.Add(newAccessorMockWithValue(1))
	c2.Add(newAccessorMockWithValue(0))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestCollectionEqualCountDiffers(t *testing.T) {
	c1 := NewCollection(testTypeInfo)
	c1.Add(newAccessorMockWithValue(0))
	c2 := NewCollection(testTypeInfo)
	c2.Add(newAccessorMockWithValue(0))
	c2.Add(newAccessorMockWithValue(0))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestCollectionEquivalent(t *testing.T) {
	c1 := NewStringCollection()
	c1.Add(NewString("Test Value"))
	c2 := NewStringCollection()
	c2.Add(NewString("test\nvalue"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}
