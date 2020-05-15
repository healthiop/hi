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
)

var decimalTypeInfo = newElementTypeInfo("decimal")

type DecimalType struct {
	value decimal.Decimal
}

type DecimalAccessor interface {
	NumberAccessor
}

func NewDecimalInt(value int32) *DecimalType {
	return newDecimal(decimal.NewFromInt32(value))
}

func NewDecimalInt64(value int64) *DecimalType {
	return newDecimal(decimal.NewFromInt(value))
}

func NewDecimalFloat64(value float64) *DecimalType {
	return newDecimal(decimal.NewFromFloat(value))
}

func ParseDecimal(value string) (*DecimalType, error) {
	if d, err := decimal.NewFromString(value); err != nil {
		return nil, fmt.Errorf("not a decimal: %s", value)
	} else {
		return newDecimal(d), nil
	}
}

func newDecimal(value decimal.Decimal) *DecimalType {
	return &DecimalType{
		value: value,
	}
}

func (t *DecimalType) DataType() DataTypes {
	return DecimalDataType
}

func (t *DecimalType) Int() int32 {
	return int32(t.value.IntPart())
}

func (t *DecimalType) Int64() int64 {
	return t.value.IntPart()
}

func (t *DecimalType) Float32() float32 {
	v, _ := t.value.Float64()
	return float32(v)
}

func (t *DecimalType) Float64() float64 {
	v, _ := t.value.Float64()
	return v
}

func (t *DecimalType) BigFloat() *big.Float {
	return t.value.BigFloat()
}

func (t *DecimalType) Decimal() decimal.Decimal {
	return t.value
}

func (e *DecimalType) TypeInfo() TypeInfoAccessor {
	return decimalTypeInfo
}

func (t *DecimalType) Negate() Accessor {
	return newDecimal(t.value.Neg())
}
