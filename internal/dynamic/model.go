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

package dynamic

import (
	"fmt"
	"github.com/healthiop/hi"
	"github.com/healthiop/hi/internal/common"
	"strconv"
)

const ResourceTypePropName = "resourceType"

type DynTypeDefRetriever interface {
	Type() common.TypeDefAccessor
}

type dyn struct {
	typeDefContainer *common.TypeDefContainer
	parent           hi.DynAccessor
}

type dynList struct {
	dyn
	itemTypeDef common.TypeDefAccessor
	items       []interface{}
}

type dynBase struct {
	dyn
}

type dynStruct struct {
	dynBase
	typeDef common.StructTypeDefAccessor
	data    map[string]interface{}
}

type dynPrimitive struct {
	dynStruct
	primitiveTypeDef common.PrimitiveTypeDefAccessor
	value            interface{}
}

func NewDynList(typeDefContainer *common.TypeDefContainer, itemTypeDef common.TypeDefAccessor, items []interface{}, parent hi.DynAccessor) (hi.DynListAccessor, error) {
	return &dynList{
		dyn{typeDefContainer, parent},
		itemTypeDef,
		items,
	}, nil
}

func NewDynPrimitive(typeDefContainer *common.TypeDefContainer, typeDef common.PrimitiveTypeDefAccessor, value interface{}, element map[string]interface{}, parent hi.DynAccessor) (hi.DynPrimitiveAccessor, error) {
	return &dynPrimitive{
		dynStruct{
			dynBase{
				dyn{typeDefContainer, parent},
			},
			typeDefContainer.ElementType(),
			element,
		},
		typeDef,
		value,
	}, nil
}

func NewDynElement(typeDefContainer *common.TypeDefContainer, typeDef common.StructTypeDefAccessor, data map[string]interface{}, parent hi.DynAccessor) (hi.DynElementAccessor, error) {
	if typeDef.TypeKind() != common.ElementTypeKind {
		panic(fmt.Sprintf("not an element type: %s", typeDef.InternalName()))
	}

	return &dynStruct{
		dynBase{
			dyn{typeDefContainer, parent},
		},
		typeDef,
		data,
	}, nil
}

func NewDynResource(typeDefContainer *common.TypeDefContainer, data map[string]interface{}, parent hi.DynAccessor) (hi.DynResourceAccessor, error) {
	resourceType, ok := data[ResourceTypePropName].(string)
	if !ok {
		return nil, fmt.Errorf("data contains no resource type")
	}

	typeDef := typeDefContainer.TypeByName(resourceType)
	if typeDef == nil {
		return nil, fmt.Errorf("resource type undefined: %s", resourceType)
	}
	structTypeDef, ok := typeDef.(common.StructTypeDefAccessor)
	if !ok {
		return nil, fmt.Errorf("no resource type: %s", resourceType)
	}
	if structTypeDef.TypeKind() != common.ResourceTypeKind {
		return nil, fmt.Errorf("no resource type: %s", resourceType)
	}

	return &dynStruct{
		dynBase{
			dyn{typeDefContainer, parent},
		},
		structTypeDef,
		data,
	}, nil
}

func (d *dyn) VersionString() string {
	return d.typeDefContainer.VersionString()
}

func (d *dyn) Parent() hi.DynAccessor {
	return d.parent
}

func (d *dynList) ItemTypeName() string {
	return d.itemTypeDef.Name()
}

func (d *dynPrimitive) TypeName() string {
	return d.primitiveTypeDef.Name()
}

func (d *dynPrimitive) BaseTypeName() string {
	base := d.primitiveTypeDef.Base()
	if base == nil {
		return ""
	}
	return base.Name()
}

func (d *dynPrimitive) Type() common.TypeDefAccessor {
	return d.primitiveTypeDef
}

func (d *dynPrimitive) SimpleType() hi.SimpleType {
	return hi.SimpleType(d.primitiveTypeDef.SimpleType())
}

func (d *dynPrimitive) NilValue() bool {
	return d.value == nil
}

func (d *dynPrimitive) StringValue() (string, error) {
	if d.primitiveTypeDef.SimpleType() != common.StringSimpleType {
		return "", fmt.Errorf("not a string primitive: %s", d.typeDef.InternalName())
	}
	if d.value == nil {
		return "", nil
	}
	if v, ok := d.value.(string); !ok {
		if f, ok := d.value.(float64); !ok {
			return "", fmt.Errorf("expected string value: %T", d.value)
		} else {
			return strconv.FormatFloat(f, 'f', -1, 64), nil
		}
	} else {
		return v, nil
	}
}

func (d *dynPrimitive) NumberValue() (float64, error) {
	if d.primitiveTypeDef.SimpleType() != common.NumberSimpleType {
		return 0.0, fmt.Errorf("not a number primitive: %s", d.typeDef.InternalName())
	}
	if d.value == nil {
		return 0.0, nil
	}
	if v, ok := d.value.(float64); !ok {
		return 0.0, fmt.Errorf("expected number value: %T", d.value)
	} else {
		return v, nil
	}
}

func (d *dynPrimitive) BoolValue() (bool, error) {
	if d.primitiveTypeDef.SimpleType() != common.BoolSimpleType {
		return false, fmt.Errorf("not a boolean primitive: %s", d.typeDef.InternalName())
	}
	if d.value == nil {
		return false, nil
	}
	if v, ok := d.value.(bool); !ok {
		return false, fmt.Errorf("expected boolean value: %T", d.value)
	} else {
		return v, nil
	}
}

func (d *dynStruct) TypeName() string {
	return d.typeDef.Name()
}

func (d *dynStruct) BaseTypeName() string {
	base := d.typeDef.Base()
	if base == nil {
		return ""
	}
	return base.Name()
}

func (d *dynStruct) Type() common.TypeDefAccessor {
	return d.typeDef
}

func (d *dynStruct) StringPropValue(name string) (string, error) {
	p, err := d.dynPrimitivePropByName(name, common.StringSimpleType)
	if p == nil || err != nil {
		return "", nil
	}
	return p.StringValue()
}

func (d *dynStruct) NumberPropValue(name string) (float64, error) {
	p, err := d.dynPrimitivePropByName(name, common.NumberSimpleType)
	if p == nil || err != nil {
		return 0.0, nil
	}
	return p.NumberValue()
}

func (d *dynStruct) BoolPropValue(name string) (bool, error) {
	p, err := d.dynPrimitivePropByName(name, common.BoolSimpleType)
	if p == nil || err != nil {
		return false, nil
	}
	return p.BoolValue()
}

func (d *dynStruct) Prop(name string) (hi.DynAccessor, error) {
	propDef, err := d.propDef(name)
	if err != nil {
		return nil, err
	}
	if propDef.Array() {
		return d.dynArrayProp(propDef)
	} else {
		return d.dynNonArrayProp(propDef, common.AnyTypeKind)
	}
}

func (d *dynStruct) ListProp(name string) (hi.DynListAccessor, error) {
	propDef, err := d.propDef(name)
	if err != nil {
		return nil, err
	}
	return d.dynArrayProp(propDef)
}

func (d *dynStruct) PrimitiveProp(name string) (hi.DynPrimitiveAccessor, error) {
	p, err := d.dynNonArrayPropByName(name, common.PrimitiveTypeKind)
	if p == nil || err != nil {
		return nil, err
	}
	return p.(hi.DynPrimitiveAccessor), nil
}

func (d *dynStruct) ElementProp(name string) (hi.DynElementAccessor, error) {
	p, err := d.dynNonArrayPropByName(name, common.ElementTypeKind)
	if p == nil || err != nil {
		return nil, err
	}
	return p.(hi.DynElementAccessor), nil
}

func (d *dynStruct) ResourceProp(name string) (hi.DynResourceAccessor, error) {
	p, err := d.dynNonArrayPropByName(name, common.ResourceTypeKind)
	if p == nil || err != nil {
		return nil, err
	}
	return p.(hi.DynResourceAccessor), nil
}

func (d *dynStruct) dynArrayProp(propDef *common.PropDef) (hi.DynListAccessor, error) {
	name := propDef.Name()
	if !propDef.Array() {
		return nil, fmt.Errorf("property %s of type %s is not a list",
			name, d.typeDef.InternalName())
	}

	if d.data == nil {
		return nil, nil
	}
	data := d.data[name]
	if data == nil {
		return nil, nil
	}

	dataArray, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("list property %s of type %s contains invalid list data with type %T",
			name, d.typeDef.InternalName(), data)
	}

	return NewDynList(d.typeDefContainer, propDef.Type(), dataArray, d)
}

func (d *dynStruct) dynPrimitivePropByName(name string, simpleType common.SimpleType) (hi.DynPrimitiveAccessor, error) {
	propDef, err := d.propDef(name)
	if err != nil {
		return nil, err
	}
	typeDef, ok := propDef.Type().(common.PrimitiveTypeDefAccessor)
	if !ok {
		return nil, fmt.Errorf("property %s of type %s has non-primitive type %s",
			name, d.typeDef.InternalName(), propDef.Type().InternalName())
	}
	if typeDef.SimpleType() != simpleType {
		return nil, fmt.Errorf("property %s of type %s has unexpected type %s",
			name, d.typeDef.InternalName(), propDef.Type().InternalName())
	}
	p, err := d.dynNonArrayProp(propDef, common.PrimitiveTypeKind)
	if p == nil || err != nil {
		return nil, err
	}
	return p.(hi.DynPrimitiveAccessor), nil
}

func (d *dynStruct) dynNonArrayPropByName(name string, kind common.TypeKind) (hi.DynAccessor, error) {
	propDef, err := d.propDef(name)
	if err != nil {
		return nil, err
	}
	return d.dynNonArrayProp(propDef, kind)
}

func (d *dynStruct) dynNonArrayProp(propDef *common.PropDef, kind common.TypeKind) (hi.DynAccessor, error) {
	if d.data == nil {
		return nil, nil
	}

	name := propDef.Name()
	propDefKind := propDef.Type().TypeKind()
	if kind != common.AnyTypeKind && propDefKind&kind == 0 {
		return nil, fmt.Errorf("property %s of type %s has type %s",
			name, d.typeDef.InternalName(), propDef.Type().InternalName())
	}
	if propDef.Array() {
		return nil, fmt.Errorf("property %s of type %s is a list",
			name, d.typeDef.InternalName())
	}

	var ok bool
	data := d.data[name]
	if propDefKind == common.PrimitiveTypeKind {
		element := d.data["_"+name]
		if data == nil && element == nil {
			return nil, nil
		}

		var elementMap map[string]interface{}
		if element != nil {
			elementMap, ok = element.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("primitive property %s of type %s contains invalid element data with type %T",
					name, d.typeDef.InternalName(), element)
			}
		}

		propType := propDef.Type().(common.PrimitiveTypeDefAccessor)
		return NewDynPrimitive(d.typeDefContainer, propType, data, elementMap, d)
	} else if data == nil {
		return nil, nil
	} else {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("struct property %s of type %s contains invalid data with type %T",
				name, d.typeDef.InternalName(), data)
		}

		if propDefKind == common.ResourceTypeKind {
			return NewDynResource(d.typeDefContainer, dataMap, d)
		} else {
			propType := propDef.Type().(common.StructTypeDefAccessor)
			return NewDynElement(d.typeDefContainer, propType, dataMap, d)
		}
	}
}

func (d *dynStruct) propDef(name string) (*common.PropDef, error) {
	propDef := d.typeDef.PropByName(name)
	if propDef == nil {
		return nil, fmt.Errorf("type %s has no property named %s", d.typeDef.InternalName(), name)
	}
	return propDef, nil
}
