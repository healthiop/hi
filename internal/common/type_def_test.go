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

package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTypeDefContainerNoElement(t *testing.T) {
	assert.Panics(t, func() {
		NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
			"Resource": NewStructTypeDef("Resource", nil, false),
		})
	})
}

func TestTypeDefContainerSymbolicVersion(t *testing.T) {
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": NewStructTypeDef("Element", nil, false),
	})
	assert.Equal(t, "R1", tdc.SymbolicVersion())
}

func TestTypeDefContainerVersionString(t *testing.T) {
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": NewStructTypeDef("Element", nil, false),
	})
	assert.Equal(t, "2.3.1", tdc.VersionString())
}

func TestNewTypeDefContainerTypeByNameUninitialized(t *testing.T) {
	tdc := TypeDefContainer{"R1", "2.3.1", nil, nil}
	assert.Nil(t, tdc.TypeByName("Element"))
}

func TestTypeDefContainerTypeByName(t *testing.T) {
	td := NewStructTypeDef("Element", nil, false)
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": td,
	})
	assert.Same(t, td, tdc.TypeByName("Element"))
}

func TestTypeDefContainerTypeByNameNotFound(t *testing.T) {
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": NewStructTypeDef("Element", nil, false),
	})
	assert.Nil(t, tdc.TypeByName("Element2"))
}

func TestTypeDefContainerMandatoryTypeByName(t *testing.T) {
	td := NewStructTypeDef("Element", nil, false)
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": td,
	})
	assert.Same(t, td, tdc.MandatoryTypeByName("Element"))
}

func TestTypeDefContainerMandatoryTypeByNameNotFound(t *testing.T) {
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": NewStructTypeDef("Element", nil, false),
	})
	assert.Panics(t, func() { tdc.MandatoryTypeByName("Element2") })
}

func TestTypeDefContainerMandatoryStructTypeByName(t *testing.T) {
	td := NewStructTypeDef("Element", nil, false)
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": td,
	})
	assert.Same(t, td, tdc.MandatoryStructTypeByName("Element"))
}

func TestTypeDefContainerMandatoryStructTypeByNameNotFound(t *testing.T) {
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": NewStructTypeDef("Element", nil, false),
	})
	assert.Panics(t, func() { tdc.MandatoryStructTypeByName("Element2") })
}

func TestTypeDefContainerMandatoryStructTypeByNamePrimitive(t *testing.T) {
	elementTypeDef := NewStructTypeDef("Element", nil, false)
	tdc := NewTypeDefContainer("R1", "2.3.1", map[string]TypeDefAccessor{
		"Element": elementTypeDef,
		"string":  NewPrimitiveTypeDef("string", elementTypeDef, nil, StringSimpleType),
	})
	assert.Panics(t, func() { tdc.MandatoryStructTypeByName("string") })
}

func TestNewStructTypeElement(t *testing.T) {
	td := NewStructTypeDef("Element", nil, false)
	assert.Equal(t, "Element", td.InternalName())
	assert.Nil(t, td.Base())
	assert.False(t, td.Anonymous())
	assert.Equal(t, ElementTypeKind, td.TypeKind())
}

func TestNewStructTypeElementBase(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity", td1, false)
	td3 := NewStructTypeDef("MoneyQuantity", td2, false)
	assert.Equal(t, "MoneyQuantity", td3.InternalName())
	assert.Same(t, td2, td3.Base())
	assert.Equal(t, ElementTypeKind, td3.TypeKind())
}

func TestNewStructTypeResource(t *testing.T) {
	td := NewStructTypeDef("Resource", nil, false)
	assert.Equal(t, "Resource", td.InternalName())
	assert.Nil(t, td.Base())
	assert.False(t, td.Anonymous())
	assert.Equal(t, ResourceTypeKind, td.TypeKind())
}

func TestNewStructTypeResourceBase(t *testing.T) {
	td1 := NewStructTypeDef("Resource", nil, false)
	td2 := NewStructTypeDef("DomainResource", td1, false)
	td3 := NewStructTypeDef("Patient", td2, false)
	assert.Equal(t, "Patient", td3.InternalName())
	assert.Same(t, td2, td3.Base())
	assert.Equal(t, ResourceTypeKind, td3.TypeKind())
}

func TestNewStructTypeInvalidRoot(t *testing.T) {
	assert.Panics(t, func() { NewStructTypeDef("Patient", nil, false) })
}

func TestTypeDefExtendsTypeNameLevel2(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity", td1, false)
	assert.True(t, td2.ExtendsTypeName("Element"))
}

func TestTypeDefExtendsTypeNameLevel2Not(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity", td1, false)
	assert.False(t, td2.ExtendsTypeName("Resource"))
}

func TestTypeDefExtendsTypeNameLevel1(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity", td1, false)
	assert.True(t, td2.ExtendsTypeName("Quantity"))
}

func TestTypeDefExtendsTypeNameLevel1Not(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity", td1, false)
	assert.False(t, td2.ExtendsTypeName("Resource"))
}

func TestNewStructTypeElementNotAnonymous(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity", td1, false)
	assert.Equal(t, "Quantity", td2.InternalName())
	assert.Equal(t, "Quantity", td2.Name())
	assert.Same(t, td1, td2.Base())
	assert.False(t, td2.Anonymous())
	assert.Equal(t, ElementTypeKind, td2.TypeKind())
}

func TestNewStructTypeElementAnonymous(t *testing.T) {
	td1 := NewStructTypeDef("Element", nil, false)
	td2 := NewStructTypeDef("Quantity_Test", td1, true)
	assert.Equal(t, "Quantity_Test", td2.InternalName())
	assert.Equal(t, "", td2.Name())
	assert.Same(t, td1, td2.Base())
	assert.True(t, td2.Anonymous())
	assert.Equal(t, ElementTypeKind, td2.TypeKind())
}
