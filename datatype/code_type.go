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

var codeTypeInfo = newElementTypeInfoWithBase("code", stringTypeInfo)

var codeRegexp = regexp.MustCompile("^[^\\s]+(\\s[^\\s]+)*$")

type CodeType struct {
	StringType
}

type CodeAccessor interface {
	StringAccessor
}

func NewCodeCollection() *CollectionType {
	return NewCollection(codeTypeInfo)
}

func NewCodeNil() *CodeType {
	return newCode(true, "")
}

func NewCode(value string) *CodeType {
	if !codeRegexp.MatchString(value) {
		panic(fmt.Sprintf("not a valid code: %s", value))
	}
	return newCode(false, value)
}

func ParseCode(value string) (*CodeType, error) {
	if !codeRegexp.MatchString(value) {
		return nil, fmt.Errorf("not a valid code: %s", value)
	}
	return newCode(false, value), nil
}

func newCode(nilValue bool, value string) *CodeType {
	return &CodeType{
		StringType{
			nilValue: nilValue,
			value:    value,
		},
	}
}

func (t *CodeType) DataType() DataTypes {
	return CodeDataType
}

func (e *CodeType) TypeInfo() TypeInfoAccessor {
	return codeTypeInfo
}
