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

package ast

import (
	"fmt"
)

//-------------------------------------
// Pos

type Pos struct {
	Line int
	Col  int
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.Line, p.Col)
}

//-------------------------------------
// Token

type TokenKind int

const (
	UNEXPECTED_CHAR TokenKind = iota
	UNEXPECTED_EOF
	badKind

	EOF

	PLUS
	MINUS
	MULT
	DIV
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	SEMICOLON
	COLON
	COMMA
	DOT

	EQ
	DBL_EQ
	NOT
	NOT_EQ
	GT
	GT_EQ
	LT
	LT_EQ
	CMP

	PIPE
	DBL_PIPE
	AMP
	DBL_AMP

	basicBegin
	NULL
	TRUE
	FALSE
	STR
	INT
	FLOAT
	basicEnd

	IDENT

	BLANK_IDENT
	IF
	ELSE
	WHILE
	BREAK
	CONTINUE
	FN
	RETURN
	CONST
	LET

	OBJ
	THIS
)

func (t TokenKind) String() string {
	switch t {
	case UNEXPECTED_CHAR:
		return "UNEXPECTED_CHAR"
	case UNEXPECTED_EOF:
		return "UNEXPECTED_EOF"
	case EOF:
		return "EOF"

	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULT:
		return "MULT"
	case DIV:
		return "DIV"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case EQ:
		return "EQ"
	case DBL_EQ:
		return "DBL_EQ"
	case SEMICOLON:
		return "SEMICOLON"
	case COLON:
		return "COLON"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"

	case NULL:
		return "NULL"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case STR:
		return "STR"
	case INT:
		return "INT"
	case FLOAT:
		return "FLOAT"

	case IDENT:
		return "IDENT"

	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case WHILE:
		return "WHILE"
	case BREAK:
		return "BREAK"
	case CONTINUE:
		return "CONTINUE"
	case FN:
		return "FN"
	case RETURN:
		return "RETURN"
	case CONST:
		return "CONST"
	case LET:
		return "LET"

	case OBJ:
		return "OBJ"
	case THIS:
		return "THIS"

	default:
		panic("unreachable")
	}
}

type Token struct {
	Kind     TokenKind
	Text     string
	Position Pos
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%v, %q, %v)", t.Kind, t.Text, t.Position)
}

func (t *Token) IsBad() bool {
	return t.Kind < badKind
}

func (t *Token) IsBasic() bool {
	return t.Kind > basicBegin && t.Kind < basicEnd
}