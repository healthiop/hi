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

package datatype

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateDataType(t *testing.T) {
	var a DateAccessor = NewDateType(time.Now())
	dataType := a.DataType()
	assert.Equal(t, DateDataType, dataType)
}

func TestDateValue(t *testing.T) {
	testTime := time.Now().Add(-time.Hour * 78)
	var a DateAccessor = NewDateType(testTime)

	value := a.Value()
	assert.Equal(t, testTime.Year(), a.Year())
	assert.Equal(t, int(testTime.Month()), a.Month())
	assert.Equal(t, testTime.Day(), a.Day())
	assert.Equal(t, DayDatePrecision, a.Precision())

	expectedTime := time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.Local)
	assert.True(t, expectedTime.Equal(value), "expected %d, got %d",
		expectedTime.UnixNano(), value.UnixNano())
}

func TestDateYMD(t *testing.T) {
	var a DateAccessor = NewDateTypeYMD(2020, 4, 23)

	assert.Equal(t, 2020, a.Year())
	assert.Equal(t, 4, a.Month())
	assert.Equal(t, 23, a.Day())
	assert.Equal(t, DayDatePrecision, a.Precision())
}

func TestParseDateValueComplete(t *testing.T) {
	dt, err := ParseDateValue("2015-02-07")
	assert.NotNil(t, dt, "expected date object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, DayDatePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateValueInvalid(t *testing.T) {
	dt, err := ParseDateValue("2015-02-0A")
	assert.Nil(t, dt, "unexpected date object")
	assert.NotNil(t, err, "expected error")
}

func TestParseDateValueNoDay(t *testing.T) {
	dt, err := ParseDateValue("2015-02")
	assert.NotNil(t, dt, "expected date object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, MonthDatePrecision, dt.Precision())

	value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateValueNoMonth(t *testing.T) {
	dt, err := ParseDateValue("2015")
	assert.NotNil(t, dt, "expected date object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, YearDatePrecision, dt.Precision())

	value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}
