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

package r4

import (
	"fmt"
	"github.com/healthiop/hi/internal/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var typeInheritanceTests = []struct {
	name      string
	baseName  string
	typeKind  common.TypeKind
	anonymous bool
}{
	{"dateTime", "Element", common.PrimitiveTypeKind, false},
	{"string", "Element", common.PrimitiveTypeKind, false},
	{"code", "string", common.PrimitiveTypeKind, false},
	{"markdown", "string", common.PrimitiveTypeKind, false},
	{"id", "string", common.PrimitiveTypeKind, false},
	{"unsignedInt", "integer", common.PrimitiveTypeKind, false},
	{"positiveInt", "integer", common.PrimitiveTypeKind, false},
	{"url", "uri", common.PrimitiveTypeKind, false},
	{"canonical", "uri", common.PrimitiveTypeKind, false},
	{"oid", "uri", common.PrimitiveTypeKind, false},
	{"uuid", "uri", common.PrimitiveTypeKind, false},
	{"Period", "Element", common.ElementTypeKind, false},
	{"BackboneElement", "Element", common.ElementTypeKind, false},
	{"Quantity", "Element", common.ElementTypeKind, false},
	{"ContactDetail", "Element", common.ElementTypeKind, false},
	{"Age", "Quantity", common.ElementTypeKind, false},
	{"Distance", "Quantity", common.ElementTypeKind, false},
	{"Duration", "Quantity", common.ElementTypeKind, false},
	{"Count", "Quantity", common.ElementTypeKind, false},
	{"Timing", "BackboneElement", common.ElementTypeKind, false},
	{"Dosage", "BackboneElement", common.ElementTypeKind, false},
	{"ElementDefinition", "BackboneElement", common.ElementTypeKind, false},
	{"Patient_Contact", "BackboneElement", common.ElementTypeKind, true},
	{common.ElementTypeName, "", common.ElementTypeKind, false},
	{"Resource", "", common.ResourceTypeKind, false},
	{"DomainResource", "Resource", common.ResourceTypeKind, false},
	{"Bundle", "Resource", common.ResourceTypeKind, false},
	{"Patient", "DomainResource", common.ResourceTypeKind, false},
}

var primitiveTypeTests = []struct {
	name       string
	pattern    string
	simpleType common.SimpleType
}{
	{"string", "^[ \\r\\n\\t\\S]+$", common.StringSimpleType},
	{"id", "^[A-Za-z0-9\\-.]{1,64}$", common.StringSimpleType},
	{"integer", "^-?([0]|([1-9][0-9]*))$", common.NumberSimpleType},
	{"boolean", "^true|false$", common.BoolSimpleType},
}

var typePropTests = []struct {
	typeName     string
	propName     string
	propTypeName string
	choice       string
	array        bool
	enum         []string
}{
	{"Patient", "language", "code", "", false, nil},
	{"Patient", "address", "Address", "", true, nil},
	{"Patient", "gender", "string", "", false, []string{
		"male",
		"female",
		"other",
		"unknown",
	}},
	{"Extension", "valueBase64Binary", "base64Binary", "value", false, nil},
	{"Extension", "valueCanonical", "canonical", "value", false, nil},
	{"Extension", "valueMarkdown", "markdown", "value", false, nil},
	{"Extension", "valueString", "string", "value", false, nil},
	{"Extension", "valueUri", "uri", "value", false, nil},
	{"Extension", "valueUrl", "url", "value", false, nil},
	{"Extension", "valueInteger", "integer", "value", false, nil},
}

func TestTypeDefContainerSingleton(t *testing.T) {
	tsc := TypeDefContainer()
	assert.Same(t, tsc, TypeDefContainer(), "typeDefContainer must be a singleton")
}

func TestTypeSpecSymbolicVersion(t *testing.T) {
	tsc := TypeDefContainer()
	assert.Equal(t, "R4", tsc.SymbolicVersion())
}

func TestTypeSpecVersionString(t *testing.T) {
	tsc := TypeDefContainer()
	assert.Equal(t, "4.0.1", tsc.VersionString())
}

func TestTypeDefContainerContainsPatient(t *testing.T) {
	tsc := TypeDefContainer()
	patientType := tsc.TypeByName("Patient")
	if assert.NotNil(t, patientType, "patient type must have been defined") {
		assert.Equal(t, "Patient", patientType.InternalName())

		if assert.Implements(t, (*common.StructTypeDefAccessor)(nil), patientType) {
			structType := patientType.(common.StructTypeDefAccessor)
			assert.Nil(t, structType.PropByName("resourceType"), "resource must not contain resource type property")
		}
	}
}

func TestTypeInheritances(t *testing.T) {
	tsc := TypeDefContainer()
	for _, tt := range typeInheritanceTests {
		t.Run(tt.name, func(t *testing.T) {
			if td := tsc.TypeByName(tt.name); td == nil {
				t.Errorf("type %s has not been defined", tt.name)
			} else {
				assert.Equal(t, tt.name, td.InternalName())
				var baseName string
				if td.Base() != nil {
					baseName = td.Base().InternalName()
				}
				assert.Equal(t, tt.baseName, baseName)
				assert.Equal(t, tt.typeKind, td.TypeKind())
				if std, ok := td.(common.StructTypeDefAccessor); ok {
					assert.Equal(t, tt.anonymous, std.Anonymous())
				} else if tt.anonymous {
					t.Errorf("type is no struct but anonymous struct expected: %T", td)
				}
			}
		})
	}
}

func TestPrimitiveTypes(t *testing.T) {
	tsc := TypeDefContainer()
	for _, tt := range primitiveTypeTests {
		t.Run(tt.name, func(t *testing.T) {
			if td := tsc.TypeByName(tt.name); td == nil {
				t.Errorf("type %s has not been defined", tt.name)
			} else {
				if ptd, ok := td.(common.PrimitiveTypeDefAccessor); !ok {
					t.Errorf("type is no primitive: %T", td)
				} else {
					if len(tt.pattern) == 0 {
						assert.Nil(t, ptd.Pattern(), "no pattern expected")
					} else if assert.NotNil(t, ptd.Pattern(), "pattern expected") {
						assert.Equal(t, tt.pattern, ptd.Pattern().String())
					}
					assert.Equal(t, common.SimpleType(tt.simpleType), ptd.SimpleType())
				}
			}
		})
	}
}

func TestTypeProps(t *testing.T) {
	tsc := TypeDefContainer()
	for _, tt := range typePropTests {
		t.Run(fmt.Sprintf("%s.%s", tt.typeName, tt.propName), func(t *testing.T) {
			if td := tsc.TypeByName(tt.typeName); td == nil {
				t.Errorf("type %s has not been defined", tt.typeName)
			} else {
				if std, ok := td.(common.StructTypeDefAccessor); !ok {
					t.Errorf("type is no structure: %T", td)
				} else {
					p := std.PropByName(tt.propName)
					if p == nil {
						t.Error("type does not define property")
					} else {
						assert.Equal(t, tt.propName, p.Name())
						if assert.NotNil(t, p.Type(), "type must have been set") {
							assert.Equal(t, tt.propTypeName, p.Type().InternalName())
						}
						assert.Equal(t, tt.choice, p.Choice())
						assert.Equal(t, tt.array, p.Array())
						assert.Equal(t, tt.enum, p.Enum())
					}
				}
			}
		})
	}
}
