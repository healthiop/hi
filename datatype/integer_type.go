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
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

var integerTypeInfo = newElementTypeInfo("integer")

type integerType struct {
	PrimitiveType
	value int32
}

type IntegerAccessor interface {
	NumberAccessor
}

func IsInteger(accessor Accessor) bool {
	dt := accessor.DataType()
	return dt == IntegerDataType ||
		dt == PositiveIntDataType ||
		dt == UnsignedIntDataType
}

func NewIntegerNil() IntegerAccessor {
	return newInteger(true, 0)
}

func NewInteger(value int32) IntegerAccessor {
	return newInteger(false, value)
}

func ParseInteger(value string) (IntegerAccessor, error) {
	if i, err := strconv.Atoi(value); err != nil {
		return nil, fmt.Errorf("not an integer: %s", value)
	} else {
		return NewInteger(int32(i)), nil
	}
}

func newInteger(nilValue bool, value int32) IntegerAccessor {
	return &integerType{
		PrimitiveType: PrimitiveType{
			nilValue: nilValue,
		},
		value: value,
	}
}

func (t *integerType) DataType() DataTypes {
	return IntegerDataType
}

func (t *integerType) Int() int32 {
	return t.value
}

func (t *integerType) Int64() int64 {
	return int64(t.value)
}

func (t *integerType) Float32() float32 {
	return float32(t.value)
}

func (t *integerType) Float64() float64 {
	return float64(t.value)
}

func (t *integerType) BigFloat() *big.Float {
	return big.NewFloat(float64(t.value))
}

func (t *integerType) Decimal() decimal.Decimal {
	return decimal.NewFromInt32(t.value)
}

func (t *integerType) TypeInfo() TypeInfoAccessor {
	return integerTypeInfo
}

func (t *integerType) Equal(accessor Accessor) bool {
	if accessor != nil && IsInteger(accessor) {
		o := accessor.(IntegerAccessor)
		return t.Nil() == o.Nil() && t.Int() == o.Int()
	}

	return decimalValueEqual(t, accessor)
}

func (t *integerType) Equivalent(accessor Accessor) bool {
	if accessor != nil && IsInteger(accessor) {
		o := accessor.(IntegerAccessor)
		return t.Nil() == o.Nil() && t.Int() == o.Int()
	}

	return decimalValueEquivalent(t, accessor)
}

func (t *integerType) String() string {
	if t.nilValue {
		return ""
	}

	return strconv.FormatInt(int64(t.value), 10)
}
