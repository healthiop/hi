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

func TestDateTimeValue(t *testing.T) {
	testTime := time.Now().Add(-time.Hour * 78)
	o := NewDateTime(testTime)
	value := o.Value()
	assert.Equal(t, NanoTimePrecision, o.Precision())
	assert.True(t, testTime.Equal(value), "expected %d, got %d",
		testTime.UnixNano(), value.UnixNano())
}

func TestParseDateTimeCompleteTzPos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+02:00")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
		time.FixedZone("+02:00", 2*60*60))
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeCompleteTzNeg(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239-05:30")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
		time.FixedZone("-05:30", -19800))
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeCompleteTzZero(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+00:00")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeCompleteTzUtc(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeNoTz(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239")
	assert.Nil(t, dt, "unexpected date/time object")
	assert.NotNil(t, err, "expected error")
}

func TestParseDateTimeFractionDigits(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.2397381239Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239738123, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeNoNanos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, SecondTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 0, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeNoTime(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, DayDatePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeNoDay(t *testing.T) {
	dt, err := ParseDateTime("2015-02")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, MonthDatePrecision, dt.Precision())

	value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseDateTimeNoMonth(t *testing.T) {
	dt, err := ParseDateTime("2015")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, YearDatePrecision, dt.Precision())

	value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeCompleteTzPos(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239+02:00")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
		time.FixedZone("+02:00", 2*60*60))
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeCompleteTzNeg(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239-05:30")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
		time.FixedZone("-05:30", -19800))
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeCompleteTzUtc(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeNoTz(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.239")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeFractionDigits(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17.2397381239Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, NanoTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 239738123, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeNoNanos(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28:17Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, SecondTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 17, 0, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeNoSeconds(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13:28Z")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, MinuteTimePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 13, 28, 0, 0, time.UTC)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeNoMinutes(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T13Z")
	assert.Nil(t, dt, "unexpected date/time object")
	assert.NotNil(t, err, "expected error")
}

func TestParseFluentDateTimeNoTime(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02-07T")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, DayDatePrecision, dt.Precision())

	value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeNoDay(t *testing.T) {
	dt, err := ParseFluentDateTime("2015-02T")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, MonthDatePrecision, dt.Precision())

	value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestParseFluentDateTimeNoMonth(t *testing.T) {
	dt, err := ParseFluentDateTime("2015T")
	assert.NotNil(t, dt, "expected date/time object")
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, YearDatePrecision, dt.Precision())

	value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, value.Equal(dt.Value()), "expected %d, got %d",
		value.UnixNano(), dt.Value().UnixNano())
}

func TestMustEvalLocationInvalid(t *testing.T) {
	assert.Panics(t, func() { mustEvalLocation("X") })
}
