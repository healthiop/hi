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
	"fmt"
	"github.com/healthiop/hi"
	"regexp"
)

type TypeKind int
type SimpleType int

const (
	Base64BinaryTypeName = "base64Binary"
	BooleanTypeName      = "boolean"
	CodeTypeName         = "code"
	DateTypeName         = "date"
	DateTimeTypeName     = "dateTime"
	DecimalTypeName      = "decimal"
	ElementTypeName      = "Element"
	IDTypeName           = "id"
	InstantTypeName      = "instant"
	IntegerTypeName      = "integer"
	MarkdownTypeName     = "markdown"
	OIDTypeName          = "oid"
	PositiveIntTypeName  = "positiveInt"
	QuantityTypeName     = "Quantity"
	ResourceTypeName     = "Resource"
	StringTypeName       = "string"
	SIDTypeName          = "sid"
	TimeTypeName         = "time"
	UnsignedIntTypeName  = "unsignedInt"
	URITypeName          = "uri"
	UUIDTypeName         = "uuid"

	StringSimpleType SimpleType = SimpleType(hi.StringSimpleType)
	NumberSimpleType SimpleType = SimpleType(hi.NumberSimpleType)
	BoolSimpleType   SimpleType = SimpleType(hi.BoolSimpleType)

	AnyTypeKind       TypeKind = 0b_00000111
	PrimitiveTypeKind TypeKind = 0b_00000011
	ElementTypeKind   TypeKind = 0b_00000010
	ResourceTypeKind  TypeKind = 0b_00000100
)

type TypeDefContainer struct {
	symbolicVersion string
	versionString   string
	typeDefs        map[string]TypeDefAccessor
	elementTypeDef  StructTypeDefAccessor
}

type typeDef struct {
	internalName string
	base         TypeDefAccessor
}

type TypeDefAccessor interface {
	InternalName() string
	Base() TypeDefAccessor
	TypeKind() TypeKind
	Anonymous() bool
	Name() string
	ExtendsTypeName(name string) bool
}

type primitiveTypeDef struct {
	typeDef
	pattern    *regexp.Regexp
	simpleType SimpleType
}

type PrimitiveTypeDefAccessor interface {
	TypeDefAccessor
	Pattern() *regexp.Regexp
	SimpleType() SimpleType
}

type structTypeDef struct {
	typeDef
	typeKind    TypeKind
	anonymous   bool
	props       []*PropDef
	propsByName map[string]*PropDef
	choiceNames map[string]interface{}
}

type StructTypeDefAccessor interface {
	TypeDefAccessor
	PropByName(name string) *PropDef
}

type InitializableStructTypeDefAccessor interface {
	StructTypeDefAccessor
	InitProps(props []*PropDef)
}

type PropDef struct {
	name    string
	typeDef TypeDefAccessor
	choice  string
	array   bool
	enum    []string
}

func NewTypeDefContainer(symbolicVersion string, versionString string, typeDefs map[string]TypeDefAccessor) *TypeDefContainer {
	elementTypeDef, ok := typeDefs[ElementTypeName].(StructTypeDefAccessor)
	if !ok {
		panic("valid element type definition must have been defined")
	}
	return &TypeDefContainer{symbolicVersion, versionString, typeDefs, elementTypeDef}
}

func (t *TypeDefContainer) SymbolicVersion() string {
	return t.symbolicVersion
}

func (t *TypeDefContainer) VersionString() string {
	return t.versionString
}

func (t *TypeDefContainer) ElementType() StructTypeDefAccessor {
	return t.elementTypeDef
}

func (t *TypeDefContainer) TypeByName(name string) TypeDefAccessor {
	if t.typeDefs == nil {
		return nil
	}
	return t.typeDefs[name]
}

func (t *TypeDefContainer) MandatoryTypeByName(name string) TypeDefAccessor {
	typeDef := t.TypeByName(name)
	if typeDef == nil {
		panic(fmt.Sprintf("mandatory type has not been defined: %s", name))
	}
	return typeDef
}

func (t *TypeDefContainer) MandatoryStructTypeByName(name string) StructTypeDefAccessor {
	typeDef := t.TypeByName(name)
	if typeDef == nil {
		panic(fmt.Sprintf("mandatory typ has not been defined: %s", name))
	}
	if structTypeDef, ok := typeDef.(StructTypeDefAccessor); !ok {
		panic(fmt.Sprintf("mandatory type is no structure: %s", name))
	} else {
		return structTypeDef
	}
}

func NewPrimitiveTypeDef(name string, base TypeDefAccessor, pattern *regexp.Regexp, simpleType SimpleType) PrimitiveTypeDefAccessor {
	return &primitiveTypeDef{
		typeDef{
			name,
			base,
		},
		pattern,
		simpleType,
	}
}

func NewStructTypeDef(internalName string, base TypeDefAccessor, anonymous bool) InitializableStructTypeDefAccessor {
	var typeKind TypeKind
	if base == nil {
		switch internalName {
		case ElementTypeName:
			typeKind = ElementTypeKind
		case ResourceTypeName:
			typeKind = ResourceTypeKind
		default:
			panic(fmt.Sprintf("not a root type: %s", internalName))
		}
	} else if base.ExtendsTypeName(ResourceTypeName) {
		typeKind = ResourceTypeKind
	} else {
		typeKind = ElementTypeKind
	}

	return &structTypeDef{
		typeDef: typeDef{
			internalName,
			base,
		},
		typeKind:  typeKind,
		anonymous: anonymous,
	}
}

func (t *typeDef) Base() TypeDefAccessor {
	return t.base
}

func (t *typeDef) InternalName() string {
	return t.internalName
}

func (t *typeDef) ExtendsTypeName(name string) bool {
	if name == t.internalName {
		return true
	}

	if t.base != nil {
		return t.base.ExtendsTypeName(name)
	}

	return false
}

func (p *primitiveTypeDef) TypeKind() TypeKind {
	return PrimitiveTypeKind
}

func (p *primitiveTypeDef) Anonymous() bool {
	return false
}

func (p *primitiveTypeDef) Name() string {
	return p.internalName
}

func (p *primitiveTypeDef) Pattern() *regexp.Regexp {
	return p.pattern
}

func (p *primitiveTypeDef) SimpleType() SimpleType {
	return p.simpleType
}

func (t *structTypeDef) InitProps(props []*PropDef) {
	if t.props != nil {
		panic(fmt.Sprintf("properties of type %s have already been initialized", t.internalName))
	}
	t.props = props
	t.propsByName = propsByName(props)
	t.choiceNames = choiceNames(props)
}

func (t *structTypeDef) TypeKind() TypeKind {
	return t.typeKind
}

func (t *structTypeDef) Anonymous() bool {
	return t.anonymous
}

func (t *structTypeDef) Name() string {
	if t.anonymous {
		return ""
	}
	return t.internalName
}

func (t *structTypeDef) PropByName(name string) *PropDef {
	if t.propsByName == nil {
		return nil
	}
	return t.propsByName[name]
}

func propsByName(props []*PropDef) map[string]*PropDef {
	propsByName := make(map[string]*PropDef)
	for _, p := range props {
		propsByName[p.name] = p
	}
	return propsByName
}

func choiceNames(props []*PropDef) map[string]interface{} {
	var choices map[string]interface{}
	for _, p := range props {
		if len(p.choice) > 0 {
			if choices == nil {
				choices = make(map[string]interface{})
			}
			choices[p.choice] = nil
		}
	}
	return choices
}

func NewPropDef(name string, typeDef TypeDefAccessor, choice string, array bool, enum []string) *PropDef {
	return &PropDef{name, typeDef, choice, array, enum}
}

func (p *PropDef) Name() string {
	return p.name
}

func (p *PropDef) Type() TypeDefAccessor {
	return p.typeDef
}

func (p *PropDef) Choice() string {
	return p.choice
}

func (p *PropDef) Array() bool {
	return p.array
}

func (p *PropDef) Enum() []string {
	return p.enum
}
