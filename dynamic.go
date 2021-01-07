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

package hi

type SimpleType int

const (
	StringSimpleType SimpleType = 1
	NumberSimpleType SimpleType = 2
	BoolSimpleType   SimpleType = 3
)

type DynAccessor interface {
	VersionString() string
	Parent() DynAccessor
}

type DynListAccessor interface {
	DynAccessor
	ItemTypeName() string
}

type DynBaseAccessor interface {
	DynAccessor
	TypeName() string
	BaseTypeName() string
}

type DynStructAccessor interface {
	DynBaseAccessor
	Prop(name string) (DynAccessor, error)
	ListProp(name string) (DynListAccessor, error)
	PrimitiveProp(name string) (DynPrimitiveAccessor, error)
	ElementProp(name string) (DynElementAccessor, error)
	ResourceProp(name string) (DynResourceAccessor, error)
	StringPropValue(name string) (string, error)
	NumberPropValue(name string) (float64, error)
	BoolPropValue(name string) (bool, error)
}

type DynElementAccessor interface {
	DynStructAccessor
}

type DynPrimitiveAccessor interface {
	DynElementAccessor
	SimpleType() SimpleType
	NilValue() bool
	StringValue() (string, error)
	NumberValue() (float64, error)
	BoolValue() (bool, error)
}

type DynResourceAccessor interface {
	DynStructAccessor
}
