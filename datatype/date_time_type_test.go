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

func TestDateTimeTypeLowestPrecision(t *testing.T) {
	o := NewDateTime(time.Now())
	assert.Equal(t, YearDatePrecision, o.LowestPrecision())
}

func TestDateTimeTypeInfo(t *testing.T) {
	o := NewDateTime(time.Now())
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.dateTime", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
		}
	}
}

func TestDateTimeNil(t *testing.T) {
	o := NewDateTimeNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", o.Time().String())
	assert.Equal(t, 1, o.Year())
	assert.Equal(t, 1, o.Month())
	assert.Equal(t, 1, o.Day())
	assert.Equal(t, NanoTimePrecision, o.Precision())
	assert.Equal(t, "", o.String())
}

func TestDateTimeValue(t *testing.T) {
	testTime := time.Date(2018, 5, 20, 17, 48, 14, 123, time.Local)
	o := NewDateTime(testTime)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	value := o.Time()
	assert.Equal(t, testTime.Year(), o.Year())
	assert.Equal(t, int(testTime.Month()), o.Month())
	assert.Equal(t, testTime.Day(), o.Day())
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
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000+02:00", dt.String())
	}
}

func TestParseDateTimeCompleteTzNeg(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239-05:30")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("-05:30", -19800))
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000-05:30", dt.String())
	}
}

func TestParseDateTimeCompleteTzZero(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+00:00")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000Z", dt.String())
	}
}

func TestParseDateTimeCompleteTzUtc(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000Z", dt.String())
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
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoNanos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 13, 28, 17, 0, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17Z", dt.String())
	}
}

func TestParseDateTimeNoTime(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, DayDatePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07", dt.String())
	}
}

func TestParseDateTimeNoDay(t *testing.T) {
	dt, err := ParseDateTime("2015-02")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, MonthDatePrecision, dt.Precision())
		assert.Equal(t, "2015-02", dt.String())
	}
}

func TestParseDateTimeNoMonth(t *testing.T) {
	dt, err := ParseDateTime("2015")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		assert.False(t, dt.Nil(), "non-nil data type expected")
		value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, YearDatePrecision, dt.Precision())
		assert.Equal(t, "2015", dt.String())
	}
}

func TestMustEvalLocationInvalid(t *testing.T) {
	assert.Panics(t, func() { mustEvalLocation("X") })
}

func TestDateTimeEqualNil(t *testing.T) {
	assert.Equal(t, false, NewDateTime(time.Now()).Equal(nil))
}

func TestDateTimeEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewDateTime(time.Now()).Equal(newAccessorMock()))
	assert.Equal(t, false, NewDateTime(time.Now()).Equivalent(newAccessorMock()))
}

func TestDateTimeEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewDateTimeNil().Equal(NewDateTime(time.Now())))
	assert.Equal(t, false, NewDateTimeNil().Equivalent(NewDateTime(time.Now())))
}

func TestDateTimeEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewDateTime(time.Now()).Equal(NewDateTimeNil()))
	assert.Equal(t, false, NewDateTime(time.Now()).Equivalent(NewDateTimeNil()))
}

func TestDateTimeEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewDateTimeNil().Equal(NewDateTimeNil()))
	assert.Equal(t, true, NewDateTimeNil().Equivalent(NewDateTimeNil()))
}

func TestDateTimeEqualEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, true, NewDateTime(now).Equal(NewDateTime(now)))
	assert.Equal(t, true, NewDateTime(now).Equivalent(NewDateTime(now)))
}

func TestDateTimeEqualPrecisionDiffer(t *testing.T) {
	dt1 := NewDateTimeWithPrecision(time.Date(2015, 2, 7, 13, 28, 17, 123, time.Local), NanoTimePrecision)
	dt2 := NewDateTimeWithPrecision(time.Date(2015, 2, 7, 13, 28, 17, 123, time.Local), SecondTimePrecision)
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equal(dt2))
		assert.Equal(t, false, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEquivalent(t *testing.T) {
	dt1 := NewDateTimeWithPrecision(time.Date(2015, 2, 7, 13, 28, 0, 123, time.Local), SecondTimePrecision)
	dt2 := NewDateTimeWithPrecision(time.Date(2015, 2, 7, 13, 28, 17, 123, time.Local), MinuteTimePrecision)
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equal(dt2))
		assert.Equal(t, true, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEqualDifferentTemporal(t *testing.T) {
	dt1 := NewDateTimeWithPrecision(time.Date(2015, 2, 7, 13, 28, 17, 123, time.Local), DayDatePrecision)
	dt2 := NewDateYMDWithPrecision(2015, 2, 20, MonthDatePrecision)
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equal(dt2))
	}
}

func TestDateTimeEquivalentDifferentTemporal(t *testing.T) {
	dt1 := NewDateTimeWithPrecision(time.Date(2015, 2, 1, 13, 28, 17, 123, time.Local), DayDatePrecision)
	dt2 := NewDateYMDWithPrecision(2015, 2, 20, MonthDatePrecision)
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, true, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEquivalentTime(t *testing.T) {
	dt1 := NewDateTimeWithPrecision(time.Date(2015, 2, 1, 10, 10, 10, 10, time.Local), SecondTimePrecision)
	dt2 := NewTimeHMSNWithPrecision(10, 10, 10, 10, SecondTimePrecision)
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEqualNotEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, false, NewDateTime(now).Equal(NewDateTime(now.Add(time.Hour))))
	assert.Equal(t, false, NewDateTime(now).Equivalent(NewDateTime(now.Add(time.Hour))))
}

func TestNewDateTimeWithPrecisionNano(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		NanoTimePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 8, v.Day())
	assert.Equal(t, 15, v.Hour())
	assert.Equal(t, 12, v.Minute())
	assert.Equal(t, 19, v.Second())
	assert.Equal(t, 982, v.Nanosecond())
	assert.Equal(t, NanoTimePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionGreaterNano(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		200)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 8, v.Day())
	assert.Equal(t, 15, v.Hour())
	assert.Equal(t, 12, v.Minute())
	assert.Equal(t, 19, v.Second())
	assert.Equal(t, 982, v.Nanosecond())
	assert.Equal(t, NanoTimePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionSecond(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		SecondTimePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 8, v.Day())
	assert.Equal(t, 15, v.Hour())
	assert.Equal(t, 12, v.Minute())
	assert.Equal(t, 19, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, SecondTimePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionMinute(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		MinuteTimePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 8, v.Day())
	assert.Equal(t, 15, v.Hour())
	assert.Equal(t, 12, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, MinuteTimePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionHour(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		HourTimePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 8, v.Day())
	assert.Equal(t, 15, v.Hour())
	assert.Equal(t, 0, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, HourTimePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionDay(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		DayDatePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 8, v.Day())
	assert.Equal(t, 0, v.Hour())
	assert.Equal(t, 0, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, DayDatePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionMonth(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		MonthDatePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 2, v.Month())
	assert.Equal(t, 1, v.Day())
	assert.Equal(t, 0, v.Hour())
	assert.Equal(t, 0, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, MonthDatePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionYear(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		YearDatePrecision)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 1, v.Month())
	assert.Equal(t, 1, v.Day())
	assert.Equal(t, 0, v.Hour())
	assert.Equal(t, 0, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, YearDatePrecision, v.Precision())
}

func TestNewDateTimeWithPrecisionLessYear(t *testing.T) {
	v := NewDateTimeWithPrecision(time.Date(2020, 2, 8, 15, 12, 19, 982, time.Local),
		-1)
	assert.Equal(t, 2020, v.Year())
	assert.Equal(t, 1, v.Month())
	assert.Equal(t, 1, v.Day())
	assert.Equal(t, 0, v.Hour())
	assert.Equal(t, 0, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, YearDatePrecision, v.Precision())
}
