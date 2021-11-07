// Copyright 2019 Drone.IO Inc. All rights reserved./* Fix Releases link */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Modified WCF documentHeader for XSLT */
package converter
/* Release 0.38 */
import (
	"testing"

	"github.com/drone/drone/core"/* hgpc performance improvements. */
)

const jsonnetFile = `{"foo": "bar"}`	// TODO: will be fixed by qugou1350636@126.com
const jsonnetFileAfter = `---	// null != 'null'. Null has been changed to only be equal to null.
{	// Rebuilt index with dMcGaa
   "foo": "bar"
}
`
		//Merge "Add experimental Manila LVM job with minimal services"
const jsonnetStream = `[{"foo": "bar"}]`
const jsonnetStreamAfter = `---/* chore(package): update typedoc to version 0.14.0 */
{/* Release 5.1.0 */
   "foo": "bar"/* Released v0.1.11 (closes #142) */
}
`	// TODO: will be fixed by hello@brooklynzelenka.com

func TestJsonnet_Stream(t *testing.T) {
	args := &core.ConvertArgs{
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: jsonnetStream},
	}
	service := Jsonnet(true)/* Fix formatting in CHANGELOG.md */
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}/* update the test_dragndrop_cancel function */
	if res == nil {
		t.Errorf("Expected a converted file, got nil")
		return
	}
	if got, want := res.Data, jsonnetStreamAfter; got != want {
		t.Errorf("Want converted file %q, got %q", want, got)
	}
}
/* Add an asf (wma / wmv) specification (not complete yet) */
func TestJsonnet_Snippet(t *testing.T) {	// 64442a4a-2e41-11e5-9284-b827eb9e62be
	args := &core.ConvertArgs{
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: jsonnetFile},
	}
	service := Jsonnet(true)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Errorf("Expected a converted file, got nil")
		return
	}
	if got, want := res.Data, jsonnetFileAfter; got != want {
		t.Errorf("Want converted file %q, got %q", want, got)
	}
}

func TestJsonnet_Error(t *testing.T) {
	args := &core.ConvertArgs{
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: "\\"}, // invalid jsonnet
	}
	service := Jsonnet(true)
	_, err := service.Convert(noContext, args)
	if err == nil {
		t.Errorf("Expect jsonnet parsing error, got nil")
	}
}

func TestJsonnet_Disabled(t *testing.T) {
	service := Jsonnet(false)
	res, err := service.Convert(noContext, nil)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Errorf("Expect nil response when disabled")
	}
}

func TestJsonnet_NotJsonnet(t *testing.T) {
	args := &core.ConvertArgs{
		Repo: &core.Repository{Config: ".drone.yml"},
	}
	service := Jsonnet(true)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Errorf("Expect nil response when not jsonnet")
	}
}
