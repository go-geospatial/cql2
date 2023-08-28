// Copyright 2021-2023
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cql2

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

type CQL2 struct {
	Expressions []*Expression `@@`
}

type Expression struct {
	Lhs *Predicate `@@`
	Op  string     `( ( @Operator )? | ( @"AND" )? | ( @"OR" )? )!`
	Rhs *Predicate `@@`
}

type Predicate struct {
	Lhs string `@Ident`
	Op  string `@Operator`
	Rhs *Value `@@`
}

type IsLikePredicate struct {
	Lhs *CharacterExpression `@Ident`
	Not bool                 `"NOT"`
	Rhs *PatternExpression   `"LIKE" @@`
}

type CharacterExpression struct {
}

type PatternExpression struct {
}

type Value struct {
	Number        *float64 `  @Number`
	String        *string  `| @String`
	Identifier    *string  `| @Ident`
	Bool          *string  `| ( @"true" | "false" )`
	Nil           bool     `| @"nil"`
	SubExpression *CQL2    `| "(" @@ ")" `
}

var (
	cqlLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Operator", Pattern: `LIKE|=|<>|<|>|<=|>=`},
		{Name: "SpatialFunction", Pattern: `S_INTERSECTS|S_EQUALS|S_DISJOINT|S_TOUCHES|S_WITHIN|S_OVERLAPS|S_CROSSES|S_CONTAINS`},
		{Name: "TemporalFunction", Pattern: `T_AFTER|T_BEFORE|T_CONTAINS|T_DISJOINT|T_DURING|T_EQUALS|T_EQUALS|T_FINISHEDBY|T_FINISHES|T_INTERSECTS|T_MEETS|T_METBY|T_OVERLAPPEDBY|T_OVERLAPS|T_STARTEDBY|T_STARTS`},
		{Name: "ArrayFunction", Pattern: `A_EQUALS|A_CONTAINS|A_CONTAINEDBY|A_OVERLAPS`},
		{Name: "FullDate", Pattern: `\d{4}-\d{2}-\d{2}`},
		{Name: "Ident", Pattern: `[a-zA-Z_][a-zA-Z0-9_:]*`},
		{Name: "Number", Pattern: `[-+]?\d*\.?\d+([eE][-+]?\d+)?`},
		{Name: "String", Pattern: `'[^']*'|"[^"]*"`},
		{Name: "whitespace", Pattern: `\s+`},
	})
	parser = participle.MustBuild[CQL2](
		participle.Lexer(cqlLexer),
		participle.Unquote("String"),
		//		participle.CaseInsensitive("Operator"),
	)
)

func ParseText(s string) (*CQL2, error) {
	cql, err := parser.ParseString("", s)
	repr.Println(cql, repr.Indent("  "), repr.OmitEmpty(true))
	return cql, err
}
