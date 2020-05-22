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
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var timeTypeInfo = newElementTypeInfo("time")

var timeRegexp = regexp.MustCompile("^([01]\\d|2[0-3]):([0-5]\\d):([0-5]\\d|60)(?:\\.(\\d+))?$")

type timeType struct {
	TemporalType
	hour       int
	minute     int
	second     int
	nanosecond int
}

type TimeAccessor interface {
	TemporalAccessor
	Hour() int
	Minute() int
	Second() int
	Nanosecond() int
}

func NewTimeNil() TimeAccessor {
	return newTime(true, 0, 0, 0, 0, NanoTimePrecision)
}

func NewTime(value time.Time) TimeAccessor {
	return NewTimeHMSN(value.Hour(), value.Minute(), value.Second(), value.Nanosecond())
}

func NewTimeHMSN(hour int, minute int, second int, nanosecond int) TimeAccessor {
	return newTime(false, hour, minute, second, nanosecond, NanoTimePrecision)
}

func NewTimeHMSNWithPrecision(hour int, minute int, second int, nanosecond int,
	precision DateTimePrecisions) TimeAccessor {
	if precision <= HourTimePrecision {
		precision = HourTimePrecision
	} else if precision > NanoTimePrecision {
		precision = NanoTimePrecision
	}

	if precision < NanoTimePrecision {
		nanosecond = 0
	}
	if precision < SecondTimePrecision {
		second = 0
	}
	if precision < MinuteTimePrecision {
		minute = 0
	}

	return newTime(false, hour, minute, second, nanosecond, precision)
}

func ParseTime(value string) (TimeAccessor, error) {
	parts := timeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid time string: %s", value)
	}
	return newTimeFromParts(parts), nil
}

func newTimeFromParts(parts []string) TimeAccessor {
	hour, _ := strconv.Atoi(parts[1])
	precision := HourTimePrecision

	minute := 0
	if parts[2] != "" {
		minute, _ = strconv.Atoi(parts[2])
		precision = MinuteTimePrecision
	}

	second := 0
	if parts[3] != "" {
		second, _ = strconv.Atoi(parts[3])
		precision = SecondTimePrecision
	}

	nanosecond := 0
	if parts[4] != "" {
		nanosecond = parseNanosecond(parts[4])
		precision = NanoTimePrecision
	}

	return newTime(false, hour, minute, second, nanosecond, precision)
}

func newTime(nilValue bool, hour int, minute int, second int, nanosecond int, precision DateTimePrecisions) TimeAccessor {
	return &timeType{
		TemporalType: TemporalType{
			PrimitiveType: PrimitiveType{
				nilValue: nilValue,
			},
			precision: precision,
		},
		hour:       hour,
		minute:     minute,
		second:     second,
		nanosecond: nanosecond,
	}
}

func parseNanosecond(value string) int {
	if value == "" {
		return 0
	}
	nanoValue := value
	if len(nanoValue) > 9 {
		nanoValue = nanoValue[0:9]
	}
	nano, _ := strconv.Atoi(nanoValue)
	nano = nano * int(math.Pow10(9-len(nanoValue)))
	return nano
}

func (t *timeType) DataType() DataTypes {
	return TimeDataType
}

func (t *timeType) LowestPrecision() DateTimePrecisions {
	return HourTimePrecision
}

func (t *timeType) Hour() int {
	return t.hour
}

func (t *timeType) Minute() int {
	return t.minute
}

func (t *timeType) Second() int {
	return t.second
}

func (t *timeType) Nanosecond() int {
	return t.nanosecond
}

func (t *timeType) TypeInfo() TypeInfoAccessor {
	return timeTypeInfo
}

func (t *timeType) Equal(accessor Accessor) bool {
	if o, ok := accessor.(TimeAccessor); !ok {
		return false
	} else {
		return t.Precision() == o.Precision() &&
			t.Hour() == o.Hour() &&
			t.Minute() == o.Minute() &&
			t.Second() == o.Second() &&
			t.Nanosecond() == o.Nanosecond()
	}
}

func (t *timeType) Equivalent(accessor Accessor) bool {
	if o, ok := accessor.(TimeAccessor); !ok {
		return false
	} else {
		return timeValueEqual(t, o)
	}
}

func timeValueEqual(t1 TimeAccessor, t2 TimeAccessor) bool {
	return t1.Nil() == t2.Nil() && t1.Hour() == t2.Hour() &&
		t1.Minute() == t2.Minute() && t1.Second() == t2.Second() &&
		t1.Nanosecond() == t2.Nanosecond()
}

func (t *timeType) String() string {
	if t.nilValue {
		return ""
	}

	var b strings.Builder
	b.Grow(19)

	writeStringBuilderInt(&b, t.hour, 2)
	if t.precision >= MinuteTimePrecision {
		b.WriteByte(':')
		writeStringBuilderInt(&b, t.minute, 2)
	}
	if t.precision >= SecondTimePrecision {
		b.WriteByte(':')
		writeStringBuilderInt(&b, t.second, 2)
	}
	if t.precision >= NanoTimePrecision {
		b.WriteByte('.')
		writeStringBuilderInt(&b, t.nanosecond, 9)
	}

	return b.String()
}
