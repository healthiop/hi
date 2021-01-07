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
	"fmt"
	"github.com/healthiop/hi"
	"github.com/healthiop/hi/internal/common"
	"github.com/healthiop/hipath/hipathsys"
)

const UCUMSystemURI = "http://unitsofmeasure.org"

type PathDynModel struct {
	typeDefContainer *common.TypeDefContainer
}

func (p *PathDynModel) ConvertToSystem(node interface{}) (interface{}, error) {
	if p, ok := node.(hi.DynPrimitiveAccessor); ok {
		if p.NilValue() {
			return nil, nil
		}
		switch p.TypeName() {
		case common.BooleanTypeName:
			v, err := p.BoolValue()
			if err != nil {
				return nil, err
			}
			return hipathsys.NewBooleanWithSource(v, p), nil
		case common.StringTypeName, common.URITypeName, common.CodeTypeName, common.OIDTypeName, common.IDTypeName, common.UUIDTypeName, common.SIDTypeName, common.MarkdownTypeName, common.Base64BinaryTypeName:
			v, err := p.StringValue()
			if err != nil {
				return nil, err
			}
			return hipathsys.NewStringWithSource(v, p), nil
		case common.IntegerTypeName, common.PositiveIntTypeName, common.UnsignedIntTypeName:
			v, err := p.NumberValue()
			if err != nil {
				return nil, err
			}
			return hipathsys.NewIntegerWithSource(int32(v), p), nil
		case common.DecimalTypeName:
			v, err := p.StringValue()
			if err != nil {
				return nil, err
			}
			return hipathsys.ParseDecimalWithSource(v, p)
		case common.DateTypeName, common.DateTimeTypeName, common.InstantTypeName:
			v, err := p.StringValue()
			if err != nil {
				return nil, err
			}
			return hipathsys.ParseDateTimeWithSource(v, p)
		case common.TimeTypeName:
			v, err := p.StringValue()
			if err != nil {
				return nil, err
			}
			return hipathsys.ParseTimeWithSource(v, p)
		default:
			return nil, nil
		}
	}
	if e, ok := node.(hi.DynElementAccessor); ok {
		return p.convertElementToSystem(e)
	}
	return p, nil
}

func (p *PathDynModel) convertElementToSystem(element hi.DynElementAccessor) (interface{}, error) {
	if element.TypeName() != common.QuantityTypeName {
		return element, nil
	}

	system, err := element.StringPropValue("system")
	if err != nil {
		return nil, fmt.Errorf("error when accessing system property of Quantity: %w", err)
	}
	if system != UCUMSystemURI {
		return element, nil
	}

	code, err := element.StringPropValue("code")
	if err != nil {
		return nil, fmt.Errorf("error when accessing code property of Quantity: %w", err)
	}
	if len(code) == 0 {
		return element, nil
	}

	valueProp, err := element.PrimitiveProp("value")
	if valueProp == nil || err != nil {
		return nil, err
	}
	value, err := p.ConvertToSystem(valueProp)
	if value == nil || err != nil {
		return nil, err
	}
	decimalValue, ok := value.(hipathsys.DecimalAccessor)
	if !ok {
		return nil, fmt.Errorf("property value of Quantity is no decimal: %T", value)
	}

	var convertedCode string
	switch code {
	case "a":
		convertedCode = "year"
	case "mo":
		convertedCode = "month"
	case "p":
		convertedCode = "day"
	case "h":
		convertedCode = "hour"
	case "min":
		convertedCode = "minute"
	case "s":
		convertedCode = "second"
	default:
		convertedCode = code
	}

	return hipathsys.NewQuantityWithSource(decimalValue, hipathsys.NewString(convertedCode), element), nil
}

func (p *PathDynModel) TypeSpec(node interface{}) hipathsys.TypeSpecAccessor {
	typeDefRetriever, ok := node.(DynTypeDefRetriever)
	if !ok {
		return nil
	}
	return &PathTypeSpec{typeDefRetriever.Type()}
}

func (p *PathDynModel) Cast(node interface{}, name hipathsys.FQTypeNameAccessor) (interface{}, error) {
	panic("implement me")
}

func (p *PathDynModel) Equal(node1 interface{}, node2 interface{}) bool {
	panic("implement me")
}

func (p *PathDynModel) Equivalent(node1 interface{}, node2 interface{}) bool {
	panic("implement me")
}

func (p *PathDynModel) Navigate(node interface{}, name string) (interface{}, error) {
	panic("implement me")
}

func (p *PathDynModel) Children(node interface{}) (hipathsys.CollectionAccessor, error) {
	panic("implement me")
}
