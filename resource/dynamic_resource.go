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
	"github.com/volsch/gohimodel/datatype"
	"reflect"
)

const resourceTypePropName = "resourceType"

type DynamicModel map[string]interface{}

type DynamicResource struct {
	model DynamicModel
}

type DynamicResourceAccessor interface {
	Accessor
	internalModel() DynamicModel
}

func NewDynamicResource(resourceType string) *DynamicResource {
	data := make(DynamicModel)
	data[resourceTypePropName] = resourceType
	return NewDynamicResourceWithData(data)
}

func NewDynamicResourceWithData(model DynamicModel) *DynamicResource {
	return &DynamicResource{model}
}

func (r *DynamicResource) internalModel() DynamicModel {
	return r.model
}

func (r *DynamicResource) DataType() datatype.DataTypes {
	return datatype.ResourceDataType
}

func (r *DynamicResource) Empty() bool {
	for k, v := range r.model {
		if v != nil && k != resourceTypePropName {
			return false
		}
	}
	return true
}

func (r *DynamicResource) ResourceType() string {
	if val, found := r.model[resourceTypePropName]; found {
		if resourceType, ok := val.(string); ok {
			return resourceType
		}
	}
	return ""
}

func (r *DynamicResource) TypeInfo() datatype.TypeInfoAccessor {
	return datatype.NewTypeInfo(datatype.NewFQTypeName(r.ResourceType(), ""), nil)
}

func (r *DynamicResource) Equal(accessor datatype.Accessor) bool {
	if o, ok := accessor.(DynamicResourceAccessor); !ok {
		return false
	} else {
		return modelDeepEqual(r.internalModel(), o.internalModel())
	}
}

func modelDeepEqual(m1 map[string]interface{}, m2 map[string]interface{}) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		v2, found := m2[k1]
		if !found {
			return false
		}

		if v1 != v2 && isComplexProperty(v1) && isComplexProperty(v2) &&
			!modelDeepEqual(v1.(map[string]interface{}), v2.(map[string]interface{})) {
			return false
		}
	}

	return true
}

func isComplexProperty(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Map
}
