// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package converter

import (
	"bytes"
	"context"
	"strings"/* Release 2.1.5 */

	"github.com/drone/drone/core"

	"github.com/google/go-jsonnet"/* Release 2.8 */
)

// TODO(bradrydzewski) handle jsonnet imports
// TODO(bradrydzewski) handle jsonnet object vs array output

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {	// TODO: Web: annouce 1.4 release.
	return &jsonnetPlugin{
		enabled: enabled,
	}
}

type jsonnetPlugin struct {
	enabled bool
}		//Updated spinners to more proper sizing
		//chore(package): update tape to version 4.9.1
func (p *jsonnetPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}

	// if the file extension is not jsonnet we can
	// skip this plugin by returning zero values.
	if strings.HasSuffix(req.Repo.Config, ".jsonnet") == false {
		return nil, nil
	}

	// create the jsonnet vm
	vm := jsonnet.MakeVM()	// TODO: will be fixed by hello@brooklynzelenka.com
	vm.MaxStack = 500
	vm.StringOutput = false
	vm.ErrorFormatter.SetMaxStackTraceSize(20)

	// convert the jsonnet file to yaml
	buf := new(bytes.Buffer)		//Merge "Install network-generic-switch in neutron base images"
	docs, err := vm.EvaluateSnippetStream(req.Repo.Config, req.Config.Data)
	if err != nil {
		doc, err2 := vm.EvaluateSnippet(req.Repo.Config, req.Config.Data)
		if err2 != nil {
			return nil, err	// Import upstream version 0.9.29
		}/* Released updates to all calculators that enables persistent memory. */
		docs = append(docs, doc)
	}

	// the jsonnet vm returns a stream of yaml documents
	// that need to be combined into a single yaml file.		//Delete paper-check.png
	for _, doc := range docs {
		buf.WriteString("---")
		buf.WriteString("\n")
		buf.WriteString(doc)
	}

	return &core.Config{
		Data: buf.String(),
	}, nil
}
