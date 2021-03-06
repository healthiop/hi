// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
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

func TestNewElementTypeSpec(t *testing.T) {
	i := newElementTypeSpec("MyType")
	assert.Equal(t, "FHIR.MyType", i.String())
	assert.Equal(t, "MyType", i.FQName().Name())
	assert.Equal(t, "FHIR", i.FQName().Namespace())
	assert.Equal(t, "FHIR.MyType", i.FQName().String())
	assert.Equal(t, "Element", i.FQBaseName().Name())
	assert.Equal(t, "FHIR", i.FQBaseName().Namespace())
	assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
}

func TestEqualNil(t *testing.T) {
	assert.Equal(t, true, Equal(nil, nil))
}

func TestEqualSame(t *testing.T) {
	a := newAccessorMockWithValue(1)
	assert.Equal(t, true, Equal(a, a))
}

func TestEqualEmpty(t *testing.T) {
	assert.Equal(t, true, Equal(newAccessorMock(), newAccessorMock()))
}

func TestEqualEmptyDiffers(t *testing.T) {
	assert.Equal(t, false, Equal(newAccessorMock(), newAccessorMockWithValue(0)))
}

func TestEqual(t *testing.T) {
	assert.Equal(t, true, Equal(newAccessorMockWithValue(10), newAccessorMockWithValue(10)))
}

func TestEqualDiffer(t *testing.T) {
	assert.Equal(t, false, Equal(newAccessorMockWithValue(10), newAccessorMockWithValue(11)))
}

func TestValueEquivalentNil(t *testing.T) {
	assert.Equal(t, true, Equivalent(nil, nil))
}

func TestValueEquivalentLeftNil(t *testing.T) {
	assert.Equal(t, false, Equivalent(nil, newAccessorMockWithValue(0)))
}

func TestValueEquivalentRightNil(t *testing.T) {
	assert.Equal(t, false, Equivalent(newAccessorMockWithValue(0), nil))
}

func TestValueEquivalentSame(t *testing.T) {
	a := newAccessorMockWithValue(1)
	assert.Equal(t, true, Equivalent(a, a))
}

func TestValueEquivalentEmpty(t *testing.T) {
	assert.Equal(t, true, Equivalent(newAccessorMock(), newAccessorMock()))
}

func TestValueEquivalentEmptyDiffers(t *testing.T) {
	assert.Equal(t, false, Equivalent(NewStringNil(), NewString("")))
}

func TestValueEquivalent(t *testing.T) {
	assert.Equal(t, true, Equivalent(NewString("TEST"), NewString("test")))
}

func TestValueEquivalentDiffer(t *testing.T) {
	assert.Equal(t, false, Equivalent(NewString("test"), NewString("type")))
}

func TestIsPrimitive(t *testing.T) {
	assert.Equal(t, true, IsPrimitive(NewStringNil()))
}

func TestIsPrimitiveNot(t *testing.T) {
	assert.Equal(t, false, IsPrimitive(NewQuantityEmpty()))
}

func TestIsPrimitiveNil(t *testing.T) {
	assert.Equal(t, false, IsPrimitive(nil))
}

func TestDataTypeEqual(t *testing.T) {
	assert.Equal(t, true, TypeEqual(NewStringNil(), NewStringNil()))
}

func TestDataTypeEqualNot(t *testing.T) {
	assert.Equal(t, false, TypeEqual(NewStringNil(), NewIDNil()))
}

func TestStringValueNil(t *testing.T) {
	assert.Equal(t, "", StringValue(nil))
}

func TestStringValueNonNil(t *testing.T) {
	assert.Equal(t, "test", StringValue(NewString("test")))
}
