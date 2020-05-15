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

import "fmt"

var booleanTypeInfo = newElementTypeInfo("boolean")

type BooleanType struct {
	nilValue bool
	value    bool
}

type BooleanAccessor interface {
	PrimitiveAccessor
	Negator
	Value() bool
}

func NewBooleanNil() *BooleanType {
	return newBoolean(true, false)
}

func NewBoolean(value bool) *BooleanType {
	return newBoolean(false, value)
}

func newBoolean(nilValue bool, value bool) *BooleanType {
	return &BooleanType{
		nilValue: nilValue,
		value:    value,
	}
}

func ParseBoolean(value string) (*BooleanType, error) {
	switch value {
	case "true":
		return NewBoolean(true), nil
	case "false":
		return NewBoolean(false), nil
	}
	return nil, fmt.Errorf("not a boolean: %s", value)
}

func (t *BooleanType) DataType() DataTypes {
	return BooleanDataType
}

func (t *BooleanType) Nil() bool {
	return t.nilValue
}

func (t *BooleanType) Value() bool {
	return t.value
}

func (e *BooleanType) TypeInfo() TypeInfoAccessor {
	return booleanTypeInfo
}

func (t *BooleanType) Negate() Accessor {
	if t.nilValue {
		return t
	}
	return NewBoolean(!t.value)
}
