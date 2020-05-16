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
	"fmt"
	"regexp"
)

var stringTypeInfo = newElementTypeInfo("string")

var stringRegexp = regexp.MustCompile("^[\\r\\n\\t\\x{0020}-\\x{FFFF}]*$")

type StringType struct {
	nilValue bool
	value    string
}

type StringAccessor interface {
	PrimitiveAccessor
}

func NewStringCollection() *CollectionType {
	return NewCollection(stringTypeInfo)
}

func NewStringNil() *StringType {
	return newString(true, "")
}

func NewString(value string) *StringType {
	if !stringRegexp.MatchString(value) {
		panic(fmt.Sprintf("not a valid string: %s", value))
	}
	return newString(false, value)
}

func ParseString(value string) (*StringType, error) {
	if !stringRegexp.MatchString(value) {
		return nil, fmt.Errorf("not a valid string: %s", value)
	}
	return newString(false, value), nil
}

func newString(nilValue bool, value string) *StringType {
	return &StringType{
		nilValue: nilValue,
		value:    value,
	}
}

func (t *StringType) Empty() bool {
	return t.Nil()
}

func (t *StringType) Nil() bool {
	return t.nilValue
}

func (t *StringType) DataType() DataTypes {
	return StringDataType
}

func (t *StringType) String() string {
	return t.value
}

func (e *StringType) TypeInfo() TypeInfoAccessor {
	return stringTypeInfo
}

func (t *StringType) Equal(accessor Accessor) bool {
	if o, ok := accessor.(PrimitiveAccessor); !ok {
		return false
	} else {
		return t.Nil() == o.Nil() && t.String() == o.String()
	}
}
