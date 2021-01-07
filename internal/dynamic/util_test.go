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
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestDynamicEqual(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false),
		"same dynamic must be equal")
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true),
		"same dynamic must be equivalent")
}

func TestDynamicEqualTypeDiffers(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["multipleBirthInteger"] = "2"
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualLeftNil(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["link"] = nil
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualRightNil(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["link"] = nil
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualLeftEmpty(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["link"] = make([]interface{}, 0)
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualRightEmpty(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["link"] = make([]interface{}, 0)
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualBothEmpty(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["link"] = make([]interface{}, 0)
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["photo"] = nil
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualNilEmpty(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["link"] = nil
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["link"] = make([]interface{}, 0)
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualMissingRight(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["birthDate"] = "1990-08-21"
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualMissingDiffer(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["telecom"] = nil
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["birthDate"] = "1990-08-21"
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualDifferentId(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["id"] = "abc123"
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["id"] = "abc124"
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualDifferentIdElement(t *testing.T) {
	id1 := make(map[string]interface{})
	id1["id"] = "abc123"
	id2 := make(map[string]interface{})
	id2["id"] = "abc124"

	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic1["_id"] = id1
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["_id"] = id2
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, true, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualArrayPropDiffers(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["name"].([]interface{})[0].(map[string]interface{})["family"] = "Brown"
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicEqualArraySizeDiffers(t *testing.T) {
	dynamic1 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2 := readTestDynamic(t, "dynamic.json.golden")
	dynamic2["name"] = append(dynamic2["name"].([]interface{}), make(map[string]interface{}))
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, false))
	assert.Equal(t, false, dynamicStructDeepEqual(dynamic1, dynamic2, true))
}

func TestDynamicDeepEqualInvalidType(t *testing.T) {
	assert.Panics(t, func() { dynamicDeepEqual(byte(1), byte(1), false) })
}

func readTestDynamic(t *testing.T, fileName string) map[string]interface{} {
	var dynamic map[string]interface{}
	if err := json.Unmarshal(readTestFile(t, fileName), &dynamic); err != nil {
		t.Fatal(err)
		return nil
	}
	return dynamic
}

func readTestFile(t *testing.T, fileName string) []byte {
	golden := filepath.Join("testdata", fileName)
	if content, err := ioutil.ReadFile(golden); err != nil {
		t.Fatal(err)
		return nil
	} else {
		return content
	}
}
