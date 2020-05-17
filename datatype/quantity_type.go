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

type QuantityComparator CodeAccessor

var (
	LessThanQuantityComparator           QuantityComparator = NewCode("<")
	LessOrEqualThanQuantityComparator    QuantityComparator = NewCode("<=")
	GreaterThanQuantityComparator        QuantityComparator = NewCode(">")
	GreaterOrEqualThanQuantityComparator QuantityComparator = NewCode(">=")
)

var UCUMSystemURI *URIType = NewURI("http://unitsofmeasure.org")

var quantityTypeInfo = newElementTypeInfo("Quantity")

type QuantityType struct {
	value      DecimalAccessor
	comparator QuantityComparator
	unit       StringAccessor
	system     URIAccessor
	code       CodeAccessor
}

type QuantityAccessor interface {
	ElementAccessor
	EqualityEvaluator
	Negator

	Value() DecimalAccessor
	Comparator() QuantityComparator
	Unit() StringAccessor
	System() URIAccessor
	Code() CodeAccessor
}

type QuantityModifier interface {
	QuantityAccessor

	SetValue(value *DecimalType) QuantityModifier
	SetComparator(value QuantityComparator) QuantityModifier
	SetUnit(value *StringType) QuantityModifier
	SetSystem(value *URIType) QuantityModifier
	SetCode(value *CodeType) QuantityModifier
}

func NewQuantityCollection() *CollectionType {
	return NewCollection(quantityTypeInfo)
}

func NewEmptyQuantity() *QuantityType {
	return &QuantityType{}
}

func NewQuantity(value DecimalAccessor, comparator QuantityComparator,
	unit StringAccessor, system URIAccessor, code CodeAccessor) *QuantityType {
	return &QuantityType{
		value:      value,
		comparator: comparator,
		unit:       unit,
		system:     system,
		code:       code,
	}
}

func (t *QuantityType) DataType() DataTypes {
	return QuantityDataType
}

func (t *QuantityType) Empty() bool {
	return t.value == nil &&
		t.comparator == nil &&
		t.unit == nil &&
		t.system == nil &&
		t.code == nil
}

func (t *QuantityType) Value() DecimalAccessor {
	return t.value
}

func (t *QuantityType) Comparator() QuantityComparator {
	return t.comparator
}

func (t *QuantityType) Unit() StringAccessor {
	return t.unit
}

func (t *QuantityType) System() URIAccessor {
	return t.system
}

func (t *QuantityType) Code() CodeAccessor {
	return t.code
}

func (t *QuantityType) SetValue(value *DecimalType) QuantityModifier {
	t.value = value
	return t
}

func (t *QuantityType) SetComparator(value QuantityComparator) QuantityModifier {
	t.comparator = value
	return t
}

func (t *QuantityType) SetUnit(value *StringType) QuantityModifier {
	t.unit = value
	return t
}

func (t *QuantityType) SetSystem(value *URIType) QuantityModifier {
	t.system = value
	return t
}

func (t *QuantityType) SetCode(value *CodeType) QuantityModifier {
	t.code = value
	return t
}

func (t *QuantityType) Negate() Accessor {
	return NewQuantity(t.value.Negate().(DecimalAccessor), t.comparator, t.unit, t.system, t.code)
}

func (e *QuantityType) TypeInfo() TypeInfoAccessor {
	return quantityTypeInfo
}

func (t *QuantityType) Equal(accessor Accessor) bool {
	if o, ok := accessor.(QuantityAccessor); !ok {
		return false
	} else {
		return Equal(t.Value(), o.Value()) &&
			Equal(t.Comparator(), o.Comparator()) &&
			Equal(t.Unit(), o.Unit()) &&
			Equal(t.System(), o.System()) &&
			Equal(t.Code(), o.Code())
	}
}

func (t *QuantityType) ValueEqual(accessor Accessor) bool {
	if o, ok := accessor.(QuantityAccessor); !ok {
		return false
	} else {
		return Equal(t.Value(), o.Value()) &&
			Equal(t.System(), o.System()) &&
			Equal(t.Code(), o.Code())
	}
}

func (t *QuantityType) ValueEquivalent(accessor Accessor) bool {
	return t.ValueEqual(accessor)
}
