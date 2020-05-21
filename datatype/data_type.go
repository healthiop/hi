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

type DataTypes int

const UndefinedDataType DataTypes = 0x0001
const CollectionDataType DataTypes = 0x0002

const ElementDataType DataTypes = 0x1000
const PrimitiveDataType = ElementDataType + 0x0200
const ComplexDataType = ElementDataType + 0x0400

const ResourceDataType DataTypes = 0x2000

const (
	BooleanDataType = iota + PrimitiveDataType
	IntegerDataType
	StringDataType
	DecimalDataType
	URIDataType
	DateDataType
	DateTimeDataType
	TimeDataType
	CodeDataType
	IDDataType
	MarkdownDataType
	UnsignedIntDataType
	PositiveIntDataType
)

const (
	QuantityDataType = iota + ComplexDataType
)

const ElementTypeName = "Element"

var fqElementTypeName = NewFQTypeName(ElementTypeName, FHIRNamespaceName)

type Accessor interface {
	DataType() DataTypes
	TypeInfo() TypeInfoAccessor
	Empty() bool
	Equal(accessor Accessor) bool
	Equivalent(accessor Accessor) bool
}

type ElementType struct {
}

type ElementAccessor interface {
	Accessor
}

type Stringifier interface {
	String() string
}

func StringValue(stringifier Stringifier) string {
	if stringifier == nil {
		return ""
	}
	return stringifier.String()
}

type PrimitiveType struct {
	ElementType
	nilValue bool
}

type PrimitiveAccessor interface {
	ElementAccessor
	Stringifier
	Nil() bool
}

func (t *PrimitiveType) Nil() bool {
	return t.nilValue
}

func (t *PrimitiveType) Empty() bool {
	return t.Nil()
}

type Comparator interface {
	Accessor
	Compare(comparator Comparator) int
}

func TypeEqual(a1 Accessor, a2 Accessor) bool {
	return a1 != nil && a2 != nil && a1.DataType() == a2.DataType()
}

func IsPrimitive(accessor Accessor) bool {
	return accessor != nil && accessor.DataType()&PrimitiveDataType == PrimitiveDataType
}

func Equal(a1 Accessor, a2 Accessor) bool {
	return a1 == a2 || (Empty(a1) && Empty(a2)) ||
		(a1 != nil && a2 != nil && a1.Equal(a2))
}

func Equivalent(a1 Accessor, a2 Accessor) bool {
	if a1 == a2 || Empty(a1) && Empty(a2) {
		return true
	}
	if a1 == nil || a2 == nil {
		return false
	}
	return a1.Equivalent(a2)
}

func Empty(a Accessor) bool {
	return a == nil || a.Empty()
}

var elementTypeInfo = NewTypeInfoWithBase(fqElementTypeName, nil)

func newElementTypeInfo(name string) *TypeInfo {
	return newElementTypeInfoWithBase(name, elementTypeInfo)
}

func newElementTypeInfoWithBase(name string, base TypeInfoAccessor) *TypeInfo {
	return NewTypeInfoWithBase(NewFQTypeName(name, FHIRNamespaceName), base)
}
