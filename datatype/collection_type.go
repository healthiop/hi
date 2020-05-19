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

var collectionTypeInfo = NewTypeInfoWithBase(NewTypeName("Collection"), nil)

type CollectionType struct {
	itemTypeInfo TypeInfoAccessor
	items        []Accessor
}

type CollectionAccessor interface {
	Accessor
	ItemTypeInfo() TypeInfoAccessor
	Count() int
	Get(i int) Accessor
}

type CollectionModifier interface {
	CollectionAccessor
	Add(accessor Accessor)
	AddUnique(accessor Accessor) bool
	AddAllUnique(collectionAccessor CollectionAccessor) int
}

func NewCollection(itemTypeInfo TypeInfoAccessor) *CollectionType {
	if itemTypeInfo == nil {
		panic("no item type has been specified")
	}
	return &CollectionType{
		itemTypeInfo: itemTypeInfo,
	}
}

func NewCollectionUndefined() *CollectionType {
	return NewCollection(undefinedTypeInfo)
}

func (c *CollectionType) DataType() DataTypes {
	return CollectionDataType
}

func (c *CollectionType) TypeInfo() TypeInfoAccessor {
	return collectionTypeInfo
}

func (c *CollectionType) ItemTypeInfo() TypeInfoAccessor {
	return c.itemTypeInfo
}

func (c *CollectionType) Empty() bool {
	return c.Count() == 0
}

func (c *CollectionType) Count() int {
	if c.items == nil {
		return 0
	}
	return len(c.items)
}

func (c *CollectionType) Get(i int) Accessor {
	if c.items == nil {
		panic("collection is empty")
	}
	return c.items[i]
}

func (c *CollectionType) Add(accessor Accessor) {
	if c.items == nil {
		c.items = make([]Accessor, 0)
	}
	c.items = append(c.items, accessor)
}

func (c *CollectionType) AddUnique(accessor Accessor) bool {
	if c.items == nil || accessor == nil {
		c.Add(accessor)
		return true
	}

	for _, item := range c.items {
		if item != nil && accessor.ValueEqual(item) {
			return false
		}
	}
	c.Add(accessor)
	return true
}

func (c *CollectionType) AddAllUnique(collectionAccessor CollectionAccessor) int {
	added := 0
	count := collectionAccessor.Count()
	for i := 0; i < count; i++ {
		if c.AddUnique(collectionAccessor.Get(i)) {
			added = added + 1
		}
	}
	return added
}

func (c *CollectionType) Equal(accessor Accessor) bool {
	if o, ok := accessor.(CollectionAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() && collectionDeepEqual(c, o)
	}
}

func (c *CollectionType) ValueEqual(accessor Accessor) bool {
	if o, ok := accessor.(CollectionAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() && collectionDeepValueEqual(c, o)
	}
}

func (c *CollectionType) ValueEquivalent(accessor Accessor) bool {
	if o, ok := accessor.(CollectionAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() && collectionDeepValueEquivalent(c, o)
	}
}

func collectionDeepEqual(c1 CollectionAccessor, c2 CollectionAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !Equal(c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func collectionDeepValueEqual(c1 CollectionAccessor, c2 CollectionAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ValueEqual(c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func collectionDeepValueEquivalent(c1 CollectionAccessor, c2 CollectionAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ValueEquivalent(c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}
