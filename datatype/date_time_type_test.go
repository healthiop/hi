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

func TestDateTimeImplementsAccessor(t *testing.T) {
	o := NewDateTime(time.Now())
	assert.Implements(t, (*DateTimeAccessor)(nil), o)
}

func TestDateTimeDataType(t *testing.T) {
	o := NewDateTime(time.Now())
	dataType := o.DataType()
	assert.Equal(t, DateTimeDataType, dataType)
}

func TestDateTimeTypeInfo(t *testing.T) {
	o := NewDateTime(time.Now())
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.dateTime", i.String())
	}
}

func TestNewDateTimeCollection(t *testing.T) {
	c := NewDateTimeCollection()
	assert.Equal(t, "FHIR.dateTime", c.ItemTypeInfo().String())
}

func TestDateTimeNil(t *testing.T) {
	o := NewDateTimeNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", o.Value().String())
	assert.Equal(t, NanoTimePrecision, o.Precision())
}

func TestDateTimeValue(t *testing.T) {
	testTime := time.Now().Add(-time.Hour * 78)
	o := NewDateTime(testTime)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	value := o.Value()
	assert.Equal(t, NanoTimePrecision, o.Precision())
	assert.True(t, testTime.Equal(value), "expected %d, got %d",
		testTime.UnixNano(), value.UnixNano())
}

func TestParseDateTimeCompleteTzPos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+02:00")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("+02:00", 2*60*60))
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeCompleteTzNeg(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239-05:30")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("-05:30", -19800))
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeCompleteTzZero(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+00:00")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeCompleteTzUtc(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoTz(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239")
	assert.Nil(t, dt, "unexpected date/time object")
	assert.NotNil(t, err, "expected error")
}

func TestParseDateTimeFractionDigits(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.2397381239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239738123, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoNanos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 0, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoTime(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, DayDatePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoDay(t *testing.T) {
	dt, err := ParseDateTime("2015-02")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, MonthDatePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoMonth(t *testing.T) {
	dt, err := ParseDateTime("2015")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, YearDatePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeCompleteTzPos(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239+02:00")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("+02:00", 2*60*60))
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeCompleteTzNeg(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239-05:30")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("-05:30", -19800))
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeCompleteTzUtc(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeNoTz(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeFractionDigits(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.2397381239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239738123, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeNoNanos(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 0, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeNoSeconds(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 0, 0, time.UTC)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, MinuteTimePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeNoMinutes(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13Z")
	assert.Nil(t, dt, "unexpected date/time object")
	assert.NotNil(t, err, "expected error")
}

func TestParseFluentDateTimeNoTime(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, DayDatePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeNoDay(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02T")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, MonthDatePrecision, dt.Precision())
	}
}

func TestParseFluentDateTimeNoMonth(t *testing.T) {
	dt, err := ParseFluentDateTime("2015T")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
			value.UnixNano(), dt.Value().UnixNano())
		assert.Equal(t, YearDatePrecision, dt.Precision())
	}
}

func TestMustEvalLocationInvalid(t *testing.T) {
	assert.Panics(t, func() { mustEvalLocation("X") })
}
