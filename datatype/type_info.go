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

const FHIRNamespaceName = "FHIR"

var undefinedTypeInfo = NewTypeInfoWithBase(nil, nil)

type FQTypeName struct {
	namespace string
	name      string
	fqName    string
}

type FQTypeNameAccessor interface {
	Namespace() string
	Name() string
	String() string
	Equal(accessor FQTypeNameAccessor) bool
}

type TypeInfo struct {
	base   TypeInfoAccessor
	fqName FQTypeNameAccessor
}

type TypeInfoAccessor interface {
	Base() TypeInfoAccessor
	FQName() FQTypeNameAccessor
	FQBaseName() FQTypeNameAccessor
	String() string
	Equal(accessor TypeInfoAccessor) bool
}

func NewFQTypeName(name string, namespace string) *FQTypeName {
	var fqName string
	if len(namespace) > 0 {
		fqName = namespace + "." + name
	} else {
		fqName = name
	}

	return &FQTypeName{
		namespace: namespace,
		name:      name,
		fqName:    fqName,
	}
}

func NewTypeName(name string) *FQTypeName {
	return &FQTypeName{
		namespace: "",
		name:      name,
		fqName:    name,
	}
}

func FQTypeNameEqual(t1 FQTypeNameAccessor, t2 FQTypeNameAccessor) bool {
	return t1 == t2 || (t1 != nil && t2 != nil && t1.Equal(t2))
}

func NewTypeInfo(fqName FQTypeNameAccessor) *TypeInfo {
	return NewTypeInfoWithBase(fqName, nil)
}

func NewTypeInfoWithBase(fqName FQTypeNameAccessor, base TypeInfoAccessor) *TypeInfo {
	return &TypeInfo{
		base:   base,
		fqName: fqName,
	}
}

func (t *FQTypeName) Namespace() string {
	return t.namespace
}

func (t *FQTypeName) Name() string {
	return t.name
}

func (t *FQTypeName) String() string {
	return t.fqName
}

func (t *FQTypeName) Equal(accessor FQTypeNameAccessor) bool {
	return t.String() == accessor.String()
}

func (t *TypeInfo) Base() TypeInfoAccessor {
	return t.base
}

func (t *TypeInfo) FQName() FQTypeNameAccessor {
	return t.fqName
}

func (t *TypeInfo) FQBaseName() FQTypeNameAccessor {
	if t.base == nil {
		return nil
	}
	return t.base.FQName()
}

func (t *TypeInfo) String() string {
	if t.fqName == nil {
		return ""
	}
	return t.fqName.String()
}

func (t *TypeInfo) Equal(accessor TypeInfoAccessor) bool {
	return FQTypeNameEqual(t.FQName(), accessor.FQName())
}

func CommonBaseType(ti1 TypeInfoAccessor, ti2 TypeInfoAccessor) TypeInfoAccessor {
	for t1 := ti1; t1 != nil; t1 = t1.Base() {
		for t2 := ti2; t2 != nil; t2 = t2.Base() {
			if t1.Equal(t2) {
				return t1
			}
		}
	}
	return nil
}
