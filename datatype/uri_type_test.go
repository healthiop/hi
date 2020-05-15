// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source URI must retain the above copyright notice, this
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

func TestURIImplementsAccessor(t *testing.T) {
	o := NewURI("test")
	assert.Implements(t, (*URIAccessor)(nil), o)
}

func TestURIDataType(t *testing.T) {
	o := NewURI("Test URI")
	dataType := o.DataType()
	assert.Equal(t, URIDataType, dataType)
}

func TestURITypeInfo(t *testing.T) {
	o := NewURI("Test URI")
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.uri", i.String())
	}
}

func TestNewURICollection(t *testing.T) {
	c := NewURICollection()
	assert.Equal(t, "FHIR.uri", c.ItemTypeInfo().String())
}

func TestURINil(t *testing.T) {
	o := NewURINil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, "", o.Value())
}

func TestURIValue(t *testing.T) {
	o := NewURI("test")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	value := o.Value()
	assert.Equal(t, "test", value)
}

func TestURIEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewURI("").Equal(newAccessorMock()))
}

func TestURIEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewURINil().Equal(NewURI("")))
}

func TestURIEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewURI("").Equal(NewURINil()))
}

func TestURIEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewURINil().Equal(NewURINil()))
}

func TestURIEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewURI("test").Equal(NewURI("test")))
}

func TestURIEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewURI("test1").Equal(NewURI("test2")))
}
