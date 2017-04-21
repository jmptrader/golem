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

type _null struct{}

var NULL Null = &_null{}

func (n *_null) TypeOf() (Type, Error) { return TNULL, nil }

func (n *_null) String() (Str, Error) { return MakeStr("null"), nil }

func (n *_null) Eq(v Value) (Bool, Error) {
	switch v.(type) {
	case *_null:
		return TRUE, nil
	default:
		return FALSE, nil
	}
}

func (n *_null) Cmp(v Value) (Int, Error) { return nil, NullValueError() }

func (n *_null) Add(v Value) (Value, Error) {
	switch t := v.(type) {

	case Str:
		return strcat([]Value{n, t})

	default:
		return nil, NullValueError()
	}
}
