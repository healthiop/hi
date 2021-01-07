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

package dynamic

import (
	"github.com/healthiop/hi/internal/common"
	"github.com/healthiop/hipath/hipathsys"
)

const PathNamespaceName = "FHIR"

type PathTypeSpec struct {
	typeDef common.TypeDefAccessor
}

func (p *PathTypeSpec) Base() hipathsys.TypeSpecAccessor {
	return basePathTypeSpec(p)
}

func (p *PathTypeSpec) FQName() hipathsys.FQTypeNameAccessor {
	return p
}

func (p *PathTypeSpec) FQBaseName() hipathsys.FQTypeNameAccessor {
	return basePathTypeSpec(p)
}

func (p *PathTypeSpec) String() string {
	if p.typeDef.Anonymous() {
		return ""
	}
	return PathNamespaceName + "." + p.typeDef.InternalName()
}

func (p *PathTypeSpec) EqualType(other hipathsys.TypeSpecAccessor) bool {
	return hipathsys.FQTypeNameEqual(p, other.FQName())
}

func (p *PathTypeSpec) ExtendsName(name hipathsys.FQTypeNameAccessor) bool {
	if name.HasNamespace() && name.Namespace() != PathNamespaceName {
		return false
	}

	n := name.Name()
	td := p.typeDef
	for {
		if n == td.InternalName() && !td.Anonymous() {
			return true
		}
		td = td.Base()
		if td == nil {
			break
		}
	}

	return false
}

func (p *PathTypeSpec) Anonymous() bool {
	return p.Anonymous()
}

func (p *PathTypeSpec) HasNamespace() bool {
	return true
}

func (p *PathTypeSpec) Namespace() string {
	return PathNamespaceName
}

func (p *PathTypeSpec) Name() string {
	return p.typeDef.Name()
}

func (p *PathTypeSpec) Equal(name hipathsys.FQTypeNameAccessor) bool {
	return name.Namespace() == PathNamespaceName &&
		name.Name() == p.typeDef.InternalName() &&
		p.typeDef.Anonymous()
}

func basePathTypeSpec(p *PathTypeSpec) *PathTypeSpec {
	base := p.typeDef.Base()
	if base == nil {
		return nil
	}
	return &PathTypeSpec{base}
}
