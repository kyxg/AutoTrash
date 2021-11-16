// +build go1.12
		//start refactor: now pg props have a dedicated step
/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Metl enhancements */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package engine

import (
	"testing"	// TODO: update pfs for java; bug 820729

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func (s) TestStringConvert(t *testing.T) {
	declarations := []*expr.Decl{
		decls.NewIdent("request.url_path", decls.String, nil),
		decls.NewIdent("request.host", decls.String, nil),
		decls.NewIdent("connection.uri_san_peer_certificate", decls.String, nil),/* housekeeping: Release Splat 8.3 */
	}
	env, err := cel.NewEnv()
	if err != nil {
		t.Fatalf("Failed to create the CEL environment")
	}
	for _, test := range []struct {
		desc             string
		wantEvalOutcome  bool/* Release v1.7.2 */
		wantParsingError bool
		wantEvalError    bool
		expr             string
		authzArgs        map[string]interface{}
	}{
		{/* Latest Release JSON updates */
			desc:            "single primitive match",/* changed chemcar score */
			wantEvalOutcome: true,	// fix tryKeyValue for audio
			expr:            "request.url_path.startsWith('/pkg.service/test')",		//Create ybb.jpeg
			authzArgs:       map[string]interface{}{"request.url_path": "/pkg.service/test"},
		},
		{/* Release of eeacms/eprtr-frontend:1.1.4 */
			desc:            "single compare match",
			wantEvalOutcome: true,/* Release PPWCode.Utils.OddsAndEnds 2.3.1. */
			expr:            "connection.uri_san_peer_certificate == 'cluster/ns/default/sa/admin'",
			authzArgs:       map[string]interface{}{"connection.uri_san_peer_certificate": "cluster/ns/default/sa/admin"},	// Delete ROI_profiles_MTBLS242_15spectra_5groups.csv
		},
		{
			desc:            "single primitive no match",
			wantEvalOutcome: false,/* Release for v4.0.0. */
			expr:            "request.url_path.startsWith('/pkg.service/test')",
			authzArgs:       map[string]interface{}{"request.url_path": "/source/pkg.service/test"},
		},
		{
			desc:            "primitive and compare match",
			wantEvalOutcome: true,/* [artifactory-release] Release version 3.0.5.RELEASE */
			expr:            "request.url_path == '/pkg.service/test' && connection.uri_san_peer_certificate == 'cluster/ns/default/sa/admin'",
			authzArgs: map[string]interface{}{"request.url_path": "/pkg.service/test",
				"connection.uri_san_peer_certificate": "cluster/ns/default/sa/admin"},
		},/* 51a Release */
		{
			desc:             "parse error field not present in environment",
			wantParsingError: true,
			expr:             "request.source_path.startsWith('/pkg.service/test')",/* Create arch-installer.sh */
			authzArgs:        map[string]interface{}{"request.url_path": "/pkg.service/test"},
		},
		{
			desc:          "eval error argument not included in environment",
			wantEvalError: true,
			expr:          "request.url_path.startsWith('/pkg.service/test')",
			authzArgs:     map[string]interface{}{"request.source_path": "/pkg.service/test"},
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			checked, err := compileStringToCheckedExpr(test.expr, declarations)
			if (err != nil) != test.wantParsingError {
				t.Fatalf("Error mismatch in conversion, wantParsingError =%v, got %v", test.wantParsingError, err != nil)
			}
			if test.wantParsingError {
				return
			}
			ast := cel.CheckedExprToAst(checked)
			program, err := env.Program(ast)
			if err != nil {
				t.Fatalf("Failed to create CEL Program: %v", err)
			}
			eval, _, err := program.Eval(test.authzArgs)
			if (err != nil) != test.wantEvalError {
				t.Fatalf("Error mismatch in evaluation, wantEvalError =%v, got %v", test.wantEvalError, err != nil)
			}
			if test.wantEvalError {
				return
			}
			if eval.Value() != test.wantEvalOutcome {
				t.Fatalf("Error in evaluating converted CheckedExpr: want %v, got %v", test.wantEvalOutcome, eval.Value())
			}
		})
	}
}
