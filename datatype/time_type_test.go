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

func TestTimeImplementsAccessor(t *testing.T) {
	o := NewTime(time.Now())
	assert.Implements(t, (*TimeAccessor)(nil), o)
}

func TestTimeDataType(t *testing.T) {
	o := NewTime(time.Now())
	dataType := o.DataType()
	assert.Equal(t, TimeDataType, dataType)
}

func TestTimeTypeInfo(t *testing.T) {
	o := NewTime(time.Now())
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.time", i.String())
	}
}

func TestNewTimeCollection(t *testing.T) {
	c := NewTimeCollection()
	assert.Equal(t, "FHIR.time", c.ItemTypeInfo().String())
}

func TestTimeNil(t *testing.T) {
	o := NewTimeNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, 0, o.Hour())
	assert.Equal(t, 0, o.Minute())
	assert.Equal(t, 0, o.Second())
	assert.Equal(t, 0, o.Nanosecond())
	assert.Equal(t, NanoTimePrecision, o.Precision())
}

func TestTimeValue(t *testing.T) {
	testTime := time.Now().Add(-time.Hour * 78)
	o := NewTime(testTime)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")

	value := o.Value()
	assert.Equal(t, testTime.Hour(), o.Hour())
	assert.Equal(t, testTime.Minute(), o.Minute())
	assert.Equal(t, testTime.Second(), o.Second())
	assert.Equal(t, testTime.Nanosecond(), o.Nanosecond())
	assert.Equal(t, NanoTimePrecision, o.Precision())

	expectedTime := time.Now()
	expectedTime = time.Date(expectedTime.Year(), expectedTime.Month(), expectedTime.Day(),
		testTime.Hour(), testTime.Minute(), testTime.Second(), testTime.Nanosecond(), testTime.Location())
	assert.True(t, expectedTime.Equal(value), "expected %d, got %d",
		expectedTime.UnixNano(), value.UnixNano())
}

func TestTimeYMD(t *testing.T) {
	o := NewTimeHMSN(16, 28, 47, 837173635)

	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, 16, o.Hour())
	assert.Equal(t, 28, o.Minute())
	assert.Equal(t, 47, o.Second())
	assert.Equal(t, 837173635, o.Nanosecond())
	assert.Equal(t, NanoTimePrecision, o.Precision())
}

func TestParseTimeComplete(t *testing.T) {
	dt, err := ParseTime("13:28:17.239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 239000000, dt.Nanosecond())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseTimeInvalid(t *testing.T) {
	dt, err := ParseTime("13:28:17.A")
	assert.Nil(t, dt, "unexpected date object")
	assert.NotNil(t, err, "expected error")
}

func TestParseTimeFractionDigits(t *testing.T) {
	dt, err := ParseTime("13:28:17.2397381239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 239738123, dt.Nanosecond())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseTimeNoNanos(t *testing.T) {
	dt, err := ParseTime("13:28:17")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
	}
}

func TestParseFluentTimeValueComplete(t *testing.T) {
	dt, err := ParseFluentTime("13:28:17.239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 239000000, dt.Nanosecond())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentTimeValueInvalid(t *testing.T) {
	dt, err := ParseFluentTime("13:28:17.A")
	assert.Nil(t, dt, "unexpected date object")
	assert.NotNil(t, err, "expected error")
}

func TestParseFluentTimeValueFractionDigits(t *testing.T) {
	dt, err := ParseFluentTime("13:28:17.2397381239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 239738123, dt.Nanosecond())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentTimeValueNoNanos(t *testing.T) {
	dt, err := ParseFluentTime("13:28:17")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
	}
}

func TestParseFluentTimeValueNoSeconds(t *testing.T) {
	dt, err := ParseFluentTime("13:28")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 0, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, MinuteTimePrecision, dt.Precision())
	}
}

func TestParseFluentTimeValueNoMinutes(t *testing.T) {
	dt, err := ParseFluentTime("13")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 0, dt.Minute())
		assert.Equal(t, 0, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, HourTimePrecision, dt.Precision())
	}
}

func TestParseNanosecondEmpty(t *testing.T) {
	assert.Equal(t, 0, parseNanosecond(""))
}
