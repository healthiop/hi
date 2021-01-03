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

package resource

import (
	"fmt"
	"reflect"
)

func modelComplexDeepEqual(model1 map[string]interface{}, model2 map[string]interface{}, equivalent bool) bool {
	var m1, m2 map[string]interface{}
	if len(model1) >= len(model2) {
		m1 = model1
		m2 = model2
	} else {
		m1 = model2
		m2 = model1
	}

	notFoundCount := 0
	for k1, v1 := range m1 {
		if !equivalent || !isIgnoredEquivalencePropName(k1) {
			v2, found := m2[k1]
			if !found {
				if !isEmptyDynamicValue(v1) {
					return false
				}
				notFoundCount = notFoundCount + 1
			}
			if found && !modelDeepEqual(v1, v2, equivalent) {
				return false
			}
		}
	}

	if notFoundCount > 0 {
		for k2, v2 := range m2 {
			_, found := m1[k2]
			if !found && !isEmptyDynamicValue(v2) {
				return false
			}
		}
	}

	return true
}

func modelCollectionDeepEqual(c1 []interface{}, c2 []interface{}, equivalent bool) bool {
	l := len(c1)
	if l != len(c2) {
		return false
	}

	if l > 0 {
		for pos, v1 := range c1 {
			if !modelDeepEqual(v1, c2[pos], equivalent) {
				return false
			}
		}
	}

	return true
}

func modelDeepEqual(v1 interface{}, v2 interface{}, equivalent bool) bool {
	if v1 == nil && v2 == nil {
		return true
	} else if (v1 == nil || v2 == nil) && isEmptyDynamicValue(v1) && isEmptyDynamicValue(v2) {
		return true
	} else {
		k := reflect.TypeOf(v1).Kind()
		if k != reflect.TypeOf(v2).Kind() {
			return false
		}
		switch k {
		case reflect.String, reflect.Float64, reflect.Bool:
			if v1 != v2 {
				return false
			}
		case reflect.Map:
			if !modelComplexDeepEqual(v1.(map[string]interface{}), v2.(map[string]interface{}), equivalent) {
				return false
			}
		case reflect.Slice:
			if !modelCollectionDeepEqual(v1.([]interface{}), v2.([]interface{}), equivalent) {
				return false
			}
		default:
			panic(fmt.Sprintf("Unhandled JSON type: %d", k))
		}
	}
	return true
}

func isEmptyDynamicValue(value interface{}) bool {
	if value == nil {
		return true
	}
	kind := reflect.TypeOf(value).Kind()
	return (kind == reflect.Map && len(value.(map[string]interface{})) == 0) ||
		(kind == reflect.Slice && len(value.([]interface{})) == 0)
}

func isIgnoredEquivalencePropName(propName string) bool {
	return propName == "id" || propName == "_id"
}
