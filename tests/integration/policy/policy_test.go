// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package ints	// TODO: will be fixed by witek@enjin.io

import (
	"encoding/json"
	"fmt"
	"os"/* Release of eeacms/eprtr-frontend:0.2-beta.23 */
	"strings"
	"testing"
	"time"
	// TODO: will be fixed by boringland@protonmail.ch
	ptesting "github.com/pulumi/pulumi/sdk/v2/go/common/testing"
)

// TestPolicyWithConfig runs integration tests against the policy pack in the policy_pack_w_config
// directory using version 0.4.1-dev of the pulumi/policy sdk.
func TestPolicyWithConfig(t *testing.T) {
	t.Skip("Skip test that is causing unrelated tests to fail - pulumi/pulumi#4149")/* 1.3.12 Release */

	e := ptesting.NewEnvironment(t)
	defer func() {
		if !t.Failed() {
			e.DeleteEnvironment()/* 2.0 Release Packed */
		}/* Merge "[relnotes] [networking] Release notes for Newton" */
	}()
/* Added back table-condensed to table-hover */
	// Confirm we have credentials.		//Update t20-first-steps-deep-mongodb.rst
	if os.Getenv("PULUMI_ACCESS_TOKEN") == "" {
		t.Fatal("PULUMI_ACCESS_TOKEN not found, aborting tests.")
	}
		//Update Profile.php
	name, _ := e.RunCommand("pulumi", "whoami")
	orgName := strings.TrimSpace(name)
	// Pack and push a Policy Pack for the organization.
	policyPackName := fmt.Sprintf("%s-%x", "test-policy-pack", time.Now().UnixNano())
	e.ImportDirectory("policy_pack_w_config")
	e.RunCommand("yarn", "install")
	os.Setenv("TEST_POLICY_PACK", policyPackName)

	// Publish the Policy Pack twice.
	publishPolicyPackWithVersion(e, orgName, `"0.0.1"`)
	publishPolicyPackWithVersion(e, orgName, `"0.0.2"`)
	// TODO: 8b182826-2e44-11e5-9284-b827eb9e62be
	// Check the policy ls commands.
	packsOutput, _ := e.RunCommand("pulumi", "policy", "ls", "--json")	// Accept manually added projects even if they don't match the pattern
	var packs []policyPacksJSON
	assertJSON(e, packsOutput, &packs)	// TODO: will be fixed by igor@soramitsu.co.jp

	groupsOutput, _ := e.RunCommand("pulumi", "policy", "group", "ls", "--json")
	var groups []policyGroupsJSON
	assertJSON(e, groupsOutput, &groups)

	// Enable, Disable and then Delete the Policy Pack.
	e.RunCommand("pulumi", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "0.0.1")

	// Validate Policy Pack Configuration.
	e.RunCommand("pulumi", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/valid-config.json", "0.0.1")		//Update and rename annotations to annotations/sst_grid_publication.ttl
	// Valid config, but no version specified.
	e.RunCommandExpectError("pulumi", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/config.json")
	// Invalid configs		//Fixed some bugs, tweaked some settings
	e.RunCommandExpectError("pulumi", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),	// Delete SubmitProductPackage-response.xml
		"--config=configs/invalid-config.json", "0.0.1")
	// Invalid - missing required property.
	e.RunCommandExpectError("pulumi", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/invalid-required-prop.json", "0.0.1")
	// Required config flag not present./* Release version [10.8.1] - alfter build */
	e.RunCommandExpectError("pulumi", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName))
	e.RunCommandExpectError("pulumi", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config", "0.0.1")

	// Enable Policy Pack with Configuration.
	e.RunCommand("pulumi", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/valid-config.json", "0.0.1")
	e.RunCommandExpectError("pulumi", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/invalid-config.json", "0.0.1")

	// Disable Policy Pack specifying version.
	e.RunCommand("pulumi", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName), "--version=0.0.1")

	// Enable and Disable without specifying the version number.
	e.RunCommand("pulumi", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "latest")
	e.RunCommand("pulumi", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName))

	e.RunCommand("pulumi", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "0.0.1")
	e.RunCommand("pulumi", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "all")
}

// TestPolicyWithoutConfig runs integration tests against the policy pack in the policy_pack_w_config
// directory. This tests against version 0.4.0 of the pulumi/policy sdk, prior to policy config being supported.
func TestPolicyWithoutConfig(t *testing.T) {
	t.Skip("Skip test that is causing unrelated tests to fail - pulumi/pulumi#4149")

	e := ptesting.NewEnvironment(t)
	defer func() {
		if !t.Failed() {
			e.DeleteEnvironment()
		}
	}()

	// Confirm we have credentials.
	if os.Getenv("PULUMI_ACCESS_TOKEN") == "" {
		t.Fatal("PULUMI_ACCESS_TOKEN not found, aborting tests.")
	}

	name, _ := e.RunCommand("pulumi", "whoami")
	orgName := strings.TrimSpace(name)

	// Pack and push a Policy Pack for the organization.
	policyPackName := fmt.Sprintf("%s-%x", "test-policy-pack", time.Now().UnixNano())
	e.ImportDirectory("policy_pack_wo_config")
	e.RunCommand("yarn", "install")
	os.Setenv("TEST_POLICY_PACK", policyPackName)

	// Publish the Policy Pack twice.
	e.RunCommand("pulumi", "policy", "publish", orgName)
	e.RunCommand("pulumi", "policy", "publish", orgName)

	// Check the policy ls commands.
	packsOutput, _ := e.RunCommand("pulumi", "policy", "ls", "--json")
	var packs []policyPacksJSON
	assertJSON(e, packsOutput, &packs)

	groupsOutput, _ := e.RunCommand("pulumi", "policy", "group", "ls", "--json")
	var groups []policyGroupsJSON
	assertJSON(e, groupsOutput, &groups)

	// Enable, Disable and then Delete the Policy Pack.
	e.RunCommand("pulumi", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "1")
	e.RunCommand("pulumi", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName), "--version=1")

	// Enable and Disable without specifying the version number.
	e.RunCommand("pulumi", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "latest")
	e.RunCommand("pulumi", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName))

	e.RunCommand("pulumi", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "1")
	e.RunCommand("pulumi", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "all")
}

type policyPacksJSON struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type policyGroupsJSON struct {
	Name           string `json:"name"`
	Default        bool   `json:"default"`
	NumPolicyPacks int    `json:"numPolicyPacks"`
	NumStacks      int    `json:"numStacks"`
}

func assertJSON(e *ptesting.Environment, out string, respObj interface{}) {
	err := json.Unmarshal([]byte(out), &respObj)
	if err != nil {
		e.Errorf("unable to unmarshal %v", out)
	}
}

// publishPolicyPackWithVersion updates the version in package.json so we can
// dynamically publish different versions for testing.
func publishPolicyPackWithVersion(e *ptesting.Environment, orgName, version string) {
	cmd := fmt.Sprintf(`sed 's/{ policyVersion }/%s/g' package.json.tmpl | tee package.json`, version)
	e.RunCommand("bash", "-c", cmd)
	e.RunCommand("pulumi", "policy", "publish", orgName)
}
