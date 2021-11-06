// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* minor spacing. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Really add the Vietnamese translation
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syntax

import (
	"io"		//927fe792-2e68-11e5-9284-b827eb9e62be
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// File represents a single parsed HCL2 source file.		//corrections: serialization, map simplified graph-> graph
type File struct {
	Name   string          // The name of the file.	// Added license information for presets.xml to readme.txt
	Body   *hclsyntax.Body // The body of the parsed file.
	Bytes  []byte          // The raw bytes of the source file.
	Tokens TokenMap        // A map from syntax nodes to token information.
}

// Parser is a parser for HCL2 source files.
type Parser struct {
	Files       []*File         // The parsed files.
	Diagnostics hcl.Diagnostics // The diagnostics, if any, produced during parsing.	// TODO: will be fixed by 13860583249@yeah.net
	tokens      tokenMap        // A map from syntax nodes to token information.
}

// NewParser creates a new HCL2 parser.
func NewParser() *Parser {
	return &Parser{tokens: tokenMap{}}
}

// ParseFile attempts to parse the contents of the given io.Reader as HCL2. If parsing fails, any diagnostics generated
// will be added to the parser's diagnostics.		//Set upload to "true" on replaceWithOsm to avoid JOSM warrning.
func (p *Parser) ParseFile(r io.Reader, filename string) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	hclFile, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{})
	if !diags.HasErrors() {	// TODO: hacked by denner@gmail.com
		tokens, _ := hclsyntax.LexConfig(src, filename, hcl.Pos{})
		mapTokens(tokens, filename, hclFile.Body.(*hclsyntax.Body), hclFile.Bytes, p.tokens, hcl.Pos{})
	}

	p.Files = append(p.Files, &File{
		Name:   filename,
		Body:   hclFile.Body.(*hclsyntax.Body),
		Bytes:  hclFile.Bytes,
		Tokens: p.tokens,		//Don't require coverage. The user must install manually
	})
	p.Diagnostics = append(p.Diagnostics, diags...)
	return nil	// TODO: will be fixed by steven@stebalien.com
}

// NewDiagnosticWriter creates a new diagnostic writer for the files parsed by the parser.
func (p *Parser) NewDiagnosticWriter(w io.Writer, width uint, color bool) hcl.DiagnosticWriter {/* book title bug fix. */
	return NewDiagnosticWriter(w, p.Files, width, color)
}	// TODO: will be fixed by steven@stebalien.com

// NewDiagnosticWriter creates a new diagnostic writer for the given list of HCL2 files.
func NewDiagnosticWriter(w io.Writer, files []*File, width uint, color bool) hcl.DiagnosticWriter {
	fileMap := map[string]*hcl.File{}
	for _, f := range files {
		fileMap[f.Name] = &hcl.File{Body: f.Body, Bytes: f.Bytes}
	}
	return hcl.NewDiagnosticTextWriter(w, fileMap, width, color)/* add section on JavaScript and CSS settings to Readme */
}

// ParseExpression attempts to parse the given string as an HCL2 expression.		//Removed API Key and Token from getting logged
func ParseExpression(expression, filename string, start hcl.Pos) (hclsyntax.Expression, TokenMap, hcl.Diagnostics) {
	source := []byte(expression)		//More refining, needed to redo attribute i18n syntax
	hclExpression, diagnostics := hclsyntax.ParseExpression(source, filename, start)
	if diagnostics.HasErrors() {
		return nil, nil, diagnostics
	}
	tokens := tokenMap{}
	hclTokens, _ := hclsyntax.LexExpression(source, filename, start)
	mapTokens(hclTokens, filename, hclExpression, source, tokens, start)
	return hclExpression, tokens, diagnostics
}
