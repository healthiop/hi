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
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestModelEqual(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model2 := readTestModel(t, "dynamic_model.json.golden")
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false),
		"same model must be equal")
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true),
		"same model must be equivalent")
}

func TestModelEqualTypeDiffers(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["multipleBirthInteger"] = "2"
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualLeftNil(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["link"] = nil
	model2 := readTestModel(t, "dynamic_model.json.golden")
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualRightNil(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["link"] = nil
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualLeftEmpty(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["link"] = make([]interface{}, 0)
	model2 := readTestModel(t, "dynamic_model.json.golden")
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualRightEmpty(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["link"] = make([]interface{}, 0)
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualBothEmpty(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["link"] = make([]interface{}, 0)
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["photo"] = nil
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualNilEmpty(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["link"] = nil
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["link"] = make([]interface{}, 0)
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualMissingRight(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["birthDate"] = "1990-08-21"
	model2 := readTestModel(t, "dynamic_model.json.golden")
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualMissingDiffer(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["telecom"] = nil
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["birthDate"] = "1990-08-21"
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualDifferentId(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["id"] = "abc123"
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["id"] = "abc124"
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualDifferentIdElement(t *testing.T) {
	id1 := make(map[string]interface{})
	id1["id"] = "abc123"
	id2 := make(map[string]interface{})
	id2["id"] = "abc124"

	model1 := readTestModel(t, "dynamic_model.json.golden")
	model1["_id"] = id1
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["_id"] = id2
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, true, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualArrayPropDiffers(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["name"].([]interface{})[0].(map[string]interface{})["family"] = "Brown"
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, true))
}

func TestModelEqualArraySizeDiffers(t *testing.T) {
	model1 := readTestModel(t, "dynamic_model.json.golden")
	model2 := readTestModel(t, "dynamic_model.json.golden")
	model2["name"] = append(model2["name"].([]interface{}), make(map[string]interface{}))
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, false))
	assert.Equal(t, false, modelComplexDeepEqual(model1, model2, true))
}

func TestModelDeepEqualInvalidType(t *testing.T) {
	assert.Panics(t, func() { modelDeepEqual(byte(1), byte(1), false) })
}

func readTestModel(t *testing.T, fileName string) map[string]interface{} {
	var model map[string]interface{}
	if err := json.Unmarshal(readTestFile(t, fileName), &model); err != nil {
		t.Fatal(err)
		return nil
	}
	return model
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
