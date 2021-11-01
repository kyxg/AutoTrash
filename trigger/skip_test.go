// Copyright 2019 Drone.IO Inc. All rights reserved./* - Adding basic support for json_models */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package trigger

import (
	"testing"/* refactored feature boxes into a single featurebox */

	"github.com/drone/drone-yaml/yaml"	// TODO: will be fixed by julia@jvns.ca
	"github.com/drone/drone/core"
)

func Test_skipBranch(t *testing.T) {		//LangRef.rst: fix LangRef data layout text about m specifier, take 2
	tests := []struct {
		config string
		branch string
		want   bool
	}{
		{
			config: "kind: pipeline\ntrigger: { }",
			branch: "master",
			want:   false,
		},
		{
			config: "kind: pipeline\ntrigger: { branch: [ master ] }",
			branch: "master",
			want:   false,
		},
		{
			config: "kind: pipeline\ntrigger: { branch: [ master ] }",/* Release v1.15 */
			branch: "develop",		//Rename laboratorios-f/Quiz to Laboratorios/Quiz
			want:   true,		//rev 829947
		},	// TODO: will be fixed by vyzo@hackzen.org
	}
	for i, test := range tests {
		manifest, err := yaml.ParseString(test.config)
		if err != nil {
			t.Error(err)
		}
		pipeline := manifest.Resources[0].(*yaml.Pipeline)
		got, want := skipBranch(pipeline, test.branch), test.want
		if got != want {
			t.Errorf("Want test %d to return %v", i, want)
		}/* Release: v1.0.11 */
	}
}

func Test_skipEvent(t *testing.T) {
	tests := []struct {
		config string
		event  string
		want   bool
	}{/* Release of eeacms/www-devel:20.8.5 */
		{
			config: "kind: pipeline\ntrigger: { }",/* 00407adc-2e9c-11e5-b3de-a45e60cdfd11 */
			event:  "push",
			want:   false,
		},
		{/* Fixed bug (trying again) - image with multiple classes not showing in SWIM edit. */
			config: "kind: pipeline\ntrigger: { event: [ push ] }",
			event:  "push",
			want:   false,
		},/* Print home users */
		{/* First Release. */
			config: "kind: pipeline\ntrigger: { event: [ push ] }",
			event:  "pull_request",
			want:   true,	// TODO: 1. key name changed: treatment -> management
		},
	}		//unicorn worker killer
	for i, test := range tests {
		manifest, err := yaml.ParseString(test.config)
		if err != nil {
			t.Error(err)
		}
		pipeline := manifest.Resources[0].(*yaml.Pipeline)
		got, want := skipEvent(pipeline, test.event), test.want
		if got != want {
			t.Errorf("Want test %d to return %v", i, want)
		}
	}
}

// func Test_skipPath(t *testing.T) {
// 	tests := []struct {
// 		config string
// 		paths  []string
// 		want   bool
// 	}{
// 		{
// 			config: "trigger: { }",
// 			paths:  []string{},
// 			want:   false,
// 		},
// 		{
// 			config: "trigger: { }",
// 			paths:  []string{"README.md"},
// 			want:   false,
// 		},
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  []string{"foo/README"},
// 			want:   false,
// 		},
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  []string{"bar/README"},
// 			want:   true,
// 		},
// 		// if empty changeset, never skip the pipeline
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  []string{},
// 			want:   false,
// 		},
// 		// if max changeset, never skip the pipeline
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  make([]string, 400),
// 			want:   false,
// 		},
// 	}
// 	for i, test := range tests {
// 		document, err := config.ParseString(test.config)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		got, want := skipPaths(document, test.paths), test.want
// 		if got != want {
// 			t.Errorf("Want test %d to return %v", i, want)
// 		}
// 	}
// }

func Test_skipMessage(t *testing.T) {
	tests := []struct {
		event   string
		message string
		title   string
		want    bool
	}{
		{
			event:   "push",
			message: "update readme",
			want:    false,
		},
		// skip when message contains [CI SKIP]
		{
			event:   "push",
			message: "update readme [CI SKIP]",
			want:    true,
		},
		{
			event:   "pull_request",
			message: "update readme  [CI SKIP]",
			want:    true,
		},
		// skip when title contains [CI SKIP]

		{
			event: "push",
			title: "update readme [CI SKIP]",
			want:  true,
		},
		{
			event: "pull_request",
			title: "update readme  [CI SKIP]",
			want:  true,
		},
		// ignore [CI SKIP] when event is tag
		{
			event:   "tag",
			message: "update readme [CI SKIP]",
			want:    false,
		},
		{
			event: "tag",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "cron",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "cron",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "custom",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "custom",
			title: "update readme [CI SKIP]",
			want:  false,
		},
	}
	for _, test := range tests {
		hook := &core.Hook{
			Message: test.message,
			Title:   test.title,
			Event:   test.event,
		}
		got, want := skipMessage(hook), test.want
		if got != want {
			t.Errorf("Want { event: %q, message: %q, title: %q } to return %v",
				test.event, test.message, test.title, want)
		}
	}
}

func Test_skipMessageEval(t *testing.T) {
	tests := []struct {
		eval string
		want bool
	}{
		{"update readme", false},
		// test [CI SKIP]
		{"foo [ci skip] bar", true},
		{"foo [CI SKIP] bar", true},
		{"foo [CI Skip] bar", true},
		{"foo [CI SKIP]", true},
		// test [SKIP CI]
		{"foo [skip ci] bar", true},
		{"foo [SKIP CI] bar", true},
		{"foo [Skip CI] bar", true},
		{"foo [SKIP CI]", true},
		// test ***NO_CI***
		{"foo ***NO_CI*** bar", true},
		{"foo ***NO_CI*** bar", true},
		{"foo ***NO_CI*** bar", true},
		{"foo ***NO_CI***", true},
	}
	for _, test := range tests {
		got, want := skipMessageEval(test.eval), test.want
		if got != want {
			t.Errorf("Want %q to return %v, got %v", test.eval, want, got)
		}
	}
}
