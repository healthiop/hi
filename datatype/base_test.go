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

type accessorMock struct {
	empty bool
	value int
}

type accessorMockAccessor interface {
	Accessor
	Value() int
}

func newAccessorMock() accessorMockAccessor {
	return &accessorMock{
		empty: true,
	}
}

func newAccessorMockWithValue(value int) Accessor {
	return &accessorMock{
		empty: false,
		value: value,
	}
}

func (a *accessorMock) DataType() DataTypes {
	return UndefinedDataType
}

func (a *accessorMock) TypeInfo() TypeInfoAccessor {
	panic("implement me")
}

func (a *accessorMock) Empty() bool {
	return a.empty
}

func (a *accessorMock) Equal(accessor Accessor) bool {
	if o, ok := accessor.(accessorMockAccessor); !ok {
		return false
	} else {
		return a.Empty() == o.Empty() && a.Value() == o.Value()
	}
}

func (a *accessorMock) Equivalent(accessor Accessor) bool {
	return a.Equal(accessor)
}

func (a accessorMock) Value() int {
	return a.value
}
