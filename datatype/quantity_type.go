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

type QuantityComparator *CodeType

var (
	LessThanQuantityComparator           QuantityComparator = NewCode("<")
	LessOrEqualThanQuantityComparator    QuantityComparator = NewCode("<=")
	GreaterThanQuantityComparator        QuantityComparator = NewCode(">")
	GreaterOrEqualThanQuantityComparator QuantityComparator = NewCode(">=")
)

var UCUMSystemURI *URIType = NewURI("http://unitsofmeasure.org")

type QuantityType struct {
	value      DecimalAccessor
	comparator QuantityComparator
	unit       StringAccessor
	system     URIAccessor
	code       CodeAccessor
}

type QuantityAccessor interface {
	ElementAccessor

	Value() DecimalAccessor
	Comparator() QuantityComparator
	Unit() StringAccessor
	System() URIAccessor
	Code() CodeAccessor
}

type QuantityModifier interface {
	QuantityAccessor

	WithValue(value *DecimalType) QuantityModifier
	WithComparator(value QuantityComparator) QuantityModifier
	WithUnit(value *StringType) QuantityModifier
	WithSystem(value *URIType) QuantityModifier
	WithCode(value *CodeType) QuantityModifier
}

func NewEmptyQuantity() *QuantityType {
	return &QuantityType{}
}

func NewQuantity(value DecimalAccessor, comparator QuantityComparator,
	unit StringAccessor, system URIAccessor, code CodeAccessor) *QuantityType {
	return &QuantityType{value, comparator, unit, system, code}
}

func (t *QuantityType) DataType() DataTypes {
	return QuantityDataType
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

func (t *QuantityType) WithValue(value *DecimalType) QuantityModifier {
	quantity := *t
	quantity.value = value
	return &quantity
}

func (t *QuantityType) WithComparator(value QuantityComparator) QuantityModifier {
	quantity := *t
	quantity.comparator = value
	return &quantity
}

func (t *QuantityType) WithUnit(value *StringType) QuantityModifier {
	quantity := *t
	quantity.unit = value
	return &quantity
}

func (t *QuantityType) WithSystem(value *URIType) QuantityModifier {
	quantity := *t
	quantity.system = value
	return &quantity
}

func (t *QuantityType) WithCode(value *CodeType) QuantityModifier {
	quantity := *t
	quantity.code = value
	return &quantity
}
