// Copyright 2017 The Golem Project Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"bytes"
	"reflect"
)

//---------------------------------------------------------------
// An ObjDef contains the information needed to instantiate an Obj
// instance.  ObjDefs are created at compile time, and
// are immutable at run time.

type ObjDef struct {
	Keys []string
}

//---------------------------------------------------------------
// _obj

type _obj struct {
	fields map[string]Value
	inited bool
}

func NewObj() Obj {
	return &_obj{nil, false}
}

func (o *_obj) Init(def *ObjDef, vals []Value) {
	o.fields = make(map[string]Value)
	for i, k := range def.Keys {
		o.fields[k] = vals[i]
	}
	o.inited = true
}

func (o *_obj) TypeOf() (Type, Error) {
	if !o.inited {
		return TOBJ, UninitializedObjError()
	}

	return TOBJ, nil
}

func (o *_obj) String() (Str, Error) {
	if !o.inited {
		return nil, UninitializedObjError()
	}

	if len(o.fields) == 0 {
		return MakeStr("obj {}"), nil
	}

	var buf bytes.Buffer
	buf.WriteString("obj {")
	idx := 0
	for k, v := range o.fields {
		if idx > 0 {
			buf.WriteString(",")
		}
		idx = idx + 1
		buf.WriteString(" ")
		buf.WriteString(k)
		buf.WriteString(": ")

		s, err := v.String()
		if err != nil {
			return nil, err
		}
		buf.WriteString(s.StrVal())
	}
	buf.WriteString(" }")
	return MakeStr(buf.String()), nil
}

func (o *_obj) Eq(v Value) (Bool, Error) {
	if !o.inited {
		return FALSE, UninitializedObjError()
	}

	switch t := v.(type) {
	case *_obj:
		return MakeBool(reflect.DeepEqual(o.fields, t.fields)), nil
	default:
		return FALSE, nil
	}
}

func (o *_obj) Cmp(v Value) (Int, Error) {
	if !o.inited {
		return nil, UninitializedObjError()
	}

	return nil, TypeMismatchError("Expected Comparable Type")
}

func (o *_obj) Add(v Value) (Value, Error) {
	if !o.inited {
		return nil, UninitializedObjError()
	}

	switch t := v.(type) {

	case Str:
		return strcat([]Value{o, t})

	default:
		return nil, TypeMismatchError("Expected Number Type")
	}
}

func (o *_obj) GetField(key Str) (Value, Error) {
	if !o.inited {
		return nil, UninitializedObjError()
	}

	v, ok := o.fields[key.StrVal()]
	if ok {
		return v, nil
	} else {
		return nil, NoSuchFieldError(key.StrVal())
	}
}

func (o *_obj) PutField(key Str, val Value) Error {
	if !o.inited {
		return UninitializedObjError()
	}

	_, ok := o.fields[key.StrVal()]
	if ok {
		o.fields[key.StrVal()] = val
		return nil
	} else {
		return NoSuchFieldError(key.StrVal())
	}
}
