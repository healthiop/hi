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
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var dateTimeTypeInfo = newElementTypeInfo("dateTime")

var timeZoneOffsetRegexp = regexp.MustCompile("^([+-])(\\d{1,2})(?::(\\d{1,2}))$")
var dateTimeRegexp = regexp.MustCompile("^(\\d(?:\\d(?:\\d[1-9]|[1-9]0)|[1-9]00)|[1-9]000)(?:-(0[1-9]|1[0-2])(?:-(0[1-9]|[1-2]\\d|3[0-1])(?:T([01]\\d|2[0-3]):([0-5]\\d):([0-5]\\d|60)(?:\\.(\\d+))?(Z|[+-](?:(?:0\\d|1[0-3]):[0-5]\\d|14:00)))?)?)?$")
var fluentDateTimeRegexp = regexp.MustCompile("^(\\d(?:\\d(?:\\d[1-9]|[1-9]0)|[1-9]00)|[1-9]000)(?:-(0[1-9]|1[0-2])(?:-(0[1-9]|[1-2]\\d|3[0-1]))?)?(?:T(?:([01]\\d|2[0-3])(?::([0-5]\\d)(?::([0-5]\\d|60)(?:\\.(\\d+))?)?)(Z|[+-](?:(?:0\\d|1[0-3]):[0-5]\\d|14:00))?)?)?$")

type DateTimeType struct {
	value     time.Time
	precision DateTimePrecisions
}

type DateTimeAccessor interface {
	PrimitiveAccessor
	Value() time.Time
	Precision() DateTimePrecisions
}

func NewDateTime(value time.Time) *DateTimeType {
	return newDateTime(value, NanoTimePrecision)
}

func ParseDateTime(value string) (*DateTimeType, error) {
	parts := dateTimeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid date/time string: %s", value)
	}
	return newDateTimeFromParts(parts), nil
}

func ParseFluentDateTime(value string) (*DateTimeType, error) {
	parts := fluentDateTimeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid fluent date/time string: %s", value)
	}
	return newDateTimeFromParts(parts), nil
}

func newDateTimeFromParts(parts []string) *DateTimeType {
	year, _ := strconv.Atoi(parts[1])
	precision := YearDatePrecision

	month := 1
	if parts[2] != "" {
		month, _ = strconv.Atoi(parts[2])
		precision = MonthDatePrecision
	}

	day := 1
	if parts[3] != "" {
		day, _ = strconv.Atoi(parts[3])
		precision = DayDatePrecision
	}

	hour := 0
	if parts[4] != "" {
		hour, _ = strconv.Atoi(parts[4])
		precision = HourTimePrecision
	}

	minute := 0
	if parts[5] != "" {
		minute, _ = strconv.Atoi(parts[5])
		precision = MinuteTimePrecision
	}

	second := 0
	if parts[6] != "" {
		second, _ = strconv.Atoi(parts[6])
		precision = SecondTimePrecision
	}

	nano := 0
	if parts[7] != "" {
		nano = parseNanosecond(parts[7])
		precision = NanoTimePrecision
	}

	location := mustEvalLocation(parts[8])
	value := time.Date(year, time.Month(month), day, hour, minute, second, nano, location)

	return newDateTime(value, precision)
}

func newDateTime(value time.Time, precision DateTimePrecisions) *DateTimeType {
	return &DateTimeType{
		value:     value,
		precision: precision,
	}
}

func mustEvalLocation(value string) *time.Location {
	if value == "" {
		return time.Local
	}
	if value == "Z" {
		return time.UTC
	}

	parts := timeZoneOffsetRegexp.FindStringSubmatch(value)
	if parts == nil {
		panic(fmt.Sprintf("not a valid time zone offset: %s", value))
	}

	hours, _ := strconv.Atoi(parts[2])
	offset := hours * 60 * 60

	if parts[3] != "" {
		minutes, _ := strconv.Atoi(parts[3])
		offset = offset + (minutes * 60)
	}

	if parts[1] == "-" {
		offset = -offset
	}

	if offset == 0 {
		return time.UTC
	}

	return time.FixedZone(string(offset), offset)
}

func (t *DateTimeType) DataType() DataTypes {
	return DateTimeDataType
}

func (t *DateTimeType) Value() time.Time {
	return t.value
}

func (t *DateTimeType) Precision() DateTimePrecisions {
	return t.precision
}

func (e *DateTimeType) TypeInfo() TypeInfoAccessor {
	return dateTimeTypeInfo
}
