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

package resource

import (
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"testing"
)

func TestDynamicResourceType(t *testing.T) {
	var dynamicResource Accessor = NewDynamicResource("Patient")
	resourceType := dynamicResource.ResourceType()
	assert.Equal(t, "Patient", resourceType)
}

func TestDynamicResourceTypeInfo(t *testing.T) {
	var dynamicResource Accessor = NewDynamicResource("Patient")
	i := dynamicResource.TypeInfo()
	if assert.NotNil(t, i.FQName(), "name expected") {
		assert.Equal(t, "", i.FQName().Namespace())
		assert.Equal(t, "Patient", i.FQName().Name())
		assert.Equal(t, "Patient", i.FQName().String())
	}
	if assert.NotNil(t, i.Base, "base type expected") {
		assert.Equal(t, "FHIR.Resource", i.Base().String())
	}
}

func TestDynamicDataType(t *testing.T) {
	var dynamicResource Accessor = NewDynamicResource("Patient")
	dataType := dynamicResource.DataType()
	assert.Equal(t, datatype.ResourceDataType, dataType)
}

func TestDynamicResourceTypeEmpty(t *testing.T) {
	model := make(map[string]interface{})
	model["resourceType"] = "Patient"

	dynamicResource := NewDynamicResourceWithData(model)
	assert.True(t, dynamicResource.Empty(), "dynamic resource contains no properties")
}

func TestDynamicResourceTypeEmptyNilProps(t *testing.T) {
	model := make(map[string]interface{})
	model["resourceType"] = "Patient"
	model["id"] = nil

	dynamicResource := NewDynamicResourceWithData(model)
	assert.True(t, dynamicResource.Empty(), "dynamic resource contains no properties")
}

func TestDynamicResourceTypeNotEmpty(t *testing.T) {
	model := make(map[string]interface{})
	model["resourceType"] = "Patient"
	model["id"] = "abc123"

	dynamicResource := NewDynamicResourceWithData(model)
	assert.False(t, dynamicResource.Empty(), "dynamic resource contains properties")
}

func TestDynamicResourceTypeUndefined(t *testing.T) {
	model := make(map[string]interface{})
	dynamicResource := NewDynamicResourceWithData(model)
	resourceType := dynamicResource.ResourceType()
	assert.Equal(t, "", resourceType)
}

func TestDynamicResourceEqualTypeDiffers(t *testing.T) {
	r := NewDynamicResource("Patient")
	assert.Equal(t, false, r.Equal(datatype.NewInteger(0)))
}

func TestDynamicResourceEqual(t *testing.T) {
	m := make(map[string]interface{})
	m["age"] = 18.0
	r1 := NewDynamicResourceWithData(m)
	m = make(map[string]interface{})
	m["age"] = 18.0
	r2 := NewDynamicResourceWithData(m)
	assert.Equal(t, true, r1.Equal(r2), "same model must equal")
}

func TestDynamicResourceEqualDiffers(t *testing.T) {
	m := make(map[string]interface{})
	m["age"] = 18.0
	r1 := NewDynamicResourceWithData(m)
	m = make(map[string]interface{})
	m["age"] = 19.0
	r2 := NewDynamicResourceWithData(m)
	assert.Equal(t, false, r1.Equal(r2), "different model must not equal")
}
