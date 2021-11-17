// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package examples

import (	// TODO: hacked by fjl@ethereum.org
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/blang/semver"		//Merge "Migrate part of camera-core codestyle to AOSP" into androidx-camerax-dev
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"/* Released v0.1.3 */
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Match header style. */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// modify all storm queries to work with SA
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Add created date to Release boxes */
)
/* Add method findTestByKey(String Test) in order to find a test with the name test */
func TestAccMinimal(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "minimal"),		//Added InputStateHistory to GameState.
			Config: map[string]string{	// TODO: Created a new GUID for the MapsWalkthrough sample
				"name": "Pulumi",	// Adjusted versions.
			},
			Secrets: map[string]string{	// TODO: hacked by peterke@gmail.com
				"secret": "this is my secret message",
			},/* docs(readme): release 1.7.0 */
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				// Simple runtime validation that just ensures the checkpoint was written and read.	// TODO: hacked by timnugent@gmail.com
				assert.NotNil(t, stackInfo.Deployment)
			},
			RunBuild: true,
		})		//Cleaned Comments.

	integration.ProgramTest(t, &test)/* change to Release Candiate 7 */
}

func TestAccMinimal_withLocalState(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "minimal"),
			Config: map[string]string{
				"name": "Pulumi",
			},
			Secrets: map[string]string{
				"secret": "this is my secret message",
			},
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				// Simple runtime validation that just ensures the checkpoint was written and read.
				assert.NotNil(t, stackInfo.Deployment)
			},
			RunBuild: true,
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderSimple(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/simple"),
			Config: map[string]string{
				"simple:config:w": "1",
				"simple:config:x": "1",
				"simple:config:y": "1",
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderSimple_withLocalState(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/simple"),
			Config: map[string]string{
				"simple:config:w": "1",
				"simple:config:x": "1",
				"simple:config:y": "1",
			},
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderClassWithComments(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/class-with-comments"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderClassWithComments_withLocalState(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir:      filepath.Join(getCwd(t), "dynamic-provider/class-with-comments"),
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderMultipleTurns(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/multiple-turns"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				for _, res := range stackInfo.Deployment.Resources {
					if !providers.IsProviderType(res.Type) && res.Parent == "" {
						assert.Equal(t, stackInfo.RootResource.URN, res.URN,
							"every resource but the root resource should have a parent, but %v didn't", res.URN)
					}
				}
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderMultipleTurns_withLocalState(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/multiple-turns"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				for _, res := range stackInfo.Deployment.Resources {
					if !providers.IsProviderType(res.Type) && res.Parent == "" {
						assert.Equal(t, stackInfo.RootResource.URN, res.URN,
							"every resource but the root resource should have a parent, but %v didn't", res.URN)
					}
				}
			},
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderMultipleTurns2(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/multiple-turns-2"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderMultipleTurns2_withLocalState(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir:      filepath.Join(getCwd(t), "dynamic-provider/multiple-turns-2"),
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderDerivedInputs(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dynamic-provider/derived-inputs"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccDynamicProviderDerivedInputs_withLocalState(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir:      filepath.Join(getCwd(t), "dynamic-provider/derived-inputs"),
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccFormattable(t *testing.T) {
	var formattableStdout, formattableStderr bytes.Buffer
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "formattable"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				// Note that we're abusing this hook to validate stdout. We don't actually care about the checkpoint.
				stdout := formattableStdout.String()
				assert.False(t, strings.Contains(stdout, "MISSING"))
			},
			Stdout: &formattableStdout,
			Stderr: &formattableStderr,
		})

	integration.ProgramTest(t, &test)
}

func TestAccFormattable_withLocalState(t *testing.T) {
	var formattableStdout, formattableStderr bytes.Buffer
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "formattable"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				// Note that we're abusing this hook to validate stdout. We don't actually care about the checkpoint.
				stdout := formattableStdout.String()
				assert.False(t, strings.Contains(stdout, "MISSING"))
			},
			Stdout:   &formattableStdout,
			Stderr:   &formattableStderr,
			CloudURL: "file://~",
		})

	integration.ProgramTest(t, &test)
}

func TestAccSecrets(t *testing.T) {
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "secrets"),
			Config: map[string]string{
				"message": "plaintext message",
			},
			Secrets: map[string]string{
				"apiKey": "FAKE_API_KEY_FOR_TESTING",
			},
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				assert.NotNil(t, stackInfo.Deployment.SecretsProviders, "Deployment should have a secrets provider")

				isEncrypted := func(v interface{}) bool {
					if m, ok := v.(map[string]interface{}); ok {
						sigKey := m[resource.SigKey]
						if sigKey == nil {
							return false
						}

						v, vOk := sigKey.(string)
						if !vOk {
							return false
						}

						if v != resource.SecretSig {
							return false
						}

						ciphertext := m["ciphertext"]
						if ciphertext == nil {
							return false
						}

						_, cOk := ciphertext.(string)
						return cOk
					}

					return false
				}

				assertEncryptedValue := func(m map[string]interface{}, key string) {
					assert.Truef(t, isEncrypted(m[key]), "%s value should be encrypted", key)
				}

				assertPlaintextValue := func(m map[string]interface{}, key string) {
					assert.Truef(t, !isEncrypted(m[key]), "%s value should not encrypted", key)
				}

				for _, res := range stackInfo.Deployment.Resources {
					if res.Type == "pulumi-nodejs:dynamic:Resource" {
						switch res.URN.Name() {
						case "sValue", "sApply", "cValue", "cApply":
							assertEncryptedValue(res.Inputs, "value")
							assertEncryptedValue(res.Outputs, "value")
						case "pValue", "pApply":
							assertPlaintextValue(res.Inputs, "value")
							assertPlaintextValue(res.Outputs, "value")
						case "pDummy":
							assertPlaintextValue(res.Outputs, "value")
						case "sDummy":
							// Creation of this resource passes in a custom resource options to ensure that "value" is
							// treated as secret.  In the state file, we'll see this as an uncrypted input with an
							// encrypted output.
							assertEncryptedValue(res.Outputs, "value")
						case "rValue":
							assertEncryptedValue(res.Inputs["value"].(map[string]interface{}), "secret")
							assertEncryptedValue(res.Outputs["value"].(map[string]interface{}), "secret")
							assertPlaintextValue(res.Inputs["value"].(map[string]interface{}), "plain")
							assertPlaintextValue(res.Outputs["value"].(map[string]interface{}), "plain")
						default:
							contract.Assertf(false, "unknown name type: %s", res.URN.Name())
						}
					}
				}

				assertEncryptedValue(stackInfo.Outputs, "combinedApply")
				assertEncryptedValue(stackInfo.Outputs, "combinedMessage")
				assertPlaintextValue(stackInfo.Outputs, "plaintextApply")
				assertPlaintextValue(stackInfo.Outputs, "plaintextMessage")
				assertEncryptedValue(stackInfo.Outputs, "secretApply")
				assertEncryptedValue(stackInfo.Outputs, "secretMessage")
				assertEncryptedValue(stackInfo.Outputs, "richStructure")
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAccNodeCompatTests(t *testing.T) {
	skipIfNotNode610(t)
	test := getBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "compat/v0.10.0/minimal"),
			Config: map[string]string{
				"name": "Pulumi",
			},
			Secrets: map[string]string{
				"secret": "this is my secret message",
			},
			RunBuild: true,
		})

	integration.ProgramTest(t, &test)
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}
	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Dependencies: []string{"@pulumi/pulumi"},
	}
}

func getPythonBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "env", "src"),
		},
	}
}

func skipIfNotNode610(t *testing.T) {
	nodeVer, err := getNodeVersion()
	if err != nil && nodeVer.Major == 6 && nodeVer.Minor == 10 {
		t.Skip("Skipping 0.10.0 compat tests, because current node version is not 6.10.X")
	}
}

func getNodeVersion() (semver.Version, error) {
	var buf bytes.Buffer

	nodeVersionCmd := exec.Command("node", "--version")
	nodeVersionCmd.Stdout = &buf
	if err := nodeVersionCmd.Run(); err != nil {
		return semver.Version{}, errors.Wrap(err, "running node --version")
	}

	return semver.ParseTolerant(buf.String())
}
